package internal

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/trinhdaiphuc/social-network/config"
	"github.com/trinhdaiphuc/social-network/pkg/models"
	"time"
)

type Claims struct {
	User interface{} `json:"user"`
	jwt.StandardClaims
}

func CreateTokenWithUser(jwtKey string, user interface{}, expirationMinute int) (string, error) {
	// Declare the expiration time of the token
	// here, we have kept it as 24 hours
	expirationTime := time.Now().Add(time.Duration(expirationMinute) * time.Minute)

	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		User: &user,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response
	stringToken, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return "", err
	}
	return stringToken, nil
}

func ParseToken(tokenString string) (*models.User, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return config.GetConfig().JwtKey, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user := claims["user"].(*models.User)
		return user, nil
	}
	return nil, err
}
