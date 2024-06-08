package main

import (
	"core"
	"fmt"
	"gateway/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

var (
	User  = core.Role{Name: "user", Scopes: []string{"read:products"}}
	Admin = core.Role{Name: "admin", Scopes: []string{"read:products", "write:products"}}
)

var routes = core.Routes{
	{
		Path:    "/api/v1/products",
		Method:  core.GET,
		Handler: handlers.GetProducts,
		Scope:   "read:products",
	},
	{
		Path:    "/api/v1/products",
		Method:  core.POST,
		Handler: handlers.CreateProduct,
		Scope:   "write:products",
	},
}

// @title Go Scope based access control
// @version 1.0
// @description Example Go access control for a REST API
// @basePath /api/v1
func main() {
	log.Println("setting up api server")

	server := gin.Default()

	auth := core.AuthHandler{
		Routes:           routes,
		Roles:            []core.Role{User, Admin},
		IdentityProvider: core.NewMockIdentityProvider(),
	}
	server.Use(auth.AuthMiddleware)

	AddRoutes(server, routes)

	server.Run(":8080")
}

func AddRoutes(e *gin.Engine, routes core.Routes) {
	for _, r := range routes {
		switch r.Method {
		case core.GET:
			e.GET(r.Path, r.Handler)
		case core.POST:
			e.POST(r.Path, r.Handler)
		case core.PUT:
			e.PUT(r.Path, r.Handler)
		case core.PATCH:
			e.PATCH(r.Path, r.Handler)
		case core.DELETE:
			e.DELETE(r.Path, r.Handler)
		case core.OPTIONS:
			e.OPTIONS(r.Path, r.Handler)
		default:
			message := fmt.Sprintf("unknown method '%s' encountered for route '%s', expected standard HTTP method", r.Method, r.Path)
			panic(message)
		}
	}
}
