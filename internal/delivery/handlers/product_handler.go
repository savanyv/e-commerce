package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	dtos "github.com/savanyv/e-commerce/internal/dto"
	"github.com/savanyv/e-commerce/internal/helpers"
	"github.com/savanyv/e-commerce/internal/usecase"
)

type ProductHandler struct {
	usecase usecase.ProductUsecase
	validator *helpers.CustomValidator
}

func NewProductHandler(usecase usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{
		usecase: usecase,
		validator: helpers.NewValidator(),
	}
}

func (h *ProductHandler) CreateHandler(c echo.Context) error {
	var product *dtos.CreateProductRequest
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	if err := h.validator.Validate(product); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	if err := h.usecase.CreateProduct(product); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "product created successfully",
	})
}
