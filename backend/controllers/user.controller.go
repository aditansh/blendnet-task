package controllers

import (
	"github.com/aditansh/blendnet-task/be/schemas"
	"github.com/aditansh/blendnet-task/be/services"
	"github.com/aditansh/blendnet-task/be/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func RegisterUser(c *fiber.Ctx) error {
	var payload schemas.RegisterUserSchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	errors := utils.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": false,
			"errors": errors,
		})
	}

	err := services.RegisterUser(&payload, false)
	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"status":  false,
			"message": err.Message,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  true,
		"message": "User registered successfully",
	})
}

func UpdateUser(c *fiber.Ctx) error {
	var payload schemas.UpdateUserSchema
	var ID *uuid.UUID

	temp := c.Locals("ID").(uuid.UUID)
	ID = &temp

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	errors := utils.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": false,
			"errors": errors,
		})
	}

	err := services.UpdateUserName(&payload, *ID)
	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"status":  false,
			"message": err.Message,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  true,
		"message": "User updated successfully",
	})
}

func Search(c *fiber.Ctx) error {
	var payload schemas.SearchSchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	errors := utils.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": false,
			"errors": errors,
		})
	}

	stocks, err := services.SearchStocks(&payload)
	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"status":  false,
			"message": err.Message,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": true,
		"data":   stocks,
	})
}

func UpdateUserWatchlist(c *fiber.Ctx) error {
	var payload schemas.UpdateWatchlistSchema
	var ID *uuid.UUID

	temp := c.Locals("ID").(uuid.UUID)
	ID = &temp

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	errors := utils.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": false,
			"errors": errors,
		})
	}

	data, err := services.UpdateWatchlist(&payload, *ID)
	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"status":  false,
			"message": err.Message,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  true,
		"message": "Watchlist updated successfully",
		"data":    data,
	})
}

func GetUserProfile(c *fiber.Ctx) error {
	ID := c.Locals("ID").(uuid.UUID)

	user, err := services.GetUserByID(ID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": true,
		"data":   user,
	})
}

// admin only routes

func GetAllUsers(c *fiber.Ctx) error {
	users, err := services.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": true,
		"data":   users,
	})
}
