package server

import (
	"context"
	auth "github.com/1hihik1/forum-auth/pkg/api/g_rpc"
	jwt "github.com/1hihik1/forum-auth/pkg/auth"
)

type AuthServer struct {
	auth.UnimplementedAuthServiceServer
}

func (s *AuthServer) ValidateToken(ctx context.Context, req *auth.TokenRequest) (*auth.TokenResponse, error) {
	_, err := jwt.ValidateToken(req.Token)
	if err != nil {
		return &auth.TokenResponse{
			Valid: false,
			Error: err.Error(),
		}, nil
	}
	return &auth.TokenResponse{Valid: true}, nil
}

func (s *AuthServer) GetUserID(ctx context.Context, req *auth.TokenRequest) (*auth.UserIDResponse, error) {
	claims, err := jwt.ValidateToken(req.Token)
	if err != nil {
		return &auth.UserIDResponse{
			Error: err.Error(),
		}, nil
	}
	return &auth.UserIDResponse{UserId: int32(claims.UserID)}, nil
}
