package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func HomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}

func PostHomePage(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(500, gin.H{
			"error": "Failed to read request body",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": string(body),
	})
}

func QueryString(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")

	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}

func PathParameterString(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")

	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}

func main() {
	fmt.Sprintf("hello world!")

	r := gin.Default()
	r.GET("/", HomePage)
	r.POST("/", PostHomePage)
	r.GET("/query", QueryString)                    // http://localhost:8080/query?name=sadık&21
	r.GET("/query/:name/:age", PathParameterString) // http://localhost:8080/query/sadık/21
	r.Run()
}
