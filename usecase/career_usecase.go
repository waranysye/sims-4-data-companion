package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"sims-daas/repository"
	"time"

	"github.com/go-redis/redis/v8"
)

type CareerUsecase struct {
	Repo  *repository.CareerRepository
	Rdb   *redis.Client
	Ctx   context.Context
}

func NewCareerUsecase(repo *repository.CareerRepository, rdb *redis.Client) *CareerUsecase {
	return &CareerUsecase{
		Repo: repo,
		Rdb:  rdb,
		Ctx:  context.Background(),
	}
}

func (u *CareerUsecase) GetRecommendations() ([]repository.CareerRecommendation, error) {
	cacheKey := "api:v1:recommendations"

	// 1. Coba ambil dari Redis Cache dulu
	cachedData, err := u.Rdb.Get(u.Ctx, cacheKey).Result()
	if err == nil {
		fmt.Println("⚡ Cache HIT: Mengambil data dari Redis!")
		var list []repository.CareerRecommendation
		if err := json.Unmarshal([]byte(cachedData), &list); err == nil {
			return list, nil
		}
	}

	// 2. Jika Cache Miss, ambil dari PostgreSQL
	fmt.Println("🐢 Cache MISS: Mengambil data dari PostgreSQL...")
	list, err := u.Repo.GetRecommendations()
	if err != nil {
		return nil, err
	}

	// 3. Simpan hasil dari Postgres ke Redis (Set EXPIRE / TTL selama 5 Menit)
	jsonData, err := json.Marshal(list)
	if err == nil {
		u.Rdb.Set(u.Ctx, cacheKey, jsonData, 5*time.Minute)
		fmt.Println("💾 Sukses menyimpan data ke Redis Cache!")
	}

	return list, nil
}