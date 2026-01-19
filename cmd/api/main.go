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

	// Inicializa os handlers
	customerHandler := handler.NewCustomerHandler(customerService)
	productHandler := handler.NewProductHandler(productService)

	// Configura o router
	r := gin.Default()

	// Rotas de clientes
	r.POST("/customers", customerHandler.Create)
	r.GET("/customers", customerHandler.List)
	r.GET("/customers/:id", customerHandler.GetByID)

	// Rotas de produtos
	r.POST("/products", productHandler.Create)
	r.GET("/products", productHandler.List)
	r.GET("/products/:id", productHandler.GetByID)

	// Inicia o servidor
	log.Println("Server starting on :8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
