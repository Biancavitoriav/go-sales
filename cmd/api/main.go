package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"vendas/internal/handler"
	"vendas/internal/repository"
	"vendas/internal/service"
)

func main() {
	customerRepo := repository.NewCustomerRepository()
	productRepo := repository.NewProductRepository()

	customerService := service.NewCustomerService(customerRepo)
	productService := service.NewProductService(productRepo)

	r := gin.Default()

	handler.NewCustomerHandler(r, customerService)
	handler.NewProductHandler(r, productService)

	log.Println("Server starting on :8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
