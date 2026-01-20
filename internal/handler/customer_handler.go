package handler

import (
	"net/http"
	"strings"
	"unicode"

	"vendas/internal/service"

	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	service *service.CustomerService
}

func NewCustomerHandler(r *gin.Engine, service *service.CustomerService) {
	customersGroup := r.Group("/customers")

	customersGroup.POST("/", createCustomer(service))
	customersGroup.GET("/", listCustomers(service))
	customersGroup.GET("/:id", getCustomerByID(service))
}

func createCustomer(svc *service.CustomerService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body struct {
			Name  string `json:"name" binding:"required"`
			Phone string `json:"phone" binding:"required"`
		}

		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if len(body.Name) < 3 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "customer name must have at least 3 characters"})
			return
		}

		cleanPhone := cleanPhoneNumber(body.Phone)

		if !isValidPhone(cleanPhone) {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid phone number"})
			return
		}

		customer, err := svc.Create(body.Name, cleanPhone)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, customer)
	}
}

func listCustomers(svc *service.CustomerService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		customers, err := svc.List()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, customers)
	}
}

func getCustomerByID(svc *service.CustomerService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var uri struct {
			ID string `uri:"id" binding:"required"`
		}

		if err := ctx.ShouldBindUri(&uri); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		customer, err := svc.GetByID(uri.ID)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, customer)
	}
}

func cleanPhoneNumber(phone string) string {
	var result strings.Builder

	for _, char := range phone {
		if unicode.IsDigit(char) {
			result.WriteRune(char)
		}
	}

	return result.String()
}

func isValidPhone(phone string) bool {
	length := len(phone)
	return length >= 10 && length <= 11
}
