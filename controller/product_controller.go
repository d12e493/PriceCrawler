package controller

import (
	"fmt"
	"net/http"
	"product-query/bo"
	"product-query/dao"
	"product-query/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FindProduct(context *gin.Context) {

	productId := context.Param("id")

	if productId != "" {
		fmt.Println("productId : " + productId)
		id, err := strconv.ParseInt(productId, 10, 64)
		if err == nil {
			var dao = dao.GetMysqlDao()
			p := &model.Product{}

			dao.Where("product_id = ?", id).First(p)

			if p.ProductId != 0 {
				bo := model.GetProductBO(p)
				context.JSON(http.StatusBadRequest, bo)
			}
		}
	} else {
		// todo other query condition
	}
}

func CreateProduct(context *gin.Context) {

	var product bo.ProductBO

	if err := context.BindJSON(&product); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "bad request parameter"})
	} else {
		var dao = dao.GetMysqlDao()
		var p = model.CreateProduct(&product)
		dao.Create(p)
		context.JSON(http.StatusOK, gin.H{"message": "create success"})
	}
}

func UpdateProduct(context *gin.Context) {

	var product bo.ProductBO

	if err := context.BindJSON(&product); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "bad request parameter"})
	} else {

		if product.ProductId == 0 {
			context.JSON(http.StatusBadRequest, gin.H{"message": "no product id"})
		}

		var dao = dao.GetMysqlDao()

		var p = &model.Product{}
		p.ProductId = product.ProductId
		dao.First(&p)
		p.Name = product.Name
		p.Price = product.Price
		p.Description = product.Description
		dao.Save(&p)
		context.JSON(http.StatusOK, gin.H{"message": "update success"})
	}
}

func DeleteProduct(context *gin.Context) {
	productId := context.Param("id")

	if len(productId) > 0 {
		id, err := strconv.ParseInt(productId, 10, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"message": "no product id"})
		} else {
			var dao = dao.GetMysqlDao()
			dao.Where("product_id = ?", id).Delete(&model.Product{})
			context.JSON(http.StatusOK, gin.H{"message": "delete success"})
		}
	} else {
		context.JSON(http.StatusBadRequest, gin.H{"message": "no product id"})
	}
}
