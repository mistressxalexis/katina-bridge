package route

import (
	"example.com/m/app/pkg/controllers"
	"github.com/gofiber/fiber/v2"
)

func LeopardRoute(r *fiber.App) {
	sen := r.Group("/leopard")
	sen.Post("/mint", controllers.LeopardMintNFT)
}
