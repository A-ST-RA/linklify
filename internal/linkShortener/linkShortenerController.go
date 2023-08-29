package linkShortener

import "github.com/gofiber/fiber/v2"

type LinkShortenerController struct {
	s *LinkShortenerService
}

func NewLinkShortenerController(s *LinkShortenerService) *LinkShortenerController {
	return &LinkShortenerController{
		s: s,
	}
}

func (c *LinkShortenerController) PingRoute(ctx *fiber.Ctx) error {
	type Response struct {
		Message string `json:"message"`
	}

	return ctx.JSON(Response{
		Message: "PONG",
	})
}
