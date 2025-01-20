package handlers

import (
	"eda/app/config"
	"eda/app/entities"
	"eda/app/service"

	"github.com/gofiber/fiber/v2"
)

type dataType struct {
	Food    []entities.Food    `json:"food"`
	Product []entities.Product `json:"product"`
}

func FindFoodAndProductAndCategories(c *fiber.Ctx) error {
	u, err := config.ParseUserSession(c)
	if err != nil {
		return c.JSON(config.BaseResult(config.GetStatus("FAIL"), nil, err.Error()))
	}

	query_parse := c.Query("str")

	if query_parse == "" {
		return c.JSON(config.BaseResult(config.GetStatus("FAIL"), nil, "Empty query"))
	}

	var data dataType

	food, err := service.FoodSearch(query_parse, u.ID)
	if err != nil {
		return c.JSON(config.BaseResult(config.GetStatus("FAIL"), nil, err.Error()))
	}

	product, err := service.ProductSearch(query_parse)
	if err != nil {
		return c.JSON(config.BaseResult(config.GetStatus("FAIL"), nil, err.Error()))
	}

	data.Food = food
	data.Product = product

	return c.JSON(config.BaseResult(config.GetStatus("OK"), data))
}
