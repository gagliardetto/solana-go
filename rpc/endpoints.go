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
	EndpointRPC_DevNet           = protocolHTTPS + hostDevNet
	EndpointRPC_TestNet          = protocolHTTPS + hostTestNet
	EndpointRPC_MainNetBeta      = protocolHTTPS + hostMainNetBeta
	EndpointRPC_MainNetBetaSerum = protocolHTTPS + hostMainNetBetaSerum
)

const (
	EndpointWS_DevNet           = protocolWSS + hostDevNet
	EndpointWS_TestNet          = protocolWSS + hostTestNet
	EndpointWS_MainNetBeta      = protocolWSS + hostMainNetBeta
	EndpointWS_MainNetBetaSerum = protocolWSS + hostMainNetBetaSerum
)
