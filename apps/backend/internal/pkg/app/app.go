package app

import (
	"linklifyBackend/internal/linkShortener"

	"github.com/gofiber/fiber/v2"
)

type App struct {
	ls   *linkShortener.LinkShortener
	fApp *fiber.App
}

func New() *App {
	fApp := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Linklify",
		AppName:       "Linklify",
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

	a.fApp.Listen(":8080")

	return nil
}
