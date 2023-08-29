package app

import (
	"linklifyBackend/internal/config"
	"linklifyBackend/internal/linkShortener"

	"github.com/gofiber/fiber/v2"
)

type App struct {
	ls   *linkShortener.LinkShortener
	fApp *fiber.App
}

func New() *App {
	if err := config.InitializeConfig(); err != nil {
		panic(err)
	}
	
	fApp := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  config.AppConfig.AppName,
		AppName:       config.AppConfig.AppName,
	})

	ls := linkShortener.New(fApp)

	return &App{
		fApp: fApp,
		ls:   ls,
	}
}

func (a *App) Run() error {
	if err := a.ls.RegisterLinkShortenerRoutes(); err != nil {
		return err
	}

	a.fApp.Listen(":" + config.AppConfig.Port)

	return nil
}
