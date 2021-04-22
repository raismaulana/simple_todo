package middleware

import (
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/raismaulana/simple_todo/helper"
	"github.com/raismaulana/simple_todo/service"
)

// NewAuthorizer returns the authorizer, uses a Casbin enforcer as input
func AuthorizerRole(e *casbin.Enforcer, jwtService service.JWTService, authService service.AuthService) gin.HandlerFunc {
	a := &BasicAuthorizer{enforcer: e, jwtService: jwtService, authService: authService}
	return func(c *gin.Context) {
		if !a.CheckPermission(c) {
			a.RequirePermission(c)
		}
	}
}

// BasicAuthorizer stores the casbin handler
type BasicAuthorizer struct {
	enforcer    *casbin.Enforcer
	jwtService  service.JWTService
	authService service.AuthService
}

// GetUserName gets the user name from the request.
// Currently, only HTTP basic authentication is supported
func (a *BasicAuthorizer) GetRole(c *gin.Context) string {
	const BEARER_SCHEMA = "Bearer "
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return "guest"
	}
	tokenString := authHeader[len(BEARER_SCHEMA):]
	token, _ := a.jwtService.ValidateToken(tokenString)
	if !token.Valid {
		return "guest"
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := claims["user_id"]
	if v, ok := userID.(string); ok {
		role := a.authService.GetRole(v)
		return role
	}
	return "guest"
}

// CheckPermission checks the user/method/path combination from the request.
// Returns true (permission granted) or false (permission forbidden)
func (a *BasicAuthorizer) CheckPermission(c *gin.Context) bool {
	user := a.GetRole(c)
	r := c.Request
	method := r.Method
	path := r.URL.Path
	ok, err := a.enforcer.Enforce(user, path, method)
	if err != nil {
		response := helper.BuildErrorResponse("Internal server error", err.Error(), nil)
		c.AbortWithStatusJSON(http.StatusInternalServerError, response)
		panic(err)
	}

	return ok
}

// RequirePermission returns the 403 Forbidden to the client
func (a *BasicAuthorizer) RequirePermission(c *gin.Context) {
	response := helper.BuildErrorResponse("Forbidden", "You don't have authorization to access this URL", nil)
	c.AbortWithStatusJSON(http.StatusForbidden, response)
}
