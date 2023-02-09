import { queryTx, sendTokens, SuperCosmosClient } from "./client";
import { MsgMultiSend, MsgSend } from "cosmjs-types/cosmos/bank/v1beta1/tx";
import {MsgCreateGameEncodeObject} from "./extension";

const TM_RPC_TESTNET = "rpc.sentry-01.theta-testnet.polypore.xyz:26657";
const TM_RPC_LOCAL = "127.0.0.1:26657";

// queryTx(TM_RPC_TESTNET, "0BB99E5D69468BE80B821834B89DA53F05BC6E67C05714761D3D3A2A85F18476").then()
// sendTokens(TM_RPC_TESTNET).then(console.log)

(async function (){
    const client = new SuperCosmosClient();
    await client.initialize(TM_RPC_LOCAL, "./keys/local.alice.key", "cosmos", "0stake");
    // const res = await client.sendToken("cosmos1wpwx3e8gw80y82gsluq7fhcant0e4eddmjcj8y",
    //     [
    //     { denom: "stake", amount: "10000" },
    // ])
    //
    // console.log(res)

    client.Q.bank.allBalances(client.address).then(console.log)
    client.Q.checker.systemInfo().then(console.log)
    client.Q.checker.params().then(console.log)

    const msg: MsgCreateGameEncodeObject = {
        typeUrl: "/checkers.checkers.MsgCreateGame",
        value: {
            creator: client.address,
            black: client.address,
            red: "cosmos1wpwx3e8gw80y82gsluq7fhcant0e4eddmjcj8y"
        },
    }
    console.log(await client.sendTx([msg]))
    console.log(await client.Q.checker.storedGame("1"))
    console.log(await client.Q.checker.storedGameAll())

})();
