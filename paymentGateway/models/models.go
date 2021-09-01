package models

// Customer incoming data for Stripe API
type Customer struct {
	Email string `json:"email"`
	Token string `json:"stripeCreditCardToken"`
}

// Payments incoming data for Stripe API
type Payments struct {
	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`
	Customer string `json:"stripeCustomerID"`
}

// Payments incoming data for Stripe API
type GetAllPayments struct {
	Customer string `json:"stripeCustomerID"`
}

// PaymentResponse outcoming data from Stripe API
type PaymentResponse struct {
	ID     string `json:"id"`
	Amount int64  `json:"amount"`
}
