package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"vendas/internal/usecase"
)

type CustomerHandler struct {
	create *usecase.CreateCustomer
}

func NewCustomerHandler(c *usecase.CreateCustomer) *CustomerHandler {
	return &CustomerHandler{create: c}
}

func (h *CustomerHandler) Create(ctx *gin.Context) {
	var body struct {
		Name string `json:"name"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customer, err := h.create.Execute(body.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, customer)
}
