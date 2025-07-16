package repositories

import (
	"context"
	"hotel-booking/internal/models"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type RoomRepository struct {
	db *pgxpool.Pool
}

func NewRoomRepository(db *pgxpool.Pool) *RoomRepository {
	return &RoomRepository{db: db}
}

func (r *RoomRepository) GetAvailableRooms(ctx context.Context, checkIn, checkOut time.Time) ([]*models.Room, error) {
	query := `
		SELECT r.id, r.room_number, r.room_type, r.description, r.price_per_night, 
		       r.capacity, r.floor, r.amenities, r.is_available, r.created_at, r.updated_at
		FROM rooms r
		WHERE r.is_available = true
		AND r.id NOT IN (
			SELECT b.room_id
			FROM bookings b
			WHERE b.check_in_date < $2 AND b.check_out_date > $1
		)
	`
	rows, err := r.db.Query(ctx, query, checkIn, checkOut)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rooms []*models.Room
	for rows.Next() {
		var room models.Room
		err := rows.Scan(
			&room.ID,
			&room.RoomNumber,
			&room.RoomType,
			&room.Description,
			&room.PricePerNight,
			&room.Capacity,
			&room.Floor,
			&room.Amenities,
			&room.IsAvailable,
			&room.CreatedAt,
			&room.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		rooms = append(rooms, &room)
	}

	return rooms, nil
}


func (r *RoomRepository)GetRoomByID(ctx context.Context, id int)(*models.Room,error){
	query := `
		SELECT id, room_number, room_type, description, price_per_night, capacity, floor, amenities, is_available, created_at, updated_at
		FROM rooms
		WHERE id = $1
	`
	var room models.Room
	err := r.db.QueryRow(ctx, query, id).Scan(
	&room.ID,
	&room.RoomNumber,
	&room.RoomType,
	&room.Description,
	&room.PricePerNight,
	&room.Capacity,
	&room.Floor,
	&room.Amenities,
	&room.IsAvailable,
	&room.CreatedAt,
	&room.UpdatedAt,
	)
	if err!=nil{
		return nil, err
	}
	return &room, nil
}

func (r *RoomRepository) CreateRoom(ctx context.Context, room *models.Room)error{
	query := `
		INSERT INTO rooms (room_number, room_type, description, price_per_night, capacity, floor, amenities, is_available)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, created_at, updated_at
	`
	return r.db.QueryRow(ctx, query,
		room.RoomNumber,
		room.RoomType,
		room.Description,
		room.PricePerNight,
		room.Capacity,
		room.Floor,
		room.Amenities,
		room.IsAvailable,
	).Scan(&room.ID, &room.CreatedAt, &room.UpdatedAt)
}

func (r *RoomRepository)UpdateRoom(ctx context.Context, room *models.Room)error{
	query := `
		UPDATE rooms 
		SET room_number = $1, room_type = $2, description = $3, price_per_night = $4, 
		    capacity = $5, floor = $6, amenities = $7, is_available = $8, updated_at = NOW()
		WHERE id = $9
		RETURNING updated_at
	`
	return r.db.QueryRow(ctx, query,
		room.RoomNumber,
		room.RoomType,
		room.Description,
		room.PricePerNight,
		room.Capacity,
		room.Floor,
		room.Amenities,
		room.IsAvailable,
		room.ID,
	).Scan(&room.UpdatedAt)
}

func (r *RoomRepository)DeleteRoom(ctx context.Context, roomID int)error{
	query := `DELETE from rooms WHERE id = $1`
	_, err := r.db.Exec(ctx, query, roomID)
	return err
}