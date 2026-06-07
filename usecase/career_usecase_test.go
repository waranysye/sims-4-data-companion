package usecase

import (
	"sims-daas/repository"
	"testing"

	"github.com/go-redis/redis/v8"
)

func TestGetRecommendations_Success(t *testing.T) {
	// 1. Siapkan data tiruan (Mock)
	mockRepo := &repository.MockCareerRepository{
		MockData: []repository.CareerRecommendation{
			{CareerName: "Astronaut", TraitName: "Genius", CompatibilityScore: 5},
		},
		MockErr: nil,
	}

	// 2. Konek ke Redis client kosong untuk testing
	fakeRedis := redis.NewClient(&redis.Options{Addr: "localhost:6379"})

	// 3. Inisialisasi Usecase dengan menginjeksikan MockRepo (SOLID implementation)
	careerUsecase := NewCareerUsecase(mockRepo, fakeRedis)

	// 4. Jalankan fungsi yang mau dites
	result, err := careerUsecase.GetRecommendations()

	// 5. Validasi hasilnya harus sukses sesuai ekspektasi
	if err != nil {
		t.Fatalf("Ekspektasi tidak error, tapi dapet error: %v", err)
	}

	if len(result) != 1 || result[0].CareerName != "Astronaut" {
		t.Errorf("Data yang dihasilkan tidak sesuai dengan mock template")
	}
}
