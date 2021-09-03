package server

import (
	"context"
	"net/http"

	pb "github.com/ebikode/peaq-challenge/challenge-3/exchange/proto/rate"
)

type analyticsServer struct {
	rateService pb.RateServiceClient
}

func NewAnalyticsServer(rateService pb.RateServiceClient) *analyticsServer {
	return &analyticsServer{rateService}
}

// GetGrowthRecords ...
// Fetch Growth Records
//
// Fetch all Growth record data saved in five minutes interval
func (s *analyticsServer) GetGrowthRecords(ctx context.Context) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		resp := NewResponse()

		from, to, format := PaginationParams(r)

		println(from, to)

		req := &pb.GetRequest{
			FromTimestamp: from,
			ToTimestamp:   to,
		}

		records, err := s.rateService.GetGrowthRecords(ctx, req)
		// fmt.Println(records, err)

		if format == jsonString {
			if err != nil {
				resp.Message(false, "Error Occurred while fetching Records.")
				resp.ErrorResponse(http.StatusBadRequest, w, r)
				return
			}

			resp.Message(true, "success")
			resp.AddCustomData("results", records.Results)
			resp.Respond(w, r)
			return
		}
		resp.Message(false, "Bad Format Supplied")
		resp.ErrorResponse(http.StatusBadRequest, w, r)
	}

}

// GetRawGrowthRecords ...
// Fetch Raw Growth Records
//
// Fetch all Raw Growth record data saved in five minutes interval
func (s *analyticsServer) GetRawGrowthRecords(ctx context.Context) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		resp := NewResponse()

		from, to, format := PaginationParams(r)

		println(from, to)

		req := &pb.GetRequest{
			FromTimestamp: from,
			ToTimestamp:   to,
		}

		records, err := s.rateService.GetRawGrowthRecords(ctx, req)
		// fmt.Println(records, err)

		if format == jsonString {
			if err != nil {
				resp.Message(false, "Error Occurred while fetching Records.")
				resp.ErrorResponse(http.StatusBadRequest, w, r)
				return
			}

			resp.Message(true, "success")
			resp.AddCustomData("raw_results", records.RawResults)
			resp.Respond(w, r)
			return
		}
		resp.Message(false, "Bad Format Supplied")
		resp.ErrorResponse(http.StatusBadRequest, w, r)
	}

}
