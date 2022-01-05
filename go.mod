module github.com/gagliardetto/solana-go

go 1.16

retract (
	v1.0.1-beta.1 // Published accidentally.
	v1.0.0 // Published accidentally.
)

require (
	contrib.go.opencensus.io/exporter/stackdriver v0.13.4 // indirect
	filippo.io/edwards25519 v1.0.0-rc.1
	github.com/AlekSi/pointer v1.1.0
	github.com/GeertJohan/go.rice v1.0.0
	github.com/aybabtme/rgbterm v0.0.0-20170906152045-cc83f3b3ce59
	github.com/buger/jsonparser v1.1.1
	github.com/davecgh/go-spew v1.1.1
	github.com/dfuse-io/logging v0.0.0-20210109005628-b97a57253f70
	github.com/fatih/color v1.7.0
	github.com/gagliardetto/binary v0.5.1
	github.com/gagliardetto/gofuzz v1.2.2
	github.com/gagliardetto/treeout v0.1.4
	github.com/google/go-cmp v0.5.1
	github.com/gorilla/rpc v1.2.0
	github.com/gorilla/websocket v1.4.2
	github.com/json-iterator/go v1.1.11
	github.com/klauspost/compress v1.13.6
	github.com/logrusorgru/aurora v2.0.3+incompatible
	github.com/magiconair/properties v1.8.1
	github.com/mostynb/zstdpool-freelist v0.0.0-20201229113212-927304c0c3b1
	github.com/mr-tron/base58 v1.2.0
	github.com/onsi/gomega v1.10.1
	github.com/pkg/errors v0.9.1
	github.com/ryanuber/columnize v0.0.0-20160712163229-9b3edd62028f
	github.com/spf13/cobra v1.1.1
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.6.1
	github.com/teris-io/shortid v0.0.0-20201117134242-e59966efd125 // indirect
	github.com/tidwall/gjson v1.9.3 // indirect
	go.opencensus.io v0.22.5 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/ratelimit v0.2.0
	go.uber.org/zap v1.16.0
	golang.org/x/crypto v0.0.0-20210513164829-c07d793c2f9a
	golang.org/x/net v0.0.0-20210510120150-4163338589ed // indirect
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
	golang.org/x/sys v0.0.0-20210514084401-e8d321eab015 // indirect
	golang.org/x/term v0.0.0-20201210144234-2321bbc49cbf // indirect
	golang.org/x/tools v0.0.0-20200601175630-2caf76543d99 // indirect
	google.golang.org/api v0.29.0
)
