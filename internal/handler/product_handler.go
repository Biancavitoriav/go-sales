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
	handler := &ProductHandler{service: service}

	r.POST("/", createProduct(service))
	r.GET("/", listProducts(handler))
	r.GET("/:id", getProductByID(handler))
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

		product, err := svc.Create(body.Name, body.Price)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, product)
	}
}

func listProducts(h *ProductHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		products, err := h.service.List()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, products)
	}
}

func getProductByID(h *ProductHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var uri struct {
			ID int64 `uri:"id" binding:"required"`
		}

		if err := ctx.ShouldBindUri(&uri); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		product, err := h.service.GetByID(uri.ID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if product.ID == 0 {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
			return
		}

		ctx.JSON(http.StatusOK, product)
	}
}
