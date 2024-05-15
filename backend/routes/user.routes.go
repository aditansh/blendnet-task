package routes

import (
	"github.com/aditansh/blendnet-task/be/controllers"
	"github.com/aditansh/blendnet-task/be/middleware"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {

	user := app.Group("/user")

	user.Post("/search", middleware.VerifyUserToken, controllers.Search)
	user.Get("/stock", middleware.VerifyUserToken, controllers.GetStock)
	user.Put("/update", middleware.VerifyUserToken, controllers.UpdateUser)
	user.Patch("/watchlist", middleware.VerifyUserToken, controllers.UpdateUserWatchlist)
	user.Get("/me", middleware.VerifyUserToken, controllers.GetUserProfile)
}
