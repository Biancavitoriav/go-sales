package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"vendas/internal/service"
)

// ProductHandler lida com as requisições HTTP relacionadas a produtos
type ProductHandler struct {
	service *service.ProductService
}

// NewProductHandler cria uma nova instância do handler
func NewProductHandler(service *service.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

// Create cria um novo produto
func (h *ProductHandler) Create(ctx *gin.Context) {
	var body struct {
		Name  string  `json:"name" binding:"required"`
		Price float64 `json:"price" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := h.service.Create(body.Name, body.Price)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, product)
}

// List retorna todos os produtos
func (h *ProductHandler) List(ctx *gin.Context) {
	products, err := h.service.List()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, products)
}

// GetByID busca um produto por ID
func (h *ProductHandler) GetByID(ctx *gin.Context) {
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
