# Solana SDK library for Go

[![GoDoc](https://pkg.go.dev/badge/github.com/gagliardetto/solana-go?status.svg)](https://pkg.go.dev/github.com/gagliardetto/solana-go@v0.3.1?tab=doc)
[![GitHub tag (latest SemVer pre-release)](https://img.shields.io/github/v/tag/gagliardetto/solana-go?include_prereleases&label=release-tag&style=flat-square)](https://github.com/gagliardetto/solana-go/releases)
[![Build Status](https://github.com/gagliardetto/solana-go/workflows/Tests/badge.svg?branch=main)](https://github.com/gagliardetto/solana-go/actions?query=branch%3Amain)
[![TODOs](https://badgen.net/https/api.tickgit.com/badgen/github.com/gagliardetto/solana-go)](https://www.tickgit.com/browse?repo=github.com/gagliardetto/solana-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/gagliardetto/solana-go)](https://goreportcard.com/report/github.com/gagliardetto/solana-go)

Go library to interface with Solana nodes's JSON-RPC interface, Solana's SPL tokens and the
[Serum DEX](https://dex.projectserum.com) instructions.  More contracts to come.

![solana-gopher-2_small](https://user-images.githubusercontent.com/15271561/126141780-cbd92d2c-e160-4385-9606-9094729b54d4.png)

## Installation

> :warning: `solana-go` works using SemVer but in 0 version, which means that the 'minor' will be changed when some broken changes are introduced into the application, and the 'patch' will be changed when a new feature with new changes is added or for bug fixing. As soon as v1.0.0 be released, `solana-go` will start to use SemVer as usual.

1. Install from https://github.com/gagliardetto/solana-go/releases

**or**

2. Build from source with:

```bash
$ go get -u -v github.com/gagliardetto/solana-go/cmd/slnc
```

# Command-line

```bash
$ slnc get balance EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v
1461600 lamports

$ slnc get account EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v
{
  "lamports": 1461600,
  "data": [
    "AQAAABzjWe1aAS4E+hQrnHUaHF6Hz9CgFhuchf/TG3jN/Nj2gCa3xLwWAAAGAQEAAAAqnl7btTwEZ5CY/3sSZRcUQ0/AjFYqmjuGEQXmctQicw==",
    "base64"
  ],
  "owner": "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
  "executable": false,
  "rentEpoch": 108
}

$ slnc spl get mint SRMuApVNdxXokk5GT7XD5cUUgXMBCoAz2LHeuAoKWRt
Supply               9999647664800000
Decimals             6
No mint authority
No freeze authority

$ slnc serum list markets
...
ALEPH/USDT   EmCzMQfXMgNHcnRoFwAdPe1i2SuiSzMj1mx6wu3KN2uA
ALEPH/WUSDC  B37pZmwrwXHjpgvd9hHDAx1yeDsNevTnbbrN9W12BoGK
BTC/USDT     8AcVjMG2LTbpkjNoyq8RwysokqZunkjy3d5JDzxC6BJa
BTC/WUSDC    CAgAeMD7quTdnr6RPa7JySQpjf3irAmefYNdTb6anemq
ETH/USDT     HfCZdJ1wfsWKfYP2qyWdXTT5PWAGWFctzFjLH48U1Hsd
ETH/WUSDC    ASKiV944nKg1W9vsf7hf3fTsjawK6DwLwrnB2LH9n61c
SOL/USDT     8mDuvJJSgoodovMRYArtVVYBbixWYdGzR47GPrRT65YJ
SOL/WUSDC    Cdp72gDcYMCLLk3aDkPxjeiirKoFqK38ECm8Ywvk94Wi
...

$ slnc serum get market 7JCG9TsCx3AErSV3pvhxiW4AbkKRcJ6ZAveRmJwrgQ16
Price    Quantity  Depth
Asks
...
527.06   444.09    ####################
393.314  443.52    ###############
463.158  443.17    ###########
200      442.63    ######
234.503  442.54    ####
50       441.86    ##
61.563   441.47    #
84.377   440.98
-------  --------
10       439.96
193.303  439.24    ##
50       438.94    ##
0.5      438.87    ##
247.891  437.65    #####
458.296  436.99    #########
452.693  435.68    ##############
372.722  435.12    ##################
0.043    431.94    ##################
...
```
# Library usage

Loading an SPL mint

```golang

import "github.com/gagliardetto/solana-go/rpc"
import "github.com/gagliardetto/solana-go/token"

	addr := solana.MustPublicKeyFromBase58("EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v")
	cli := rpc.NewClient("http://api.mainnet-beta.solana.com")

	var m token.Mint
	err := cli.GetAccountDataIn(context.Background(), addr, &m)
	// handle `err`

	json.NewEncoder(os.Stdout).Encode(m)
	// {"OwnerOption":1,
	//  "Owner":"2wmVCSfPxGPjrnMMn7rchp4uaeoTqN39mXFC2zhPdri9",
	//  "Decimals":128,
	//  "IsInitialized":true}

```


Getting any account's data:

```golang

import "github.com/gagliardetto/solana-go/rpc"
import "github.com/gagliardetto/solana-go/token"

	addr := solana.MustPublicKeyFromBase58("EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v")
	cli := rpc.NewClient("http://api.mainnet-beta.solana.com")

	acct, err := cli.GetAccountInfo(context.Background(), addr)
	// handle `err`

	json.NewEncoder(os.Stdout).Encode(m)
// {
//   "context": {
//     "Slot": 47836700
//   },
//   "value": {
//     "lamports": 1461600,
//     "data": {
//       "data": "AQAAABzjWe1aAS4E+hQrnHUaHF6Hz9CgFhuchf/TG3jN/Nj2gCa3xLwWAAAGAQEAAAAqnl7btTwEZ5CY/3sSZRcUQ0/AjFYqmjuGEQXmctQicw==",
//       "encoding": "base64"
//     },
//     "owner": "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
//     "executable": false,
//     "rentEpoch": 109
//   }
// }

```

# Contributing

Any contributions are welcome, use your standard GitHub-fu to pitch in and improve.

# License

[Apache 2.0](LICENSE)

# Credits

- Gopher logo was originally created by Takuya Ueda (https://twitter.com/tenntenn). Licensed under the Creative Commons 3.0 Attributions license.
