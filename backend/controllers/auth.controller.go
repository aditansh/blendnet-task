package controllers

import (
	"github.com/aditansh/blendnet-task/be/schemas"
	"github.com/aditansh/blendnet-task/be/services"
	"github.com/aditansh/blendnet-task/be/utils"
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {

	var err *fiber.Error
	var token string

	var payload schemas.LoginSchema

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

	token, err = services.LoginEmail(&payload)

	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"status":  false,
			"message": err.Message,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  true,
		"message": "Login successful",
		"data":   token,
	})
}

func UpdatePassword(c *fiber.Ctx) error {
	var payload schemas.UpdatePasswordSchema

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

	email := c.Locals("Email").(string)

	err := services.UpdatePassword(&payload, email)
	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"status":  false,
			"message": err.Message,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  true,
		"message": "Password updated successfully",
	})
}

func RefreshToken(c *fiber.Ctx) error {
	var payload schemas.RefreshTokenSchema

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

	token, err := services.RefreshAccessToken(&payload)
	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"status":  false,
			"message": err.Message,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  true,
		"message": "Token refreshed successfully",
		"data":   token,
	})
}

func Logout(c *fiber.Ctx) error {
	var payload schemas.LogoutSchema

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

	err := services.Logout(&payload)
	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"status":  false,
			"message": err.Message,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  true,
		"message": "Logged out successfully",
	})
}

func DeleteAccount(c *fiber.Ctx) error {
	email := c.Locals("Email").(string)

	err := services.DeleteAccount(email)
	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"status":  false,
			"message": err.Message,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  true,
		"message": "Account deleted successfully",
	})
}
