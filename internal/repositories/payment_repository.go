package repositories

import (
	"context"
	"hotel-booking/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PaymentRepository struct {
	db *pgxpool.Pool
}

func NewPaymentRepository(db *pgxpool.Pool) *PaymentRepository {
	return &PaymentRepository{db: db}
}

func (r *PaymentRepository) GetPaymentByID(ctx context.Context, id int) (*models.Payment, error) {
	query := `
		SELECT id, booking_id, amount, payment_method, transaction_id, status,
		       card_last4, card_brand, receipt_url, processed_at, created_at
		FROM payments
		WHERE id = $1
	`

	row := r.db.QueryRow(ctx, query, id)

	var p models.Payment
	err := row.Scan(
		&p.ID, &p.BookingID, &p.Amount, &p.PaymentMethod, &p.TransactionID, &p.Status,
		&p.CardLast4, &p.CardBrand, &p.ReceiptURL, &p.ProcessedAt, &p.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
