/**
 * @fileoverview gRPC-Web generated client stub for service
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.service = require('./req_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.service.ServerServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.service.ServerServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.service.AddReq,
 *   !proto.service.Confirmation>}
 */
const methodDescriptor_ServerService_AddCustomer = new grpc.web.MethodDescriptor(
  '/service.ServerService/AddCustomer',
  grpc.web.MethodType.UNARY,
  proto.service.AddReq,
  proto.service.Confirmation,
  /**
   * @param {!proto.service.AddReq} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.service.Confirmation.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.service.AddReq,
 *   !proto.service.Confirmation>}
 */
const methodInfo_ServerService_AddCustomer = new grpc.web.AbstractClientBase.MethodInfo(
  proto.service.Confirmation,
  /**
   * @param {!proto.service.AddReq} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.service.Confirmation.deserializeBinary
);


/**
 * @param {!proto.service.AddReq} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.service.Confirmation)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.service.Confirmation>|undefined}
 *     The XHR Node Readable Stream
 */
proto.service.ServerServiceClient.prototype.addCustomer =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/service.ServerService/AddCustomer',
      request,
      metadata || {},
      methodDescriptor_ServerService_AddCustomer,
      callback);
};


/**
 * @param {!proto.service.AddReq} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.service.Confirmation>}
 *     A native promise that resolves to the response
 */
proto.service.ServerServicePromiseClient.prototype.addCustomer =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/service.ServerService/AddCustomer',
      request,
      metadata || {},
      methodDescriptor_ServerService_AddCustomer);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.service.GetReq,
 *   !proto.service.Customer>}
 */
const methodDescriptor_ServerService_RetrieveCustomer = new grpc.web.MethodDescriptor(
  '/service.ServerService/RetrieveCustomer',
  grpc.web.MethodType.UNARY,
  proto.service.GetReq,
  proto.service.Customer,
  /**
   * @param {!proto.service.GetReq} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.service.Customer.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.service.GetReq,
 *   !proto.service.Customer>}
 */
const methodInfo_ServerService_RetrieveCustomer = new grpc.web.AbstractClientBase.MethodInfo(
  proto.service.Customer,
  /**
   * @param {!proto.service.GetReq} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.service.Customer.deserializeBinary
);


/**
 * @param {!proto.service.GetReq} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.service.Customer)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.service.Customer>|undefined}
 *     The XHR Node Readable Stream
 */
proto.service.ServerServiceClient.prototype.retrieveCustomer =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/service.ServerService/RetrieveCustomer',
      request,
      metadata || {},
      methodDescriptor_ServerService_RetrieveCustomer,
      callback);
};


/**
 * @param {!proto.service.GetReq} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.service.Customer>}
 *     A native promise that resolves to the response
 */
proto.service.ServerServicePromiseClient.prototype.retrieveCustomer =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/service.ServerService/RetrieveCustomer',
      request,
      metadata || {},
      methodDescriptor_ServerService_RetrieveCustomer);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.service.DeleteReq,
 *   !proto.service.Confirmation>}
 */
const methodDescriptor_ServerService_DeleteCustomer = new grpc.web.MethodDescriptor(
  '/service.ServerService/DeleteCustomer',
  grpc.web.MethodType.UNARY,
  proto.service.DeleteReq,
  proto.service.Confirmation,
  /**
   * @param {!proto.service.DeleteReq} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.service.Confirmation.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.service.DeleteReq,
 *   !proto.service.Confirmation>}
 */
const methodInfo_ServerService_DeleteCustomer = new grpc.web.AbstractClientBase.MethodInfo(
  proto.service.Confirmation,
  /**
   * @param {!proto.service.DeleteReq} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.service.Confirmation.deserializeBinary
);


/**
 * @param {!proto.service.DeleteReq} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.service.Confirmation)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.service.Confirmation>|undefined}
 *     The XHR Node Readable Stream
 */
proto.service.ServerServiceClient.prototype.deleteCustomer =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/service.ServerService/DeleteCustomer',
      request,
      metadata || {},
      methodDescriptor_ServerService_DeleteCustomer,
      callback);
};


/**
 * @param {!proto.service.DeleteReq} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.service.Confirmation>}
 *     A native promise that resolves to the response
 */
proto.service.ServerServicePromiseClient.prototype.deleteCustomer =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/service.ServerService/DeleteCustomer',
      request,
      metadata || {},
      methodDescriptor_ServerService_DeleteCustomer);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.service.GetAllReq,
 *   !proto.service.Customers>}
 */
const methodDescriptor_ServerService_GetAllCustomers = new grpc.web.MethodDescriptor(
  '/service.ServerService/GetAllCustomers',
  grpc.web.MethodType.UNARY,
  proto.service.GetAllReq,
  proto.service.Customers,
  /**
   * @param {!proto.service.GetAllReq} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.service.Customers.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.service.GetAllReq,
 *   !proto.service.Customers>}
 */
const methodInfo_ServerService_GetAllCustomers = new grpc.web.AbstractClientBase.MethodInfo(
  proto.service.Customers,
  /**
   * @param {!proto.service.GetAllReq} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.service.Customers.deserializeBinary
);


/**
 * @param {!proto.service.GetAllReq} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.service.Customers)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.service.Customers>|undefined}
 *     The XHR Node Readable Stream
 */
proto.service.ServerServiceClient.prototype.getAllCustomers =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/service.ServerService/GetAllCustomers',
      request,
      metadata || {},
      methodDescriptor_ServerService_GetAllCustomers,
      callback);
};


/**
 * @param {!proto.service.GetAllReq} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.service.Customers>}
 *     A native promise that resolves to the response
 */
proto.service.ServerServicePromiseClient.prototype.getAllCustomers =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/service.ServerService/GetAllCustomers',
      request,
      metadata || {},
      methodDescriptor_ServerService_GetAllCustomers);
};


module.exports = proto.service;

