package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"vendas/internal/handler"
	"vendas/internal/repository"
	"vendas/internal/service"
)

func main() {
	// Inicializa os repositórios
	customerRepo := repository.NewCustomerRepository()
	productRepo := repository.NewProductRepository()

	// Inicializa os serviços
	customerService := service.NewCustomerService(customerRepo)
	productService := service.NewProductService(productRepo)

	r := gin.Default()

	// Inicializa os handlers
	handler.NewCustomerHandler(r, customerService)
	handler.NewProductHandler(r, productService)

	// Inicia o servidor
	log.Println("Server starting on :8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
