import {ChangeEvent, Component, MouseEvent} from "react"
import styles from '@/styles/Home.module.css'
import {Window as KeplrWindow} from "@keplr-wallet/types";
import {StargateClient, SigningStargateClient, Coin} from "@cosmjs/stargate"
import {AccountData} from "@cosmjs/proto-signing";

declare global {
    interface Window extends KeplrWindow {
    }
}

interface FaucetSenderState {
    denom: string
    faucetBalance: string
    myAddress: string
    myBalance: string
    toSend: string
}

export interface FaucetSenderProps {
    faucetAddress: string
    rpcUrl: string
}

export class FaucetSender extends Component<FaucetSenderProps, FaucetSenderState> {


    // Set the initial state
    constructor(props: FaucetSenderProps) {
        super(props)
        this.state = {
            denom: "Loading...",
            faucetBalance: "Loading...",
            myAddress: "Click first",
            myBalance: "Click first",
            toSend: "0",
        }
    }

    init = async () => this.updateFaucetBalance(await StargateClient.connect(this.props.rpcUrl))

    // Get the faucet's balance
    updateFaucetBalance = async (client: StargateClient) => {
        const balances: readonly Coin[] = await client.getAllBalances(this.props.faucetAddress)
        const first: Coin = balances[0]
        const offlineSigner =
            window.getOfflineSigner!("theta-testnet-001")
        const signingClient = await SigningStargateClient.connectWithSigner(
            this.props.rpcUrl,
            offlineSigner,
        )
        // Get the address and balance of your user
        const myAcc: AccountData = (await offlineSigner.getAccounts())[0]
        const amt: string = (await signingClient.getBalance(myAcc.address, first.denom)).amount
        this.setState((prev) => {
            return {
                ...prev,
                denom: first.denom,
                faucetBalance: first.amount,
                myAddress: myAcc.address,
                myBalance: amt,
            }
        })
    }

    // Store changed token amount to state
    onToSendChanged = (e: ChangeEvent<HTMLInputElement>) => {
        const { value } = e.currentTarget
        this.setState((prev) => {
            return {
                ...prev,
                toSend: value,
            }
        })
    }

    // When the user clicks the "send to faucet button"
    onSendClicked = async (e: MouseEvent<HTMLButtonElement>) => {
        const {keplr} = window
        if (!keplr) {
            alert("You need to install Keplr")
            return
        }

        const {denom, toSend} = this.state
        const {faucetAddress, rpcUrl} = this.props

        const offlineSigner = window.getOfflineSigner!("theta-testnet-001")
        const signingClient = await SigningStargateClient.connectWithSigner(
            rpcUrl,
            offlineSigner,
        )
        // Get the address and balance of your user
        const account: AccountData = (await offlineSigner.getAccounts())[0]
        // Submit the transaction to send tokens to the faucet
        const sendResult = await signingClient.sendTokens(
            account.address,
            faucetAddress,
            [
                {
                    denom: denom,
                    amount: toSend,
                },
            ],
            {
                amount: [{denom: "uatom", amount: "500"}],
                gas: "200000",
            },
        )
        // Print the result to the console
        // console.log(sendResult)

        // Update the balance in the user interface
        const afterBalance = (await signingClient.getBalance(account.address, denom)).amount
        const afterFaucetBalance = (await signingClient.getBalance(faucetAddress, denom)).amount
        this.setState((prev) => {
            return {
                ...prev,
                myBalance: afterBalance,
                faucetBalance: afterFaucetBalance,
            }
        })
    }

    // The render function that draws the component at init and at state change
    render() {
        const {denom, faucetBalance, myAddress, myBalance, toSend} = this.state
        const {faucetAddress} = this.props

        return <div>
            <fieldset className={styles.card}>
                <legend>Faucet</legend>
                <p>Address: {faucetAddress}</p>
                <p>Balance: {faucetBalance}</p>
            </fieldset>
            <fieldset className={styles.card}>
                <legend>You</legend>
                <p>Address: {myAddress}</p>
                <p>Balance: {myBalance}</p>
            </fieldset>
            <fieldset className={styles.card}>
                <legend>Send</legend>
                <p>To faucet:</p>
                <input value={toSend} type="number" onChange={this.onToSendChanged}/> {denom}
                <br/>
                <button onClick={this.onSendClicked}>Send to faucet</button>
            </fieldset>
        </div>
    }

    async componentDidMount() {
        if (!window.getOfflineSigner || !window.keplr) {
            alert("Please install keplr extension");
        } else {
            if (window.keplr.experimentalSuggestChain) {
                try {
                    await window.keplr.experimentalSuggestChain({
                        chainId: "theta-testnet-001",
                        chainName: "theta-testnet-001",
                        rpc: "https://rpc.sentry-01.theta-testnet.polypore.xyz/",
                        rest: "https://rest.sentry-01.theta-testnet.polypore.xyz/",
                        bip44: {
                            coinType: 118,
                        },
                        bech32Config: {
                            bech32PrefixAccAddr: "cosmos",
                            bech32PrefixAccPub: "cosmos" + "pub",
                            bech32PrefixValAddr: "cosmos" + "valoper",
                            bech32PrefixValPub: "cosmos" + "valoperpub",
                            bech32PrefixConsAddr: "cosmos" + "valcons",
                            bech32PrefixConsPub: "cosmos" + "valconspub",
                        },
                        currencies: [
                            {
                                coinDenom: "ATOM",
                                coinMinimalDenom: "uatom",
                                coinDecimals: 6,
                                coinGeckoId: "cosmos",
                            },
                            {
                                coinDenom: "THETA",
                                coinMinimalDenom: "theta",
                                coinDecimals: 0,
                            },
                            {
                                coinDenom: "LAMBDA",
                                coinMinimalDenom: "lambda",
                                coinDecimals: 0,
                            },
                            {
                                coinDenom: "RHO",
                                coinMinimalDenom: "rho",
                                coinDecimals: 0,
                            },
                            {
                                coinDenom: "EPSILON",
                                coinMinimalDenom: "epsilon",
                                coinDecimals: 0,
                            },
                        ],
                        feeCurrencies: [
                            {
                                coinDenom: "ATOM",
                                coinMinimalDenom: "uatom",
                                coinDecimals: 6,
                                coinGeckoId: "cosmos",
                                gasPriceStep: {
                                    low: 0.1,
                                    average: 1,
                                    high: 10,
                                },
                            },
                        ],
                        stakeCurrency: {
                            coinDenom: "ATOM",
                            coinMinimalDenom: "uatom",
                            coinDecimals: 6,
                            coinGeckoId: "cosmos",
                        },
                        features: ["stargate", "ibc-transfer", "no-legacy-stdTx"],
                    });
                } catch(err) {
                    alert(`Failed to suggest the chain ${err}`);
                }
            } else {
                alert("Please use the recent version of keplr extension");
            }
        }

        // Initialize balance updater
        await this.init()
    }
}
