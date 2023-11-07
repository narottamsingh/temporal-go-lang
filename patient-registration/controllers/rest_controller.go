package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"patient-registration/app"

	"github.com/gofiber/fiber/v2"
	"go.temporal.io/sdk/client"
)

func RegisterPatient(e *fiber.Ctx) error {
	//	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//	defer cancel() // Ensure the context is canceled when the function returns

	// Execute the Temporal Workflow to start the subscription.

	options := client.StartWorkflowOptions{
		ID:        "Patient_Creation-108",
		TaskQueue: app.RegistrationTaskQueue,
	}

	var data app.PatientDetails

	data.Name = "Alex"
	data.Age = "30"
	data.Contact = "123456789"
	data.Address = "Texas"

	fmt.Printf("11111111111111111")
	fmt.Printf("%s", app.TemporalClient)

	we, err := app.TemporalClient.ExecuteWorkflow(context.Background(), options, app.PatientRegistrationWorkflow, data)
	fmt.Printf("222222222222")

	if err != nil {
		log.Fatalln("Unable to complete workflow\n", err)

	}

	var workflowOutput app.PatientDetails
	err = we.Get(context.Background(), &workflowOutput)

	//log.Fatalln("workflow workflowOutput: ", workflowOutput)
	if err != nil {
		log.Fatalln("Unable to get workflow result", err)
	}

	return e.Status(http.StatusBadRequest).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "success",
		"data":    "ok",
	})

}

func RegisterPatient1(e *fiber.Ctx) error {
	return e.Status(http.StatusBadRequest).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "success",
		"data":    "ok",
	})

}
