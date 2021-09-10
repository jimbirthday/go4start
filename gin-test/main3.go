package main

//
//import (
//	"github.com/gin-gonic/gin"
//	"github.com/gin-gonic/gin/binding"
//	"net/http"
//)
//
////解决验证器关闭后,想再次开启的问题
//type User struct {
//	Email string `json:"email" form:"e-mail" binding:"required"`
//	Name  string `json:"name" binding:"required"`
//	Org   string `json:"org" binding:"required"`
//}
//
//func main() {
//	route := gin.Default()
//	route.POST("/user", validateUser)
//	route.GET("/validate", validateUsers)
//	route.Run(":8085")
//}
//
//func validateUser(c *gin.Context) {
//	var u User
//	if err := c.ShouldBindJSON(&u); err == nil {
//		c.JSON(http.StatusOK, gin.H{"message": "User validation successful."})
//	} else {
//		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
//		return
//	}
//}
//func validateUsers(c *gin.Context) {
//	var u []User
//	backup := binding.Validator
//	defer func() { binding.Validator = backup }()
//	gin.DisableBindValidation()
//	//Note: As soon as this route gets executed the validateUser route will stop validating the json
//	if err := c.ShouldBindJSON(&u); err == nil {
//		c.JSON(http.StatusOK, gin.H{"message": "User validation successful."})
//	} else {
//		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
//		return
//	}
//}
