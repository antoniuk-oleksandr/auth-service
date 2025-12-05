package errorhandler

import (
	"errors"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

func New(ctx *fiber.Ctx, err error) error {
	status := getHTTPStatus(err)

	return ctx.Status(status).JSON(fiber.Map{
		"error":     err.Error(),
		"timestamp": time.Now(),
		"path":      ctx.Path(),
	})
}

func getHTTPStatus(err error) int {
	for domainErr, status := range errorStatusMap {
		if errors.Is(err, domainErr) {
			return status
		}
	}

	var fiberErr *fiber.Error
	if errors.As(err, &fiberErr) {
		return fiberErr.Code
	}

	return http.StatusInternalServerError
}
