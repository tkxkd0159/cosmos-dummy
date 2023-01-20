# cosmos-dummy
Built dummy chain using Tendermint Core &amp; Cosmos SDK for practicing

# Setup
## Install
```shell
curl https://get.ignite.com/cli! | bash
ignite version

# Build from source (optional)
rm $(which ignite)
git clone https://github.com/ignite/cli --depth=1
cd cli && make install
```
## Build
`field`의 경우 <field_name>:<type>으로 설정 가능. `type`생략 시 string으로 설정됨. (e.g. `ignite scaffold message dummy toX:uint toY:uint`)
)
```shell
# 1317 : REST API, 26657 : Tendermint API
# 3000 : GUI for chain, 4500 : token faucet

# 이미 scaffold한 거 overwrite하려면 각 명령어에 --clear-cache

ignite scaffold chain [name] [--address-prefix <string>]  # default prefix is "cosmos"
                                                          # name is go module name                                                        

ignite scaffold module <module_name> [--ibc] [--params p1,p2...] [--dep account,bank...]  --require-registration

# Module's message
# signer field의 경우 설정하지 않으면 creator라는 이름으로 자동 생성
# -r 플래그는 tx message의 proto response field
# <name>:[type] 형태로 각 field 정의 가능하고 타입 정의안할 경우 string이 기본
# supported types : string, bool, int, uint, coin, array.string, array.int, array.uint, array.coin
ignite scaffold message <msg_name> [field1] [field2] ... [--response field1,field2,...] --signer <signer_field_name> --module <target_module> [--no-simulation]

## ex) implement the logic for storing and interacting with data stored as a list in the blockchain state
## list가 아닌 singleton 값 저장이나 map 형태로 저장을 원할 경우 single/map으로 변경해서 실행
ignite scaffold list pool amount:coin tags:array.string height:int


# Module's query. Message와 만들어지는 원리 같음
ignite scaffold query <query_name> [field1] [field2] ... [-r field1,field2,...] --module <target_module>

# Scaffold an IBC packet in a specific IBC-enabled Cosmos SDK module
ignite scaffold packet <packet_name> [filed1] [field2] ... [--ack field1,field2] --module <target_module> 

# proto 수정 후 업데이트
# 단, proto 빌드만 다시 하는거므로 기존에 생성되었던 연관 메서드들은 수동으로 수정 필요
ignite generate proto-go
```

```shell
# 개발용. config.yml 조절을 통해 초기 상태를 제어할 수 있음
# <name>d로 go install 
ignite chain init  # production에서는 바이너리 직접 빌드 후 init, add-genesis-account, gentx, collect-gentx를 수동으로 진행하는 것을 추천
ignite chain serve # chain init + start node
                   # Build proto, install dependencies, compile codes, 
                   # Initialize the node with a single validator(first account), Add accounts based on config.yml

# Production
ignite chain build [--skip-proto] [--output dist]

# GUI
cd vue
npm install
npm run dev
```

## Check
```shell
<app>d status 2>&1 | jq
```

## Directory Structure
* `app`: a folder for the application.
* `cmd`: a folder for the command-line interface commands.
* `proto`: a folder for the Protobuf objects definitions.
* `vue`: a folder for the auto-generated UI.
* `x`: a folder for all your modules

# Resources
* [ABCI Spec](https://github.com/tendermint/spec/blob/c939e15/spec/abci/abci.md)
* [ABCI Go Interface](https://github.com/tendermint/tendermint/tree/main/abci)