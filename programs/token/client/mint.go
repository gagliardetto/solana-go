package client

import (
	"errors"

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
	err = tokenInfo.Decode(data)
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
func (c *Client) CreateToken(args *TokenArgs, payerPrivateKey solana.PrivateKey, destination solana.PublicKey, supply uint64) (*tkn.Mint, error) {
	payer := payerPrivateKey.PublicKey()

	mintPrivateKey, err := solana.NewRandomPrivateKey()
	if err != nil {
		return nil, err
	}
	mint := mintPrivateKey.PublicKey()

	blockHash, err := c.latest_blockhash()
	if err != nil {
		return nil, err
	}

	inst_1, err := c.create_token_instruction_create_account(payer, mint)
	if err != nil {
		return nil, err
	}
	inst_2, err := c.create_token_instruction_initialize_mint(args.Decimals, mint, args.MintAuthority, args.FreezeAuthority)
	if err != nil {
		return nil, err
	}
	inst_3, err := c.create_token_create_account_destination(payer, destination)
	if err != nil {
		return nil, err
	}
	inst_4 := c.create_token_initialize_destination(mint, destination, payer)

	var list []solana.Instruction
	if 0 < supply {
		list = []solana.Instruction{
			inst_1, inst_2, inst_3, inst_4, c.create_token_mint(mint, args.MintAuthority, destination, supply),
		}
	} else {
		list = []solana.Instruction{
			inst_1, inst_2, inst_3, inst_4,
		}
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
		return nil
	})
	if err != nil {
		return nil, err
	}

	sig, err := c.rpc.SendTransactionWithOpts(c.ctx, tx, true, c.commitment)
	if err != nil {
		return nil, err
	}

	err = c.block_on_tx_processing(sig)
	if err != nil {
		return nil, err
	}

	return c.GetMint(mint)
}

func (c *Client) create_token_mint(mint solana.PublicKey, mintAuthority solana.PublicKey, destination solana.PublicKey, supply uint64) *tkn.Instruction {
	return tkn.NewMintToInstructionBuilder().SetAmount(supply).SetDestinationAccount(destination).SetMintAccount(mint).SetAuthorityAccount(mintAuthority).Build()
}

func (c *Client) create_token_instruction_create_account(payer solana.PublicKey, mint solana.PublicKey) (*system.Instruction, error) {
	space := uint64(82)
	minLamports, err := c.rpc.GetMinimumBalanceForRentExemption(c.ctx, space, c.commitment)
	if err != nil {
		return nil, err
	}
	return system.NewCreateAccountInstructionBuilder().SetSpace(space).SetLamports(minLamports).SetOwner(tkn.ProgramID).SetFundingAccount(payer).SetNewAccount(mint).Build(), nil
}

func (c *Client) create_token_instruction_initialize_mint(decimals uint8, mint solana.PublicKey, mintAuthority solana.PublicKey, freezeAuthority solana.PublicKey) (*tkn.Instruction, error) {
	return tkn.NewInitializeMintInstructionBuilder().SetDecimals(decimals).SetMintAccount(mint).SetMintAuthority(mintAuthority).SetFreezeAuthority(freezeAuthority).Build(), nil
}

func (c *Client) create_token_create_account_destination(payer solana.PublicKey, destination solana.PublicKey) (*system.Instruction, error) {
	destinationSpace := uint64(165)
	destinationLamports, err := c.rpc.GetMinimumBalanceForRentExemption(c.ctx, destinationSpace, c.commitment)
	if err != nil {
		return nil, err
	}
	return system.NewCreateAccountInstructionBuilder().SetSpace(destinationSpace).SetLamports(destinationLamports).SetOwner(tkn.ProgramID).SetNewAccount(destination).SetFundingAccount(payer).Build(), nil
}

func (c *Client) create_token_initialize_destination(mint solana.PublicKey, destination solana.PublicKey, owner solana.PublicKey) *tkn.Instruction {
	return tkn.NewInitializeAccountInstructionBuilder().SetSysVarRentPubkeyAccount(solana.SysVarRentPubkey).SetAccount(destination).SetMintAccount(mint).SetOwnerAccount(owner).Build()
}
