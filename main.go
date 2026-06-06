package main

import (
	"os" // 👈 1. Kita tambahkan package os di sini
	"sims-daas/config"
	"sims-daas/handler"
	"sims-daas/repository"
	"sims-daas/usecase"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	dbPool := config.ConnectDB()
	defer dbPool.Close()

	// 👈 2. Cari host Redis dari environment variable (untuk Docker)
	// Jika kosong (artinya jalan di laptop langsung), dia otomatis pakai default "localhost:6379"
	redisHost := os.Getenv("REDIS_HOST")
	if redisHost == "" {
		redisHost = "localhost:6379"
	}

	// Inisialisasi Redis Client menggunakan variabel yang dinamis
	rdb := redis.NewClient(&redis.Options{
		Addr: redisHost,
	})

	// Daftarkan GlobalErrorHandler saat membuat instance Fiber baru
	app := fiber.New(fiber.Config{
		ErrorHandler: handler.GlobalErrorHandler,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))
	app.Post("/api/v1/auth/login", handler.Login)

	// Dependency Injection manual ala Clean Architecture
	careerRepo := repository.NewCareerRepository(dbPool)
	careerUsecase := usecase.NewCareerUsecase(careerRepo, rdb)
	careerHandler := handler.NewCareerHandler(careerUsecase)

	app.Get("/api/v1/recommendations", handler.JWTMiddleware, careerHandler.FetchRecommendations)

	app.Listen(":8080")
}
