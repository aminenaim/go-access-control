package core

import "github.com/gin-gonic/gin"

type HttpMethod string

/** Main HTTP methods */
const (
	GET     HttpMethod = "GET"
	POST    HttpMethod = "POST"
	PUT     HttpMethod = "PUT"
	PATCH   HttpMethod = "PATCH"
	DELETE  HttpMethod = "DELETE"
	OPTIONS HttpMethod = "OPTIONS"
	HEAD    HttpMethod = "HEAD"
	TRACE   HttpMethod = "TRACE"
	CONNECT HttpMethod = "CONNECT"
)

/** Route structure describes an API route */

type Route struct {
	Path    string
	Method  HttpMethod
	Handler gin.HandlerFunc
	Scope   string
}

type Routes []Route

type RouteInterface interface {
	GetScopes(method HttpMethod, path string) []Route
}

func (routes Routes) GetScope(method HttpMethod, path string) string {
	for _, r := range routes {
		if r.Path == path && r.Method == method {
			return r.Scope
		}
	}
	return ""
}
