package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("foo", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"foo":    "foo",
			"foofoo": "foofoo",
		})
	})
	router.GET("bar", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"bar":  "bar",
			"beer": "beer",
		})
	})
	router.Run(":3000")
}
