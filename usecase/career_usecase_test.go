package usecase

import (
	"sims-daas/repository"
	"testing"

	"github.com/go-redis/redis/v8"
)

func newTestUsecase(mockData []repository.CareerRecommendation, mockErr error) *CareerUsecase {
	mockRepo := &repository.MockCareerRepository{
		MockData: mockData,
		MockErr:  mockErr,
	}
	fakeRedis := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	return NewCareerUsecase(mockRepo, fakeRedis)
}

func TestGetRecommendations_Success(t *testing.T) {
	uc := newTestUsecase([]repository.CareerRecommendation{
		{CareerName: "Astronaut", TraitName: "Genius", CompatibilityScore: 5},
	}, nil)

	result, err := uc.GetRecommendations()
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	if len(result) != 1 || result[0].CareerName != "Astronaut" {
		t.Errorf("Data tidak sesuai dengan mock template")
	}
}

func TestCreateRecommendation_Success(t *testing.T) {
	uc := newTestUsecase(nil, nil)
	req := RecommendationRequest{
		CareerID:           1,
		TraitID:            2,
		CompatibilityScore: 5,
		Reason:             "Test Create",
	}
	if err := uc.CreateRecommendation(req); err != nil {
		t.Fatalf("Expected no error on create, got: %v", err)
	}
}

func TestUpdateRecommendation_Success(t *testing.T) {
	uc := newTestUsecase(nil, nil)
	req := UpdateRecommendationRequest{
		CompatibilityScore: 4,
		Reason:             "Test Update",
	}
	if err := uc.UpdateRecommendation(1, 2, req); err != nil {
		t.Fatalf("Expected no error on update, got: %v", err)
	}
}

func TestDeleteRecommendation_Success(t *testing.T) {
	uc := newTestUsecase(nil, nil)
	if err := uc.DeleteRecommendation(1, 2); err != nil {
		t.Fatalf("Expected no error on delete, got: %v", err)
	}
}

