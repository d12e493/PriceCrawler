package bo

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type PageableBO struct {
	Page   int
	Size   int
	Start  int
	Orders []Order
}

type Order struct {
	Name      string
	Direction string
}

func ParseContextToPage(context *gin.Context) PageableBO {
	pageable := PageableBO{
		Page:  1,
		Size:  10,
		Start: 0,
	}

	if contextPage := context.Query("page"); len(contextPage) > 0 {
		intPage, _ := strconv.Atoi(contextPage)
		pageable.Page = intPage
	}

	if contextSize := context.Query("size"); len(contextSize) > 0 {
		intSize, _ := strconv.Atoi(contextSize)
		pageable.Size = intSize
	}

	pageable.Start = (pageable.Page - 1) * pageable.Size
	return pageable
}
