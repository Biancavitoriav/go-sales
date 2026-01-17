package main

import (
	"github.com/gin-gonic/gin"

	"vendas/internal/handler"
	"vendas/internal/infra/memory"
	"vendas/internal/usecase"
)

func main() {
	r := gin.Default()

	customerRepo := memory.NewCustomerRepo()
	createCustomer := usecase.NewCreateCustomer(customerRepo)
	customerHandler := handler.NewCustomerHandler(createCustomer)

	productRepo := memory.NewProductRepo()
	createProduct := usecase.NewCreateProduct(productRepo)
	productHandler := handler.NewProductHandler(createProduct)

	r.POST("/customers", customerHandler.Create)
	r.POST("/products", productHandler.Create)

	r.Run(":8080")
}
