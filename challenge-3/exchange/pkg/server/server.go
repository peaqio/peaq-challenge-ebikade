package server

// import (
// 	"fmt"
// 	"math/big"
// )

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/ebikode/peaq-challenge/challenge-3/exchange/jobs"
	"github.com/ebikode/peaq-challenge/challenge-3/exchange/pkg/growth"
	pb "github.com/ebikode/peaq-challenge/challenge-3/exchange/proto/rate"
	storage "github.com/ebikode/peaq-challenge/challenge-3/exchange/storage/mysql"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func InitServer() error {
	address := os.Getenv("SERVER_ADDRESS")
	// address := ":50051"

	// initialize DB
	dbConfig := storage.New()
	mdb, err := dbConfig.InitDB()

	// if an error occurred while initialising db
	if err != nil {
		return err
	}

	growthStorage := storage.NewGrowthRecordStorage(mdb)

	repo := growth.NewService(growthStorage)

	service := NewExchangeServiceServer(repo)

	// Set-up our gRPC server.
	s := grpc.NewServer()

	// Register service with the gRPC server
	pb.RegisterRateServiceServer(s, service)

	// Register reflection service on gRPC server.
	reflection.Register(s)

	// Init Jobs
	jobs.Init(mdb)

	lis, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	err = s.Serve(lis)
	if err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}
	log.Println("Running on :", address)

	return nil

}
