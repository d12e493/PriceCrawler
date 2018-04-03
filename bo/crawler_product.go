package bo

type CrawlerProductBO struct {
	ProductId     int    `json:"ProductId"`
	Name          string `json:"Name" binding:"required"`
	Website       string `json:"WebSite" binding:"required"`
	SiteProductId string `json:"SiteProductId" binding:"required"`
	Link          string `json:"Link" binding:"required"`
	Price         int    `json:"Price" binding:"required"`
	Picture       string `json:"Picture" binding:"required"`
	Description   string `json:"Description"`
}
