# Solana SDK library for Go

[![GoDoc](https://pkg.go.dev/badge/github.com/gagliardetto/solana-go?status.svg)](https://pkg.go.dev/github.com/gagliardetto/solana-go@v0.3.2?tab=doc)
[![GitHub tag (latest SemVer pre-release)](https://img.shields.io/github/v/tag/gagliardetto/solana-go?include_prereleases&label=release-tag)](https://github.com/gagliardetto/solana-go/releases)
[![Build Status](https://github.com/gagliardetto/solana-go/workflows/tests/badge.svg?branch=main)](https://github.com/gagliardetto/solana-go/actions?query=branch%3Amain)
[![TODOs](https://badgen.net/https/api.tickgit.com/badgen/github.com/gagliardetto/solana-go/main)](https://www.tickgit.com/browse?repo=github.com/gagliardetto/solana-go&branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/gagliardetto/solana-go)](https://goreportcard.com/report/github.com/gagliardetto/solana-go)

Go library to interface with Solana JSON RPC and WebSocket interfaces.

Clients for Solana native programs, Solana Program Library (SPL), and [Serum DEX](https://dex.projectserum.com) are in development.

More contracts to come.

<div align="center">
    <img src="https://user-images.githubusercontent.com/15271561/126141780-cbd92d2c-e160-4385-9606-9094729b54d4.png" margin="auto" height="175"/>
</div>

## Contents

- [Solana SDK library for Go](#solana-sdk-library-for-go)
  - [Contents](#contents)
  - [Features](#features)
  - [Current development status](#current-development-status)
  - [Requirements](#requirements)
  - [Installation](#installation)
  - [RPC usage examples](#rpc-usage-examples)
  - [Contributing](#contributing)
  - [License](#license)
  - [Credits](#credits)
  - [Installation](#installation)

## Features

- [x] Full JSON RPC API
- [x] Full WebSocket JSON streaming API
- [ ] Wallet, account, and keys management
- [ ] Clients for native programs
- [ ] Clients for Solana Program Library
- [ ] Client for Serum
- [ ] More programs

## Current development status

There is currently **no stable release**. The SDK is actively developed and latest is `v0.3.2` which is an `alpha` release.

The RPC and WS client implementation is based on [this RPC spec](https://github.com/solana-labs/solana/blob/dff9c88193da142693cabebfcd3bf68fa8e8b873/docs/src/developing/clients/jsonrpc-api.md).

## Requirements

- Go 1.16 or later

## Installation

> :warning: `solana-go` works using SemVer but in 0 version, which means that the 'minor' will be changed when some broken changes are introduced into the application, and the 'patch' will be changed when a new feature with new changes is added or for bug fixing. As soon as v1.0.0 be released, `solana-go` will start to use SemVer as usual.

```bash
$ cd my-project
$ go get github.com/gagliardetto/solana-go@latest
```

## Examples

### Create account

```go
package main

import (
  "context"
  "fmt"

  "github.com/gagliardetto/solana-go"
  "github.com/gagliardetto/solana-go/rpc"
)

func main() {
  // Create a new account:
  account := solana.NewAccount()
  fmt.Println("account private key:", account.PrivateKey)
  fmt.Println("account public key:", account.PublicKey())

  // Create a new RPC client:
  client := rpc.New(rpc.TestNet_RPC)

  // Airdrop 1 sol to the new account:
  out, err := client.RequestAirdrop(
    context.TODO(),
    account.PublicKey(),
    solana.LAMPORTS_PER_SOL,
    "",
  )
  if err != nil {
    panic(err)
  }
  fmt.Println("airdrop transaction signature:", out)
}
```

### Send Sol from one account to another


## RPC usage examples

TODO

## Contributing

We encourage everyone to contribute, submit issues, PRs, discuss. Every kind of help is welcome.

## License

[Apache 2.0](LICENSE)

## Credits

- Gopher logo was originally created by Takuya Ueda (https://twitter.com/tenntenn). Licensed under the Creative Commons 3.0 Attributions license.
- Pit Vipers https://www.pitviper.com/
