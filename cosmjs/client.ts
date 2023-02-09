import {
    IndexedTx, StargateClient, SigningStargateClient, SigningStargateClientOptions, QueryClient, DeliverTxResponse, GasPrice,
    setupAuthExtension, setupBankExtension, setupStakingExtension, setupDistributionExtension, setupTxExtension, AuthExtension, BankExtension, StakingExtension, DistributionExtension, TxExtension,
} from "@cosmjs/stargate"
import {Tendermint34Client} from "@cosmjs/tendermint-rpc"
import {OfflineDirectSigner, EncodeObject} from "@cosmjs/proto-signing"
import {MsgSend} from "cosmjs-types/cosmos/bank/v1beta1/tx"
import {Tx} from "cosmjs-types/cosmos/tx/v1beta1/tx"
import {Coin} from "cosmjs-types/cosmos/base/v1beta1/coin"
import {WalletFromMnemonic} from "./wallet"
import {CheckerExtension, setupCheckerExtension, genCustomRegistry} from "./extension"

const queryTx = async (rpc: string, txid: string): Promise<void> => {
    const client = await StargateClient.connect(rpc)
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
    console.log("=".repeat(60))
    {
        const rawLog = JSON.parse(faucetTx.rawLog)
        console.log("Raw log:", JSON.stringify(rawLog, null, 4))
        const faucet: string = rawLog[0].events
            .find((eventEl: any) => eventEl.type === "coin_spent")
            .attributes.find((attribute: any) => attribute.key === "spender").value
        console.log("Faucet address from raw log:", faucet)
    }
}

const sendTokens = async (rpc: string): Promise<DeliverTxResponse> => {
    const signer = await WalletFromMnemonic("./keys/testnet.dummy.mnemonic.key", "cosmos")
    const signerAddr = (await signer.getAccounts())[0].address
    const client = await SigningStargateClient.connectWithSigner(rpc, signer)
    console.log("With client, chain id:", await client.getChainId(), ", height:", await client.getHeight())
    console.log("My balances", await client.getAllBalances(signerAddr))

    const recipient = await WalletFromMnemonic("./keys/testnet.recipient.mnemonic.key", "cosmos")
    const recipAddr = (await recipient.getAccounts())[0].address
    console.log("Recipient balances", await client.getAllBalances(recipAddr))


    const sentAmt = [{denom: "uatom", amount: "5000"}]
    const fee = {
        amount: [{denom: "uatom", amount: "500"}],
        gas: "100000",
    }
    const res = await client.sendTokens(signerAddr, recipAddr, sentAmt, fee, "memo")
    console.log("My balances after send", await client.getAllBalances(signerAddr))
    console.log("Recipient balances after send", await client.getAllBalances(recipAddr))

    return res
}

type CustomQuerier = QueryClient & CheckerExtension & AuthExtension & BankExtension & StakingExtension & DistributionExtension & TxExtension

class SuperCosmosClient {
    signer!: OfflineDirectSigner
    address!: string
    rpc!: string
    private txclient!: SigningStargateClient
    private querier!: CustomQuerier

    constructor() {
    }

    // gasPrice: price per gas. e.g. 0.0025stake
    async initialize(rpc: string, mnemonicPath: string, prefix: string, gasPrice?: string) {
        this.rpc = rpc
        this.signer = await WalletFromMnemonic(mnemonicPath, prefix)
        this.address = (await this.signer.getAccounts())[0].address

        let txClientOpts: SigningStargateClientOptions
        if (gasPrice !== undefined) {
            txClientOpts = {registry: genCustomRegistry(), gasPrice: GasPrice.fromString(gasPrice)}
        } else {
            txClientOpts = {registry: genCustomRegistry()}
        }
        const client = await SigningStargateClient.connectWithSigner(rpc, this.signer, txClientOpts)
        this.txclient = client
        this.querier = QueryClient.withExtensions(await Tendermint34Client.connect(rpc),
            setupAuthExtension, setupBankExtension, setupCheckerExtension, setupStakingExtension, setupDistributionExtension, setupTxExtension)
    }

    sendToken(recipient: string, amt: Coin[]): Promise<DeliverTxResponse> {
        const msg1: MsgSend = {
            fromAddress: this.address,
            toAddress: recipient,
            amount: amt
        }
        return this.sendTx([
            {
                typeUrl: "/cosmos.bank.v1beta1.MsgSend",
                value: msg1,
            }
        ])
    }

    sendTx(msgs: EncodeObject[]): Promise<DeliverTxResponse> {
        return this.txclient.signAndBroadcast(
            this.address,
            msgs,
            "auto"
        )
    }

    get Q(): CustomQuerier {
        return this.querier
    }
}

export {
    queryTx,
    sendTokens,
    SuperCosmosClient,
}