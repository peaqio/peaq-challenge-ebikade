package jobs

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ebikode/peaq-challenge/challenge-3/exchange/models"
	"github.com/ebikode/peaq-challenge/challenge-3/exchange/pkg/growth"
	"github.com/ebikode/peaq-challenge/challenge-3/exchange/pkg/rate"
	"github.com/ebikode/peaq-challenge/challenge-3/exchange/utils"
)

const (
	timeFormat = "2006-01-02 15:04:05"
	url        = "https://bittrex.com/api/v1.1/public/getmarketsummaries"
	btcAda     = "BTC-ADA"
	ethAda     = "ETH-ADA"
)

var allowedMarket = map[string]string{
	btcAda: btcAda,
	ethAda: ethAda,
}

type marketJobHandler struct {
	RateService   rate.Service
	GrowthService growth.Service
}

func NewMarketJobHandler(rateService rate.Service, growthService growth.Service) *marketJobHandler {
	return &marketJobHandler{
		rateService,
		growthService,
	}
}

func (mjh *marketJobHandler) getMarketData() {

	fmt.Println("Get Market Data called!")

	respData := &models.MarketResponse{}

	err := utils.SendRequestAndParseResponse(http.MethodGet, url, nil, respData)

	if err != nil {
		fmt.Printf(`an error occurred while fetching market data: "%s"`, err.Error())
		return
	}

	for _, data := range respData.Result {
		market := data.MarketName
		// Screen required data
		if _, ok := allowedMarket[market]; ok {

			mjh.processMarketData(data, market)
		}
	}
}

func (mjh *marketJobHandler) processMarketData(data models.ResultPayload, market string) {
	oldRate := mjh.RateService.GetByMarketName(market)

	newRate, err := mjh.saveRate(data)

	if err != nil {
		fmt.Printf(`an error occurred while creating market rate: "%s"`, err.Error())

	}

	if oldRate != nil && err == nil {

		err := mjh.saveGrowthRate(newRate, oldRate)

		if err != nil {
			fmt.Printf(`an error occurred while creating market growth rate: "%s"`, err.Error())
		}
	}
}

func (mjh *marketJobHandler) saveRate(data models.ResultPayload) (*models.Rate, error) {
	now := time.Now()
	market := data.MarketName
	volume := data.Volume
	high := data.High
	low := data.Low

	rate := models.Rate{
		MarketName: market,
		High:       high,
		Low:        low,
		Volume:     volume,
		Timestamp:  now,
	}

	newRate, err := mjh.RateService.CreateRate(rate)
	return newRate, err
}

func (mjh *marketJobHandler) saveGrowthRate(newRate, oldRate *models.Rate) error {
	now := time.Now()

	volumeGrowth := utils.CalculatePercentageDifference(newRate.Volume, oldRate.Volume)
	highGrowth := utils.CalculatePercentageDifference(newRate.High, oldRate.High)
	lowGrowth := utils.CalculatePercentageDifference(newRate.Low, oldRate.Low)

	record := models.GrowthRecord{
		FromRateID:   oldRate.ID,
		ToRateID:     newRate.ID,
		FromDate:     oldRate.Timestamp.Unix(),
		ToDate:       now.Unix(),
		VolumeGrowth: utils.ToFloat(volumeGrowth),
		LowGrowth:    utils.ToFloat(lowGrowth),
		HighGrowth:   utils.ToFloat(highGrowth),
	}

	// b, _ := json.Marshal(record)

	// fmt.Println("======", string(b), "======")
	// fmt.Println(record.VolumeGrowth)
	// fmt.Println(record.HighGrowth)
	// fmt.Println(record.LowGrowth)

	_, err := mjh.GrowthService.CreateGrowthRecord(record)

	return err
}
