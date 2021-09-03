package server

import (
	"context"
	"os"
	"strconv"

	pb "github.com/ebikode/peaq-challenge/challenge-3/exchange/proto/rate"
)

type exchangeServiceServer struct {
	repository Repository
	pb.UnimplementedRateServiceServer
}

func NewExchangeServiceServer(repo Repository) *exchangeServiceServer {
	return &exchangeServiceServer{repository: repo}
}

// FetchGrowths ...
func (h *exchangeServiceServer) GetGrowthRecords(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {

	numberOfCrytoMarket := os.Getenv("NUMBER_MARKET")
	numMarket, _ := strconv.ParseInt(numberOfCrytoMarket, 10, 64)

	records, err := h.repository.GetByTimestamp(req.FromTimestamp, req.ToTimestamp)

	var response pb.Response
	response.Results = UnmarshalProtoResponseData(records, int(numMarket))

	// fmt.Println("GetGrowthRecords ", response.Results)

	return &response, err
}

// FetchGrowths ...
func (h *exchangeServiceServer) GetRawGrowthRecords(ctx context.Context, req *pb.GetRequest) (*pb.RawResponse, error) {

	numberOfCrytoMarket := os.Getenv("NUMBER_MARKET")
	numMarket, _ := strconv.ParseInt(numberOfCrytoMarket, 10, 64)

	records, err := h.repository.GetByTimestamp(req.FromTimestamp, req.ToTimestamp)

	var response pb.RawResponse

	response.RawResults = UnmarshalProtoRawResponseData(records, int(numMarket))

	// fmt.Println("GetGrowthRecords ", response.RawResults)

	return &response, err
}
