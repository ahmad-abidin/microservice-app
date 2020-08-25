import * as pb_1 from "google-protobuf";
import * as grpc_1 from "grpc";
export namespace proto {
    export class Identity extends pb_1.Message {
        constructor(data?: any[] | {
            username?: string;
            password?: string;
        }) {
            super();
            pb_1.Message.initialize(this, Array.isArray(data) && data, 0, -1, [], null);
            if (!Array.isArray(data) && typeof data == "object") {
                this.username = data.username;
                this.password = data.password;
            }
        }
        get username(): string | undefined {
            return pb_1.Message.getFieldWithDefault(this, 1, undefined) as string | undefined;
        }
        set username(value: string) {
            pb_1.Message.setField(this, 1, value);
        }
        get password(): string | undefined {
            return pb_1.Message.getFieldWithDefault(this, 2, undefined) as string | undefined;
        }
        set password(value: string) {
            pb_1.Message.setField(this, 2, value);
        }
        toObject() {
            return {
                username: this.username,
                password: this.password
            };
        }
        serialize(w?: pb_1.BinaryWriter): Uint8Array | undefined {
            const writer = w || new pb_1.BinaryWriter();
            if (this.username !== undefined)
                writer.writeString(1, this.username);
            if (this.password !== undefined)
                writer.writeString(2, this.password);
            if (!w)
                return writer.getResultBuffer();
        }
        serializeBinary(): Uint8Array { throw new Error("Method not implemented."); }
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): Identity {
            const reader = bytes instanceof Uint8Array ? new pb_1.BinaryReader(bytes) : bytes, message = new Identity();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    case 1:
                        message.username = reader.readString();
                        break;
                    case 2:
                        message.password = reader.readString();
                        break;
                    default: reader.skipField();
                }
            }
            return message;
        }
    }
    export class FullIdentity extends pb_1.Message {
        constructor(data?: any[] | {
            name?: string;
            email?: string;
            address?: string;
        }) {
            super();
            pb_1.Message.initialize(this, Array.isArray(data) && data, 0, -1, [], null);
            if (!Array.isArray(data) && typeof data == "object") {
                this.name = data.name;
                this.email = data.email;
                this.address = data.address;
            }
        }
        get name(): string | undefined {
            return pb_1.Message.getFieldWithDefault(this, 1, undefined) as string | undefined;
        }
        set name(value: string) {
            pb_1.Message.setField(this, 1, value);
        }
        get email(): string | undefined {
            return pb_1.Message.getFieldWithDefault(this, 2, undefined) as string | undefined;
        }
        set email(value: string) {
            pb_1.Message.setField(this, 2, value);
        }
        get address(): string | undefined {
            return pb_1.Message.getFieldWithDefault(this, 3, undefined) as string | undefined;
        }
        set address(value: string) {
            pb_1.Message.setField(this, 3, value);
        }
        toObject() {
            return {
                name: this.name,
                email: this.email,
                address: this.address
            };
        }
        serialize(w?: pb_1.BinaryWriter): Uint8Array | undefined {
            const writer = w || new pb_1.BinaryWriter();
            if (this.name !== undefined)
                writer.writeString(1, this.name);
            if (this.email !== undefined)
                writer.writeString(2, this.email);
            if (this.address !== undefined)
                writer.writeString(3, this.address);
            if (!w)
                return writer.getResultBuffer();
        }
        serializeBinary(): Uint8Array { throw new Error("Method not implemented."); }
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): FullIdentity {
            const reader = bytes instanceof Uint8Array ? new pb_1.BinaryReader(bytes) : bytes, message = new FullIdentity();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    case 1:
                        message.name = reader.readString();
                        break;
                    case 2:
                        message.email = reader.readString();
                        break;
                    case 3:
                        message.address = reader.readString();
                        break;
                    default: reader.skipField();
                }
            }
            return message;
        }
    }
    export class Credential extends pb_1.Message {
        constructor(data?: any[] | {
            token?: string;
        }) {
            super();
            pb_1.Message.initialize(this, Array.isArray(data) && data, 0, -1, [], null);
            if (!Array.isArray(data) && typeof data == "object") {
                this.token = data.token;
            }
        }
        get token(): string | undefined {
            return pb_1.Message.getFieldWithDefault(this, 1, undefined) as string | undefined;
        }
        set token(value: string) {
            pb_1.Message.setField(this, 1, value);
        }
        toObject() {
            return {
                token: this.token
            };
        }
        serialize(w?: pb_1.BinaryWriter): Uint8Array | undefined {
            const writer = w || new pb_1.BinaryWriter();
            if (this.token !== undefined)
                writer.writeString(1, this.token);
            if (!w)
                return writer.getResultBuffer();
        }
        serializeBinary(): Uint8Array { throw new Error("Method not implemented."); }
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): Credential {
            const reader = bytes instanceof Uint8Array ? new pb_1.BinaryReader(bytes) : bytes, message = new Credential();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    case 1:
                        message.token = reader.readString();
                        break;
                    default: reader.skipField();
                }
            }
            return message;
        }
    }
    export var Auth = {
        Authentication: {
            path: "/proto.Auth/Authentication",
            requestStream: false,
            responseStream: false,
            requestType: proto.Identity,
            responseType: proto.Credential,
            requestSerialize: (message: Identity) => Buffer.from(message.serialize()),
            requestDeserialize: (bytes: Buffer) => Identity.deserialize(new Uint8Array(bytes)),
            responseSerialize: (message: Credential) => Buffer.from(message.serialize()),
            responseDeserialize: (bytes: Buffer) => Credential.deserialize(new Uint8Array(bytes))
        },
        Authorization: {
            path: "/proto.Auth/Authorization",
            requestStream: false,
            responseStream: false,
            requestType: proto.Credential,
            responseType: proto.FullIdentity,
            requestSerialize: (message: Credential) => Buffer.from(message.serialize()),
            requestDeserialize: (bytes: Buffer) => Credential.deserialize(new Uint8Array(bytes)),
            responseSerialize: (message: FullIdentity) => Buffer.from(message.serialize()),
            responseDeserialize: (bytes: Buffer) => FullIdentity.deserialize(new Uint8Array(bytes))
        }
    };
    export class AuthClient extends grpc_1.makeGenericClientConstructor(Auth, "Auth", {}) {
        constructor(address: string, credentials: grpc_1.ChannelCredentials) {
            super(address, credentials)
        }
    }
}
