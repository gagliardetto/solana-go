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
	EndpointRPCDevNet           = protocolHTTPS + hostDevNet
	EndpointRPCTestNet          = protocolHTTPS + hostTestNet
	EndpointRPCMainNetBeta      = protocolHTTPS + hostMainNetBeta
	EndpointRPCMainNetBetaSerum = protocolHTTPS + hostMainNetBetaSerum
)

const (
	EndpointWSDevNet           = protocolWSS + hostDevNet
	EndpointWSTestNet          = protocolWSS + hostTestNet
	EndpointWSMainNetBeta      = protocolWSS + hostMainNetBeta
	EndpointWSMainNetBetaSerum = protocolWSS + hostMainNetBetaSerum
)
