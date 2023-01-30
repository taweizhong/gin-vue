package middleware

import (
	"awesomeProject1/commen"
	"awesomeProject1/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			context.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			context.Abort()
			return
		}
		tokenString = tokenString[7:]
		token, claims, err := commen.ParseToken(tokenString)
		if err != nil || !token.Valid {
			context.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "token解析失败"})
			context.Abort()
			return
		}
		userId := claims.UserId
		Db := commen.GetDb()
		var user model.User
		Db.First(&user, userId)

		if user.ID == 0 {
			context.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "用户不存在"})
			context.Abort()
			return
		}

		context.Set("user", user)
		context.Next()
	}
}
