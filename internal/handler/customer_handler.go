package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"vendas/internal/service"
)

// CustomerHandler lida com as requisições HTTP relacionadas a clientes
type CustomerHandler struct {
	service *service.CustomerService
}

// NewCustomerHandler cria uma nova instância do handler
func NewCustomerHandler(service *service.CustomerService) *CustomerHandler {
	return &CustomerHandler{service: service}
}

// Create cria um novo cliente
func (h *CustomerHandler) Create(ctx *gin.Context) {
	var body struct {
		Name string `json:"name" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customer, err := h.service.Create(body.Name)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, customer)
}

// List retorna todos os clientes
func (h *CustomerHandler) List(ctx *gin.Context) {
	customers, err := h.service.List()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, customers)
}

// GetByID busca um cliente por ID
func (h *CustomerHandler) GetByID(ctx *gin.Context) {
	var uri struct {
		ID int64 `uri:"id" binding:"required"`
	}

	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customer, err := h.service.GetByID(uri.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if customer.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "customer not found"})
		return
	}

	ctx.JSON(http.StatusOK, customer)
}
