package routes

import (
	"net/http"
	"net/http/httputil"

	"github.com/gin-gonic/gin"
)

func ReverseProxy() gin.HandlerFunc {
	target := "localhost:5000"

	return func(c *gin.Context) {
		director := func(req *http.Request) {
			req.URL.Scheme = "http"
			req.URL.Host, req.Host = target, target
		}
		proxy := &httputil.ReverseProxy{Director: director}
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

func ExpressReverseProxy(r *gin.RouterGroup) {
	r.GET("/", ReverseProxy())
}
