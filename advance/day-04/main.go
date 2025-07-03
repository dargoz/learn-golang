package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/dargoz/day04/config"
	"github.com/dargoz/day04/data/local/db"
	"github.com/dargoz/day04/data/remote/pb"
	"github.com/dargoz/day04/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	conf, err := config.LoadConfig(".env")
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}

	conn, err := sql.Open("postgres", conf.DBSource)
	if err != nil {
		log.Fatalf("cannot connect to db: %v", err)
	}
	conn.SetMaxOpenConns(10)
	conn.SetMaxIdleConns(5)

	store := db.NewStore(conn)
	server := grpc.NewServer()
	pb.RegisterTransferServiceServer(server, &service.TransferServer{
		Store: store,
	})
	reflection.Register(server)

	listener, err := net.Listen("tcp", conf.GRPCAddress)
	if err != nil {
		log.Fatalf("failed to create listener: %v", err)
	}

	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	log.Printf("gRPC server is running on %s", conf.GRPCAddress)

	defer conn.Close()
}
