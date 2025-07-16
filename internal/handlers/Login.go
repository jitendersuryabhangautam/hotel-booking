package handlers

import (
	"hotel-booking/internal/models"
	"hotel-booking/internal/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *UserHandler) Login(c *gin.Context) {
	var req models.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil || req.Email == "" || req.Password == "" {
		response.JSON(c, http.StatusBadRequest, false, "Invalid email or password", nil, "Missing email or password")
		return
	}

	loginResp, err := h.service.LoginUser(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		response.JSON(c, http.StatusUnauthorized, false, "Login failed", nil, err.Error())
		return
	}

	response.JSON(c, http.StatusOK, true, "Login successful", loginResp, "")
}
