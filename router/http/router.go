package router

import (
	"itmx-test/domain/customers"
	customerRoutes "itmx-test/router/http/customers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewHTTPHandler(customveSvc customers.CustomerService) http.Handler {
	router := gin.Default()

	api := router.Group("/api")

	customerGroup := api.Group("/customers")
	customerRoutes.NewRoutesFactory(customerGroup)(customveSvc)
	return router
}
