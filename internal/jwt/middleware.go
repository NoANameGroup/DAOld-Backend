package jwt

import (
	"errors"
	"github.com/NoANameGroup/DAOld-Backend/internal/consts"
	"github.com/NoANameGroup/DAOld-Backend/pkg/log"
	"github.com/NoANameGroup/DAOld-Backend/pkg/response"
	"github.com/gin-gonic/gin"
	"strings"
)

// AuthMiddleware 使用 response.PostProcess 统一响应，并记录日志
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			err := errors.New("missing token")
			log.CtxError(c.Request.Context(), "AuthMiddleware: %v", err)
			response.PostProcess(c, nil, nil, err)
			c.Abort()
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		userID, err := ExtractUserID(tokenStr)
		if err != nil {
			log.CtxError(c.Request.Context(), "AuthMiddleware: failed to extract userID: %v", err)
			response.PostProcess(c, nil, nil, err)
			c.Abort()
			return
		}

		// 验证成功
		c.Set(consts.ContextUserID, userID)
		c.Next()
	}
}
