# Solana library for Go

Go library to interface with Solana nodes's JSON-RPC interface, Solana's SPL tokens and the
[https://dex.projectserum.com](Serum DEX) instructions.  More contracts to come.


# Command-line

```

$ slnc serum markets
...
SUSHI/USDC -> 7LVJtqSrF6RudMaz5rKGTmR3F3V5TKoDcN6bnk68biYZ
SXP/USDC -> 13vjJ8pxDMmzen26bQ5UrouX8dkXYPW1p3VLVDjxXrKR
MSRM/USDC -> AwvPwwSprfDZ86beBJDNH5vocFvuw4ZbVQ6upJDbSCXZ
FTT/USDC -> FfDb3QZUdMW2R2aqJQgzeieys4ETb3rPrFFfPSemzq7R
YFI/USDC -> 4QL5AQvXdMSCVZmnKXiuMMU83Kq3LCwVfU8CyznqZELG
LINK/USDC -> 7JCG9TsCx3AErSV3pvhxiW4AbkKRcJ6ZAveRmJwrgQ16
HGET/USDC -> 3otQFkeQ7GNUKT3i2p3aGTQKS2SAw6NLYPE5qxh3PoqZ
CREAM/USDC -> 2M8EBxFbLANnCoHydypL1jupnRHG782RofnvkatuKyLL
...

$ slnc serum market 7JCG9TsCx3AErSV3pvhxiW4AbkKRcJ6ZAveRmJwrgQ16


```



# Library

```go
solana.NewClient("https://mainnet.solana.dfuse.io")

```



# Examples




# Contributing

Any contributions are welcome, use your standard GitHub-fu to pitch in and improve.


License
-------

Apache-2
