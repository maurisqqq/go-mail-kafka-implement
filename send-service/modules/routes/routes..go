package routes

import (
	"send-service/app/config"
	"send-service/app/connections"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	sendMailHandler "send-service/modules/utilities/send-mail/handlers"
)

func SetupRoutes(app *fiber.App) {
	// General variable
	conf, _ := config.Init()
	kafka, _ := connections.KafkaPublisherClient()
	// Set handler
	sendMailHandler := sendMailHandler.NewSendMailHandler(kafka)

	// Set cors middleware
	app.Use(cors.New())
	app.Use(recover.New())

	// Set routes for API
	api := app.Group("/api/v1", logger.New())
	api.Get("/send-mail/:mail", sendMailHandler.SendMail)

	app.Listen(conf.App.Port)
}
