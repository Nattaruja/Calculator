package router

import (
	"app/customer"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func GenerateRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers"},
	}))
	customerV1 := router.Group("/v1/customer")
	customerV1.Use()
	{
		customerV1.POST("/calc", customer.TestCalHandler)
		customerV1.GET("/customers", customer.TestCalHandler)
	}

	return router
}
