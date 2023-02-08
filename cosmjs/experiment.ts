import { queryTx, sendTokens, genClient } from "./client";

const TM_RPC_TESTNET = "rpc.sentry-01.theta-testnet.polypore.xyz:26657"
const TM_RPC_LOCAL = "127.0.0.1:26657"

// queryTx(TM_RPC_TESTNET, "0BB99E5D69468BE80B821834B89DA53F05BC6E67C05714761D3D3A2A85F18476").then()
sendTokens(TM_RPC_TESTNET).then(console.log)

// genClient("./keys/local.alice.key", "cosmos").then(suite => {
//    return suite.signlingClient.signAndBroadcast(
//        suite.signerAddr,
//        [
//            {
//                typeUrl: "/cosmos.bank.v1beta1.MsgSend",
//                value: {
//                    fromAddress: suite.signerAddr,
//                    toAddress: "cosmos1wpwx3e8gw80y82gsluq7fhcant0e4eddmjcj8y",
//                    amount: [
//                        { denom: "stake", amount: "10000" },
//                    ]
//                }
//            }
//        ],
//        "auto"
//    )
// }).then(console.log)