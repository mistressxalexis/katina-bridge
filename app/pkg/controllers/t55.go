package controllers

import (
	"example.com/m/app/pkg/services"
	"github.com/gofiber/fiber/v2"
)

func MintT55NFTs(c *fiber.Ctx) error {
	var body struct {
		Address string `json:"address"`
		URL     string `json:"url"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := services.MintT55NFTs(body.Address, body.URL); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
	})
}

func DepositT55(c *fiber.Ctx) error {
	var body struct {
		Address string `json:"address"`
		Key     string `json:"key"`
		TokenID int64  `json:"tokenID"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	err := services.DepositT55(body.Key, body.TokenID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
	})

}

func ClaimFromT54(c *fiber.Ctx) error {
	var body struct {
		Address string `json:"address"`
		TokenId int64  `json:"tokenId"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	err := services.ClaimT55NFTs(body.Address, body.TokenId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": true,
	})
}
