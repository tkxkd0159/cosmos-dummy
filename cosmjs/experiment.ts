import { queryTx, sendTokens, genClient } from "./client";

// queryTx("540484BDD342702F196F84C2FD42D63FA77F74B26A8D7383FAA5AB46E4114A9B").then()
// queryTx("")
// sendTokens().then(console.log)

genClient("./keys/local.alice.key").then(suite => {
   return suite.signlingClient.signAndBroadcast(
       suite.signerAddr,
       [
           {
               typeUrl: "/cosmos.bank.v1beta1.MsgSend",
               value: {
                   fromAddress: suite.signerAddr,
                   toAddress: "cosmos1wpwx3e8gw80y82gsluq7fhcant0e4eddmjcj8y",
                   amount: [
                       { denom: "stake", amount: "10000" },
                   ]
               }
           }
       ],
       "auto"
   )
}).then(console.log)