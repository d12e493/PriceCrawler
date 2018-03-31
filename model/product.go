package model

import (
	"product-query/bo"
)

type Product struct {
	ProductId   int `gorm:"primary_key"`
	Name        string
	Price       int
	Description string
}

func (Product) TableName() string {
	return "product"
}

func CreateProduct(bo *bo.ProductBO) Product {
	var p = Product{}

	p.ProductId = bo.ProductId
	p.Name = bo.Name
	p.Price = bo.Price
	p.Description = bo.Description

	return p
}

func GetProductBO(p *Product) bo.ProductBO {
	bo := bo.ProductBO{}
	bo.ProductId = p.ProductId
	bo.Name = p.Name
	bo.Price = p.Price
	bo.Description = p.Description
	return bo
}
