package models

import "time"

type Payment struct {
	ID            int        `json:"id"`
	BookingID     int        `json:"booking_id"`
	Amount        float64    `json:"amount"`
	PaymentMethod string     `json:"payment_method"`
	TransactionID string     `json:"transaction_id"`
	Status        string     `json:"status"`
	CardLast4     string     `json:"card_last4"`
	CardBrand     string     `json:"card_brand"`
	ReceiptURL    string     `json:"receipt_url"`
	ProcessedAt   *time.Time `json:"processed_at"`
	CreatedAt     time.Time  `json:"created_at"`
}
