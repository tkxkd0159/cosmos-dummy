/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "checkers.checkers";

export interface SystemInfo {
  nextId: Long;
}

export interface StoredGame {
  index: string;
  board: string;
  turn: string;
  black: string;
  red: string;
  moveCount: Long;
}

function createBaseSystemInfo(): SystemInfo {
  return { nextId: Long.UZERO };
}

export const SystemInfo = {
  encode(message: SystemInfo, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (!message.nextId.isZero()) {
      writer.uint32(8).uint64(message.nextId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SystemInfo {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSystemInfo();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.nextId = reader.uint64() as Long;
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): SystemInfo {
    return { nextId: isSet(object.nextId) ? Long.fromValue(object.nextId) : Long.UZERO };
  },

  toJSON(message: SystemInfo): unknown {
    const obj: any = {};
    message.nextId !== undefined && (obj.nextId = (message.nextId || Long.UZERO).toString());
    return obj;
  },

  create<I extends Exact<DeepPartial<SystemInfo>, I>>(base?: I): SystemInfo {
    return SystemInfo.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<SystemInfo>, I>>(object: I): SystemInfo {
    const message = createBaseSystemInfo();
    message.nextId = (object.nextId !== undefined && object.nextId !== null)
      ? Long.fromValue(object.nextId)
      : Long.UZERO;
    return message;
  },
};

function createBaseStoredGame(): StoredGame {
  return { index: "", board: "", turn: "", black: "", red: "", moveCount: Long.UZERO };
}

export const StoredGame = {
  encode(message: StoredGame, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.board !== "") {
      writer.uint32(18).string(message.board);
    }
    if (message.turn !== "") {
      writer.uint32(26).string(message.turn);
    }
    if (message.black !== "") {
      writer.uint32(34).string(message.black);
    }
    if (message.red !== "") {
      writer.uint32(42).string(message.red);
    }
    if (!message.moveCount.isZero()) {
      writer.uint32(48).uint64(message.moveCount);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): StoredGame {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseStoredGame();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.board = reader.string();
          break;
        case 3:
          message.turn = reader.string();
          break;
        case 4:
          message.black = reader.string();
          break;
        case 5:
          message.red = reader.string();
          break;
        case 6:
          message.moveCount = reader.uint64() as Long;
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): StoredGame {
    return {
      index: isSet(object.index) ? String(object.index) : "",
      board: isSet(object.board) ? String(object.board) : "",
      turn: isSet(object.turn) ? String(object.turn) : "",
      black: isSet(object.black) ? String(object.black) : "",
      red: isSet(object.red) ? String(object.red) : "",
      moveCount: isSet(object.moveCount) ? Long.fromValue(object.moveCount) : Long.UZERO,
    };
  },

  toJSON(message: StoredGame): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.board !== undefined && (obj.board = message.board);
    message.turn !== undefined && (obj.turn = message.turn);
    message.black !== undefined && (obj.black = message.black);
    message.red !== undefined && (obj.red = message.red);
    message.moveCount !== undefined && (obj.moveCount = (message.moveCount || Long.UZERO).toString());
    return obj;
  },

  create<I extends Exact<DeepPartial<StoredGame>, I>>(base?: I): StoredGame {
    return StoredGame.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<StoredGame>, I>>(object: I): StoredGame {
    const message = createBaseStoredGame();
    message.index = object.index ?? "";
    message.board = object.board ?? "";
    message.turn = object.turn ?? "";
    message.black = object.black ?? "";
    message.red = object.red ?? "";
    message.moveCount = (object.moveCount !== undefined && object.moveCount !== null)
      ? Long.fromValue(object.moveCount)
      : Long.UZERO;
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Long ? string | number | Long : T extends Array<infer U> ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
