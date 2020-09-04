// package: proto
// file: proto/auth.proto

import * as proto_auth_pb from "../proto/auth_pb";
import {grpc} from "@improbable-eng/grpc-web";

type AuthAuthentication = {
  readonly methodName: string;
  readonly service: typeof Auth;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_auth_pb.Credential;
  readonly responseType: typeof proto_auth_pb.Token;
};

type AuthAuthorization = {
  readonly methodName: string;
  readonly service: typeof Auth;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_auth_pb.Token;
  readonly responseType: typeof proto_auth_pb.Identity;
};

export class Auth {
  static readonly serviceName: string;
  static readonly Authentication: AuthAuthentication;
  static readonly Authorization: AuthAuthorization;
}

export type ServiceError = { message: string, code: number; metadata: grpc.Metadata }
export type Status = { details: string, code: number; metadata: grpc.Metadata }

interface UnaryResponse {
  cancel(): void;
}
interface ResponseStream<T> {
  cancel(): void;
  on(type: 'data', handler: (message: T) => void): ResponseStream<T>;
  on(type: 'end', handler: (status?: Status) => void): ResponseStream<T>;
  on(type: 'status', handler: (status: Status) => void): ResponseStream<T>;
}
interface RequestStream<T> {
  write(message: T): RequestStream<T>;
  end(): void;
  cancel(): void;
  on(type: 'end', handler: (status?: Status) => void): RequestStream<T>;
  on(type: 'status', handler: (status: Status) => void): RequestStream<T>;
}
interface BidirectionalStream<ReqT, ResT> {
  write(message: ReqT): BidirectionalStream<ReqT, ResT>;
  end(): void;
  cancel(): void;
  on(type: 'data', handler: (message: ResT) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'end', handler: (status?: Status) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'status', handler: (status: Status) => void): BidirectionalStream<ReqT, ResT>;
}

export class AuthClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  authentication(
    requestMessage: proto_auth_pb.Credential,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_auth_pb.Token|null) => void
  ): UnaryResponse;
  authentication(
    requestMessage: proto_auth_pb.Credential,
    callback: (error: ServiceError|null, responseMessage: proto_auth_pb.Token|null) => void
  ): UnaryResponse;
  authorization(
    requestMessage: proto_auth_pb.Token,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_auth_pb.Identity|null) => void
  ): UnaryResponse;
  authorization(
    requestMessage: proto_auth_pb.Token,
    callback: (error: ServiceError|null, responseMessage: proto_auth_pb.Identity|null) => void
  ): UnaryResponse;
}

