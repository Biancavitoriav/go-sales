package handler

import (
	"net/http"

	"vendas/internal/service"

	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	service *service.CustomerService
}

func NewCustomerHandler(r *gin.Engine, service *service.CustomerService) {
	customersGroup := r.Group("/customers")
	customersGroup.POST("/", create(service))
	customersGroup.GET("/", list(service))
	customersGroup.GET("/:id", getByID(service))
}

func create(svc *service.CustomerService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body struct {
			Name string `json:"name" binding:"required"`
		}

		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		customer, err := svc.Create(body.Name)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, customer)
	}
}

func list(svc *service.CustomerService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		customers, err := svc.List()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, customers)
	}
}

func getByID(svc *service.CustomerService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var uri struct {
			ID int64 `uri:"id" binding:"required"`
		}

		if err := ctx.ShouldBindUri(&uri); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		customer, err := svc.GetByID(uri.ID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		if customer.ID == 0 {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "customer not found"})
		}

		ctx.JSON(http.StatusOK, customer)
	}
}
