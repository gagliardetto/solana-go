#!/bin/bash

set -euxo pipefail

source $HOME/.cargo/env

decho(){
    1>&2 echo $@
}


new(){
    decho "setting up solana test environment on localhost"
    solana-keygen new -o $KEYPAIR_FILE_PATH
    solana config set --url localhost
    solana config set --keypair $KEYPAIR_FILE_PATH
    ( ( </dev/null solana-test-validator 1>/dev/null 2>/dev/null & )  & )
    sleep 3
    decho "solana test validator has started"

    solana airdrop 10
}

build(){
    yarn
    anchor build
}

deploy(){
    decho "building serum program"
    ( cd deps/serum-dex/dex/ && cargo build-bpf  ) 
    export DEX_PID=$(solana program deploy ./deps/serum-dex/dex/target/deploy/serum_dex.so | grep Program | perl -e '$x=<STDIN>;if($x=~m#Program Id\: (\w+)#){print "$1";}else{die "no match";}' )
    decho "dex id is ${DEX_PID}"
}

test(){
    if [[ -z "${DEX_PID}" ]]; then
        decho "DEX_PID not defined"
        exit 1
    fi
    anchor test
}




CMD=$1

case $1 in
    all)
        new
        deploy
        build
        test
    ;;
    new)
        new
    ;;
    build)
        build
    ;;
    deploy)
        deploy
    ;;
    test)
        test
    ;;
    *)
        decho "bad choice"
        exit 1
    ;;
esac