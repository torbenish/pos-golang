package dto

type CreateProductImput struct {
	Name string `json:"name"`
	Price float64 `json:"price"`
}