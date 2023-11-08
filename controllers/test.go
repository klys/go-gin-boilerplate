package controllers

import (
	"net/http"
  	"github.com/gin-gonic/gin"
	"strconv"
	"fmt"
)

import "gintest/accesors"
import "gintest/models"

func Test (c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "testing",
	  })
}

func Ping (c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	  })
}

func TestPost (c *gin.Context) {
	var newTest models.Test
	if err := c.BindJSON(&newTest); err != nil {
        return
    }
	newTest.ID = accesors.TestCreate(newTest.Test)
	
	c.IndentedJSON(http.StatusCreated, newTest)
}

func TestGet (c *gin.Context) {
	id := c.Param("id")
	test := accesors.TestConsult(id)
	c.IndentedJSON(http.StatusOK, test)
}

func TestPut (c *gin.Context) {
	id := c.Param("id")
	var updateTest models.Test
	if err := c.BindJSON(&updateTest); err != nil {
        return
    }
	value, err := strconv.ParseInt(id, 10, 64)
	if (err != nil) {
		fmt.Println(err)
        return
	}
	updateTest.ID = value
	result := accesors.TestUpdate(updateTest)
	if (result == true) {
		c.IndentedJSON(http.StatusOK, updateTest)
		
	} else {
		c.JSON(http.StatusOK, gin.H{
			"response": "error",
		  })
	}
	
}

func TestDelete (c *gin.Context) {
	id := c.Param("id")
	result := accesors.TestDelete(id)
	if (result == true) {
		c.JSON(http.StatusOK, gin.H{
			"result": "true",
		  })
	} else {
		c.JSON(http.StatusOK, gin.H{
			"result": "false",
		  })
	}
	
}