package gapi

import (
	"context"

	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

const (
	grpcGatewayUserAgentHeader = "grpcgateway-user-agent"
	userAgentHeader            = "user-agent"
	xForwardedForHeader        = "x-forwarded-for"
)

type Metadata struct {
	UserAgent string
	ClientIp  string
}

func (s *Server) extractMetadata(ctx context.Context) *Metadata {
	mtdt := &Metadata{}

	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if ua := md.Get(grpcGatewayUserAgentHeader); len(ua) > 0 {
			mtdt.UserAgent = ua[0]
		}
		if ua := md.Get(userAgentHeader); len(ua) > 0 {
			mtdt.UserAgent = ua[0]
		}
		if ip := md.Get(xForwardedForHeader); len(ip) > 0 {
			mtdt.ClientIp = ip[0]
		}
	}

	if p, ok := peer.FromContext(ctx); ok {
		if p.Addr != nil {
			mtdt.ClientIp = p.Addr.String()
		}
	}

	return mtdt
}
