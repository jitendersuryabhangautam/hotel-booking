package services

import (
	"context"
	"encoding/json"
	"fmt"
	"hotel-booking/internal/cache"
	"hotel-booking/internal/models"
	"hotel-booking/internal/repositories"
	"time"
)

type RoomService struct {
	repo *repositories.RoomRepository
}

func NewRoomService(repo *repositories.RoomRepository) *RoomService {
	return &RoomService{repo: repo}
}

// func (s *RoomService) GetAvailableRooms(ctx context.Context) ([]*models.Room, error) {
// 	return s.repo.GetAvailableRooms(ctx)
// }
func (s *RoomService) GetAvailableRooms(ctx context.Context, checkIn, checkOut time.Time) ([]*models.Room, error) {
	cacheKey := fmt.Sprintf("available_rooms:%s:%s", checkIn.Format("2006-01-02"), checkOut.Format("2006-01-02"))

	cached, err := cache.Client.Get(ctx, cacheKey).Result()
	if err == nil {
		var rooms []*models.Room
		if err := json.Unmarshal([]byte(cached), &rooms); err == nil {
			fmt.Println("✅ Cache hit")
			return rooms, nil
		}
	}

	fmt.Println("❌ Cache miss ➔ Fetching from DB")
	rooms, err := s.repo.GetAvailableRooms(ctx, checkIn, checkOut)
	if err != nil {
		return nil, err
	}

	roomsJSON, err := json.Marshal(rooms)
	if err == nil {
		_ = cache.Client.Set(ctx, cacheKey, roomsJSON, 10*time.Minute).Err()
	}

	return rooms, nil
}

func (s *RoomService) GetRoomByID(ctx context.Context, id int)(*models.Room, error){
	return s.repo.GetRoomByID(ctx, id)
}

func (s *RoomService) CreateRoom(ctx context.Context, room *models.Room)(*models.Room, error){
	err:=s.repo.CreateRoom(ctx, room)
	if err!=nil{
		return nil, err
	}
	return room, nil
}

func (s *RoomService)UpdateRoom(ctx context.Context, room *models.Room)(*models.Room, error){
	err:=s.repo.UpdateRoom(ctx, room)
	if err!=nil{
		return nil, err
	}
	return room, nil
}

func (s *RoomService)DeleteRoom(ctx context.Context, roomID int)error{
	return s.repo.DeleteRoom(ctx, roomID)
}