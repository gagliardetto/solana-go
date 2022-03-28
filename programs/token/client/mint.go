package client

import (
	"errors"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/system"
	tkn "github.com/gagliardetto/solana-go/programs/token"
	"github.com/gagliardetto/solana-go/rpc"
)

func (c *Client) GetMint(mint solana.PublicKey) (*tkn.Mint, error) {
	result, err := c.rpc.GetAccountInfoWithOpts(c.ctx, mint, &rpc.GetAccountInfoOpts{Commitment: c.commitment, Encoding: "base64"})
	if err != nil {
		return nil, err
	}
	data, err := getBinaryDataFromAccountInfo(result)
	if err != nil {
		return nil, err
	}

	tokenInfo := new(tkn.Mint)

	err = tokenInfo.UnmarshalWithDecoder(bin.NewBinDecoder(data))
	//err = tokenInfo.Decode(data)
	if err != nil {
		return nil, err
	}
	return tokenInfo, nil
}

func getBinaryDataFromAccountInfo(result *rpc.GetAccountInfoResult) ([]byte, error) {
	v := result.Value
	if v == nil {
		return nil, errors.New("no value")
	}
	data := v.Data
	if data == nil {
		return nil, errors.New("no data")
	}
	return data.GetBinary(), nil
}

type TokenArgs struct {
	Decimals        uint8
	InitialSupply   uint64
	MintAuthority   solana.PublicKey
	FreezeAuthority solana.PublicKey
}

// create a token; optionally mint tokens to a destination
func (c *Client) CreateToken(args *TokenArgs, payerPrivateKey solana.PrivateKey, owner solana.PublicKey) (*tkn.Mint, error) {
	keyMap := make(map[string]*solana.PrivateKey)

	payer := payerPrivateKey.PublicKey()
	keyMap[payer.String()] = &payerPrivateKey

	destinationForTokensPrivateKey, err := solana.NewRandomPrivateKey()
	if err != nil {
		return nil, err
	}
	destination := destinationForTokensPrivateKey.PublicKey()
	keyMap[destination.String()] = &destinationForTokensPrivateKey

	mintPrivateKey, err := solana.NewRandomPrivateKey()
	if err != nil {
		return nil, err
	}
	mint := mintPrivateKey.PublicKey()
	keyMap[mint.String()] = &mintPrivateKey

	if payer == mint {
		return nil, errors.New("payer cannot be the same as mint")
	}

	blockHash, err := c.latest_blockhash()
	if err != nil {
		return nil, err
	}
	list := []solana.Instruction{}

	inst_1, err := c.instruction_create_account_for_mint(payer, mint)
	if err != nil {
		return nil, err
	}
	list = append(list, inst_1)
	// spl_token::instruction::initialize_mint
	inst_2, err := c.instruction_initialize_mint(args.Decimals, mint, args.MintAuthority, args.FreezeAuthority)
	if err != nil {
		return nil, err
	}
	list = append(list, inst_2)

	inst_3, err := c.instruction_create_account_for_token(payer, destination)
	if err != nil {
		return nil, err
	}
	list = append(list, inst_3)
	// spl_token::instruction::initialize_account
	inst_4, err := c.instruction_initialize_account_for_token(mint, destination, owner)
	if err != nil {
		return nil, err
	}
	list = append(list, inst_4)

	if 0 < args.InitialSupply {
		inst_n, err := c.instruction_mint_to(mint, args.MintAuthority, destination, args.InitialSupply)
		if err != nil {
			return nil, err
		}
		list = append(list, inst_n)
	}

	tx, err := solana.NewTransaction(
		list,
		blockHash,
		solana.TransactionPayer(payer),
	)
	if err != nil {
		return nil, err
	}

	_, err = tx.Sign(func(key solana.PublicKey) *solana.PrivateKey {
		k, present := keyMap[key.String()]
		if present {
			return k
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	sig, err := c.rpc.SendTransactionWithOpts(c.ctx, tx, false, c.commitment)
	if err != nil {
		return nil, err
	}

	err = c.block_on_tx_processing(sig)
	if err != nil {
		return nil, err
	}

	return c.GetMint(mint)
}

// spl_token::instruction::mint_to
func (c *Client) instruction_mint_to(mint solana.PublicKey, mintAuthority solana.PublicKey, destination solana.PublicKey, supply uint64) (*tkn.Instruction, error) {
	return tkn.NewMintToInstructionBuilder().SetAmount(supply).SetDestinationAccount(destination).SetMintAccount(mint).SetAuthorityAccount(mintAuthority).ValidateAndBuild()
}

// system program; create an account with space to initialize a mint
func (c *Client) instruction_create_account_for_mint(payer solana.PublicKey, mint solana.PublicKey) (*system.Instruction, error) {
	return c.instruction_create_generic_account(82, payer, tkn.ProgramID, mint)
}

// spl_token::instruction::initialize_mint
func (c *Client) instruction_initialize_mint(decimals uint8, mint solana.PublicKey, mintAuthority solana.PublicKey, freezeAuthority solana.PublicKey) (*tkn.Instruction, error) {
	return tkn.NewInitializeMintInstructionBuilder().SetDecimals(decimals).SetMintAccount(mint).SetMintAuthority(mintAuthority).SetFreezeAuthority(freezeAuthority).ValidateAndBuild()
}

func (c *Client) instruction_create_generic_account(space uint64, payer solana.PublicKey, owner solana.PublicKey, destination solana.PublicKey) (*system.Instruction, error) {
	minLamports, err := c.rpc.GetMinimumBalanceForRentExemption(c.ctx, space, c.commitment)
	if err != nil {
		return nil, err
	}
	return system.NewCreateAccountInstructionBuilder().SetSpace(space).SetLamports(minLamports).SetOwner(owner).SetFundingAccount(payer).SetNewAccount(destination).ValidateAndBuild()
}

// system program; create an account with space to initialize an account that receives tokens
func (c *Client) instruction_create_account_for_token(payer solana.PublicKey, destination solana.PublicKey) (*system.Instruction, error) {
	return c.instruction_create_generic_account(165, payer, tkn.ProgramID, destination)
}

func (c *Client) instruction_initialize_account_for_token(mint solana.PublicKey, destination solana.PublicKey, owner solana.PublicKey) (*tkn.Instruction, error) {
	return tkn.NewInitializeAccountInstructionBuilder().SetSysVarRentPubkeyAccount(solana.SysVarRentPubkey).SetAccount(destination).SetMintAccount(mint).SetOwnerAccount(owner).ValidateAndBuild()
}
