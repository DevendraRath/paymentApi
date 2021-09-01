package main

import (
	"paymentGateway/controllers"

	"github.com/gofiber/fiber"
)

func setupRoutes(app *fiber.App) {
	app.Post("/api/customer", controllers.AddCustomer)
	app.Post("/api/payments", controllers.AddPayments)
	app.Get("/api/payments/:stripeCustomerID", controllers.GetAllPayments)
}

func main() {
	app := fiber.New()

	setupRoutes(app)
	app.Listen(3008)
}
