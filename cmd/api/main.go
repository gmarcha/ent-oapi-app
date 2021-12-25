package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	s "github.com/gmarcha/goswagger-ent-app/internal/api/controllers"
	"github.com/gmarcha/goswagger-ent-app/internal/api/database"
	"github.com/gmarcha/goswagger-ent-app/internal/api/errors"
	api "github.com/gmarcha/goswagger-ent-app/internal/api/gen"
	_ "github.com/gmarcha/goswagger-ent-app/internal/api/validators"
	_ "github.com/lib/pq"
)

func main() {

	tokenMap := make(map[string]string, 0)

	tokenMap["42"] = "1"

	var server s.Server

	server.Client = database.Connect()
	defer server.Client.Close()

	router := gin.Default()
	router.Use(func(c *gin.Context) {
		if c.Request.URL.Path == "/auth" {
			return
		}
		sessionToken, err := c.Cookie("token")
		if err != nil || sessionToken == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errors.ErrJSON(http.StatusUnauthorized, "auth: missing token"))
			return
		}
		accessToken, exist := tokenMap[sessionToken]
		if !exist {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errors.ErrJSON(http.StatusUnauthorized, "auth: invalid token"))
			return
		}
		tokenPath := fmt.Sprintf("https://api.intra.42.fr/oauth/token/info?access_token=%v", accessToken)
		res, err := http.Get(tokenPath)
		if err != nil || res.StatusCode != http.StatusOK {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errors.ErrJSON(http.StatusUnauthorized, "auth: expired token"))
			return
		}
		var tokenRes map[string]interface{}
		json.NewDecoder(res.Body).Decode(&tokenRes)
		c.AbortWithStatusJSON(http.StatusUnauthorized, tokenRes)
	})

	router = api.RegisterHandlers(router, server)
	router.Run("0.0.0.0:8080")
}
