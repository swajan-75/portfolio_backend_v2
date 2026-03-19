package middleware

import (
	"net/http"

	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
)

func Auth_Middleware(authClient *auth.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idToken, err := ctx.Cookie("access_token")
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Authentication cookie missing. Please log in.",
			})
			ctx.Abort() 
			return
		}

		token, err := authClient.VerifyIDToken(ctx.Request.Context(), idToken)
		if err != nil {
			ctx.SetCookie("access_token", "", -1, "/", "", true, true)
			
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid or expired session. Please log in again.",
			})
			ctx.Abort()
			return
		}


		if token.UID == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user identity"})
			ctx.Abort()
			return
		}

		ctx.Set("user_id", token.UID)

		ctx.Next()
	}
}