package auth

import (
	"errors"
	"fmt"
	"log/slog"
	"open_gamba/internal/model"
	"open_gamba/internal/user"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(user *user.User) (string, error) {
	if user == nil {
		return "", errors.New("user cannot be nil")
	}

	secret := os.Getenv("AUTH_JWT_SECRET_KEY")
	if secret == "" {
		return "", errors.New("JWT secret key not configured")
	}

	expirationTime := time.Now().Add(30 * 24 * time.Hour)

	claims := model.JwtCustomClaims{
		Username: user.Username,
		Roles:    user.AccessRoles,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   fmt.Sprintf("%d", user.ID),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    "opengamba",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		slog.Error("Failed to sign token", "error", err)
		return "", err
	}
	slog.Info("Genrated token for user", "user", user.Username)
	return tokenString, nil
}

func ValidateToken(tokenString string) (*model.JwtCustomClaims, error) {
	secret := os.Getenv("AUTH_JWT_SECRET_KEY")
	if secret == "" {
		errMsg := "JWT secret key not configured"
		slog.Error(errMsg)
		return nil, errors.New(errMsg)
	}

	token, err := jwt.ParseWithClaims(tokenString, &model.JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			errMsg := fmt.Sprintf("Unexpected signing method: %v", token.Header["alg"])
			slog.Error("Invalid JWT signing method", "error", errMsg)
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		slog.Error("Failed to parse JWT token", "error", err.Error())
		return nil, fmt.Errorf("Failed to parse token: %w", err)
	}

	if claims, ok := token.Claims.(*model.JwtCustomClaims); ok && token.Valid {
		slog.Info("JWT token validated successfully",
			"username", claims.Username,
			"roles", claims.Roles)
		return claims, nil
	}

	slog.Error("Error validating token claims", "error", "Invalid token claims or token not valid")
	return nil, errors.New("invalid token claims")
}
