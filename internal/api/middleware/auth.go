package middleware

import (
	"context"
	"net/http"

	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
)

func Auth_Middleware(auth_client *auth.Client) gin.HandlerFunc {
    return func(ctx *gin.Context) {
       
        id_token, err := ctx.Cookie("access_token")
        
        if err != nil {
           
            ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication cookie missing. Please log in."})
            ctx.Abort()
            return
        }

      
        token, err := auth_client.VerifyIDToken(context.Background(), id_token)
        if err != nil {
            ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired session"})
            ctx.Abort()
            return
        }

      
        ctx.Set("user_id", token.UID)
        ctx.Next()
    }
}