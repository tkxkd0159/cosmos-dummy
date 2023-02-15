import {
    createProtobufRpcClient, QueryClient, defaultRegistryTypes, DeliverTxResponse, MsgSendEncodeObject,
    setupAuthExtension, setupBankExtension, setupStakingExtension, setupDistributionExtension, setupTxExtension, AuthExtension, BankExtension, StakingExtension, DistributionExtension, TxExtension,
    SigningStargateClient, SigningStargateClientOptions, StargateClient, StargateClientOptions,
    } from "@cosmjs/stargate"
import { StdFee } from "@cosmjs/amino"
import {EncodeObject, GeneratedType, Registry, OfflineSigner } from "@cosmjs/proto-signing"
import { Tendermint34Client } from "@cosmjs/tendermint-rpc";
import { Coin } from "cosmjs-types/cosmos/base/v1beta1/coin";

import { QueryClientImpl, QueryAllStoredGameResponse } from "@checker/query"
import type { Params } from "@checker/params"
import type { StoredGame, SystemInfo } from "@checker/types"
import type { PageRequest } from "@third-proto/base/query/v1beta1/pagination"
import { MsgCreateGame, MsgRejectGame, MsgPlayMove } from "@checker/tx";


function assertDefinedAndNotNull<T>(value: T | undefined | null, msg?: string): asserts value is T {
    if (value === undefined || value === null) {
        throw new Error(msg ?? "value is undefined or null");
    }
}

interface CheckerExtension {
    readonly checker: {
        readonly params: () => Promise<Params>
        readonly systemInfo: () => Promise<SystemInfo>
        readonly storedGame: (index: string) => Promise<StoredGame>
        readonly storedGameAll: (pagination?: PageRequest) => Promise<QueryAllStoredGameResponse>
    }
}

function setupCheckerExtension(base: QueryClient): CheckerExtension {
    const rpc = createProtobufRpcClient(base);
    const querier = new QueryClientImpl(rpc);

    return {
        checker: {
            params: async () => {
                const {params} = await querier.Params({})
                assertDefinedAndNotNull(params)
                return params
            },

            systemInfo: async() => {
                const {SystemInfo} = await querier.SystemInfo({})
                assertDefinedAndNotNull(SystemInfo)
                return SystemInfo
            },

            storedGame: async(index) => {
                const {storedGame} = await querier.StoredGame({index})
                assertDefinedAndNotNull(storedGame)
                return storedGame
            },

            storedGameAll: async(pagination)=> {
                return await querier.StoredGameAll({pagination})
            }
        }
    }
}

interface MsgCreateGameEncodeObject extends EncodeObject {
    readonly typeUrl: "/checkers.checkers.MsgCreateGame";
    readonly value: MsgCreateGame;
}

interface MsgRejectGameEncodeObject extends EncodeObject {
    readonly typeUrl: "/checkers.checkers.MsgRejectGame";
    readonly value: MsgRejectGame;
}

interface MsgPlayMoveEncodeObject extends EncodeObject {
    readonly typeUrl: "/checkers.checkers.MsgPlayMove";
    readonly value: MsgPlayMove;
}

const checkerTypes: ReadonlyArray<[string, GeneratedType]> = [
    ["/checkers.checkers.MsgCreateGame", MsgCreateGame],
    ["/checkers.checkers.MsgRejectGame", MsgRejectGame],
    ["/checkers.checkers.MsgPlayMove", MsgPlayMove],
]

const customRegistryTypes: ReadonlyArray<[string, GeneratedType]> = [
    ...defaultRegistryTypes,
    ...checkerTypes,
]

const genCustomRegistry = (): Registry => new Registry(customRegistryTypes)


// Test class with inheritance
class MyStargateClient extends StargateClient {
    public readonly myQueryClient: CustomQuerier | undefined

    public static async connect(
        endpoint: string,
        options: StargateClientOptions = {},
    ): Promise<MyStargateClient> {
        const tmClient = await Tendermint34Client.connect(endpoint)
        return new MyStargateClient(tmClient, options)
    }

    protected constructor(tmClient: Tendermint34Client | undefined, options: StargateClientOptions) {
        super(tmClient, options)
        if (tmClient) {
            this.myQueryClient = QueryClient.withExtensions(tmClient,  setupAuthExtension, setupBankExtension, setupStakingExtension, setupDistributionExtension, setupTxExtension, setupCheckerExtension)
        }
    }
}

type CustomQuerier = QueryClient & CheckerExtension & AuthExtension & BankExtension & StakingExtension & DistributionExtension & TxExtension

class MySigningStargateClient extends SigningStargateClient {
    public readonly myQueryClient: CustomQuerier | undefined

    public static async connectWithSigner(
        endpoint: string,
        signer: OfflineSigner,
        options: SigningStargateClientOptions = {}
    ): Promise<MySigningStargateClient> {
        const tmClient = await Tendermint34Client.connect(endpoint)
        return new MySigningStargateClient(tmClient, signer, {
            registry: genCustomRegistry(),
            ...options,
        })
    }

    protected constructor(tmClient: Tendermint34Client | undefined, signer: OfflineSigner, options: SigningStargateClientOptions) {
        super(tmClient, signer, options)
        if (tmClient) {
            this.myQueryClient = QueryClient.withExtensions(tmClient,  setupAuthExtension, setupBankExtension, setupStakingExtension, setupDistributionExtension, setupTxExtension, setupCheckerExtension)
        }
    }

    public async sendTokens(
        senderAddress: string,
        recipientAddress: string,
        amount: readonly Coin[],
        fee: StdFee | "auto" | number,
        memo = "",
    ): Promise<DeliverTxResponse> {
        const sendMsg: MsgSendEncodeObject = {
            typeUrl: "/cosmos.bank.v1beta1.MsgSend",
            value: {
                fromAddress: senderAddress,
                toAddress: recipientAddress,
                amount: [...amount],
            },
        };
        return this.signAndBroadcast(senderAddress, [sendMsg], fee, memo);
    }
}


export {
    setupCheckerExtension,
    CheckerExtension,
    MsgCreateGameEncodeObject,
    MsgRejectGameEncodeObject,
    MsgPlayMoveEncodeObject,
    genCustomRegistry,
    CustomQuerier,
}