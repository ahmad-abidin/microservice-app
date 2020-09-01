// package: model
// file: model/auth.proto

import * as jspb from "google-protobuf";

export class Identity extends jspb.Message {
  getUsername(): string;
  setUsername(value: string): void;

  getPassword(): string;
  setPassword(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Identity.AsObject;
  static toObject(includeInstance: boolean, msg: Identity): Identity.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Identity, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Identity;
  static deserializeBinaryFromReader(message: Identity, reader: jspb.BinaryReader): Identity;
}

export namespace Identity {
  export type AsObject = {
    username: string,
    password: string,
  }
}

export class FullIdentity extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  getEmail(): string;
  setEmail(value: string): void;

  getAddress(): string;
  setAddress(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): FullIdentity.AsObject;
  static toObject(includeInstance: boolean, msg: FullIdentity): FullIdentity.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: FullIdentity, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): FullIdentity;
  static deserializeBinaryFromReader(message: FullIdentity, reader: jspb.BinaryReader): FullIdentity;
}

export namespace FullIdentity {
  export type AsObject = {
    name: string,
    email: string,
    address: string,
  }
}

export class Credential extends jspb.Message {
  getToken(): string;
  setToken(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Credential.AsObject;
  static toObject(includeInstance: boolean, msg: Credential): Credential.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Credential, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Credential;
  static deserializeBinaryFromReader(message: Credential, reader: jspb.BinaryReader): Credential;
}

export namespace Credential {
  export type AsObject = {
    token: string,
  }
}

