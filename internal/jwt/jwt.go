package jwt

import (
	"errors"
	"strings"
	"time"

	"github.com/NoANameGroup/DAOld-Backend/internal/consts"
	"github.com/NoANameGroup/DAOld-Backend/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// GenerateToken generates a JWT token for a given UserID.
func GenerateToken(userId bson.ObjectID) (string, error) {
	claims := jwt.MapClaims{
		"userId": userId.Hex(),
		"exp":    time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(consts.JWTSecret))
}

// ParseToken parses a JWT token and returns the claims.
func ParseToken(tokenStr string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(consts.JWTSecret), nil
	})

	if err != nil {
		log.Error("ParseToken error: %v", err)
		return nil, err
	}

	return token, nil
}

// ExtractUserID extracts the user ID from a JWT token.
func ExtractUserID(tokenStr string) (bson.ObjectID, error) {
	token, err := ParseToken(tokenStr)
	if err != nil || !token.Valid {
		log.Error("ExtractUserID failed, invalid token: %v", err)
		return bson.NilObjectID, err
	}

	claims := token.Claims.(jwt.MapClaims)
	userIdStr, ok := claims["userId"].(string)
	if !ok {
		err = errors.New("invalid userId in token")
		log.Error("ExtractUserID error: %v", err)
		return bson.NilObjectID, err
	}

	objectId, err := bson.ObjectIDFromHex(userIdStr)
	if err != nil {
		log.Error("ExtractUserID invalid ObjectID: %v", err)
		return bson.NilObjectID, err
	}

	return objectId, nil
}

// ExtractUserIDFromContext 从 gin.Context 中提取用户ID
func ExtractUserIDFromContext(c *gin.Context) bson.ObjectID {
	authHeader := c.GetHeader("Authorization")
	if !strings.HasPrefix(authHeader, "Bearer ") {
		err := errors.New("missing or malformed authorization header")
		log.CtxError(c.Request.Context(), "ExtractUserIDFromContext: %v", err)
		return bson.NilObjectID
	}

	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
	userId, _ := ExtractUserID(tokenStr)
	return userId
}
