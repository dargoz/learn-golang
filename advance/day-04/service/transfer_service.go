package service

import (
	"context"

	"github.com/dargoz/day04/data/local/db"
	"github.com/dargoz/day04/data/remote/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type TransferServer struct {
	pb.UnimplementedTransferServiceServer
	Store db.Store
}

func (s *TransferServer) CreateTransfer(ctx context.Context, req *pb.TransferRequest) (*pb.TransferResponse, error) {
	// Validate request
	if req.FromAccountId == 0 || req.ToAccountId == 0 || req.Amount <= 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid transfer request")
	}

	// Create transfer
	transfer, err := s.Store.TransferTx(ctx, db.TransferTxParams{
		FromAccountID: req.FromAccountId,
		ToAccountID:   req.ToAccountId,
		Amount:        req.Amount,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create transfer: %v", err)
	}

	return &pb.TransferResponse{
		Transfer: &pb.Transfer{
			Id:            transfer.Transfer.ID,
			FromAccountId: transfer.Transfer.FromAccountID,
			ToAccountId:   transfer.Transfer.ToAccountID,
			Amount:        transfer.Transfer.Amount,
			CreatedAt:     timestamppb.Now(),
		},
		Message: "Transfer successful",
	}, nil
}

func (s *TransferServer) GetTransferByID(ctx context.Context, req *pb.GetTransferByIDRequest) (*pb.Transfer, error) {
	if req.Id == 0 {
		return nil, status.Error(codes.InvalidArgument, "transfer ID is required")
	}

	transferTx, err := s.Store.GetTransferByID(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get transfer: %v", err)
	}

	return &pb.Transfer{
		Id:            transferTx.Transfer.ID,
		FromAccountId: transferTx.Transfer.FromAccountID,
		ToAccountId:   transferTx.Transfer.FromAccountID,
		CreatedAt:     timestamppb.Now(),
	}, nil
}
