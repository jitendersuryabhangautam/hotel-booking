package services

import (
	"context"
	"errors"
	"hotel-booking/internal/models"
	"hotel-booking/internal/repositories"
	"time"
)

type BookingService struct {
	repo *repositories.BookingRepository
}

func NewBookingService(repo *repositories.BookingRepository) *BookingService {
	return &BookingService{repo: repo}
}

func (s *BookingService) CreateBooking(ctx context.Context, req *models.CreateBookingRequest) (*models.Booking, error) {
	// ✅ Parse check-in and check-out dates
	checkIn, err := time.Parse("2006-01-02", req.CheckInDate)
	if err != nil {
		return nil, errors.New("invalid check-in date format (expected YYYY-MM-DD)")
	}

	checkOut, err := time.Parse("2006-01-02", req.CheckOutDate)
	if err != nil {
		return nil, errors.New("invalid check-out date format (expected YYYY-MM-DD)")
	}

	// ✅ Validate that check-out is after check-in
	if !checkOut.After(checkIn) {
		return nil, errors.New("check-out date must be after check-in date")
	}

	// ✅ Calculate duration and total amount (you can later get price from DB)
	days := int(checkOut.Sub(checkIn).Hours() / 24)
	if days < 1 {
		return nil, errors.New("booking must be for at least one night")
	}
	pricePerNight := 100.0 // Example: You should fetch real price from room table
	totalAmount := float64(days) * pricePerNight

	// ✅ Create booking struct with correct types
	booking := &models.Booking{
		UserID:          req.UserID,
		RoomID:          req.RoomID,
		CheckInDate:     checkIn,
		CheckOutDate:    checkOut,
		Adults:          req.Adults,
		Children:        req.Children,
		TotalAmount:     totalAmount,
		Status:          "confirmed",
		PaymentStatus:   "pending",
		SpecialRequests: req.SpecialRequests,
	}

	// ✅ Save booking to DB
	err = s.repo.CreateBooking(ctx, booking)
	if err != nil {
		return nil, err
	}

	return booking, nil
}

func (s *BookingService) GetBookingByID(ctx context.Context, bookingID int) (*models.Booking, error) {
	return s.repo.GetBookingByID(ctx, bookingID)
}

func (s *BookingService)GetAllBookings(ctx context.Context)([]*models.BookingResponse, error){
	return s.repo.GetAllBookings(ctx)
}