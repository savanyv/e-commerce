package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/savanyv/e-commerce/internal/database"
	"github.com/savanyv/e-commerce/internal/delivery/handlers"
	"github.com/savanyv/e-commerce/internal/repository"
	"github.com/savanyv/e-commerce/internal/usecase"
)

func productRoutes(e *echo.Group) {
	repo := repository.NewProductRepository(database.DB)
	brandRepo := repository.NewBrandRepository(database.DB)
	usecase := usecase.NewProductRepository(repo, brandRepo)
	handler := handlers.NewProductHandler(usecase)

	e.POST("/products", handler.CreateHandler)
	e.GET("/products", handler.GetAllProducts)
}
