# cosmos-dummy
Built dummy chain using Tendermint Core &amp; Cosmos SDK for practicing

# Setup
## 1) Ignite CLI
### Install
```shell
curl https://get.ignite.com/cli! | bash
ignite version

# Build from source (optional)
rm $(which ignite)
git clone https://github.com/ignite/cli --depth=1
cd cli && make install
```
### Build
```shell
# scaffold한 거 overwrite하려면 --clear-cache
ignite scaffold chain [chain_name] [--address-prefix <string>] # default prefix is "cosmos"
ignite scaffold module <module_name> [--ibc] [--params p1,p2...] [--dep account,bank...]  --require-registration
ignite scaffold message <msg_name> [field1] [field2] ... [-r field1,field2,...] --signer <signer_field_name> --module <target_module> 
ignite scaffold query <query_name> [field1] [field2] ... [-r field1,field2,...] --module <target_module>
# Scaffold an IBC packet in a specific IBC-enabled Cosmos SDK module
ignite scaffold packet <packet_name> [filed1] [field2] ... [--ack field1,field2] --module <target_module> 
ignite chain serve # Start a blockchain

```

# Resources
* [ABCI Spec](https://github.com/tendermint/spec/blob/c939e15/spec/abci/abci.md)
* [ABCI Go Interface](https://github.com/tendermint/tendermint/tree/main/abci)