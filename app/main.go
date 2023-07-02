package main

import (
	"example.com/m/app/config"
	"example.com/m/app/pkg/engine"
)

func main() {
	config.Config()
	eng := engine.New()
	eng.Run()
}
