package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	openApiMdware "github.com/go-openapi/runtime/middleware"
	"github.com/rakyll/statik/fs"

	_ "github.com/ebikode/peaq-challenge/challenge-3/exchange/proto/statik"

	pb "github.com/ebikode/peaq-challenge/challenge-3/exchange/proto/rate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
)

func InitServer() error {
	serverAddr := os.Getenv("EXCHANGE_HOST")
	host := fmt.Sprintf(":%s", os.Getenv("PORT"))
	// host := ":50052"
	// serverAddr := ":50051"
	// var opts []grpc.DialOption

	fmt.Printf("gPRC Server started on %s \n\n", serverAddr)
	conn, err := SetUpGrpcClient(serverAddr)
	go ShutDownConnection(conn)

	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}
	defer conn.Close()

	ctx := context.Background()

	client := pb.NewRateServiceClient(conn)

	analyticServer := NewAnalyticsServer(client)

	fmt.Printf("Server started on %s \n\n", host)
	http.HandleFunc("/export/analytics", analyticServer.GetGrowthRecords(ctx))
	http.HandleFunc("/export/analytics/raw", analyticServer.GetRawGrowthRecords(ctx))

	// Serve the swagger-ui and swagger file
	statikFS, err := fs.New()
	if err != nil {
		return err
	}
	fileServer := http.FileServer(statikFS)

	// Added swagger setup and routes
	redocOpts := openApiMdware.RedocOpts{SpecURL: "/public/server.swagger.json"}
	redoc := openApiMdware.Redoc(redocOpts, nil)

	http.Handle("/docs", redoc)
	http.Handle("/public/", http.StripPrefix("/public", fileServer))

	http.ListenAndServe(host, nil)

	return nil

}

var (
	Params = grpc.ConnectParams{
		Backoff: backoff.Config{
			BaseDelay:  1.0 * time.Second,
			Jitter:     0.2,
			Multiplier: 1.2,
			MaxDelay:   10 * time.Second,
		},
	}
)

func SetUpGrpcClient(address string) (conn *grpc.ClientConn, err error) {
	// Set up a connection to the server.
	conn, err = grpc.DialContext(
		context.Background(),
		address,
		grpc.WithConnectParams(Params),
		grpc.WithInsecure())
	if err != nil {
		return conn, err
	}
	return conn, nil
}

// Shuts down grpc connection
func ShutDownConnection(conn *grpc.ClientConn) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			log.Println("shutting down gRPC connection...")
			// sig is a ^C, handle it
			conn.Close()
		}
	}()
}
