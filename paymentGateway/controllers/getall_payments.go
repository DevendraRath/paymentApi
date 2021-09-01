package controllers

import (
	"encoding/json"
	"log"
	"os"
	"paymentGateway/models"

	"github.com/gofiber/fiber"
	"github.com/joho/godotenv"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentintent"
)

func GetAllPayments(c *fiber.Ctx) {

	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	apiKey := os.Getenv("SK_TEST_KEY")
	stripe.Key = apiKey
	var body models.GetAllPayments
	err = c.BodyParser(&body)

	if err != nil {
		println("Error while parsing", err)
	}

	params := &stripe.PaymentIntentListParams{Customer: stripe.String(body.Customer)}

	response := paymentintent.List(params)
	paymentDetails := []models.PaymentResponse{}

	for response.Next() {
		payments := models.PaymentResponse{}
		pi := response.PaymentIntent()
		payments.Amount = *(&pi.Amount)
		payments.ID = *(&pi.ID)
		paymentDetails = append(paymentDetails, payments)

	}
	json.Marshal(paymentDetails)
	c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success":  true,
		"payments": paymentDetails,
	})

}
