package utils

import (
	"github.com/MgHtinLynn/final-year-project-mcc/service/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GeneratePaginationFromRequest(c *gin.Context) models.Pagination {
	// Initializing default
	limit := 10
	page := 1
	sort := "id desc"
	query := c.Request.URL.Query()

	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)
			break
		case "page":
			page, _ = strconv.Atoi(queryValue)
			break
		case "sort":
			sort = queryValue
			break

		}
	}
	return models.Pagination{
		Limit: limit,
		Page:  page,
		Sort:  sort,
	}

}
