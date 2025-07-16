package handlers

import (
	"hotel-booking/internal/response"
	"hotel-booking/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	service *services.PaymentService
}

func NewPaymentHandler(service *services.PaymentService) *PaymentHandler {
	return &PaymentHandler{service: service}
}

func (h *PaymentHandler) GetPaymentByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.JSON(c, http.StatusBadRequest, false, "Invalid payment ID", nil, err.Error())
		return
	}

	payment, err := h.service.GetPaymentByID(c.Request.Context(), id)
	if err != nil {
		response.JSON(c, http.StatusNotFound, false, "Payment not found", nil, err.Error())
		return
	}

	response.JSON(c, http.StatusOK, true, "Payment fetched successfully", payment, "")
}
