package routes

import (
	"github.com/aditansh/blendnet-task/be/controllers"
	"github.com/aditansh/blendnet-task/be/middleware"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {

	user := app.Group("/user")
	user.Post("/register", controllers.RegisterUser)

	user.Post("/update", middleware.VerifyUserToken, controllers.UpdateUser)
	user.Get("/me", middleware.VerifyUserToken, controllers.GetUserProfile)
}
