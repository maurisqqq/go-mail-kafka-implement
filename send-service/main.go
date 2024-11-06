package main

import (
	"send-service/app/config"
	router "send-service/modules/routes"
	"send-service/pkg/logs"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func main() {
	conf, _ := config.Init()
	app := fiber.New()

	// Format loging
	logs.SetFormatter()

	// Running Apps
	logrus.Infof("🚀 Succeed Running On http://localhost%v", conf.App.Port)
	router.SetupRoutes(app)
}
