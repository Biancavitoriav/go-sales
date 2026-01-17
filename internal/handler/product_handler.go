package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"vendas/internal/usecase"
)

type ProductHandler struct {
	create *usecase.CreateProduct
}

func NewProductHandler(p *usecase.CreateProduct) *ProductHandler {
	return &ProductHandler{create: p}
}

func (h *ProductHandler) Create(ctx *gin.Context) {
	var body struct {
		Name  string  `json:"name"`
		Price float64 `json:"price"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := h.create.Execute(body.Name, body.Price)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, product)
}
