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
	EndpointRPCDevNet           = protocolHTTPS + "api.devnet.solana.com"
	EndpointRPCTestNet          = protocolHTTPS + "api.testnet.solana.com"
	EndpointRPCMainNetBeta      = protocolHTTPS + "api.mainnet-beta.solana.com"
	EndpointRPCMainNetBetaSerum = protocolHTTPS + "solana-api.projectserum.com"
)

const (
	EndpointWSDevNet           = protocolWSS + "api.devnet.solana.com"
	EndpointWSTestNet          = protocolWSS + "api.testnet.solana.com"
	EndpointWSMainNetBeta      = protocolWSS + "api.mainnet-beta.solana.com"
	EndpointWSMainNetBetaSerum = protocolWSS + "solana-api.projectserum.com"
)
