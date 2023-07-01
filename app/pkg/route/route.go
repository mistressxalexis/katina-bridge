package route

import (
	controllers2 "example.com/m/app/pkg/controllers"
	"github.com/gofiber/fiber/v2"
)

func DepositNFTsRoute(r *fiber.App) {
	r.Put("/deposit/t55", controllers2.DepositT55)
	r.Put("/deposit/t54", controllers2.DepositT54)
}

func MintNFTsRoute(r *fiber.App) {
	r.Put("/mint/t54", controllers2.MintT54NFTs)
	r.Put("/mint/t55", controllers2.MintT55NFTs)
}

func ClaimNFTsRoute(r *fiber.App) {
	r.Put("/claim/t55", controllers2.ClaimFromT54)
	r.Put("/claim/t54", controllers2.ClaimFromT55)
}
