package main

import (
	"binanceWebsocket/models"
	"log"
	"strconv"
)

func sumDepth(depth models.Depth) models.DepthResponse {
	depthResponse := models.DepthResponse{
		Asks: make([]models.PriceQuantity, 15),
		Bids: make([]models.PriceQuantity, 15),
	}
	for i := 0; i < 15; i++ {
		for j := 0; j < 2; j++ {
			floatAsk, err := strconv.ParseFloat(depth.Asks[i][j], 64)
			if err != nil {
				log.Fatalf("bad ask float value")
			}
			floatBid, err := strconv.ParseFloat(depth.Bids[i][j], 64)
			if err != nil {
				log.Fatalf("bad bid float value")
			}
			switch j {
			case 0:
				depthResponse.Asks[i].Price = floatAsk
				depthResponse.AskPriceSum += floatAsk
				depthResponse.Bids[i].Price = floatBid
				depthResponse.BidPriceSum += floatBid
			case 1:
				depthResponse.Asks[i].Quantity = floatAsk
				depthResponse.AskQuantSum += floatAsk
				depthResponse.Bids[i].Quantity = floatBid
				depthResponse.BidQuantSum += floatBid
			}
		}
	}
	//depthResponse.AskPriceSum = math.Round(depthResponse.AskPriceSum*1000000) / 1000000
	//depthResponse.AskQuantSum = math.Round(depthResponse.AskQuantSum*1000000) / 1000000
	//depthResponse.BidPriceSum = math.Round(depthResponse.BidPriceSum*1000000) / 1000000
	//depthResponse.BidQuantSum = math.Round(depthResponse.BidQuantSum*1000000) / 1000000
	return depthResponse
}
