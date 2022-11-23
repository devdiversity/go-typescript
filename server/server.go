package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Typescript: interface
type HTTPResponse struct {
	Data  interface{} `json:"data"`
	Error interface{} `json:"error"`
}

// Typescript: interface
type FormRequest struct {
	Req   string `json:"req"`
	Count int    `json:"count"`
}

// Typescript: interface
type FormResponse struct {
	Test string `json:"test"`
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:5173")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, cache-control, Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS", c.Request.Method, c.Request.RequestURI)
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

func Server() {
	r := gin.Default()

	r.Use(CORSMiddleware())

	// Typescript: TSEndpoint= path=/ping; name=ping; method=GET; response=string
	r.GET("/ping", func(c *gin.Context) {
		response := HTTPResponse{Data: "pong", Error: nil}
		c.JSON(http.StatusOK, response)
	})

	// Typescript: TSEndpoint= path=/test/:len/*lat;  name=test; method=GET; response=string
	r.GET("/test/:len/*lat", func(c *gin.Context) {
		p := c.Param("len")
		response := HTTPResponse{Data: p, Error: nil}
		c.JSON(http.StatusOK, response)
	})

	// Typescript: TSEndpoint= path=/testpost;  name=testPost; method=POST; request=FormRequest; response=string
	r.POST("/post", func(c *gin.Context) {
		p := c.Param("len")
		response := HTTPResponse{Data: p, Error: nil}
		c.JSON(http.StatusOK, response)
	})

	// Typescript: TSEndpoint= path=/postTest;  name=postTest; method=POST; request=FormRequest; response=FormResponse
	r.POST("/postTest", func(c *gin.Context) {
		var requestBody FormRequest
		if err := c.BindJSON(&requestBody); err != nil {
			response := HTTPResponse{Data: nil, Error: "wrongData"}
			c.JSON(http.StatusOK, response)
		}
		response := HTTPResponse{Data: FormResponse{Test: fmt.Sprintf("%d", requestBody.Count)}, Error: nil}
		c.JSON(http.StatusOK, response)
	})
	r.Run()
}
