package services

import (
	"time"

	"github.com/aditansh/blendnet-task/be/database"
	"github.com/aditansh/blendnet-task/be/models"
	"github.com/aditansh/blendnet-task/be/schemas"
	"github.com/aditansh/blendnet-task/be/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := database.DB.Find(&users)
	if result.Error != nil {
		return []models.User{}, result.Error
	}

	return users, nil
}

func GetUserByID(id uuid.UUID) (models.User, error) {
	var user models.User
	result := database.DB.Where("id = ?", id).First(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil
}

func GetUserByEmail(email string) (models.User, error) {
	var user models.User
	result := database.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil
}

func RegisterUser(payload *schemas.RegisterUserSchema, isAdmin bool) *fiber.Error {

	_, check1 := GetUserByEmail(payload.Email)
	if check1 == nil {
		return fiber.NewError(fiber.StatusConflict, "Email already exists")
	}

	hashedPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Error hashing password")
	}

	var role string
	if isAdmin {
		role = "admin"
	} else {
		role = "user"
	}

	newUser := models.User{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: hashedPassword,
		Role:     role,
	}

	result := database.DB.Create(&newUser)
	if result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError, result.Error.Error())
	}

	return nil
}

func UpdateUserName(payload *schemas.UpdateUserSchema, ID uuid.UUID) *fiber.Error {
	user, err := GetUserByID(ID)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	updates := make(map[string]interface{})
	updates["name"] = payload.Name
	updates["updated_at"] = time.Now()

	result := database.DB.Model(&user).Updates(updates)
	if result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Error updating user")
	}

	return nil
}

func UpdateUserWatchlist(payload *schemas.UpdateWatchlistSchema, ID uuid.UUID) *fiber.Error {
	user, err := GetUserByID(ID)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	updates := make(map[string]interface{})
	updates["watchlist"] = payload.Watchlist
	updates["updated_at"] = time.Now()

	result := database.DB.Model(&user).Updates(updates)
	if result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Error updating user")
	}

	return nil
}

func UpdateRole(user *models.User, role string) *fiber.Error {
	updates := make(map[string]interface{})
	updates["role"] = role
	updates["updated_at"] = time.Now()

	result := database.DB.Model(&user).Updates(updates)
	if result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Error updating user")
	}

	return nil
}
