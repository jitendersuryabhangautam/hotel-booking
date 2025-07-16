package handlers

import (
	"hotel-booking/internal/models"
	"hotel-booking/internal/response"
	"hotel-booking/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookingHandler struct {
	service *services.BookingService
}

func NewBookingHandler(service *services.BookingService) *BookingHandler {
	return &BookingHandler{service: service}
}

func (h *BookingHandler) CreateBooking(c *gin.Context) {
	var req models.CreateBookingRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.JSON(c, http.StatusBadRequest, false, "Invalid request", nil, err.Error())
		return
	}

	booking, err := h.service.CreateBooking(c.Request.Context(), &req)
	if err != nil {
		response.JSON(c, http.StatusBadRequest, false, "Failed to create booking", nil, err.Error())
		return
	}

	response.JSON(c, http.StatusCreated, true, "Booking created successfully", booking, "")
}

func (h *BookingHandler) GetBookingByID(c *gin.Context) {
	idStr := c.Param("id")
	bookingID, err := strconv.Atoi(idStr)
	if err != nil {
		response.JSON(c, http.StatusBadRequest, false, "Invalid booking ID", nil, err.Error())
		return
	}

	booking, err := h.service.GetBookingByID(c.Request.Context(), bookingID)
	if err != nil {
		response.JSON(c, http.StatusNotFound, false, "Booking not found", nil, err.Error())
		return
	}

	response.JSON(c, http.StatusOK, true, "Booking fetched successfully", booking, "")
}

func (h *BookingHandler)GetAllBookings(c *gin.Context){
	bookings, err:=h.service.GetAllBookings(c.Request.Context())
	if err!=nil{
		response.JSON(c, http.StatusInternalServerError, false, "failed to fetch bookings data", nil, err.Error())
		return
	}
	response.JSON(c, http.StatusOK, true, "Bookings Fetched successfully", bookings,"")
}