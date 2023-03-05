/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "checkers.checkers";

export interface EventCreateGame {
  creator: string;
  gameIndex: string;
  black: string;
  red: string;
}

export interface EventMove {
  creator: string;
  gameIndex: string;
  capturedX: number;
  capturedY: number;
  winner: string;
  board: string;
}

export interface EventRejectGame {
  creator: string;
  gameIndex: string;
}

export interface EventForfeitGame {
  gameIndex: string;
  winner: string;
  board: string;
}

function createBaseEventCreateGame(): EventCreateGame {
  return { creator: "", gameIndex: "", black: "", red: "" };
}

export const EventCreateGame = {
  encode(message: EventCreateGame, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.gameIndex !== "") {
      writer.uint32(18).string(message.gameIndex);
    }
    if (message.black !== "") {
      writer.uint32(26).string(message.black);
    }
    if (message.red !== "") {
      writer.uint32(34).string(message.red);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): EventCreateGame {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEventCreateGame();
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
          message.black = reader.string();
          break;
        case 4:
          message.red = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EventCreateGame {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      gameIndex: isSet(object.gameIndex) ? String(object.gameIndex) : "",
      black: isSet(object.black) ? String(object.black) : "",
      red: isSet(object.red) ? String(object.red) : "",
    };
  },

  toJSON(message: EventCreateGame): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.gameIndex !== undefined && (obj.gameIndex = message.gameIndex);
    message.black !== undefined && (obj.black = message.black);
    message.red !== undefined && (obj.red = message.red);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<EventCreateGame>, I>>(object: I): EventCreateGame {
    const message = createBaseEventCreateGame();
    message.creator = object.creator ?? "";
    message.gameIndex = object.gameIndex ?? "";
    message.black = object.black ?? "";
    message.red = object.red ?? "";
    return message;
  },
};

function createBaseEventMove(): EventMove {
  return { creator: "", gameIndex: "", capturedX: 0, capturedY: 0, winner: "", board: "" };
}

export const EventMove = {
  encode(message: EventMove, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.gameIndex !== "") {
      writer.uint32(18).string(message.gameIndex);
    }
    if (message.capturedX !== 0) {
      writer.uint32(24).int64(message.capturedX);
    }
    if (message.capturedY !== 0) {
      writer.uint32(32).int64(message.capturedY);
    }
    if (message.winner !== "") {
      writer.uint32(42).string(message.winner);
    }
    if (message.board !== "") {
      writer.uint32(50).string(message.board);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): EventMove {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEventMove();
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
          message.capturedX = longToNumber(reader.int64() as Long);
          break;
        case 4:
          message.capturedY = longToNumber(reader.int64() as Long);
          break;
        case 5:
          message.winner = reader.string();
          break;
        case 6:
          message.board = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EventMove {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      gameIndex: isSet(object.gameIndex) ? String(object.gameIndex) : "",
      capturedX: isSet(object.capturedX) ? Number(object.capturedX) : 0,
      capturedY: isSet(object.capturedY) ? Number(object.capturedY) : 0,
      winner: isSet(object.winner) ? String(object.winner) : "",
      board: isSet(object.board) ? String(object.board) : "",
    };
  },

  toJSON(message: EventMove): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.gameIndex !== undefined && (obj.gameIndex = message.gameIndex);
    message.capturedX !== undefined && (obj.capturedX = Math.round(message.capturedX));
    message.capturedY !== undefined && (obj.capturedY = Math.round(message.capturedY));
    message.winner !== undefined && (obj.winner = message.winner);
    message.board !== undefined && (obj.board = message.board);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<EventMove>, I>>(object: I): EventMove {
    const message = createBaseEventMove();
    message.creator = object.creator ?? "";
    message.gameIndex = object.gameIndex ?? "";
    message.capturedX = object.capturedX ?? 0;
    message.capturedY = object.capturedY ?? 0;
    message.winner = object.winner ?? "";
    message.board = object.board ?? "";
    return message;
  },
};

function createBaseEventRejectGame(): EventRejectGame {
  return { creator: "", gameIndex: "" };
}

export const EventRejectGame = {
  encode(message: EventRejectGame, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.gameIndex !== "") {
      writer.uint32(18).string(message.gameIndex);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): EventRejectGame {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEventRejectGame();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.gameIndex = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EventRejectGame {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      gameIndex: isSet(object.gameIndex) ? String(object.gameIndex) : "",
    };
  },

  toJSON(message: EventRejectGame): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.gameIndex !== undefined && (obj.gameIndex = message.gameIndex);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<EventRejectGame>, I>>(object: I): EventRejectGame {
    const message = createBaseEventRejectGame();
    message.creator = object.creator ?? "";
    message.gameIndex = object.gameIndex ?? "";
    return message;
  },
};

function createBaseEventForfeitGame(): EventForfeitGame {
  return { gameIndex: "", winner: "", board: "" };
}

export const EventForfeitGame = {
  encode(message: EventForfeitGame, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.gameIndex !== "") {
      writer.uint32(10).string(message.gameIndex);
    }
    if (message.winner !== "") {
      writer.uint32(18).string(message.winner);
    }
    if (message.board !== "") {
      writer.uint32(26).string(message.board);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): EventForfeitGame {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEventForfeitGame();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.gameIndex = reader.string();
          break;
        case 2:
          message.winner = reader.string();
          break;
        case 3:
          message.board = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EventForfeitGame {
    return {
      gameIndex: isSet(object.gameIndex) ? String(object.gameIndex) : "",
      winner: isSet(object.winner) ? String(object.winner) : "",
      board: isSet(object.board) ? String(object.board) : "",
    };
  },

  toJSON(message: EventForfeitGame): unknown {
    const obj: any = {};
    message.gameIndex !== undefined && (obj.gameIndex = message.gameIndex);
    message.winner !== undefined && (obj.winner = message.winner);
    message.board !== undefined && (obj.board = message.board);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<EventForfeitGame>, I>>(object: I): EventForfeitGame {
    const message = createBaseEventForfeitGame();
    message.gameIndex = object.gameIndex ?? "";
    message.winner = object.winner ?? "";
    message.board = object.board ?? "";
    return message;
  },
};

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
