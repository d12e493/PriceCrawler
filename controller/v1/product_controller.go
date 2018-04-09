package controller

import (
	"net/http"
	"product-query/bo"
	"product-query/dao"
	. "product-query/global"
	"product-query/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FindProduct(context *gin.Context) {

	productId := context.Param("id")

	if productId != "" {
		Logger.Debug("productId : " + productId)
		id, err := strconv.ParseInt(productId, 10, 64)
		if err == nil {
			var dao = dao.GetMysqlDao()
			p := &model.Product{}

			dao.Where("product_id = ?", id).First(p)

			if p.ProductId != 0 {
				bo := model.GetProductBO(p)
				context.JSON(http.StatusOK, bo)
			}
		}
	} else {
		// other query condition
		productName := context.Query("product_name")
		pageable := bo.ParseContextToPage(context)

		products := &[]model.Product{}

		var dao = dao.GetMysqlDao()
		if len(productName) > 0 {
			dao = dao.Where("name LIKE ?", "%"+productName+"%")
			Logger.Debug("Query product name : " + productName)
		}
		Logger.Debug("Query pageable")
		Logger.Debug(pageable)
		dao.Limit(pageable.Size).Offset((pageable.Page - 1) * pageable.Size).Find(&products)

		context.JSON(http.StatusOK, products)
	}
}

func CreateProduct(context *gin.Context) {

	var product bo.CrawlerProductBO

	if err := context.BindJSON(&product); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "bad request parameter"})
	} else {
		var dao = dao.GetMysqlDao()
		var p = model.CreateProduct(&product)
		dao.Create(&p)
		context.JSON(http.StatusOK, gin.H{"message": "create success"})
	}
}

func UpdateProduct(context *gin.Context) {

	var product bo.CrawlerProductBO

	if err := context.BindJSON(&product); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "bad request parameter"})
	} else {

		if product.ProductId == 0 {
			context.JSON(http.StatusBadRequest, gin.H{"message": "no product id"})
		} else {
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
}

// sync crawler product.
// It'll create / update
func SyncProduct(context *gin.Context) {
	var bo bo.CrawlerProductBO

	if err := context.BindJSON(&bo); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "bad request parameter"})
	} else {

		if len(bo.Website) == 0 || len(bo.SiteProductId) == 0 {
			context.JSON(http.StatusBadRequest, gin.H{"message": "no product id"})
		} else {
			var dao = dao.GetMysqlDao()

			p := &model.Product{}

			dao.Where("website = ? and site_product_id = ?", bo.Website, bo.SiteProductId).First(p)

			if p.ProductId != 0 {
				// update
				dao.Model(&p).Updates(bo)

				context.JSON(http.StatusOK, gin.H{"message": "update success"})
			} else {
				// create new one
				p := model.CreateProduct(&bo)
				dao.Create(&p)

				context.JSON(http.StatusOK, gin.H{"message": "create success"})
			}
		}
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
