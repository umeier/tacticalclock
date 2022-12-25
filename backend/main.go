package main

import (
	"embed"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

//go:embed frontend/*
var frontendFolder embed.FS

func main() {
	app := fiber.New(fiber.Config{
		AppName:      "TacticalClock alpha",
		ServerHeader: "TacticalClock",
	})

	app.Use("/", filesystem.New(filesystem.Config{
		Root:       http.FS(frontendFolder),
		PathPrefix: "frontend",
	}))

	// Listen from a different goroutine
	go func() {
		if err := app.Listen(":3000"); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	_ = <-c // This blocks the main thread until an interrupt is received
	fmt.Println("Gracefully shutting down...")
	_ = app.Shutdown()

	fmt.Println("Running cleanup tasks...")

	// Your cleanup tasks go here
	// db.Close()
	// redisConn.Close()
	fmt.Println("Fiber was successful shutdown.")
}
