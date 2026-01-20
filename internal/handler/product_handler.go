package handler

import (
	"net/http"

	"vendas/internal/service"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(r *gin.Engine, service *service.ProductService) {
	productsGroup := r.Group("/products")

	productsGroup.POST("/", createProduct(service))
	productsGroup.GET("/", listProducts(service))
	productsGroup.GET("/:id", getProductByID(service))
}

func createProduct(svc *service.ProductService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body struct {
			Name  string  `json:"name" binding:"required"`
			Price float64 `json:"price" binding:"required"`
		}

		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if len(body.Name) < 3 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "product name must have at least 3 characters"})
			return
		}

		if body.Price < 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "price cannot be negative"})
			return
		}

		product, err := svc.Create(body.Name, body.Price)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, product)
	}
}

func listProducts(svc *service.ProductService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		products, err := svc.List()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, products)
	}
}

func getProductByID(svc *service.ProductService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var uri struct {
			ID string `uri:"id" binding:"required"`
		}

		if err := ctx.ShouldBindUri(&uri); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		product, err := svc.GetByID(uri.ID)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, product)
	}
}
