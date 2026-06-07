package usecase

import (
	"context"
	"encoding/json"
	"net/http"
	"sims-daas/repository"
	"time"

	"github.com/go-redis/redis/v8"
)

type TravelPackage struct {
	Destination          string `json:"destination"`
	Description          string `json:"description"`
	Price                int    `json:"price"`
	IdealMood            string `json:"ideal_mood"`
	RecommendedForSalary int    `json:"recommended_for_salary"`
}

type TravelAPIResponse struct {
	Success bool            `json:"success"`
	Data    []TravelPackage `json:"data"`
}

type EnrichedRecommendation struct {
	repository.CareerRecommendation
	RecommendedVacation *TravelPackage `json:"recommended_vacation"`
}

type CareerUsecaseInterface interface {
	GetRecommendations() ([]EnrichedRecommendation, error)
}

type CareerUsecase struct {
	Repo repository.CareerRepositoryInterface
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

func (u *CareerUsecase) fetchTravelPackages() ([]TravelPackage, error) {
	// Panggil Node.js Microservice
	resp, err := http.Get("http://travel_api:3000/api/travel-packages")
	if err != nil {
		// Fallback for local development if not in docker
		resp, err = http.Get("http://localhost:3000/api/travel-packages")
		if err != nil {
			return nil, err
		}
	}
	defer resp.Body.Close()

	var travelResp TravelAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&travelResp); err != nil {
		return nil, err
	}
	return travelResp.Data, nil
}

func (u *CareerUsecase) GetRecommendations() ([]EnrichedRecommendation, error) {
	cacheKey := "api:v1:enriched_recommendations"

	cachedData, err := u.Rdb.Get(u.Ctx, cacheKey).Result()
	if err == nil {
		var list []EnrichedRecommendation
		if err := json.Unmarshal([]byte(cachedData), &list); err == nil {
			return list, nil
		}
	}

	careers, err := u.Repo.GetRecommendations()
	if err != nil {
		return nil, err
	}

	// Fetch Travel Packages from Node.js service
	travelPackages, err := u.fetchTravelPackages()
	
	var enrichedList []EnrichedRecommendation
	for _, career := range careers {
		enriched := EnrichedRecommendation{
			CareerRecommendation: career,
		}

		// Basic matching logic: find a vacation that matches salary and mood
		if err == nil {
			for _, pkg := range travelPackages {
				if career.BaseSalary >= pkg.RecommendedForSalary {
					// Pointers loop issue fix by copying
					pkgCopy := pkg
					enriched.RecommendedVacation = &pkgCopy
					if career.IdealMood == pkg.IdealMood {
						break // Perfect match!
					}
				}
			}
		}

		enrichedList = append(enrichedList, enriched)
	}

	jsonData, err := json.Marshal(enrichedList)
	if err == nil {
		u.Rdb.Set(u.Ctx, cacheKey, jsonData, 5*time.Minute)
	}

	return enrichedList, nil
}

