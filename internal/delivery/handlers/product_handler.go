package handlers

import (
	"net/http"
	"strconv"

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

func (h *ProductHandler) GetAllProducts(c echo.Context) error {
	pageStr := c.QueryParam("page")
	limitStr := c.QueryParam("limit")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}

	products, total, err := h.usecase.GetAllProduct(page, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data":    products,
		"total":   total,
		"message": "products retrieved successfully",
	})
}

func (h *ProductHandler) GetByIDProduct(c echo.Context) error {
	idStr := c.Param("id")
	ID, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	product, err := h.usecase.GetByIDProduct(ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data":    product,
		"message": "product retrieved successfully",
	})
}

func (h *ProductHandler) UpdateProduct(c echo.Context) error {
	idStr := c.Param("id")
	ID, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	var req *dtos.UpdateProductRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	if err := h.validator.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	if err := h.usecase.UpdateProduct(ID, req); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "product updated successfully",
	})
}

func (h *ProductHandler) DeleteProduct(c echo.Context) error {
	idStr := c.Param("id")
	ID, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	if err := h.usecase.DeleteProduct(ID); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "product deleted successfully",
	})
}
