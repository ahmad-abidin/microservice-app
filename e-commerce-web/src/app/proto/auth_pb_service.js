// package: proto
// file: proto/auth.proto

var proto_auth_pb = require("../proto/auth_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var Auth = (function () {
  function Auth() {}
  Auth.serviceName = "proto.Auth";
  return Auth;
}());

Auth.Authentication = {
  methodName: "Authentication",
  service: Auth,
  requestStream: false,
  responseStream: false,
  requestType: proto_auth_pb.Credential,
  responseType: proto_auth_pb.Token
};

Auth.Authorization = {
  methodName: "Authorization",
  service: Auth,
  requestStream: false,
  responseStream: false,
  requestType: proto_auth_pb.Token,
  responseType: proto_auth_pb.Identity
};

exports.Auth = Auth;

function AuthClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

AuthClient.prototype.authentication = function authentication(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Auth.Authentication, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

AuthClient.prototype.authorization = function authorization(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Auth.Authorization, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

exports.AuthClient = AuthClient;

