package controllers

import (
  "net/http"
  "github.com/gin-gonic/gin"
)

func HomeIndex(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{ "home": "index" })
}

func HomeShow(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{ "home": "show" })
}
