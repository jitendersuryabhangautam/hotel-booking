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

type PaymentService struct {
	repo *repositories.PaymentRepository
}

func NewPaymentService(repo *repositories.PaymentRepository) *PaymentService {
	return &PaymentService{repo: repo}
}

func (s *PaymentService) GetPaymentByID(ctx context.Context, id int) (*models.Payment, error) {
	cacheKey := fmt.Sprintf("payment:%d", id)

	// 1️⃣ Try fetching from cache
	cached, err := cache.Client.Get(ctx, cacheKey).Result()
	if err == nil {
		var payment models.Payment
		if err := json.Unmarshal([]byte(cached), &payment); err == nil {
			fmt.Println("✅ Cache hit (payment)")
			return &payment, nil
		}
	}

	// 2️⃣ Fallback to DB
	fmt.Println("❌ Cache miss ➔ Fetching payment from DB")
	payment, err := s.repo.GetPaymentByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// 3️⃣ Store in Redis
	paymentJSON, err := json.Marshal(payment)
	if err == nil {
		_ = cache.Client.Set(ctx, cacheKey, paymentJSON, 10*time.Minute).Err()
	}

	return payment, nil
}
