package utils

import (
	"encoding/json"
	"net/http"

	"github.com/aditansh/blendnet-task/be/schemas"
	"github.com/gofiber/fiber/v2"
)

func FetchDataFromURL(url string) (*schemas.SearchStocksSchema, *fiber.Error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	defer resp.Body.Close()

	var response schemas.SearchStocksSchema
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	if len(response.BestMatches) == 0 {
		return nil, fiber.NewError(fiber.StatusNotFound, "no matches found")
	}

	return &response, nil
}
