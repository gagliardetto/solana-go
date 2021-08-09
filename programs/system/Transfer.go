package system

import (
	bin "github.com/dfuse-io/binary"
	solana "github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/text"
	. "github.com/gagliardetto/solana-go/text"
	"github.com/gagliardetto/treeout"
)

func NewTransferInstruction(
	lamports uint64,
	from solana.PublicKey,
	to solana.PublicKey,
) *Instruction {
	return &Instruction{
		BaseVariant: bin.BaseVariant{

			TypeID: Instruction_Transfer,

			Impl: &Transfer{
				Lamports: bin.Uint64(lamports),

				AccountMetaSlice: []*solana.AccountMeta{
					solana.Meta(from).WRITE().SIGNER(),
					solana.Meta(to).WRITE(),
				},
			},
		},
	}
}

type Transfer struct {
	Lamports bin.Uint64

	// [0] = [WRITE, SIGNER] Funding account.
	// [1] = [WRITE] Recipient account.
	solana.AccountMetaSlice `bin:"-" text:"-"`
}

func (inst *Transfer) EncodeToTree(parent treeout.Branches) {
	parent.Child(Sf(CC(IndigoBG(Bold("Program")), ": System (%s)"), text.ColorizeBG(PROGRAM_ID.String()))).
		ParentFunc(func(programBranch treeout.Branches) {
			programBranch.Child(CC(Purple(Bold("Instruction")), Bold(": Transfer"))).
				ParentFunc(func(instructionBranch treeout.Branches) {
					instructionBranch.Child("Params").ParentFunc(func(paramsBranch treeout.Branches) {
						paramsBranch.Child(Sf(CC(Shakespeare("Lamports"), ": %v"), Lime(S(inst.Lamports))))
					})
					instructionBranch.Child("Accounts:").ParentFunc(func(accountsBranch treeout.Branches) {
						accountsBranch.Child(Sf(CC(Shakespeare("Funding account"), ": %s"), text.ColorizeBG(inst.AccountMetaSlice[0].PublicKey.String())))
						accountsBranch.Child(Sf(CC(Shakespeare("Recipient account"), ": %s"), text.ColorizeBG(inst.AccountMetaSlice[1].PublicKey.String())))
					})
				})
		})
}
