/* eslint-disable */
import {
  MembershipStatus,
  membershipStatusFromJSON,
  membershipStatusToJSON,
} from "../membership/member";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "cdbo.brain.membership";

/** EventMemberEnrolled is an event emitted when a new member joins The Denom */
export interface EventMemberEnrolled {
  member_address: string;
}

/**
 * EventMemberStatusChanged is an event emitted when a member's citizenship
 * status changes
 */
export interface EventMemberStatusChanged {
  member_address: string;
  status: MembershipStatus;
  previous_status: MembershipStatus;
}

const baseEventMemberEnrolled: object = { member_address: "" };

export const EventMemberEnrolled = {
  encode(
    message: EventMemberEnrolled,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.member_address !== "") {
      writer.uint32(10).string(message.member_address);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): EventMemberEnrolled {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseEventMemberEnrolled } as EventMemberEnrolled;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.member_address = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EventMemberEnrolled {
    const message = { ...baseEventMemberEnrolled } as EventMemberEnrolled;
    if (object.member_address !== undefined && object.member_address !== null) {
      message.member_address = String(object.member_address);
    } else {
      message.member_address = "";
    }
    return message;
  },

  toJSON(message: EventMemberEnrolled): unknown {
    const obj: any = {};
    message.member_address !== undefined &&
      (obj.member_address = message.member_address);
    return obj;
  },

  fromPartial(object: DeepPartial<EventMemberEnrolled>): EventMemberEnrolled {
    const message = { ...baseEventMemberEnrolled } as EventMemberEnrolled;
    if (object.member_address !== undefined && object.member_address !== null) {
      message.member_address = object.member_address;
    } else {
      message.member_address = "";
    }
    return message;
  },
};

const baseEventMemberStatusChanged: object = {
  member_address: "",
  status: 0,
  previous_status: 0,
};

export const EventMemberStatusChanged = {
  encode(
    message: EventMemberStatusChanged,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.member_address !== "") {
      writer.uint32(10).string(message.member_address);
    }
    if (message.status !== 0) {
      writer.uint32(16).int32(message.status);
    }
    if (message.previous_status !== 0) {
      writer.uint32(24).int32(message.previous_status);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): EventMemberStatusChanged {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseEventMemberStatusChanged,
    } as EventMemberStatusChanged;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.member_address = reader.string();
          break;
        case 2:
          message.status = reader.int32() as any;
          break;
        case 3:
          message.previous_status = reader.int32() as any;
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EventMemberStatusChanged {
    const message = {
      ...baseEventMemberStatusChanged,
    } as EventMemberStatusChanged;
    if (object.member_address !== undefined && object.member_address !== null) {
      message.member_address = String(object.member_address);
    } else {
      message.member_address = "";
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = membershipStatusFromJSON(object.status);
    } else {
      message.status = 0;
    }
    if (
      object.previous_status !== undefined &&
      object.previous_status !== null
    ) {
      message.previous_status = membershipStatusFromJSON(
        object.previous_status
      );
    } else {
      message.previous_status = 0;
    }
    return message;
  },

  toJSON(message: EventMemberStatusChanged): unknown {
    const obj: any = {};
    message.member_address !== undefined &&
      (obj.member_address = message.member_address);
    message.status !== undefined &&
      (obj.status = membershipStatusToJSON(message.status));
    message.previous_status !== undefined &&
      (obj.previous_status = membershipStatusToJSON(message.previous_status));
    return obj;
  },

  fromPartial(
    object: DeepPartial<EventMemberStatusChanged>
  ): EventMemberStatusChanged {
    const message = {
      ...baseEventMemberStatusChanged,
    } as EventMemberStatusChanged;
    if (object.member_address !== undefined && object.member_address !== null) {
      message.member_address = object.member_address;
    } else {
      message.member_address = "";
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = object.status;
    } else {
      message.status = 0;
    }
    if (
      object.previous_status !== undefined &&
      object.previous_status !== null
    ) {
      message.previous_status = object.previous_status;
    } else {
      message.previous_status = 0;
    }
    return message;
  },
};

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
