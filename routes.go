package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func testGet(c *gin.Context) {
	c.JSON(http.StatusOK, "success")
}

func insertDoctorDB(c *gin.Context) {
	c.JSON(http.StatusOK, insert())
}

func queryDoctorDB(c *gin.Context) {
	c.JSON(http.StatusOK, query())
}

func getStats(c *gin.Context) {
	c.JSON(http.StatusOK, Stats())
}
