package gapi

import (
	"fmt"

	db "github.com/TheDP66/simple_bank_go/db/sqlc"
	"github.com/TheDP66/simple_bank_go/pb"
	"github.com/TheDP66/simple_bank_go/token"
	"github.com/TheDP66/simple_bank_go/util"
)

// Server serves gRPC requests for our banking service
type Server struct {
	pb.UnimplementedSimpleBankServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

// NewServer creates a new gRPC server
func NewServer(config util.Config, store db.Store) (*Server, error) {
	// ? Create Token using JWT
	// tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)

	// ? Create Token using Paseto (recommend)
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
