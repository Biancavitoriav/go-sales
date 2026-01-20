package main

import (
	"log"
	"os"

	"vendas/internal/database"
	"vendas/internal/handler"
	"vendas/internal/repository"
	"vendas/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default values")
	}

	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
	}

	mongoDBName := os.Getenv("MONGODB_DATABASE")
	if mongoDBName == "" {
		mongoDBName = "go_sales"
	}

	db, err := database.ConnectMongoDB(mongoURI, mongoDBName)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	customerRepo := repository.NewCustomerRepository(db)
	productRepo := repository.NewProductRepository(db)

	customerService := service.NewCustomerService(customerRepo)
	productService := service.NewProductService(productRepo)

	r := gin.Default()

	handler.NewCustomerHandler(r, customerService)
	handler.NewProductHandler(r, productService)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Server starting on :%s...\n", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
