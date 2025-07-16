package dtos

type CreateBrandRequest struct {
	Name string `json:"name" validate:"required"`
}

type ProductResponse struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	Price float64 `json:"price"`
	Quantity int `json:"quantity"`
}

type BrandResponse struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Products []ProductResponse `json:"products"`
}
