package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.TrustedProxies = []string{"::1"}
	r.GET("/ping", func(c *gin.Context) {
		//type test struct {
		//	weight int "json:weight"
		//}
		//var t test
		//err := c.BindJSON(t)
		//if err != nil {
		//	fmt.Println(err)
		//}
		fmt.Println(c.RemoteIP())
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
