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
# 이미 scaffold한 거 overwrite하려면 각 명령어에 --clear-cache

ignite scaffold chain [name] [--address-prefix <string>]  # default prefix is "cosmos"
                                                          # name is go module name                                                        
ignite scaffold module <module_name> [--ibc] [--params p1,p2...] [--dep account,bank...]  --require-registration
ignite scaffold message <msg_name> [field1] [field2] ... [-r field1,field2,...] --signer <signer_field_name> --module <target_module> 
ignite scaffold query <query_name> [field1] [field2] ... [-r field1,field2,...] --module <target_module>

# Scaffold an IBC packet in a specific IBC-enabled Cosmos SDK module
ignite scaffold packet <packet_name> [filed1] [field2] ... [--ack field1,field2] --module <target_module> 

# 개발용. config.yml 조절을 통해 초기 상태를 제어할 수 있음
# <name>d로 go install 
ignite chain init  # production에서는 init, add-genesis-account, gentx, collect-gentx를 수동으로 진행하는 것을 추천
ignite chain serve # Build proto, install dependencies, compile codes, 
                   # Initialize the node with a single validator(first account), Add accounts based on config.yml

# Production
ignite chain build

# GUI
cd vue
npm install
npm run dev
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