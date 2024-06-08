package core

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	Routes           Routes
	Roles            Roles
	IdentityProvider IdentityProvider
}

type AuthHandlerInterface interface {
	AuthMiddleware(c *gin.Context)
}

// AuthMiddleware is a middleware function that checks if the user
// is authenticated and authorized to access the requested route.
func (a AuthHandler) AuthMiddleware(c *gin.Context) {

	var path = c.FullPath()
	var method = HttpMethod(c.Request.Method)

	// Checks if the current request has a user-id cookie
	userID, err := c.Cookie("user-id")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "Unauthorized, missing user-id cookie",
		})
		return
	}

	// Checks if the user-id cookie is valid
	user, err := a.IdentityProvider.GetUserFromID(userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "Unauthorized, invalid user-id cookie",
		})
		return
	}

	role := a.Roles.GetRole(user.Role)
	permission := a.Routes.GetScope(method, path)

	if !role.HasPermissions(permission) {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"status":  http.StatusForbidden,
			"message": "Forbidden, user has no permission to access this route",
		})
		fmt.Printf("\t%s %s from userID[%s] with role[%s] : denied\n", method, path, userID, role.Name)
		return
	}

	c.Next()

	fmt.Printf("\t%s %s from userID[%s] with role[%s] : granted\n", method, path, userID, role.Name)

}
