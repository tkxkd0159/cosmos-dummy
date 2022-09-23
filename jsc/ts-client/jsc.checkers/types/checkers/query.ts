/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../checkers/params";

export const protobufPackage = "jsc.checkers";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryMyqRequest {
  req1: string;
  req2: string;
}

export interface QueryMyqResponse {
  res1: string;
  res2: string;
}

const baseQueryParamsRequest: object = {};

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },
};

const baseQueryParamsResponse: object = {};

export const QueryParamsResponse = {
  encode(
    message: QueryParamsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },
};

const baseQueryMyqRequest: object = { req1: "", req2: "" };

export const QueryMyqRequest = {
  encode(message: QueryMyqRequest, writer: Writer = Writer.create()): Writer {
    if (message.req1 !== "") {
      writer.uint32(10).string(message.req1);
    }
    if (message.req2 !== "") {
      writer.uint32(18).string(message.req2);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryMyqRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryMyqRequest } as QueryMyqRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.req1 = reader.string();
          break;
        case 2:
          message.req2 = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryMyqRequest {
    const message = { ...baseQueryMyqRequest } as QueryMyqRequest;
    if (object.req1 !== undefined && object.req1 !== null) {
      message.req1 = String(object.req1);
    } else {
      message.req1 = "";
    }
    if (object.req2 !== undefined && object.req2 !== null) {
      message.req2 = String(object.req2);
    } else {
      message.req2 = "";
    }
    return message;
  },

  toJSON(message: QueryMyqRequest): unknown {
    const obj: any = {};
    message.req1 !== undefined && (obj.req1 = message.req1);
    message.req2 !== undefined && (obj.req2 = message.req2);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryMyqRequest>): QueryMyqRequest {
    const message = { ...baseQueryMyqRequest } as QueryMyqRequest;
    if (object.req1 !== undefined && object.req1 !== null) {
      message.req1 = object.req1;
    } else {
      message.req1 = "";
    }
    if (object.req2 !== undefined && object.req2 !== null) {
      message.req2 = object.req2;
    } else {
      message.req2 = "";
    }
    return message;
  },
};

const baseQueryMyqResponse: object = { res1: "", res2: "" };

export const QueryMyqResponse = {
  encode(message: QueryMyqResponse, writer: Writer = Writer.create()): Writer {
    if (message.res1 !== "") {
      writer.uint32(10).string(message.res1);
    }
    if (message.res2 !== "") {
      writer.uint32(18).string(message.res2);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryMyqResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryMyqResponse } as QueryMyqResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.res1 = reader.string();
          break;
        case 2:
          message.res2 = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryMyqResponse {
    const message = { ...baseQueryMyqResponse } as QueryMyqResponse;
    if (object.res1 !== undefined && object.res1 !== null) {
      message.res1 = String(object.res1);
    } else {
      message.res1 = "";
    }
    if (object.res2 !== undefined && object.res2 !== null) {
      message.res2 = String(object.res2);
    } else {
      message.res2 = "";
    }
    return message;
  },

  toJSON(message: QueryMyqResponse): unknown {
    const obj: any = {};
    message.res1 !== undefined && (obj.res1 = message.res1);
    message.res2 !== undefined && (obj.res2 = message.res2);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryMyqResponse>): QueryMyqResponse {
    const message = { ...baseQueryMyqResponse } as QueryMyqResponse;
    if (object.res1 !== undefined && object.res1 !== null) {
      message.res1 = object.res1;
    } else {
      message.res1 = "";
    }
    if (object.res2 !== undefined && object.res2 !== null) {
      message.res2 = object.res2;
    } else {
      message.res2 = "";
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a list of Myq items. */
  Myq(request: QueryMyqRequest): Promise<QueryMyqResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request("jsc.checkers.Query", "Params", data);
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  Myq(request: QueryMyqRequest): Promise<QueryMyqResponse> {
    const data = QueryMyqRequest.encode(request).finish();
    const promise = this.rpc.request("jsc.checkers.Query", "Myq", data);
    return promise.then((data) => QueryMyqResponse.decode(new Reader(data)));
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
