package server

import (
	"strconv"
	"time"

	"github.com/ebikode/peaq-challenge/challenge-3/exchange/models"
	pb "github.com/ebikode/peaq-challenge/challenge-3/exchange/proto/rate"
	"github.com/ebikode/peaq-challenge/challenge-3/exchange/utils"
)

const (
	timeFormat = "2006-01-02 15:04:05"
)

type Repository interface {
	GetByTimestamp(int64, int64) ([]*models.GrowthRecord, error)
}

// UnmarshalProtoResponseData ...
func UnmarshalProtoResponseData(records []*models.GrowthRecord, numberOfCrytoMarket int) []*pb.ResponseData {

	responseDataList := []*pb.ResponseData{}

	for i := 0; i < len(records); i++ {

		responseMarketData := []*pb.ResponseMarketData{}

		fromDate := time.Unix(records[i].FromDate, 0).UTC()
		toDate := time.Unix(records[i].ToDate, 0).UTC()

		for n := 0; n < numberOfCrytoMarket; n++ {
			record := records[i+n]
			marketData := processMarketData(record)
			responseMarketData = append(responseMarketData, marketData)
		}
		i += (numberOfCrytoMarket - 1)

		response := &pb.ResponseData{
			From:       fromDate.Format(timeFormat),
			To:         toDate.Format(timeFormat),
			MarketData: responseMarketData,
		}

		responseDataList = append(responseDataList, response)
	}
	return responseDataList
}

func UnmarshalProtoRawResponseData(records []*models.GrowthRecord, numberOfCrytoMarket int) []*pb.GrowthRecord {

	responseGrowthDataList := []*pb.GrowthRecord{}

	for i := 0; i < len(records); i++ {

		for n := 0; n < numberOfCrytoMarket; n++ {
			record := records[i+n]
			growthRecord := processGrowthRecord(record)

			responseGrowthDataList = append(responseGrowthDataList, growthRecord)
		}
		i += (numberOfCrytoMarket - 1)

	}
	return responseGrowthDataList
}

func processMarketData(record *models.GrowthRecord) *pb.ResponseMarketData {
	marketData := &pb.ResponseMarketData{
		MarketPair: record.FromRate.MarketName,
	}

	growthData := &pb.GrowthData{}

	growthData.VolumeGrowth = utils.Round4Decimal(record.VolumeGrowth)
	growthData.HighGrowth = utils.Round4Decimal(record.HighGrowth)
	growthData.LowGrowth = utils.Round4Decimal(record.LowGrowth)
	marketData.GrowthData = growthData

	return marketData
}

func processGrowthRecord(record *models.GrowthRecord) *pb.GrowthRecord {

	growthRecord := &pb.GrowthRecord{}

	growthRecord.VolumeGrowth = utils.Round4Decimal(record.VolumeGrowth)
	growthRecord.HighGrowth = utils.Round4Decimal(record.HighGrowth)
	growthRecord.LowGrowth = utils.Round4Decimal(record.LowGrowth)
	growthRecord.To = record.ToDate
	growthRecord.From = record.FromDate
	growthRecord.ToRateId = int32(record.ToRateID)
	growthRecord.FromRateId = int32(record.FromRateID)

	growthRecord.FromRate = assignRate(record.FromRate)
	growthRecord.ToRate = assignRate(record.ToRate)
	return growthRecord
}

func assignRate(rate models.Rate) *pb.Rate {
	return &pb.Rate{
		Id:         int32(rate.ID),
		MarketName: rate.MarketName,
		Volume:     utils.Round4Decimal(rate.Volume),
		High:       utils.Round4Decimal(rate.High),
		Low:        utils.Round4Decimal(rate.Low),
		Timestamp:  strconv.Itoa(int(rate.Timestamp.Unix())),
	}
}
