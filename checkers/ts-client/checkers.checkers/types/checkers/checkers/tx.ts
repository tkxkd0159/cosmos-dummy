/* eslint-disable */
import Long from "long";
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

export interface MsgPlayMove {
  creator: string;
  gameIndex: string;
  fromX: number;
  fromY: number;
  toX: number;
  toY: number;
}

export interface MsgPlayMoveResponse {
  capturedX: number;
  capturedY: number;
  winner: string;
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

function createBaseMsgPlayMove(): MsgPlayMove {
  return { creator: "", gameIndex: "", fromX: 0, fromY: 0, toX: 0, toY: 0 };
}

export const MsgPlayMove = {
  encode(message: MsgPlayMove, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.gameIndex !== "") {
      writer.uint32(18).string(message.gameIndex);
    }
    if (message.fromX !== 0) {
      writer.uint32(24).uint64(message.fromX);
    }
    if (message.fromY !== 0) {
      writer.uint32(32).uint64(message.fromY);
    }
    if (message.toX !== 0) {
      writer.uint32(40).uint64(message.toX);
    }
    if (message.toY !== 0) {
      writer.uint32(48).uint64(message.toY);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgPlayMove {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgPlayMove();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.gameIndex = reader.string();
          break;
        case 3:
          message.fromX = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.fromY = longToNumber(reader.uint64() as Long);
          break;
        case 5:
          message.toX = longToNumber(reader.uint64() as Long);
          break;
        case 6:
          message.toY = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgPlayMove {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      gameIndex: isSet(object.gameIndex) ? String(object.gameIndex) : "",
      fromX: isSet(object.fromX) ? Number(object.fromX) : 0,
      fromY: isSet(object.fromY) ? Number(object.fromY) : 0,
      toX: isSet(object.toX) ? Number(object.toX) : 0,
      toY: isSet(object.toY) ? Number(object.toY) : 0,
    };
  },

  toJSON(message: MsgPlayMove): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.gameIndex !== undefined && (obj.gameIndex = message.gameIndex);
    message.fromX !== undefined && (obj.fromX = Math.round(message.fromX));
    message.fromY !== undefined && (obj.fromY = Math.round(message.fromY));
    message.toX !== undefined && (obj.toX = Math.round(message.toX));
    message.toY !== undefined && (obj.toY = Math.round(message.toY));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgPlayMove>, I>>(object: I): MsgPlayMove {
    const message = createBaseMsgPlayMove();
    message.creator = object.creator ?? "";
    message.gameIndex = object.gameIndex ?? "";
    message.fromX = object.fromX ?? 0;
    message.fromY = object.fromY ?? 0;
    message.toX = object.toX ?? 0;
    message.toY = object.toY ?? 0;
    return message;
  },
};

function createBaseMsgPlayMoveResponse(): MsgPlayMoveResponse {
  return { capturedX: 0, capturedY: 0, winner: "" };
}

export const MsgPlayMoveResponse = {
  encode(message: MsgPlayMoveResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.capturedX !== 0) {
      writer.uint32(8).int32(message.capturedX);
    }
    if (message.capturedY !== 0) {
      writer.uint32(16).int32(message.capturedY);
    }
    if (message.winner !== "") {
      writer.uint32(26).string(message.winner);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgPlayMoveResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgPlayMoveResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.capturedX = reader.int32();
          break;
        case 2:
          message.capturedY = reader.int32();
          break;
        case 3:
          message.winner = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgPlayMoveResponse {
    return {
      capturedX: isSet(object.capturedX) ? Number(object.capturedX) : 0,
      capturedY: isSet(object.capturedY) ? Number(object.capturedY) : 0,
      winner: isSet(object.winner) ? String(object.winner) : "",
    };
  },

  toJSON(message: MsgPlayMoveResponse): unknown {
    const obj: any = {};
    message.capturedX !== undefined && (obj.capturedX = Math.round(message.capturedX));
    message.capturedY !== undefined && (obj.capturedY = Math.round(message.capturedY));
    message.winner !== undefined && (obj.winner = message.winner);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgPlayMoveResponse>, I>>(object: I): MsgPlayMoveResponse {
    const message = createBaseMsgPlayMoveResponse();
    message.capturedX = object.capturedX ?? 0;
    message.capturedY = object.capturedY ?? 0;
    message.winner = object.winner ?? "";
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateGame(request: MsgCreateGame): Promise<MsgCreateGameResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  PlayMove(request: MsgPlayMove): Promise<MsgPlayMoveResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.CreateGame = this.CreateGame.bind(this);
    this.PlayMove = this.PlayMove.bind(this);
  }
  CreateGame(request: MsgCreateGame): Promise<MsgCreateGameResponse> {
    const data = MsgCreateGame.encode(request).finish();
    const promise = this.rpc.request("checkers.checkers.Msg", "CreateGame", data);
    return promise.then((data) => MsgCreateGameResponse.decode(new _m0.Reader(data)));
  }

  PlayMove(request: MsgPlayMove): Promise<MsgPlayMoveResponse> {
    const data = MsgPlayMove.encode(request).finish();
    const promise = this.rpc.request("checkers.checkers.Msg", "PlayMove", data);
    return promise.then((data) => MsgPlayMoveResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
