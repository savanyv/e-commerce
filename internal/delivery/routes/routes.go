package routes

import "github.com/labstack/echo/v4"

func SetupRoutes(e *echo.Echo) error {
	api := e.Group("/api")

	brandRoutes(api)

	return nil
}
