package services

import (
	"time"

	"github.com/aditansh/blendnet-task/be/cache"
	"github.com/aditansh/blendnet-task/be/database"
	"github.com/aditansh/blendnet-task/be/schemas"
	"github.com/aditansh/blendnet-task/be/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func LoginEmail(payload *schemas.LoginSchema) (string, *fiber.Error) {

	user, err := GetUserByEmail(payload.Email)
	if err != nil {
		return "", fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	if !utils.CheckPasswordHash(payload.Password, user.Password) {
		return "", fiber.NewError(fiber.StatusUnauthorized, "Incorrect password")
	}

	token, err := utils.GenerateRefreshToken(user.ID, user.Email)
	if err != nil {
		return "", fiber.NewError(fiber.StatusInternalServerError, "Error generating token")
	}

	err = cache.SetValue(token, user.ID.String(), 0)
	if err != nil {
		return "", fiber.NewError(fiber.StatusInternalServerError, "Error storing token")
	}

	return token, nil
}

func UpdatePassword(payload *schemas.UpdatePasswordSchema, email string) *fiber.Error {
	user, err := GetUserByEmail(email)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	if !utils.CheckPasswordHash(payload.OldPassword, user.Password) {
		return fiber.NewError(fiber.StatusUnauthorized, "Incorrect password")
	}

	if payload.OldPassword == payload.NewPassword {
		return fiber.NewError(fiber.StatusBadRequest, "New Password cannot be the same as old password")
	}

	newPass, err := utils.HashPassword(payload.NewPassword)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Error hashing password")
	}

	updates := make(map[string]interface{})
	updates["password"] = newPass
	updates["updated_at"] = time.Now()

	result := database.DB.Model(&user).Updates(updates)
	if result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Error updating user")
	}

	return nil
}

func RefreshAccessToken(payload *schemas.RefreshTokenSchema) (string, *fiber.Error) {
	ID, err := cache.GetValue(payload.RefreshToken)
	if err != nil {
		return "", fiber.NewError(fiber.StatusUnauthorized, "Invalid refresh token")
	}

	userID, err := uuid.Parse(ID)
	if err != nil {
		return "", fiber.NewError(fiber.StatusUnauthorized, "Error parsing user id")
	}

	user, err := GetUserByID(userID)
	if err != nil {
		return "", fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	token, err := utils.GenerateAccessToken(user.ID, user.Email)
	if err != nil {
		return "", fiber.NewError(fiber.StatusInternalServerError, "Error generating token")
	}

	return token, nil
}

func Logout(payload *schemas.LogoutSchema) *fiber.Error {
	id, err := cache.GetValue(payload.RefreshToken)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid refresh token")
	}

	_, err = uuid.Parse(id)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid refresh token")
	}

	err = cache.DeleteValue(payload.RefreshToken)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid refresh token")
	}

	return nil
}

func DeleteAccount(email string) *fiber.Error {
	user, err := GetUserByEmail(email)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	result := database.DB.Delete(&user)

	if result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Error updating user")
	}

	return nil
}
