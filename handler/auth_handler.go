package handler

import (
	"time"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// JWT_SECRET adalah kunci rahasia server untuk menandatangani token.
// Di aplikasi asli, ini wajib ditaruh di os.Getenv("JWT_SECRET")
var JWT_SECRET = []byte("SuperSecretSims4EngineKey")

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *fiber.Ctx) error {
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Format request tidak valid")
	}

	// 🔐 Hardcoded credential sederhana untuk simulasi login admin
	if req.Username != "admin" || req.Password != "sims4pro" {
		return fiber.NewError(fiber.StatusUnauthorized, "Username atau password salah")
	}

	// 🎫 Membuat klaim data di dalam JWT token
	claims := jwt.MapClaims{
		"username": req.Username,
		"exp":      time.Now().Add(time.Minute * 15).Unix(), // Token kedaluwarsa dalam 15 menit
	}

	// Buat token dengan algoritma HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(JWT_SECRET)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Gagal membuat token keamanan")
	}

	// Kembalikan token dalam format JSON rapi
	return c.JSON(fiber.Map{
		"success": true,
		"token":   t,
	})
}