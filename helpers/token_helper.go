package helpers

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	fiber "github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt/v5"
	config "github.com/heru-oktafian/fiber-apotek/configs"
)

// TokenValidation validate token
func TokenValidation(c *fiber.Ctx, key string) error {
	// Get token value from header Authorization
	token := c.Get("Authorization")
	// Remove prefix "Bearer " if exist
	if strings.HasPrefix(token, "Bearer ") {
		token = token[len("Bearer "):]
	}

	// Check if token is empty
	if token == "" {
		return JSONResponse(c, fiber.StatusUnauthorized, "Missing token", "Insert valid token to access this endpoint!")
	}

	// Check token in blacklist Redis
	ctx := context.Background()
	redisKey := fmt.Sprintf("blacklist:%s", token)
	rdb := config.RDB
	isBlacklisted, err := rdb.Exists(ctx, redisKey).Result()

	if err != nil {
		log.Printf("Error checking token in Redis: %v", err)
		return JSONResponse(c, fiber.StatusInternalServerError, "Token verification failed", "Server error!")
	}

	if isBlacklisted > 0 {
		return JSONResponse(c, fiber.StatusUnauthorized, "Using token failed", "Token was revoked, please login again!")
	}

	// Verify token using secret key
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		secretKey := []byte(os.Getenv("JWT_SECRET"))
		return secretKey, nil
	})

	if err != nil || !parsedToken.Valid {
		return JSONResponse(c, fiber.StatusUnauthorized, "Invalid token", "Try to login again!")
	}

	// Check klaim token (opsional, misalnya validasi user_id, role, dll.)
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || claims[key] == nil {
		return JSONResponse(c, fiber.StatusUnauthorized, "Invalid token claims", "Try to login again!")
	}

	// Go to next middleware
	return c.Next()
}
