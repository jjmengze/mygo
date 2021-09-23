package http

import (
	"github.com/labstack/echo/v4"
)

func SetRoutes(e *echo.Echo, h *Handler) {
	basic := e.Group("")
	basic.GET("/", func(c echo.Context) error {
		return c.NoContent(200)
	})

	v1APIPublic := e.Group("/public/apis/v1")
	{
		v1APIPublic.GET("/health", func(c echo.Context) error {
			return c.NoContent(200)
		})

		v1APIPublic.POST("/auth/register", h.registerEndpoint)

		v1APIPublic.POST("/auth/login", h.loginEndpoint)
	}
}
