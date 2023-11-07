package routes

import (
	"patient-registration/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func PatientRoute(app *fiber.App) {
	app.Get("/patient", controllers.RegisterPatient) // get all employee

}
