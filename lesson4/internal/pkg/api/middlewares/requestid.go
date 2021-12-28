package middlewares

import (
	contextUtils "backendOneLessons/lesson4/internal/pkg/context"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RequestIDMiddleware() echo.MiddlewareFunc {
	cfg := middleware.DefaultRequestIDConfig
	cfg.Generator = func() string {
		return uuid.New().String()
	}

	cfg.RequestIDHandler = func(ectx echo.Context, requestID string) {
		newCtx := contextUtils.SetRequestID(ectx.Request().Context(), requestID)
		ectx.SetRequest(ectx.Request().WithContext(newCtx))
	}

	return middleware.RequestIDWithConfig(cfg)
}
