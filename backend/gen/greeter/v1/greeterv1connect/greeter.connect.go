// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: greeter/v1/greeter.proto

package greeterv1connect

import (
	v1 "backend/gen/greeter/v1"
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// GreeterServiceName is the fully-qualified name of the GreeterService service.
	GreeterServiceName = "greeter.v1.GreeterService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// GreeterServiceGreetProcedure is the fully-qualified name of the GreeterService's Greet RPC.
	GreeterServiceGreetProcedure = "/greeter.v1.GreeterService/Greet"
)

// GreeterServiceClient is a client for the greeter.v1.GreeterService service.
type GreeterServiceClient interface {
	Greet(context.Context, *connect.Request[v1.GreetRequest]) (*connect.Response[v1.GreetResponse], error)
}

// NewGreeterServiceClient constructs a client for the greeter.v1.GreeterService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewGreeterServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) GreeterServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	greeterServiceMethods := v1.File_greeter_v1_greeter_proto.Services().ByName("GreeterService").Methods()
	return &greeterServiceClient{
		greet: connect.NewClient[v1.GreetRequest, v1.GreetResponse](
			httpClient,
			baseURL+GreeterServiceGreetProcedure,
			connect.WithSchema(greeterServiceMethods.ByName("Greet")),
			connect.WithClientOptions(opts...),
		),
	}
}

// greeterServiceClient implements GreeterServiceClient.
type greeterServiceClient struct {
	greet *connect.Client[v1.GreetRequest, v1.GreetResponse]
}

// Greet calls greeter.v1.GreeterService.Greet.
func (c *greeterServiceClient) Greet(ctx context.Context, req *connect.Request[v1.GreetRequest]) (*connect.Response[v1.GreetResponse], error) {
	return c.greet.CallUnary(ctx, req)
}

// GreeterServiceHandler is an implementation of the greeter.v1.GreeterService service.
type GreeterServiceHandler interface {
	Greet(context.Context, *connect.Request[v1.GreetRequest]) (*connect.Response[v1.GreetResponse], error)
}

// NewGreeterServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewGreeterServiceHandler(svc GreeterServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	greeterServiceMethods := v1.File_greeter_v1_greeter_proto.Services().ByName("GreeterService").Methods()
	greeterServiceGreetHandler := connect.NewUnaryHandler(
		GreeterServiceGreetProcedure,
		svc.Greet,
		connect.WithSchema(greeterServiceMethods.ByName("Greet")),
		connect.WithHandlerOptions(opts...),
	)
	return "/greeter.v1.GreeterService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case GreeterServiceGreetProcedure:
			greeterServiceGreetHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedGreeterServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedGreeterServiceHandler struct{}

func (UnimplementedGreeterServiceHandler) Greet(context.Context, *connect.Request[v1.GreetRequest]) (*connect.Response[v1.GreetResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("greeter.v1.GreeterService.Greet is not implemented"))
}
