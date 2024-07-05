package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	helper "github.com/rohit123sinha456/digitalSignage/helper"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")
		if clientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("No Authorization Header Provided")})
			c.Abort()
			return
		}

		userid := c.Request.Header.Get("userid")
		if userid == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("No USerID Header Provided")})
			c.Abort()
			return
		}

		claims, err := helper.ValidateToken(clientToken)
		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}

		uuiderr := uuid.Validate(userid)
		if uuiderr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": uuiderr})
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Set("uid", claims.Uid)
		c.Next()
	}
}
