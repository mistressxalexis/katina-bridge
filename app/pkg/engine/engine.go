package engine

import (
	"log"

	"example.com/m/app/pkg/route"

	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Engine struct {
	app *fiber.App
}

func New() *Engine {
	return &Engine{
		app: fiber.New(),
	}
}

func (e *Engine) Run() {
	e.app.Use(cors.New(cors.Config{
		AllowHeaders: "Origin, Content-Type, Accept, Content-Length, " +
			"Accept-Language, Accept-Encoding, Connection, Access-Control-Allow-Origin",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))
	e.app.Use(logger.New())
	////////////////////////
	//route.MintNFTsRoute(e.app)
	//route.DepositNFTsRoute(e.app)
	//route.ClaimNFTsRoute(e.app)
	route.LeopardRoute(e.app)
	////////////////////////
	//go ListenClaimT54NFTs()
	//go ListenClaimT55NFTs()
	//go ListenT54Bridge()
	//go ListenT55Bridge()
	//go ListenT54MintNFTs()
	//go ListenT55MintNFTs()

	err := e.app.Listen("localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
}
