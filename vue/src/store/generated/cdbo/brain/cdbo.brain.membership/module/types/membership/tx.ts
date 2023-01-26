/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

export const protobufPackage = "cdbo.brain.membership";

export interface MsgEnroll {
  member_address: string;
  nickname: string;
}

export interface MsgEnrollResponse {}

export interface MsgUpdateStatus {
  creator: string;
  address: string;
  status: string;
}

export interface MsgUpdateStatusResponse {}

const baseMsgEnroll: object = { member_address: "", nickname: "" };

export const MsgEnroll = {
  encode(message: MsgEnroll, writer: Writer = Writer.create()): Writer {
    if (message.member_address !== "") {
      writer.uint32(10).string(message.member_address);
    }
    if (message.nickname !== "") {
      writer.uint32(18).string(message.nickname);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgEnroll {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgEnroll } as MsgEnroll;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.member_address = reader.string();
          break;
        case 2:
          message.nickname = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgEnroll {
    const message = { ...baseMsgEnroll } as MsgEnroll;
    if (object.member_address !== undefined && object.member_address !== null) {
      message.member_address = String(object.member_address);
    } else {
      message.member_address = "";
    }
    if (object.nickname !== undefined && object.nickname !== null) {
      message.nickname = String(object.nickname);
    } else {
      message.nickname = "";
    }
    return message;
  },

  toJSON(message: MsgEnroll): unknown {
    const obj: any = {};
    message.member_address !== undefined &&
      (obj.member_address = message.member_address);
    message.nickname !== undefined && (obj.nickname = message.nickname);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgEnroll>): MsgEnroll {
    const message = { ...baseMsgEnroll } as MsgEnroll;
    if (object.member_address !== undefined && object.member_address !== null) {
      message.member_address = object.member_address;
    } else {
      message.member_address = "";
    }
    if (object.nickname !== undefined && object.nickname !== null) {
      message.nickname = object.nickname;
    } else {
      message.nickname = "";
    }
    return message;
  },
};

const baseMsgEnrollResponse: object = {};

export const MsgEnrollResponse = {
  encode(_: MsgEnrollResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgEnrollResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgEnrollResponse } as MsgEnrollResponse;
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

  fromJSON(_: any): MsgEnrollResponse {
    const message = { ...baseMsgEnrollResponse } as MsgEnrollResponse;
    return message;
  },

  toJSON(_: MsgEnrollResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgEnrollResponse>): MsgEnrollResponse {
    const message = { ...baseMsgEnrollResponse } as MsgEnrollResponse;
    return message;
  },
};

const baseMsgUpdateStatus: object = { creator: "", address: "", status: "" };

export const MsgUpdateStatus = {
  encode(message: MsgUpdateStatus, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    if (message.status !== "") {
      writer.uint32(26).string(message.status);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateStatus {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUpdateStatus } as MsgUpdateStatus;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.address = reader.string();
          break;
        case 3:
          message.status = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateStatus {
    const message = { ...baseMsgUpdateStatus } as MsgUpdateStatus;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = String(object.status);
    } else {
      message.status = "";
    }
    return message;
  },

  toJSON(message: MsgUpdateStatus): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.address !== undefined && (obj.address = message.address);
    message.status !== undefined && (obj.status = message.status);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgUpdateStatus>): MsgUpdateStatus {
    const message = { ...baseMsgUpdateStatus } as MsgUpdateStatus;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = object.status;
    } else {
      message.status = "";
    }
    return message;
  },
};

const baseMsgUpdateStatusResponse: object = {};

export const MsgUpdateStatusResponse = {
  encode(_: MsgUpdateStatusResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateStatusResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgUpdateStatusResponse,
    } as MsgUpdateStatusResponse;
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

  fromJSON(_: any): MsgUpdateStatusResponse {
    const message = {
      ...baseMsgUpdateStatusResponse,
    } as MsgUpdateStatusResponse;
    return message;
  },

  toJSON(_: MsgUpdateStatusResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgUpdateStatusResponse>
  ): MsgUpdateStatusResponse {
    const message = {
      ...baseMsgUpdateStatusResponse,
    } as MsgUpdateStatusResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  Enroll(request: MsgEnroll): Promise<MsgEnrollResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  UpdateStatus(request: MsgUpdateStatus): Promise<MsgUpdateStatusResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Enroll(request: MsgEnroll): Promise<MsgEnrollResponse> {
    const data = MsgEnroll.encode(request).finish();
    const promise = this.rpc.request(
      "cdbo.brain.membership.Msg",
      "Enroll",
      data
    );
    return promise.then((data) => MsgEnrollResponse.decode(new Reader(data)));
  }

  UpdateStatus(request: MsgUpdateStatus): Promise<MsgUpdateStatusResponse> {
    const data = MsgUpdateStatus.encode(request).finish();
    const promise = this.rpc.request(
      "cdbo.brain.membership.Msg",
      "UpdateStatus",
      data
    );
    return promise.then((data) =>
      MsgUpdateStatusResponse.decode(new Reader(data))
    );
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
