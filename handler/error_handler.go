package handler

import (
	"github.com/gofiber/fiber/v2"
)

// GlobalErrorHandler menangkap semua error yang terjadi di aplikasi secara terpusat
func GlobalErrorHandler(c *fiber.Ctx, err error) error {
	// Status default adalah 500 Internal Server Error
	code := fiber.StatusInternalServerError

	// Jika error-nya berasal dari bawaan Fiber (misal 404 atau 400), ambil status aslinya
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	// Balikkan response dengan format JSON yang sangat rapi standar Enterprise
	return c.Status(code).JSON(fiber.Map{
		"success": false,
		"error": fiber.Map{
			"code":    code,
			"message": err.Error(),
			"trace_id": "ERR-SIMS4-DATA-ENGINE", // Di industri asli, ini diisi kode UUID unik log
		},
	})
}