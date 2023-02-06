import {IndexedTx, StargateClient, SigningStargateClient, DeliverTxResponse, GasPrice} from "@cosmjs/stargate"
import {OfflineDirectSigner} from "@cosmjs/proto-signing"
import {MsgSend} from "cosmjs-types/cosmos/bank/v1beta1/tx"
import {Tx} from "cosmjs-types/cosmos/tx/v1beta1/tx"
import {WalletFromMnemonic} from "./wallet"

const TM_RPC_TESTNET = "rpc.sentry-01.theta-testnet.polypore.xyz:26657"
const TM_RPC_LOCAL = "127.0.0.1:26657"

const queryTx = async (txid: string): Promise<void> => {
    const client = await StargateClient.connect(TM_RPC_LOCAL)
    console.log("With client, chain id:", await client.getChainId(), ", height:", await client.getHeight())

    const faucetTx: IndexedTx = (await client.getTx(
        txid
    ))!
    console.log("Faucet Tx:", faucetTx)
    const decodedTx: Tx = Tx.decode(faucetTx.tx)
    console.log("DecodedTx:", decodedTx)
    console.log("Decoded messages:", decodedTx.body!.messages)
    const sendMessage: MsgSend = MsgSend.decode(decodedTx.body!.messages[0].value)
    console.log("Sent message:", sendMessage)
    console.log("Gas fee:", decodedTx.authInfo!.fee!.amount)
    console.log("Gas limit:", decodedTx.authInfo!.fee!.gasLimit.toString(10))

    const faucet: string = sendMessage.fromAddress
    console.log("Faucet balances:", await client.getAllBalances(faucet))

    // Get the faucet address another way
    {
        const rawLog = JSON.parse(faucetTx.rawLog)
        console.log("Raw log:", JSON.stringify(rawLog, null, 4))
        const faucet: string = rawLog[0].events
            .find((eventEl: any) => eventEl.type === "coin_spent")
            .attributes.find((attribute: any) => attribute.key === "spender").value
        console.log("Faucet address from raw log:", faucet)
    }
}

type signerSuite = { signlingClient: SigningStargateClient, signer: OfflineDirectSigner, signerAddr: string }

const genClient = async (mnemonicPath: string): Promise<signerSuite> => {
    const signer = await WalletFromMnemonic(mnemonicPath)
    const signerAddr = (await signer.getAccounts())[0].address
    const client = await SigningStargateClient.connectWithSigner(TM_RPC_LOCAL, signer, {gasPrice: GasPrice.fromString("0.0025stake")})
    return {
        signer,
        signerAddr,
        signlingClient: client
    }
}

const sendTokens = async (): Promise<DeliverTxResponse> => {
    const signer = await WalletFromMnemonic("./keys/local.alice.key")
    const signerAddr = (await signer.getAccounts())[0].address
    const client = await SigningStargateClient.connectWithSigner(TM_RPC_LOCAL, signer)
    console.log("With client, chain id:", await client.getChainId(), ", height:", await client.getHeight())
    console.log("My balances", await client.getAllBalances(signerAddr))

    const recipient = await WalletFromMnemonic("./keys/local.bob.key")
    const recipAddr = (await recipient.getAccounts())[0].address
    console.log("Recipient balances", await client.getAllBalances(recipAddr))


    const sentAmt = [{denom: "stake", amount: "100000"}]
    const fee = {
        amount: [{denom: "stake", amount: "500"}],
        gas: "200000",
    }
    const res = await client.sendTokens(signerAddr, recipAddr, sentAmt, fee, "memo")
    console.log("My balances after send", await client.getAllBalances(signerAddr))
    console.log("Recipient balances after send", await client.getAllBalances(recipAddr))

    return res
}

export {
    queryTx,
    sendTokens,
    genClient,
}