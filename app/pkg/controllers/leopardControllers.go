package controllers

import (
	"example.com/m/app/pkg/services"
	"github.com/gofiber/fiber/v2"
)

func LeopardMintNFT(c *fiber.Ctx) error {
	var body struct {
		Address string `json:"address"`
		URL     string `json:"url"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	services.LeopardMintNFT(body.Address, body.URL)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
	})
}

func DepositToLeopardBridge(c *fiber.Ctx) error {
	return nil
}

func ClaimFromLeopardsBridge(c *fiber.Ctx) error {
	return nil
}
