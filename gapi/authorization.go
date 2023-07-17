package gapi

import (
	"context"
	"fmt"
	"simple_bank/token"
	"strings"

	"google.golang.org/grpc/metadata"
)

const (
	authHeaderKey = "authorization"
)

func (server *Server) authorizeUser(ctx context.Context) (*token.Payload, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("metadata is not provided")
	}
	values := md.Get(authHeaderKey)
	if len(values) == 0 {
		return nil, fmt.Errorf("authorization token is not provided")
	}

	authHeader := values[0]
	fileds := strings.Fields(authHeader)
	if len(fileds) < 2 {
		return nil, fmt.Errorf("invalid authorization token format")
	}
	authType := strings.ToLower(fileds[0])
	if authType != "bearer" {
		return nil, fmt.Errorf("unsupported authorization type %s", authType)
	}

	accessToken := fileds[1]
	payload, err := server.tokenMaker.VerifyToken(accessToken)
	if err != nil {
		return nil, fmt.Errorf("invalid access token: %w", err)
	}

	return payload, nil
}
