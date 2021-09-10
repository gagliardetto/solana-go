package solana

// NewInstruction creates a generic instruction with the provided
// programID, accounts, and data bytes.
func NewInstruction(
	programID PublicKey,
	accounts AccountMetaSlice,
	data []byte,
) *GenericInstruction {
	return &GenericInstruction{
		AccountValues: accounts,
		ProgID:        programID,
		DataBytes:     data,
	}
}

var _ Instruction = &GenericInstruction{}

type GenericInstruction struct {
	AccountValues AccountMetaSlice
	ProgID        PublicKey
	DataBytes     []byte
}

func (in *GenericInstruction) ProgramID() PublicKey {
	return in.ProgID
}

func (in *GenericInstruction) Accounts() []*AccountMeta {
	return in.AccountValues
}

func (in *GenericInstruction) Data() ([]byte, error) {
	return in.DataBytes, nil
}
