/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "checkers.checkers";

export interface MsgCreateGame {
  creator: string;
  black: string;
  red: string;
}

export interface MsgCreateGameResponse {
  gameIndex: string;
}

function createBaseMsgCreateGame(): MsgCreateGame {
  return { creator: "", black: "", red: "" };
}

export const MsgCreateGame = {
  encode(message: MsgCreateGame, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.black !== "") {
      writer.uint32(18).string(message.black);
    }
    if (message.red !== "") {
      writer.uint32(26).string(message.red);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateGame {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateGame();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.black = reader.string();
          break;
        case 3:
          message.red = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateGame {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      black: isSet(object.black) ? String(object.black) : "",
      red: isSet(object.red) ? String(object.red) : "",
    };
  },

  toJSON(message: MsgCreateGame): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.black !== undefined && (obj.black = message.black);
    message.red !== undefined && (obj.red = message.red);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateGame>, I>>(object: I): MsgCreateGame {
    const message = createBaseMsgCreateGame();
    message.creator = object.creator ?? "";
    message.black = object.black ?? "";
    message.red = object.red ?? "";
    return message;
  },
};

function createBaseMsgCreateGameResponse(): MsgCreateGameResponse {
  return { gameIndex: "" };
}

export const MsgCreateGameResponse = {
  encode(message: MsgCreateGameResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.gameIndex !== "") {
      writer.uint32(10).string(message.gameIndex);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateGameResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateGameResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.gameIndex = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateGameResponse {
    return { gameIndex: isSet(object.gameIndex) ? String(object.gameIndex) : "" };
  },

  toJSON(message: MsgCreateGameResponse): unknown {
    const obj: any = {};
    message.gameIndex !== undefined && (obj.gameIndex = message.gameIndex);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateGameResponse>, I>>(object: I): MsgCreateGameResponse {
    const message = createBaseMsgCreateGameResponse();
    message.gameIndex = object.gameIndex ?? "";
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  CreateGame(request: MsgCreateGame): Promise<MsgCreateGameResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.CreateGame = this.CreateGame.bind(this);
  }
  CreateGame(request: MsgCreateGame): Promise<MsgCreateGameResponse> {
    const data = MsgCreateGame.encode(request).finish();
    const promise = this.rpc.request("checkers.checkers.Msg", "CreateGame", data);
    return promise.then((data) => MsgCreateGameResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
