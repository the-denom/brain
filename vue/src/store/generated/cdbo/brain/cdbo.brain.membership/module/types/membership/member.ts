/* eslint-disable */
import { BaseAccount } from "../cosmos/auth/v1beta1/auth";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "cdbo.brain.membership";

/**
 * MembershipStatus enumerates the valid membership states for a citizen of The
 * Denom
 */
export enum MembershipStatus {
  /** MEMBERSHIP_STATUS_UNDEFINED - MEMBERSHIP_STATUS_UNDEFINED defines a no-op status */
  MEMBERSHIP_STATUS_UNDEFINED = 0,
  /** MEMBERSHIP_STATUS_ELECTORATE - MEMBERSHIP_STATUS_ELECTORATE defines this member as being an active citizen */
  MEMBERSHIP_STATUS_ELECTORATE = 1,
  /** MEMBERSHIP_STATUS_INACTIVE - MEMBERSHIP_STATUS_INACTIVE defines this member as being an inactive citizen */
  MEMBERSHIP_STATUS_INACTIVE = 2,
  /** MEMBERSHIP_STATUS_RECALLED - MEMBERSHIP_STATUS_RECALLED defines this member as being recalled */
  MEMBERSHIP_STATUS_RECALLED = 3,
  /** MEMBERSHIP_STATUS_EXPULSED - MEMBERSHIP_STATUS_EXPULSED defines this member as being expulsed */
  MEMBERSHIP_STATUS_EXPULSED = 4,
  UNRECOGNIZED = -1,
}

export function membershipStatusFromJSON(object: any): MembershipStatus {
  switch (object) {
    case 0:
    case "MEMBERSHIP_STATUS_UNDEFINED":
      return MembershipStatus.MEMBERSHIP_STATUS_UNDEFINED;
    case 1:
    case "MEMBERSHIP_STATUS_ELECTORATE":
      return MembershipStatus.MEMBERSHIP_STATUS_ELECTORATE;
    case 2:
    case "MEMBERSHIP_STATUS_INACTIVE":
      return MembershipStatus.MEMBERSHIP_STATUS_INACTIVE;
    case 3:
    case "MEMBERSHIP_STATUS_RECALLED":
      return MembershipStatus.MEMBERSHIP_STATUS_RECALLED;
    case 4:
    case "MEMBERSHIP_STATUS_EXPULSED":
      return MembershipStatus.MEMBERSHIP_STATUS_EXPULSED;
    case -1:
    case "UNRECOGNIZED":
    default:
      return MembershipStatus.UNRECOGNIZED;
  }
}

export function membershipStatusToJSON(object: MembershipStatus): string {
  switch (object) {
    case MembershipStatus.MEMBERSHIP_STATUS_UNDEFINED:
      return "MEMBERSHIP_STATUS_UNDEFINED";
    case MembershipStatus.MEMBERSHIP_STATUS_ELECTORATE:
      return "MEMBERSHIP_STATUS_ELECTORATE";
    case MembershipStatus.MEMBERSHIP_STATUS_INACTIVE:
      return "MEMBERSHIP_STATUS_INACTIVE";
    case MembershipStatus.MEMBERSHIP_STATUS_RECALLED:
      return "MEMBERSHIP_STATUS_RECALLED";
    case MembershipStatus.MEMBERSHIP_STATUS_EXPULSED:
      return "MEMBERSHIP_STATUS_EXPULSED";
    default:
      return "UNKNOWN";
  }
}

/**
 * Member is a specialisation of BaseAccount that adds Member Status and
 * Nickname
 */
export interface Member {
  base_account: BaseAccount | undefined;
  status: MembershipStatus;
  nickname: string;
}

const baseMember: object = { status: 0, nickname: "" };

export const Member = {
  encode(message: Member, writer: Writer = Writer.create()): Writer {
    if (message.base_account !== undefined) {
      BaseAccount.encode(
        message.base_account,
        writer.uint32(10).fork()
      ).ldelim();
    }
    if (message.status !== 0) {
      writer.uint32(16).int32(message.status);
    }
    if (message.nickname !== "") {
      writer.uint32(26).string(message.nickname);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Member {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMember } as Member;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.base_account = BaseAccount.decode(reader, reader.uint32());
          break;
        case 2:
          message.status = reader.int32() as any;
          break;
        case 3:
          message.nickname = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Member {
    const message = { ...baseMember } as Member;
    if (object.base_account !== undefined && object.base_account !== null) {
      message.base_account = BaseAccount.fromJSON(object.base_account);
    } else {
      message.base_account = undefined;
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = membershipStatusFromJSON(object.status);
    } else {
      message.status = 0;
    }
    if (object.nickname !== undefined && object.nickname !== null) {
      message.nickname = String(object.nickname);
    } else {
      message.nickname = "";
    }
    return message;
  },

  toJSON(message: Member): unknown {
    const obj: any = {};
    message.base_account !== undefined &&
      (obj.base_account = message.base_account
        ? BaseAccount.toJSON(message.base_account)
        : undefined);
    message.status !== undefined &&
      (obj.status = membershipStatusToJSON(message.status));
    message.nickname !== undefined && (obj.nickname = message.nickname);
    return obj;
  },

  fromPartial(object: DeepPartial<Member>): Member {
    const message = { ...baseMember } as Member;
    if (object.base_account !== undefined && object.base_account !== null) {
      message.base_account = BaseAccount.fromPartial(object.base_account);
    } else {
      message.base_account = undefined;
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = object.status;
    } else {
      message.status = 0;
    }
    if (object.nickname !== undefined && object.nickname !== null) {
      message.nickname = object.nickname;
    } else {
      message.nickname = "";
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
