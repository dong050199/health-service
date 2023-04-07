package handler

import (
	"fmt"
	"health_service/service/model/entity"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type OrderHandler struct {
}

func NewOrderHandler() OrderHandler {
	return OrderHandler{}
}

func composeConfigHandlerName(name string) string {
	return fmt.Sprintf("config_handler_%s", name)
}

type Product struct {
	Products []entity.Products `json:"products"`
}

var productsMock = Product{
	Products: []entity.Products{
		{
			ID:               1,
			CategoryID:       1,
			Price:            10000,
			Quantity:         10,
			Name:             "Sản phẩm 1",
			Description:      "Mô tả sản phẩm dàiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiii",
			ShortDescription: "Mô tả ngắn",
			ImageURL:         "https://thinkingvn.b-cdn.net/wp-content/uploads/2021/02/phat-trien-y-tuong-san-pham-moi.jpg",
			Category: entity.Categories{
				Name:             "Ngành hàng 1",
				Description:      "Mô tả ngành hàng 1",
				ShortDescription: "Mô tả ngắn NH1",
			},
		},
		{
			ID:               2,
			CategoryID:       1,
			Price:            10000,
			Quantity:         10,
			Name:             "Sản phẩm 2",
			Description:      "Mô tả sản phẩm dàiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiii",
			ShortDescription: "Mô tả ngắn",
			ImageURL:         "https://thinkingvn.b-cdn.net/wp-content/uploads/2021/02/phat-trien-y-tuong-san-pham-moi.jpg",
			Category: entity.Categories{
				Name:             "Ngành hàng 1",
				Description:      "Mô tả ngành hàng 1",
				ShortDescription: "Mô tả ngắn NH1",
			},
		},
		{
			ID:               3,
			CategoryID:       1,
			Price:            10000,
			Quantity:         10,
			Name:             "Sản phẩm 3",
			Description:      "Mô tả sản phẩm dàiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiii",
			ShortDescription: "Mô tả ngắn",
			ImageURL:         "https://thinkingvn.b-cdn.net/wp-content/uploads/2021/02/phat-trien-y-tuong-san-pham-moi.jpg",
			Category: entity.Categories{
				Name:             "Ngành hàng 1",
				Description:      "Mô tả ngành hàng 1",
				ShortDescription: "Mô tả ngắn NH1",
			},
		},
		{
			ID:               4,
			CategoryID:       1,
			Price:            10000,
			Quantity:         10,
			Name:             "Sản phẩm 4",
			Description:      "Mô tả sản phẩm dàiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiii",
			ShortDescription: "Mô tả ngắn",
			ImageURL:         "https://thinkingvn.b-cdn.net/wp-content/uploads/2021/02/phat-trien-y-tuong-san-pham-moi.jpg",
			Category: entity.Categories{
				Name:             "Ngành hàng 1",
				Description:      "Mô tả ngành hàng 1",
				ShortDescription: "Mô tả ngắn NH1",
			},
		},
		{
			ID:               5,
			CategoryID:       1,
			Price:            10000,
			Quantity:         10,
			Name:             "Sản phẩm 5",
			Description:      "Mô tả sản phẩm dàiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiii",
			ShortDescription: "Mô tả ngắn",
			ImageURL:         "https://thinkingvn.b-cdn.net/wp-content/uploads/2021/02/phat-trien-y-tuong-san-pham-moi.jpg",
			Category: entity.Categories{
				Name:             "Ngành hàng 1",
				Description:      "Mô tả ngành hàng 1",
				ShortDescription: "Mô tả ngắn NH1",
			},
		},
	},
}

func (h *OrderHandler) GetList() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, productsMock)
	}
}

func (h *OrderHandler) GetDetails() gin.HandlerFunc {
	return func(c *gin.Context) {
		productID := c.Param("id")
		var dataResponse *entity.Products
		for _, product := range productsMock.Products {
			if cast.ToInt(productID) == product.ID {
				dataResponse = &product
			}
		}
		c.JSON(http.StatusOK, dataResponse)
	}
}
