package handlers

import (
	"fmt"
	"hotel-booking/internal/models"
	"hotel-booking/internal/response"
	"hotel-booking/internal/services"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type RoomHandler struct {
	service *services.RoomService
}

func NewRoomHandler(service *services.RoomService) *RoomHandler {
	return &RoomHandler{service: service}
}

var validate = validator.New()

func (h *RoomHandler) GetAvailableRooms(c *gin.Context) {
	checkInStr := c.Query("check_in")
	checkOutStr := c.Query("check_out")

	if checkInStr == "" || checkOutStr == "" {
		response.JSON(c, http.StatusBadRequest, false, "Missing check_in or check_out", nil, "Missing query parameters")
		return
	}

	checkIn, err1 := time.Parse("2006-01-02", checkInStr)
	checkOut, err2 := time.Parse("2006-01-02", checkOutStr)

	if err1 != nil || err2 != nil || !checkOut.After(checkIn) {
		response.JSON(c, http.StatusBadRequest, false, "Invalid dates", nil, "Invalid check_in or check_out")
		return
	}

	rooms, err := h.service.GetAvailableRooms(c.Request.Context(), checkIn, checkOut)
	if err != nil {
		response.JSON(c, http.StatusInternalServerError, false, "Failed to fetch rooms", nil, err.Error())
		return
	}

	response.JSON(c, http.StatusOK, true, "Available rooms fetched successfully", rooms, "")
}



func (h *RoomHandler)GetRoomByID(c *gin.Context){
	idParam:=c.Param("id")
	id, err:= strconv.Atoi(idParam)
	if err!=nil{
		response.JSON(c, http.StatusBadRequest, false, "Invalid room ID", nil, "ID must be a number")
	}
	room, err := h.service.GetRoomByID(c.Request.Context(), id)
	if err!=nil{
		response.JSON(c, http.StatusNotFound, false, "Room not found", nil, err.Error())
		return
	}
	response.JSON(c, http.StatusOK, true, "Room fetched successfully", room, "")
}

// func (h *RoomHandler)CreateRoom(c *gin.Context){
// 	var req models.CreateRoomRequest
// 	if err:= c.ShouldBindJSON(&req); err!=nil{
// 		response.JSON(c, http.StatusBadRequest, false, "Invalid request", nil, err.Error())
// 		return
// 	}
// 	room:= &models.Room{
// 		RoomNumber: req.RoomNumber,
// 		RoomType: req.RoomType,
// 		Description: req.Description,
// 		PricePerNight: req.PricePerNight,
// 		Capacity: req.Capacity,
// 		Floor: req.Floor,
// 		Amenities: req.Amenities,
// 		IsAvailable: req.IsAvailable,
// 	}

// 	createdRoom, err:= h.service.CreateRoom(c.Request.Context(), room)
// 	if err!=nil{
// 		response.JSON(c, http.StatusInternalServerError, false, "Failed to create room", nil, err.Error())
// 		return
// 	}
// 	response.JSON(c, http.StatusCreated, true, "Room created successfully", createdRoom, "")
// }

func (h *RoomHandler) CreateRoom(c *gin.Context) {
	var req models.CreateRoomRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.JSON(c, http.StatusBadRequest, false, "Invalid JSON", nil, err.Error())
		return
	}

	if err := validate.Struct(req); err != nil {
		var errors []string
		for _, e := range err.(validator.ValidationErrors) {
			errors = append(errors, fmt.Sprintf("Field '%s' failed on '%s'", e.Field(), e.ActualTag()))
		}
		response.JSON(c, http.StatusBadRequest, false, "Validation failed", nil, strings.Join(errors, ", "))
		return
	}

	room := &models.Room{
		RoomNumber:    req.RoomNumber,
		RoomType:      req.RoomType,
		Description:   req.Description,
		PricePerNight: req.PricePerNight,
		Capacity:      req.Capacity,
		Floor:         req.Floor,
		Amenities:     req.Amenities,
		IsAvailable:   req.IsAvailable,
	}

	createdRoom, err := h.service.CreateRoom(c.Request.Context(), room)
	if err != nil {
		response.JSON(c, http.StatusInternalServerError, false, "Failed to create room", nil, err.Error())
		return
	}
	response.JSON(c, http.StatusCreated, true, "Room created successfully", createdRoom, "")
}

func (h *RoomHandler)UpdateRoom(c *gin.Context){
	idParam:=c.Param("id")
	if idParam == ""{
		response.JSON(c, http.StatusBadRequest, false, "Invalid room ID", nil, "ID must be a number")
		return
	}

	var req models.UpdateRoomRequest
	if err:= c.ShouldBindJSON(&req); err!=nil{
		response.JSON(c, http.StatusBadRequest, false, "Invalid request", nil, err.Error())
		return
	}

	roomID, err := strconv.Atoi(idParam)
	if err!=nil{
		response.JSON(c, http.StatusBadRequest, false, "Invalid room ID", nil, "ID must be a number")
		return
	}
	room := &models.Room{
		ID:            roomID,
		RoomNumber:    req.RoomNumber,
		RoomType:      req.RoomType,
		Description:   req.Description,
		PricePerNight: req.PricePerNight,
		Capacity:      req.Capacity,
		Floor:         req.Floor,
		Amenities:     req.Amenities,
		IsAvailable:   req.IsAvailable,
	}

	updatedRoom, err:= h.service.UpdateRoom(c.Request.Context(), room)
	if err!=nil{
		response.JSON(c, http.StatusInternalServerError, false, "Failed to update room", nil, err.Error())
		return
	}
	response.JSON(c, http.StatusOK, true, "Room updated successfully", updatedRoom, "")
}


func (h *RoomHandler)DeleteRoom(c *gin.Context){
	idParam := c.Param("id")
	if idParam == "" {
		response.JSON(c, http.StatusBadRequest, false, "Room ID is required", nil, "Missing room ID")
		return
	}

	roomID, err:= strconv.Atoi(idParam)
	if err!=nil{
		response.JSON(c, http.StatusBadRequest, false, "Invalid room ID", nil, "ID must be a number")
		return
	}
	err = h.service.DeleteRoom(c.Request.Context(), roomID)
	if err!=nil{
		response.JSON(c, http.StatusInternalServerError, false, "Failed to delete room", nil, err.Error())
		return
	}
	response.JSON(c, http.StatusOK, true, "Room deleted successfully", nil, "")
}