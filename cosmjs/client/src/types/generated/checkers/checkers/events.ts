/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "checkers.checkers";

export interface EventCreateGame {
  creator: string;
  gameIndex: string;
  block: string;
  red: string;
}

export interface EventMove {
  creator: string;
  gameIndex: string;
  capturedX: Long;
  capturedY: Long;
  winner: string;
}

export interface EventRejectGame {
  creator: string;
  gameIndex: string;
}

function createBaseEventCreateGame(): EventCreateGame {
  return { creator: "", gameIndex: "", block: "", red: "" };
}

export const EventCreateGame = {
  encode(message: EventCreateGame, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.gameIndex !== "") {
      writer.uint32(18).string(message.gameIndex);
    }
    if (message.block !== "") {
      writer.uint32(26).string(message.block);
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
          message.block = reader.string();
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
      block: isSet(object.block) ? String(object.block) : "",
      red: isSet(object.red) ? String(object.red) : "",
    };
  },

  toJSON(message: EventCreateGame): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.gameIndex !== undefined && (obj.gameIndex = message.gameIndex);
    message.block !== undefined && (obj.block = message.block);
    message.red !== undefined && (obj.red = message.red);
    return obj;
  },

  create<I extends Exact<DeepPartial<EventCreateGame>, I>>(base?: I): EventCreateGame {
    return EventCreateGame.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<EventCreateGame>, I>>(object: I): EventCreateGame {
    const message = createBaseEventCreateGame();
    message.creator = object.creator ?? "";
    message.gameIndex = object.gameIndex ?? "";
    message.block = object.block ?? "";
    message.red = object.red ?? "";
    return message;
  },
};

function createBaseEventMove(): EventMove {
  return { creator: "", gameIndex: "", capturedX: Long.ZERO, capturedY: Long.ZERO, winner: "" };
}

export const EventMove = {
  encode(message: EventMove, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.gameIndex !== "") {
      writer.uint32(18).string(message.gameIndex);
    }
    if (!message.capturedX.isZero()) {
      writer.uint32(24).int64(message.capturedX);
    }
    if (!message.capturedY.isZero()) {
      writer.uint32(32).int64(message.capturedY);
    }
    if (message.winner !== "") {
      writer.uint32(42).string(message.winner);
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
          message.capturedX = reader.int64() as Long;
          break;
        case 4:
          message.capturedY = reader.int64() as Long;
          break;
        case 5:
          message.winner = reader.string();
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
      capturedX: isSet(object.capturedX) ? Long.fromValue(object.capturedX) : Long.ZERO,
      capturedY: isSet(object.capturedY) ? Long.fromValue(object.capturedY) : Long.ZERO,
      winner: isSet(object.winner) ? String(object.winner) : "",
    };
  },

  toJSON(message: EventMove): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.gameIndex !== undefined && (obj.gameIndex = message.gameIndex);
    message.capturedX !== undefined && (obj.capturedX = (message.capturedX || Long.ZERO).toString());
    message.capturedY !== undefined && (obj.capturedY = (message.capturedY || Long.ZERO).toString());
    message.winner !== undefined && (obj.winner = message.winner);
    return obj;
  },

  create<I extends Exact<DeepPartial<EventMove>, I>>(base?: I): EventMove {
    return EventMove.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<EventMove>, I>>(object: I): EventMove {
    const message = createBaseEventMove();
    message.creator = object.creator ?? "";
    message.gameIndex = object.gameIndex ?? "";
    message.capturedX = (object.capturedX !== undefined && object.capturedX !== null)
      ? Long.fromValue(object.capturedX)
      : Long.ZERO;
    message.capturedY = (object.capturedY !== undefined && object.capturedY !== null)
      ? Long.fromValue(object.capturedY)
      : Long.ZERO;
    message.winner = object.winner ?? "";
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

  create<I extends Exact<DeepPartial<EventRejectGame>, I>>(base?: I): EventRejectGame {
    return EventRejectGame.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<EventRejectGame>, I>>(object: I): EventRejectGame {
    const message = createBaseEventRejectGame();
    message.creator = object.creator ?? "";
    message.gameIndex = object.gameIndex ?? "";
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
