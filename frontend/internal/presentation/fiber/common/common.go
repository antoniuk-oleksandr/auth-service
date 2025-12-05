package common

import (
	"context"
	"encoding/json"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
)

func Page(ctx *fiber.Ctx, tmpl templ.Component) error {
	ctx.Type("html")
	return tmpl.Render(context.Background(), ctx.Response().BodyWriter())
}

func JSON(ctx *fiber.Ctx, status int, body any) error {
	return ctx.Status(status).JSON(body)
}

func Trigger(events map[string]any) string {
	b, _ := json.Marshal(events)
	return string(b)
}

func HTMXTrigger(ctx *fiber.Ctx, events map[string]any) {
	ctx.Set("HX-Trigger", Trigger(events))
}

func HTMXRedirect(ctx *fiber.Ctx, url string) {
	ctx.Set("HX-Redirect", url)
}

func TriggerError(ctx *fiber.Ctx, name string, message string) {
	ctx.Set("HX-Trigger", Trigger(map[string]any{
		name: map[string]any{
			"error": message,
		},
	}))
}
