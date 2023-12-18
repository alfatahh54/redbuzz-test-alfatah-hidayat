package routes

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

type MainRouteType struct {
	Engine *gin.RouterGroup
}

var MainRoute MainRouteType

func Router(r *gin.RouterGroup) {
	MainRoute.Engine = r
}

func (r *MainRouteType) NewRoute(method string, endpoint string, handlers ...gin.HandlerFunc) {
	for {
		if r.Engine != nil {
			break
		}
	}
	r.Engine.Handle(method, endpoint, handlers...)
}

func ReverseProxy(c *gin.Context) {
	remote, _ := url.Parse("http://localhost:3000")
	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.Director = func(req *http.Request) {
		req.Header = c.Request.Header
		req.Host = remote.Host
		req.URL = c.Request.URL
		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
	}

	proxy.ServeHTTP(c.Writer, c.Request)
}
