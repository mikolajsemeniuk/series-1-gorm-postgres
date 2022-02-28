package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Order interface {
	All(*gin.Context)
	Find(*gin.Context)
	Add(*gin.Context)
	Update(*gin.Context)
	Remove(*gin.Context)
}

type order struct{}

func (*order) All(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "All",
	})
}

func (*order) Find(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "Find",
	})
}

func (*order) Add(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "Add",
	})
}

func (*order) Update(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "Update",
	})
}

func (*order) Remove(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "Remove",
	})
}

func NewOrder() Order {
	return &order{}
}
