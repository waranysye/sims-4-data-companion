package usecase

import (
	"context"
	"encoding/json"
	"sims-daas/repository"
	"time"

	"github.com/go-redis/redis/v8"
)

// 👈 Kita buat juga interface untuk Usecase-nya agar nanti bisa di-Mock saat testing!
type CareerUsecaseInterface interface {
	GetRecommendations() ([]repository.CareerRecommendation, error)
}

type CareerUsecase struct {
	Repo repository.CareerRepositoryInterface // 👈 Merujuk ke Interface, bukan struct langsung
	Rdb  *redis.Client
	Ctx  context.Context
}

func NewCareerUsecase(repo repository.CareerRepositoryInterface, rdb *redis.Client) *CareerUsecase {
	return &CareerUsecase{
		Repo: repo,
		Rdb:  rdb,
		Ctx:  context.Background(),
	}
}

func (u *CareerUsecase) GetRecommendations() ([]repository.CareerRecommendation, error) {
	cacheKey := "api:v1:recommendations"

	cachedData, err := u.Rdb.Get(u.Ctx, cacheKey).Result()
	if err == nil {
		var list []repository.CareerRecommendation
		if err := json.Unmarshal([]byte(cachedData), &list); err == nil {
			return list, nil
		}
	}

	list, err := u.Repo.GetRecommendations()
	if err != nil {
		return nil, err
	}

	jsonData, err := json.Marshal(list)
	if err == nil {
		u.Rdb.Set(u.Ctx, cacheKey, jsonData, 5*time.Minute)
	}

	return list, nil
}
