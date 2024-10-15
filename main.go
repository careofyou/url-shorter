package main

import (
	"fmt"

	"github.com/careofyou/url-shorter/handler"
	"github.com/careofyou/url-shorter/store"
	"github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Hello there go shortner !",
        })
    })

    r.POST("/create-short-url", func(c *gin.Context) {
        handler.CreateShortUrl(c)
    })

    r.GET("/:shortUrl", func(c *gin.Context){
        handler.HandleShortUrlRedirect(c)
    })

    store.InitializeStore()

    err := r.Run(":9808")
    if err != nil {
        panic(fmt.Sprintf("Failed to start the server - err: %v", err))
    }
}
