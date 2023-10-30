package gapi

import (
	"context"
	"fmt"
	"time"

	db "github.com/TheDP66/simple_bank_go/db/sqlc"
	"github.com/TheDP66/simple_bank_go/pb"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) RenewAccess(ctx context.Context, req *pb.RenewAccessRequest) (*pb.RenewAccessResponse, error) {
	var violations []*errdetails.BadRequest_FieldViolation

	refreshPayload, err := server.tokenMaker.VerifyToken(req.RefreshToken)
	if err != nil {
		return nil, invalidArgumentError(append(violations, fieldViolation("refresh_token", err)))
	}

	session, err := server.store.GetSession(ctx, refreshPayload.ID)
	if err != nil {
		if db.ErrorCode(err) == db.ErrRecordNotFound.Error() {
			return nil, status.Errorf(codes.NotFound, "%s", err)
		}

		return nil, status.Errorf(codes.Internal, "%s", err)
	}

	if session.IsBlocked {
		err := fmt.Errorf("blocked session")
		return nil, status.Errorf(codes.Unauthenticated, "%s", err)
	}

	if session.Username != refreshPayload.Username {
		err := fmt.Errorf("incorect session user")
		return nil, status.Errorf(codes.Unauthenticated, "%s", err)
	}

	if session.RefreshToken != req.RefreshToken {
		err := fmt.Errorf("mismatched session token")
		return nil, status.Errorf(codes.Unauthenticated, "%s", err)
	}

	if time.Now().After(session.ExpiresAt) {
		err := fmt.Errorf("expired session")
		return nil, status.Errorf(codes.Unauthenticated, "%s", err)
	}

	accessToken, accessPayload, err := server.tokenMaker.CreateToken(refreshPayload.Username, refreshPayload.Role, server.config.AccessTokenDuration)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%s", err)
	}

	rsp := &pb.RenewAccessResponse{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: timestamppb.New(accessPayload.ExpiredAt),
	}
	return rsp, nil
}
