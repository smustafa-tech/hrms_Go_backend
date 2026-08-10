package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetQueriesUnreadCount(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"count": 0})
}

func GetMyQueries(c *gin.Context) {
	c.JSON(http.StatusOK, []interface{}{})
}

func GetAllQueries(c *gin.Context) {
	c.JSON(http.StatusOK, []interface{}{})
}

func SubmitQuery(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "query submitted"})
}

func ReplyToQuery(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "reply sent"})
}

func CloseQuery(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "query closed"})
}

func DeleteQuery(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "query deleted"})
}
