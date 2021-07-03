package middlewares

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/golang-jwt/jwt"
	customErrors "github.com/wevr-in/common/custom-errors"
	"os"
	"strings"
)

type AccessDetails struct {
	AccessUuid string
	UserId     [12]byte
}

var client *redis.Client

func AuthMiddleware(cl *redis.Client) gin.HandlerFunc {
	client = cl
	return func(c *gin.Context) {
		tm, err := ExtractTokenMetadata(c)
		if err != nil {
			c.Error(errors.New("authorization required")).SetType(customErrors.ErrorTypeUnauthorized)
			c.Abort()
			return
		}
		uid, err := FetchAuth(tm)
		if err != nil {
			c.Error(errors.New("authorization required")).SetType(customErrors.ErrorTypeUnauthorized)
			c.Abort()
			return
		}
		fmt.Println(uid)
		c.Set("userId", uid)
	}
}

func extractToken(c *gin.Context) string {
	token := c.Request.Header.Get("Authorization")
	strArr := strings.Split(token, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func verifyToken(c *gin.Context) (*jwt.Token, error) {
	tokenString := extractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func TokenValid(c *gin.Context) error {
	token, err := verifyToken(c)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

func ExtractTokenMetadata(c *gin.Context) (*AccessDetails, error) {
	token, err := verifyToken(c)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}
		userIdStr := fmt.Sprintf("%.f", claims["user_id"])
		if userIdStr != "" {
			return nil, err
		}
		var userId [12]byte
		copy(userId[:], userIdStr)
		return &AccessDetails{
			AccessUuid: accessUuid,
			UserId:     userId,
		}, nil
	}
	return nil, err
}

func FetchAuth(authD *AccessDetails) ([12]byte, error) {
	userId, err := client.Get(authD.AccessUuid).Result()
	if err != nil {
		return [12]byte{}, err
	}
	var userID [12]byte
	copy(userID[:], userId)
	return userID, nil
}
