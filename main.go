package main

import (
	"fmt"
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

	v1 := app.Group("/api/v1")
	v1.Get("/recommendations", handler.JWTMiddleware, careerHandler.FetchRecommendations)
	v1.Post("/recommendations", handler.JWTMiddleware, careerHandler.CreateRecommendation)
	v1.Put("/recommendations/:career_id/:trait_id", handler.JWTMiddleware, careerHandler.UpdateRecommendation)
	v1.Delete("/recommendations/:career_id/:trait_id", handler.JWTMiddleware, careerHandler.DeleteRecommendation)

	// Melayani file Swagger UI dan spesifikasi kontrak API
	app.Static("/swagger", "./swagger.html")
	app.Static("/api-contract.yaml", "./api-contract.yaml")

	// 📡 Inisialisasi dan jalankan gRPC Server di Goroutine (Background Thread)
	grpcServer := handler.NewGrpcServer(careerUsecase)
	go func() {
		if err := grpcServer.Run(); err != nil {
			fmt.Printf("Gagal menyalakan server gRPC: %v\n", err)
		}
	}()

	app.Listen(":8888")
}
