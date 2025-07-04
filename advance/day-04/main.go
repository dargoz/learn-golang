package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/dargoz/day04/config"
	"github.com/dargoz/day04/data/local/db"
	"github.com/dargoz/day04/data/remote/pb"
	"github.com/dargoz/day04/service"
	_ "github.com/lib/pq" // PostgreSQL driver
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres@localhost:5432/postgres?sslmode=disable"
)

func main() {
	conf, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}
	log.Println("Config loaded successfully")

	conn, err := sql.Open(dbDriver, conf.DBSource)
	if err != nil {
		log.Fatalf("cannot connect to db: %v", err)
	}
	// conn.SetMaxOpenConns(10)
	// conn.SetMaxIdleConns(5)
	log.Println("Connected to database successfully")

	store := db.NewStore(conn)
	server := grpc.NewServer()
	pb.RegisterTransferServiceServer(server, &service.TransferServer{
		Store: store,
	})
	pb.RegisterAccountServiceServer(server, &service.AccountServer{
		Store: store,
	})
	reflection.Register(server)

	listener, err := net.Listen("tcp", conf.GRPCAddress)
	if err != nil {
		log.Fatalf("failed to create listener: %v", err)
	}
	log.Println("gRPC server is listening on", conf.GRPCAddress)

	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	log.Printf("gRPC server is running on %s\n", conf.GRPCAddress)

	defer conn.Close()
}
