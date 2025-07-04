package service

import (
	"context"

	"github.com/dargoz/day04/data/local/db"
	"github.com/dargoz/day04/data/remote/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AccountServer struct {
	pb.UnimplementedAccountServiceServer
	Store db.Store
}

func (s *AccountServer) CreateAccount(ctx context.Context, req *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error) {
	// Validate request
	if req.Owner == "" || req.Balance <= 0 || req.Currency == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid account creation request")
	}

	// Create account
	account, err := s.Store.CreateAccountTx(ctx, db.CreateAccountTxParams{
		Owner:    req.Owner,
		Balance:  req.Balance,
		Currency: req.Currency,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create account: %v", err)
	}

	return &pb.CreateAccountResponse{
		Id:        account.Account.ID,
		Owner:     account.Account.Owner,
		Balance:   account.Account.Balance,
		Currency:  account.Account.Currency,
		CreatedAt: timestamppb.Now(),
	}, nil
}
