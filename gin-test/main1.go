package main

// 解决 404页面不需要鉴权中间件的问题
// 增加分组
//func main() {
//	e := gin.Default()
//
//	e.NoRoute(func(c *gin.Context) {
//		c.JSON(404, "Page not Found.")
//	})
//
//	group := e.Group("/")
//	group.Use(func(c *gin.Context) {
//		if c.Request.URL.Path != "/not-login" {
//			// check login
//			c.Abort()
//			c.JSON(200, "need login")
//			return
//		}
//	})
//	group.GET("/not-login", func(c *gin.Context) {
//		c.JSON(200, "not-login")
//	})
//	group.GET("/required-login", func(c *gin.Context) {
//		c.JSON(200, "required-login")
//	})
//	_ = e.Run()
//}
