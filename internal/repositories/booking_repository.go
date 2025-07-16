package repositories

import (
	"context"
	"hotel-booking/internal/models"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type BookingRepository struct {
	db *pgxpool.Pool
}

func NewBookingRepository(db *pgxpool.Pool)*BookingRepository{
	return &BookingRepository{db: db}
}

func (r *BookingRepository)CreateBooking (ctx context.Context, booking *models.Booking)error{
	query := `
		INSERT INTO bookings 
		(user_id, room_id, check_in_date, check_out_date, adults, children, total_amount, status, payment_status, special_requests)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING id, created_at, updated_at
	`
	err:= r.db.QueryRow(ctx, query, 
		booking.UserID,
		booking.RoomID,
		booking.CheckInDate,
		booking.CheckOutDate,
		booking.Adults,
		booking.Children,
		booking.TotalAmount,
		booking.Status,
		booking.PaymentStatus,
		booking.SpecialRequests,
	).Scan(&booking.ID, &booking.CreatedAt, &booking.UpdatedAt)
	return err
}

func (r *BookingRepository) GetBookingByID(ctx context.Context, bookingID int) (*models.Booking, error) {
	query := `
		SELECT id, user_id, room_id, check_in_date, check_out_date, adults, children, total_amount, status, payment_status, special_requests, created_at, updated_at
		FROM bookings
		WHERE id = $1
	`

	var booking models.Booking
	var specialRequests pgtype.Text

	err := r.db.QueryRow(ctx, query, bookingID).Scan(
		&booking.ID,
		&booking.UserID,
		&booking.RoomID,
		&booking.CheckInDate,
		&booking.CheckOutDate,
		&booking.Adults,
		&booking.Children,
		&booking.TotalAmount,
		&booking.Status,
		&booking.PaymentStatus,
		&specialRequests,        // ✅ nullable field
		&booking.CreatedAt,
		&booking.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	if specialRequests.Valid {
		booking.SpecialRequests = specialRequests.String
	} else {
		booking.SpecialRequests = ""
	}

	return &booking, nil
}

func (r *BookingRepository)GetAllBookings(ctx context.Context)([]*models.BookingResponse, error){
	query := `
		SELECT 
			b.id, b.check_in_date, b.check_out_date, b.adults, b.children,
			b.total_amount, b.status, b.payment_status, b.special_requests,
			b.created_at, b.updated_at,
			u.id, u.name, u.email,
			r.id, r.room_number, r.room_type
		FROM bookings b
		JOIN users u ON b.user_id = u.id
		JOIN rooms r ON b.room_id = r.id
		ORDER BY b.created_at DESC
	`
	rows, err:=r.db.Query(ctx, query)
	if err!=nil{
		return nil, err
	}
	defer rows.Close()

	var bookings []*models.BookingResponse
	
	for rows.Next(){
		var b models.BookingResponse
			err := rows.Scan(
			&b.ID, &b.CheckInDate, &b.CheckOutDate, &b.Adults, &b.Children,
			&b.TotalAmount, &b.Status, &b.PaymentStatus, &b.SpecialRequests,
			&b.CreatedAt, &b.UpdatedAt,
			&b.User.ID, &b.User.Name, &b.User.Email,
			&b.Room.ID, &b.Room.RoomNumber, &b.Room.RoomType,
		)

		if err!=nil{
			return nil, err
		}

		bookings = append(bookings, &b)
	}

	return bookings , nil

}


