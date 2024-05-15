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
	data, err := utils.FetchDataFromURL(url)
	if err != nil {
		return nil, err
	}

	return data, nil

}
