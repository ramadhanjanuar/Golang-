package main

import (
	"context"
	"log"

	"google.golang.org/grpc/metadata"

	"google.golang.org/grpc"
)

// AuthInterceptor struct
type AuthInterceptor struct {
	accessToken string
}

// NewAuthInterceptor returns new authinterceptor
func NewAuthInterceptor(accessToken string) *AuthInterceptor {
	return &AuthInterceptor{accessToken}
}

// Unary unary
func (interceptor *AuthInterceptor) Unary() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		log.Printf("--> unary interceptor: %s", method)
		if len(interceptor.accessToken) >= 0 {
			return invoker(interceptor.attachToken(ctx), method, req, reply, cc, opts...)
		}
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func (interceptor *AuthInterceptor) attachToken(ctx context.Context) context.Context {
	return metadata.AppendToOutgoingContext(ctx, "authorization", interceptor.accessToken)
}
