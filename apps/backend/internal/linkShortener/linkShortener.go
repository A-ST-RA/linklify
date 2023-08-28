package linkShortener

import "github.com/gofiber/fiber/v2"

type LinkShortener struct {
	c    *LinkShortenerController
	fApp *fiber.App
	s    *LinkShortenerService
}

func New(fApp *fiber.App) *LinkShortener {
	s := NewLinkShortenerService()
	c := NewLinkShortenerController(s)

	return &LinkShortener{
		c:    c,
		s:    s,
		fApp: fApp,
	}
}

func (l *LinkShortener) RegisterLinkShortenerRoutes() error {
	l.fApp.Get("/ping", l.c.PingRoute)

	return nil
}
