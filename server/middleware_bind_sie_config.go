package server

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jeonghoikun/colagom.com/site"
	"github.com/jeonghoikun/colagom.com/store"
)

func bindSiteConfig(c *fiber.Ctx) error {
	m := fiber.Map{
		"Site": fiber.Map{
			"Config": site.Config,
			"Store": fiber.Map{
				"Categories": store.ListAllCategories(),
			},
		},
	}
	if err := c.Bind(m); err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	return c.Next()
}
