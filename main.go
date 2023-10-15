package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/ss497254/keyboard-auto-typer/sendkeys"

	"github.com/gofiber/fiber/v2"
)

func parseArgs() map[string]string {
	args := map[string]string{
		"port": "51212",
		"host": "localhost",
	}

	for i, arg := range os.Args {
		if arg == "--host" {
			args["host"] = os.Args[i+1]
		} else if arg == "--port" {
			args["port"] = os.Args[i+1]
		}
	}

	return args
}

func main() {
	args := parseArgs()
	isRunning := false
	serverAddress := fmt.Sprintf("%s:%s", args["host"], args["port"])

	app := fiber.New(fiber.Config{
		AppName: "super-duper",
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"success": true,
			"message": "Hi",
		})
	})

	app.Post("/send", func(c *fiber.Ctx) error {
		type Data struct {
			Text       []string      `json:"text"`
			StartAfter time.Duration `json:"startAfter"`
			Delay      time.Duration `json:"delay"`
		}

		data := &Data{
			StartAfter: 5000,
			Delay:      10,
		}

		if err := c.BodyParser(data); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"success": false,
				"message": "invalid data received",
				"error":   err.Error(),
			})
		}

		if isRunning {
			return c.Status(400).JSON(fiber.Map{
				"success": false,
				"message": "data is being sent, wait",
				"error":   "INUSE",
			})
		}

		if data.StartAfter > 0 {
			time.Sleep(data.StartAfter * time.Millisecond)
		}

		go func() {
			isRunning = true
			err := sendkeys.SendStrings(data.Text, data.Delay*time.Millisecond)

			if err != nil {
				fmt.Println("-- Unable to send --", err)
			}
			isRunning = false
		}()

		return c.Status(200).JSON(fiber.Map{
			"success": true,
			"message": "text send successfully",
		})

	})

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		<-c

		fmt.Println("Gracefully shutting down")
		app.Shutdown()
	}()

	log.Fatal(app.Listen(serverAddress))
}
