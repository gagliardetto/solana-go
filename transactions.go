package solana

type Transaction struct {
	Signatures []Signature `json:"signatures"`
	Message    Message     `json:"message"`
}

type Message struct {
	Header          MessageHeader `json:"header"`
	AccountKeys     []PublicKey   `json:"accountKeys"`
	RecentBlockhash PublicKey/* TODO: change to Hash */ `json:"recentBlockhash"`
	Instructions    []CompiledInstruction `json:"instructions"`
}

type MessageHeader struct {
	NumRequiredSignatures       uint8 `json:"numRequiredSignatures"`
	NumReadonlySignedAccounts   uint8 `json:"numReadonlySignedAccounts"`
	NumReadonlyunsignedAccounts uint8 `json:"numReadonlyUnsignedAccounts"`
}

type CompiledInstruction struct {
	ProgramIDIndex uint8   `json:"programIdIndex"`
	AccountsCount  uint16  `json:"-",struc:"sizeof=Accounts,little"`
	Accounts       []uint8 `json:"accounts"`
	DataLength     uint16  `json:"-",struc:"sizeof=Data,little"`
	Data           Base58  `json:"data"`
}
