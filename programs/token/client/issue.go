package client

import (
	"github.com/gagliardetto/solana-go"
)

func (c *Client) CreateAccountAndIssue(payerPrivateKey solana.PrivateKey, mintFull *MintWithAddress, owner solana.PublicKey, amount uint64) (*solana.PublicKey, error) {
	keyMap := make(map[string]*solana.PrivateKey)

	payer := payerPrivateKey.PublicKey()
	keyMap[payer.String()] = &payerPrivateKey

	mint := mintFull.Address

	//s := mintFull.State

	list := []solana.Instruction{}

	destinationPrivateKey, err := solana.NewRandomPrivateKey()
	if err != nil {
		return nil, err
	}
	destination := destinationPrivateKey.PublicKey()
	keyMap[destination.String()] = &destinationPrivateKey

	inst_1, err := c.instruction_create_account_for_token(payer, destination)
	if err != nil {
		return nil, err
	}
	list = append(list, inst_1)
	// spl_token::instruction::initialize_account
	inst_2, err := c.instruction_initialize_account_for_token(mint, destination, owner)
	if err != nil {
		return nil, err
	}
	list = append(list, inst_2)

	err = c.send_tx_generic(keyMap, payer, list)
	if err != nil {
		return nil, err
	}

	return &destination, nil

}
