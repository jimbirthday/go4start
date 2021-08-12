package gtest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestEnsureAuthRedirectsToLogin(t *testing.T) {
	cookie := http.Cookie{
		Name:       "",
		Value:      "",
		Path:       "/login",
		Domain:     "",
		Expires:    time.Time{},
		RawExpires: "",
		MaxAge:     0,
		Secure:     false,
		HttpOnly:   false,
		SameSite:   0,
		Raw:        "",
		Unparsed:   nil,
	}
	w := httptest.NewRecorder()
	ctx, engine := gin.CreateTestContext(w)
	ctx.Set("UserData", cookie)
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, "/", nil)

	// What do I do with `ctx`? Is there a way to inject this into my test?

	engine.Use(ensureAuth())
	engine.ServeHTTP(w, req)

	assert.Equal(t, 302, w.Result().StatusCode)
	assert.Equal(t, "/login", w.Result().Header.Get("Path"))
}

func ensureAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		get, exists := c.Get("UserData")
		if !exists {
			fmt.Println("UserData not exists")
			c.Abort()
			return
		}
		cookie := get.(http.Cookie)
		path := cookie.Path
		fmt.Println("cookie path is", path)
		if c.FullPath() != "/login" {
			c.Abort()
			c.Redirect(302, "/login")
			return
		}
		c.Next()
	}
}
