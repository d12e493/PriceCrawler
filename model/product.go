package model

import (
	"product-query/bo"

	"github.com/jinzhu/copier"
)

type Product struct {
	ProductId     int `gorm:"primary_key"`
	Name          string
	Website       string
	SiteProductId string
	Link          string
	Price         int
	Picture       string
	Description   string
}

func (Product) TableName() string {
	return "product"
}

func CreateProduct(bo *bo.CrawlerProductBO) Product {
	var p = Product{}
	copier.Copy(&p, bo)
	return p
}

func GetProductBO(p *Product) bo.CrawlerProductBO {
	var bo = bo.CrawlerProductBO{}
	copier.Copy(&bo, p)
	return bo
}
