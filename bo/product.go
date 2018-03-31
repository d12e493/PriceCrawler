package bo

type ProductBO struct {
	ProductId   int    `json:"productId"`
	Name        string `json:"name" binding:"required"`
	Price       int    `json:"price" binding:"required"`
	Description string `json:"description"`
}
