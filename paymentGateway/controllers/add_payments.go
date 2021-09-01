package controllers

import (
	"log"
	"os"
	"paymentGateway/models"

	"github.com/gofiber/fiber"
	"github.com/joho/godotenv"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentintent"
)

func AddPayments(c *fiber.Ctx) {

	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	apiKey := os.Getenv("SK_TEST_KEY")
	stripe.Key = apiKey
	var json models.Payments
	err = c.BodyParser(&json)

	if err != nil {
		println("Error while parsing :", err)
	}

	response, err := paymentintent.New(&stripe.PaymentIntentParams{
		Amount:   stripe.Int64(json.Amount),
		Customer: stripe.String(json.Customer),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		PaymentMethodTypes: []*string{
			stripe.String("card"),
		},
	})

	if err != nil {
		println(err)
	}

	c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success":         true,
		"paymentIntentID": response.ID,
	})

}
