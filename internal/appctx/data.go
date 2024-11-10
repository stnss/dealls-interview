package appctx

import "github.com/gofiber/fiber/v2"

type Data struct {
	Ctx    *fiber.Ctx
	Config *Config
}
