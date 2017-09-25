package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HomeIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"home": "index"})
}

func HomeShow(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"home": "show"})
}
