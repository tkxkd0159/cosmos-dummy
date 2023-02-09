import {createProtobufRpcClient, QueryClient, defaultRegistryTypes } from "@cosmjs/stargate"
import {EncodeObject, GeneratedType, Registry} from "@cosmjs/proto-signing"

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
    readonly value: Partial<MsgCreateGame>;
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

export {
    setupCheckerExtension,
    CheckerExtension,
    MsgCreateGameEncodeObject,
    genCustomRegistry,
}