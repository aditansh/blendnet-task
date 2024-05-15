package routes

import (
	"github.com/aditansh/blendnet-task/be/controllers"
	"github.com/aditansh/blendnet-task/be/middleware"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {

	auth := app.Group("/auth")

	auth.Post("/register", controllers.RegisterUser)
	auth.Post("/login", controllers.Login)
	auth.Post("/refresh", controllers.RefreshToken)

	auth.Put("/updatepassword", middleware.VerifyToken, controllers.UpdatePassword)
	auth.Post("/logout", middleware.VerifyToken, controllers.Logout)
	auth.Delete("/delete", middleware.VerifyToken, controllers.DeleteAccount)
}
