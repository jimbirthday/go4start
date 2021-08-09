package main

// 解决gin 绑定结构体响应结构问题
// 转换错误类型,自定义响应体结构

//
//import (
//	"errors"
//	"fmt"
//	"github.com/gin-test-gonic/gin-test"
//	"github.com/gin-test-gonic/gin-test/binding"
//	"github.com/go-playground/validator/v10"
//	"net/http"
//)
//
//type User struct {
//	Email string `json:"email" binding:"required,bookabledate"`
//	Name  string `json:"name" binding:"required"`
//}
//
//func main() {
//	route := gin-test.Default()
//
//	var bookableDate validator.Func = func(fl validator.FieldLevel) bool {
//		_, ok := fl.Field().Interface().(string)
//		return ok
//	}
//
//	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
//		v.RegisterValidation("bookabledate", bookableDate)
//	}
//	route.POST("/user", validateUser)
//	route.Run(":8085")
//}
//func validateUser(c *gin-test.Context) {
//	var u User
//	if err := c.ShouldBindJSON(&u); err == nil {
//		c.JSON(http.StatusOK, gin-test.H{"message": "User validation successful."})
//	} else {
//		var verr validator.ValidationErrors
//		if errors.As(err, &verr) {
//			c.JSON(http.StatusBadRequest, gin-test.H{"errors": Simple(verr)})
//			return
//		} else {
//			c.JSON(http.StatusBadRequest, gin-test.H{"errors": err.Error()})
//		}
//	}
//}
//func Simple(verr validator.ValidationErrors) map[string]string {
//	errs := make(map[string]string)
//	for _, f := range verr {
//		err := f.ActualTag()
//		if f.Param() != "" {
//			err = fmt.Sprintf("%s=%s", err, f.Param())
//		}
//		errs[f.Field()] = err
//	}
//	return errs
//}
