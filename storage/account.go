package storage

import (
	bin "github.com/dfuse-io/binary"
	"github.com/dfuse-io/solana-go"
)

type StoredMeta struct {
	Version           bin.Uint64
	DataLength        bin.Uint64 `bin:"sizeof=Data"`
	PubKey            solana.PublicKey
	Lamport           bin.Uint64
	RentEpoch         bin.Uint64
	Owner             solana.PublicKey
	ExecutablePadding [7]byte
	Executable        bool

	Hash solana.Hash
	Data []byte
}

//c8 49 cbfee98c362c717fe8ad198d43e26db4965c28ab60f1aeb9597bcd3efcf4
//73 65 72 75 6d 03 00 00 00 00 00 00 00 45 5d 51 f8 c1 d9 eb d8 0f 70 f0 53 4d 89 cc 95 3b 13 73 4f 63 42 ba 34 4e 22 45 d7 39 c2 6f 30 00 00 00 00 00 00 00 00 12 4a 27 e3 1c 9c 77 74 35 79 cd f7 dd 8d 94 15 3d b2 4f 27 24 25 2c 19 b3 fb 53 db 78 23 07 c9 9a a1 82 88 ee 84 3d 83 f6 3d 09 1f fd 54 53 bf 0e c9 8d 7b c8 8c f1 49 cb 4e 90 47 50 aa d5 67 72 ea f8 5c 66 f3 73 b9 f3 f3 8d 86 7e 83 c6 43 d5 13 f1 5d 00 d0 db e6 7f f4 8b 63 ae 77 e8 ae 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 d6 98 d0 09 fc e4 25 09 75 6e e7 7c 4b 1a e1 d2 df 93 74 62 8e 98 7c f2 4b d1 aa 4b e6 01 11 13 7a 49 c9 04 10 00 00 00 00 00 00 00 00 00 00 00 64 00 00 00 00 00 00 00 2c 04 f0 79 78 bf ea d5 2d bd 42 17 12 ae 65 7f 42 3f 56 f5 40 e0 98 58 d8 ae 37 43 b2 a2 39 d3 78 37 1a dd 23 17 d4 8a fa e1 84 16 ba 7e a8 17 7c d1 64 ec 97 56 f7 56 c0 3b 06 71 d6 86 7b 8e bf 07 1e 2e 81 74 bb 06 5a 6f 41 d4 0e 2f 24 42 38 c9 bc 26 12 83 8e 27 61 bc 9b 43 72 2d 4d 65 cf 98 7a 11 2f f5 a8 3c a8 aa 97 b5 09 7c 4e a6 9f 1f ba fd c0 1c f1 3e 37 98 fc e2 c9 20 c6 8b a0 86 01 00 00 00 00 00 64 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 70 61 64 64 69 6e 67
//
//{
//"lamports": 3591360,
//"data": {
//"data": "c2VydW0DAAAAAAAAAEVdUfjB2evYD3DwU02JzJU7E3NPY0K6NE4iRdc5wm8wAAAAAAAAAAASSifjHJx3dDV5zffdjZQVPbJPJyQlLBmz+1PbeCMHyZqhgojuhD2D9j0JH/1UU78OyY17yIzxSctOkEdQqtVncur4XGbzc7nz842GfoPGQ9UT8V0A0Nvmf/SLY6536K4AAAAAAAAAAAAAAAAAAAAA1pjQCfzkJQl1bud8Sxrh0t+TdGKOmHzyS9GqS+YBERN6SckEEAAAAAAAAAAAAAAAZAAAAAAAAAAsBPB5eL/q1S29QhcSrmV/Qj9W9UDgmFjYrjdDsqI503g3Gt0jF9SK+uGEFrp+qBd80WTsl1b3VsA7BnHWhnuOvwceLoF0uwZab0HUDi8kQjjJvCYSg44nYbybQ3ItTWXPmHoRL/WoPKiql7UJfE6mnx+6/cAc8T43mPziySDGi6CGAQAAAAAAZAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAcGFkZGluZw==",
//"encoding": "base64"
//},
//"owner": "EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o",
//"executable": false,
//"rentEpoch": 138
//}
