package handler

import (
	"sims-daas/usecase"

	"github.com/gofiber/fiber/v2"
)

type CareerHandler struct {
	Usecase usecase.CareerUsecaseInterface // 👈 Pakai Interface juga
}

func NewCareerHandler(u usecase.CareerUsecaseInterface) *CareerHandler {
	return &CareerHandler{Usecase: u}
}

func (h *CareerHandler) FetchRecommendations(c *fiber.Ctx) error {
	data, err := h.Usecase.GetRecommendations()
	if err != nil {
		return err // 👈 Cukup return error-nya, biar Global Error Handler yang mengurus formatnya!
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    data,
	})
}
