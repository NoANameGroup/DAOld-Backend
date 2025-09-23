package jwt

import (
	"errors"
	"time"

	"github.com/NoANameGroup/DAOld-Backend/consts"
	"github.com/NoANameGroup/DAOld-Backend/pkg/log"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GenerateToken generates a JWT token for a given UserID.
func GenerateToken(userId primitive.ObjectID) (string, error) {
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
func ExtractUserID(tokenStr string) (primitive.ObjectID, error) {
	token, err := ParseToken(tokenStr)
	if err != nil || !token.Valid {
		log.Error("ExtractUserID failed, invalid token: %v", err)
		return primitive.NilObjectID, err
	}

	claims := token.Claims.(jwt.MapClaims)
	userIdStr, ok := claims["userId"].(string)
	if !ok {
		err = errors.New("invalid userId in token")
		log.Error("ExtractUserID error: %v", err)
		return primitive.NilObjectID, err
	}

	objectId, err := primitive.ObjectIDFromHex(userIdStr)
	if err != nil {
		log.Error("ExtractUserID invalid ObjectID: %v", err)
		return primitive.NilObjectID, err
	}

	return objectId, nil
}
