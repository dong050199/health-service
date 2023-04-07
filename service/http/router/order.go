package router

import (
	"health_service/service/http/handler"

	"github.com/gin-gonic/gin"
)

type OrderRouter interface {
	Register(rg *gin.RouterGroup)
}

type orderRouter struct {
	orderHandler handler.OrderHandler
}

func NewOrderRouter(
	orderHandler handler.OrderHandler,
) OrderRouter {
	return &orderRouter{
		orderHandler: orderHandler,
	}
}

func (sr *orderRouter) Register(r *gin.RouterGroup) {
	orderGroup := r.Group("/product")
	{
		orderGroup.GET("/list", sr.orderHandler.GetList())
		orderGroup.GET("/:id", sr.orderHandler.GetDetails())
	}
}
