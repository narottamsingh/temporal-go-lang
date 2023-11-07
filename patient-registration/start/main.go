package main

import (
	"fmt"
	"patient-registration/app"
	"patient-registration/app/routes"

	"github.com/gofiber/fiber/v2"
	"go.temporal.io/sdk/client"
)

func main() {

	apiCreate := fiber.New()

	apiCreate.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{"data": "Welcome to golang rest with biber and mongodb"})
	})
	routes.PatientRoute(apiCreate)

	clientT, err := client.Dial(client.Options{
		HostPort: client.DefaultHostPort,
	})

	if err != nil {
		panic(err)
	}
	fmt.Printf("Not null................")
	app.SetTemporalClient(&clientT)

	apiCreate.Listen(":8081")

}
