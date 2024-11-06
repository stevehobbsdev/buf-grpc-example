package interceptors

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
)

func NewLoggerInterceptor() connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			fmt.Printf("Request received: %s %s %s\n", req.HTTPMethod(), req.Peer().Addr, req.Peer().Query)
			return next(ctx, req)
		})
	}

	return connect.UnaryInterceptorFunc(interceptor)
}
