package dtos

type CreateBrandRequest struct {
	Name string `json:"name" validate:"required"`
}

type ResponseProduct struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Price float64 `json:"price"`
	Quantity int `json:"quantity"`
}

type BrandResponse struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Products []*ResponseProduct `json:"products"`
}
