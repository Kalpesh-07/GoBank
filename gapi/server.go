package gapi

import (
	"fmt"

	db "github.com/Kalpesh-07/GoBank/db/sqlc"
	"github.com/Kalpesh-07/GoBank/pb"
	"github.com/Kalpesh-07/GoBank/token"
	"github.com/Kalpesh-07/GoBank/util"
	//"github.com/Kalpesh-07/GoBank//worker"
)

// Server serves gRPC requests for our banking service.
type Server struct {
	pb.UnimplementedSimpleBankServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	//	taskDistributor worker.TaskDistributor
}

// NewServer creates a new gRPC server.  taskDistributor worker.TaskDistributor
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
		//taskDistributor: taskDistributor,
	}

	return server, nil
}
