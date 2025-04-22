package solana

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_compareString(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				a: "111g",
				b: "111G",
			},
			want: -1,
		},
		{
			name: "shorter length takes priority when meet loose case",
			args: args{
				a: "111g1",
				b: "111G",
			},
			want: 1,
		},
		{
			args: args{
				a: "B11g1",
				b: "A11G",
			},
			want: 1,
		},
		{
			args: args{
				a: "A11g11",
				b: "A11G1",
			},
			want: 1,
		},
		{
			args: args{
				a: "011g11",
				b: "A11G1",
			},
			want: -1,
		},
		{
			args: args{
				a: "a11g11",
				b: "A11G1",
			},
			want: 1,
		},
		{
			args: args{
				a: "9FZEjbygsCBB9fPDaEeq7Jq7DcoF7V1JEcowpi6ZbBbf",
				b: "So11111111111111111111111111111111111111112",
			},
			want: -1,
		},
		{
			args: args{
				a: "ComputeBudget111111111111111111111111111111",
				b: "CPMMoo8L3F4NbTegBCKVNunggL7H1ZpdTHKxQB5qKP1C",
			},
			want: -1,
		},
		{
			args: args{
				a: "SO11111111111111111111111111111111111111112",
				b: "So11111111111111111111111111111111111111112",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, compareString(tt.args.a, tt.args.b), "compareString(%v, %v)", tt.args.a, tt.args.b)
		})
	}
}
