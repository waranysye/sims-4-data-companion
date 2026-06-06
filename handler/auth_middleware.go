package handler

import (
	"strings"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTMiddleware(c *fiber.Ctx) error {
	// Ambil token dari HTTP Header bernama "Authorization"
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "Akses ditolak: Token tidak ditemukan")
	}

	// Format header biasanya: "Bearer <TOKEN_DISINI>". Kita potong teks "Bearer " nya
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader {
		return fiber.NewError(fiber.StatusUnauthorized, "Format token harus: Bearer <token>")
	}

	// Validasi keaslian token menggunakan kunci JWT_SECRET kita
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return JWT_SECRET, nil
	})

	if err != nil || !token.Valid {
		return fiber.NewError(fiber.StatusUnauthorized, "Akses ditolak: Token palsu atau sudah kedaluwarsa")
	}

	// Jika lolos pemeriksaan, silakan lanjut ke handler utama (FetchRecommendations)
	return c.Next()
}