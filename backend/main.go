package main

import (
	"embed"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"log"
	"net/http"
)

//go:embed dist/*
var distFolder embed.FS

func main() {
	app := fiber.New()

	app.Use("/", filesystem.New(filesystem.Config{
		Root:       http.FS(distFolder),
		PathPrefix: "dist",
	}))

	log.Fatal(app.Listen(":3000"))
}
