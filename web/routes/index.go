package routes

import (
	"github.com/ggmolly/belfast/misc"
	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"title":   "Belfast",
		"commits": misc.GetCommits(),
		"page":    "index",
	}, "layouts/main")
}
