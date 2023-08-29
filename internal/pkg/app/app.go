package app

import (
	"linklifyBackend/internal/config"
	"linklifyBackend/internal/linkShortener"
	"log/slog"
	"os"

	"github.com/gofiber/fiber/v2"
)

type App struct {
	ls     *linkShortener.LinkShortener
	fApp   *fiber.App
	logger *slog.Logger
}

func New() *App {
	if err := config.InitializeConfig(); err != nil {
		panic(err)
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	fApp := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  config.AppConfig.AppName,
		AppName:       config.AppConfig.AppName,
	})

	ls := linkShortener.New(fApp)

	return &App{
		fApp:   fApp,
		ls:     ls,
		logger: logger,
	}
}

func (a *App) Run() error {
	if err := a.ls.RegisterLinkShortenerRoutes(); err != nil {
		return err
	}

	a.logger.Info("Logger initialized")

	a.fApp.Listen(":" + config.AppConfig.Port)

	return nil
}
