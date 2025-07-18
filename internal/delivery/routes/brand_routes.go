package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/savanyv/e-commerce/internal/database"
	"github.com/savanyv/e-commerce/internal/delivery/handlers"
	"github.com/savanyv/e-commerce/internal/repository"
	"github.com/savanyv/e-commerce/internal/usecase"
)

func brandRoutes(e *echo.Group) {
	repo := repository.NewBrandRepository(database.DB)
	usecase := usecase.NewBrandUsecase(repo)
	handler := handlers.NewBrandHandler(usecase)

	e.POST("/brands", handler.CreateBrand)
	e.DELETE("/brands/:id", handler.DeleteBrand)
	e.GET("/brands", handler.GetAllBrands)
}
