package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	dtos "github.com/savanyv/e-commerce/internal/dto"
	"github.com/savanyv/e-commerce/internal/helpers"
	"github.com/savanyv/e-commerce/internal/usecase"
)

type BrandHandler struct {
	usecase usecase.BrandUsecase
	validator *helpers.CustomValidator
}

func NewBrandHandler(usecase usecase.BrandUsecase) *BrandHandler {
	return &BrandHandler{
		usecase: usecase,
		validator: helpers.NewValidator(),
	}
}

func (h *BrandHandler) Create(c echo.Context) error {
	req := &dtos.CreateBrandRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	if err := h.validator.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	if err := h.usecase.Create(req); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "brand created successfully",
	})
}

func (h *BrandHandler) Delete(c echo.Context) error {
	idStr := c.Param("id")
	ID, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	if err := h.usecase.Delete(ID); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "brand deleted successfully",
	})
}

func (h *BrandHandler) GetAllBrands(c echo.Context) error {
	brands, err := h.usecase.GetAllBrands()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": brands,
		"message": "brands retrieved successfully",
	})
}
