// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: account.proto

/*
Package grpc is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package grpc

import (
	"context"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = metadata.Join

func request_LendingBooksService_AccountInfo_0(ctx context.Context, marshaler runtime.Marshaler, client LendingBooksServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq Account
	var metadata runtime.ServerMetadata

	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.AccountInfo(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_LendingBooksService_AccountInfo_0(ctx context.Context, marshaler runtime.Marshaler, server LendingBooksServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq Account
	var metadata runtime.ServerMetadata

	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.AccountInfo(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterLendingBooksServiceHandlerServer registers the http handlers for service LendingBooksService to "mux".
// UnaryRPC     :call LendingBooksServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterLendingBooksServiceHandlerFromEndpoint instead.
// GRPC interceptors will not work for this type of registration. To use interceptors, you must use the "runtime.WithMiddlewares" option in the "runtime.NewServeMux" call.
func RegisterLendingBooksServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server LendingBooksServiceServer) error {

	mux.Handle("POST", pattern_LendingBooksService_AccountInfo_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/myapp.LendingBooksService/AccountInfo", runtime.WithHTTPPathPattern("/v1/accountinfo/books"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_LendingBooksService_AccountInfo_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_LendingBooksService_AccountInfo_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterLendingBooksServiceHandlerFromEndpoint is same as RegisterLendingBooksServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterLendingBooksServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.NewClient(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterLendingBooksServiceHandler(ctx, mux, conn)
}

// RegisterLendingBooksServiceHandler registers the http handlers for service LendingBooksService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterLendingBooksServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterLendingBooksServiceHandlerClient(ctx, mux, NewLendingBooksServiceClient(conn))
}

// RegisterLendingBooksServiceHandlerClient registers the http handlers for service LendingBooksService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "LendingBooksServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "LendingBooksServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "LendingBooksServiceClient" to call the correct interceptors. This client ignores the HTTP middlewares.
func RegisterLendingBooksServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client LendingBooksServiceClient) error {

	mux.Handle("POST", pattern_LendingBooksService_AccountInfo_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/myapp.LendingBooksService/AccountInfo", runtime.WithHTTPPathPattern("/v1/accountinfo/books"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_LendingBooksService_AccountInfo_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_LendingBooksService_AccountInfo_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_LendingBooksService_AccountInfo_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v1", "accountinfo", "books"}, ""))
)

var (
	forward_LendingBooksService_AccountInfo_0 = runtime.ForwardResponseMessage
)