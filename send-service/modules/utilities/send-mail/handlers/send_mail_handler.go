package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func (h *SendMailHandler) SendMail(c *fiber.Ctx) error {
	mail := c.Params("mail")
	result := h.SendMailService.SendMail(mail)
	if result != nil {
		return c.Status(500).JSON(result)
	}
	return c.Status(200).JSON(result)
}
