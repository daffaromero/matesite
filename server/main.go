package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/daffaromero/matesite/server/config"
	"github.com/daffaromero/matesite/server/controller"
	"github.com/daffaromero/matesite/server/helper/logger"
	"github.com/daffaromero/matesite/server/repository"
	"github.com/daffaromero/matesite/server/repository/query"
	"github.com/daffaromero/matesite/server/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/requestid"
)

var logs = logger.New("main")

func webServer() error {
	app := fiber.New()
	app.Use(requestid.New())

	serverConfig := config.NewServerConfig()
	dbConfig := config.NewPostgresDatabase()
	store := repository.NewStore(dbConfig)
	validate := validator.New()

	issueQuery := query.NewIssueQuery(dbConfig)
	issueRepo := repository.NewIssueRepository(store, issueQuery)
	issueService := service.NewIssueService(issueRepo, logs)
	issueController := controller.NewIssueController(validate, issueService)

	logs.Log(fmt.Sprintf("Starting HTTP issue server on %s", serverConfig.HTTP))
	app.Use(cors.New())
	issueController.Route(app)

	err := app.Listen(serverConfig.HTTP, fiber.ListenConfig{
		DisableStartupMessage: true,
	})
	if err != nil {
		logs.Error("Failed to start issue server")
		return err
	}
	return nil
}

func main() {
	if err := webServer(); err != nil {
		logs.Error(err)
	}

	logs.Log("Issue server started")
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	sig := <-sigchan
	logs.Log(fmt.Sprintf("Received signal: %s. Shutting down gracefully...", sig))
}
