/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

export const protobufPackage = "jsc.checkers";

export interface MsgCreateGame {
  creator: string;
  red: string;
  black: string;
}

export interface MsgCreateGameResponse {
  idValue: string;
}

export interface MsgEndGame {
  creator: string;
  sig: string;
}

export interface MsgEndGameResponse {
  res: string;
  hash: string;
}

const baseMsgCreateGame: object = { creator: "", red: "", black: "" };

export const MsgCreateGame = {
  encode(message: MsgCreateGame, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.red !== "") {
      writer.uint32(18).string(message.red);
    }
    if (message.black !== "") {
      writer.uint32(26).string(message.black);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateGame {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateGame } as MsgCreateGame;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.red = reader.string();
          break;
        case 3:
          message.black = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateGame {
    const message = { ...baseMsgCreateGame } as MsgCreateGame;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.red !== undefined && object.red !== null) {
      message.red = String(object.red);
    } else {
      message.red = "";
    }
    if (object.black !== undefined && object.black !== null) {
      message.black = String(object.black);
    } else {
      message.black = "";
    }
    return message;
  },

  toJSON(message: MsgCreateGame): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.red !== undefined && (obj.red = message.red);
    message.black !== undefined && (obj.black = message.black);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateGame>): MsgCreateGame {
    const message = { ...baseMsgCreateGame } as MsgCreateGame;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.red !== undefined && object.red !== null) {
      message.red = object.red;
    } else {
      message.red = "";
    }
    if (object.black !== undefined && object.black !== null) {
      message.black = object.black;
    } else {
      message.black = "";
    }
    return message;
  },
};

const baseMsgCreateGameResponse: object = { idValue: "" };

export const MsgCreateGameResponse = {
  encode(
    message: MsgCreateGameResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.idValue !== "") {
      writer.uint32(10).string(message.idValue);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateGameResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateGameResponse } as MsgCreateGameResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.idValue = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateGameResponse {
    const message = { ...baseMsgCreateGameResponse } as MsgCreateGameResponse;
    if (object.idValue !== undefined && object.idValue !== null) {
      message.idValue = String(object.idValue);
    } else {
      message.idValue = "";
    }
    return message;
  },

  toJSON(message: MsgCreateGameResponse): unknown {
    const obj: any = {};
    message.idValue !== undefined && (obj.idValue = message.idValue);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateGameResponse>
  ): MsgCreateGameResponse {
    const message = { ...baseMsgCreateGameResponse } as MsgCreateGameResponse;
    if (object.idValue !== undefined && object.idValue !== null) {
      message.idValue = object.idValue;
    } else {
      message.idValue = "";
    }
    return message;
  },
};

const baseMsgEndGame: object = { creator: "", sig: "" };

export const MsgEndGame = {
  encode(message: MsgEndGame, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.sig !== "") {
      writer.uint32(18).string(message.sig);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgEndGame {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgEndGame } as MsgEndGame;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.sig = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgEndGame {
    const message = { ...baseMsgEndGame } as MsgEndGame;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.sig !== undefined && object.sig !== null) {
      message.sig = String(object.sig);
    } else {
      message.sig = "";
    }
    return message;
  },

  toJSON(message: MsgEndGame): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.sig !== undefined && (obj.sig = message.sig);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgEndGame>): MsgEndGame {
    const message = { ...baseMsgEndGame } as MsgEndGame;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.sig !== undefined && object.sig !== null) {
      message.sig = object.sig;
    } else {
      message.sig = "";
    }
    return message;
  },
};

const baseMsgEndGameResponse: object = { res: "", hash: "" };

export const MsgEndGameResponse = {
  encode(
    message: MsgEndGameResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.res !== "") {
      writer.uint32(10).string(message.res);
    }
    if (message.hash !== "") {
      writer.uint32(18).string(message.hash);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgEndGameResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgEndGameResponse } as MsgEndGameResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.res = reader.string();
          break;
        case 2:
          message.hash = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgEndGameResponse {
    const message = { ...baseMsgEndGameResponse } as MsgEndGameResponse;
    if (object.res !== undefined && object.res !== null) {
      message.res = String(object.res);
    } else {
      message.res = "";
    }
    if (object.hash !== undefined && object.hash !== null) {
      message.hash = String(object.hash);
    } else {
      message.hash = "";
    }
    return message;
  },

  toJSON(message: MsgEndGameResponse): unknown {
    const obj: any = {};
    message.res !== undefined && (obj.res = message.res);
    message.hash !== undefined && (obj.hash = message.hash);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgEndGameResponse>): MsgEndGameResponse {
    const message = { ...baseMsgEndGameResponse } as MsgEndGameResponse;
    if (object.res !== undefined && object.res !== null) {
      message.res = object.res;
    } else {
      message.res = "";
    }
    if (object.hash !== undefined && object.hash !== null) {
      message.hash = object.hash;
    } else {
      message.hash = "";
    }
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateGame(request: MsgCreateGame): Promise<MsgCreateGameResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  EndGame(request: MsgEndGame): Promise<MsgEndGameResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  CreateGame(request: MsgCreateGame): Promise<MsgCreateGameResponse> {
    const data = MsgCreateGame.encode(request).finish();
    const promise = this.rpc.request("jsc.checkers.Msg", "CreateGame", data);
    return promise.then((data) =>
      MsgCreateGameResponse.decode(new Reader(data))
    );
  }

  EndGame(request: MsgEndGame): Promise<MsgEndGameResponse> {
    const data = MsgEndGame.encode(request).finish();
    const promise = this.rpc.request("jsc.checkers.Msg", "EndGame", data);
    return promise.then((data) => MsgEndGameResponse.decode(new Reader(data)));
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
