package dtos

type CreateProductRequest struct {
	Name string `json:"name" validate:"required"`
	Price float64 `json:"price" validate:"required"`
	Quantity int `json:"quantity" validate:"required"`
	IDBrand uint `json:"id_brand" validate:"required"`
}

type UpdateProductRequest struct {
	Name string `json:"name"`
	Price float64 `json:"price"`
	Quantity int `json:"quantity"`
	IDBrand uint `json:"id_brand"`
}

type BrandSimple struct {
	ID uint `json:"id"`
	Name string `json:"name"`
}

type ProductResponse struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Price float64 `json:"price"`
	Quantity int `json:"quantity"`
	Brand BrandSimple `json:"brand"`
}
