package main

// import (
// 	"fmt"

// 	"github.com/gin-gonic/gin"
// )

// import "fmt"

import "github.com/eulbyvan/go-enigma-laundry/delivery"

func main() {
	delivery.Server().Run()
	// fmt.Print("Hello Docker!")
	// route := gin.Default()
	// route.GET("/ping", func(ctx *gin.Context) {
	// 	ctx.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	// err := route.Run(":8081")
	// if err != nil {
	// 	return
	// }
}
