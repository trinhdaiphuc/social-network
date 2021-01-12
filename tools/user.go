package tools

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/trinhdaiphuc/social-network/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ForUserContext finds the user from the context. REQUIRES Middleware to have run.
func ForUserContext(ctx context.Context) (*models.User, error) {
	claims, ok := ctx.Value("user").(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("not found user context")
	}

	userClaim, ok := claims["user"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("Parse token error", claims)
	}
	id, _ := userClaim["id"].(string)
	oid, _ := primitive.ObjectIDFromHex(id)
	email, _ := userClaim["email"].(string)
	userName, _ := userClaim["username"].(string)
	user := &models.User{
		ID:       oid,
		Email:    email,
		Username: userName,
	}
	return user, nil
}
