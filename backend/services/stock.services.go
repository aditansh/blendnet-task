package services

import (
	"github.com/aditansh/blendnet-task/be/schemas"
	"github.com/aditansh/blendnet-task/be/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func SearchStocks(payload *schemas.SearchSchema) (*schemas.SearchStocksSchema, *fiber.Error) {
	// var stocks []string
	apikey := viper.GetString("API_KEY")

	if apikey == "" {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "API Key not set")
	}

	url := "https://www.alphavantage.co/query?function=SYMBOL_SEARCH&keywords=" + payload.Key + "&apikey=" + apikey
	// url = "https://www.alphavantage.co/query?function=SYMBOL_SEARCH&keywords=tesco&apikey=demo"
	data, err := utils.FetchSearchDataFromURL(url)
	if err != nil {
		return nil, err
	}

	return data, nil

}

func GetStock(key string, interval string) (*schemas.GetStocksResponseSchema, *fiber.Error) {
	apikey := viper.GetString("API_KEY")

	if apikey == "" {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "API Key not set")
	}

	url := "https://www.alphavantage.co/query?function=TIME_SERIES_INTRADAY&symbol=" + key + "&interval=" + interval + "&apikey" + apikey
	// url = "https://www.alphavantage.co/query?function=TIME_SERIES_INTRADAY&symbol=IBM&interval=5min&apikey=demo"
	data, err := utils.FetchStockDataFromURL(url)
	if err != nil {
		return nil, err
	}

	return data, nil
}
