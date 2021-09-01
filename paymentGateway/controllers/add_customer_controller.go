package controllers

import (
	"log"
	"os"
	"paymentGateway/models"

	"github.com/gofiber/fiber"
	"github.com/joho/godotenv"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/customer"
)

func AddCustomer(c *fiber.Ctx) {

	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	apiKey := os.Getenv("SK_TEST_KEY")
	stripe.Key = apiKey
	var json models.Customer
	err = c.BodyParser(&json)

	if err != nil {
		println("Error while parsing :", err)
	}

	response, err := customer.New(&stripe.CustomerParams{
		Email: stripe.String(json.Email),
		Token: stripe.String(json.Token),
	})

	if err != nil {
		println(err)
	}
	c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success":          true,
		"stripeCustomerID": response.ID,
	})

}
