package rpc

// See more: https://docs.solana.com/cluster/rpc-endpoints

const (
	protocolHTTPS = "https://"
	protocolWSS   = "wss://"
)

const (
	hostDevNet           = "api.devnet.solana.com"
	hostTestNet          = "api.testnet.solana.com"
	hostMainNetBeta      = "api.mainnet-beta.solana.com"
	hostMainNetBetaSerum = "solana-api.projectserum.com"
)

const (
	DevNet_RPC           = protocolHTTPS + hostDevNet
	TestNet_RPC          = protocolHTTPS + hostTestNet
	MainNetBeta_RPC      = protocolHTTPS + hostMainNetBeta
	MainNetBetaSerum_RPC = protocolHTTPS + hostMainNetBetaSerum
	LocalNet_RPC         = "http://127.0.0.1:8899"
)

const (
	DevNet_WS           = protocolWSS + hostDevNet
	TestNet_WS          = protocolWSS + hostTestNet
	MainNetBeta_WS      = protocolWSS + hostMainNetBeta
	MainNetBetaSerum_WS = protocolWSS + hostMainNetBetaSerum
)
