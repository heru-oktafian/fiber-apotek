package helpers

import (
	context "context"
	fmt "fmt"
	log "log"
	os "os"
	strings "strings"

	fiber "github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt/v5"
	config "github.com/heru-oktafian/fiber-apotek/configs"
)

// GetClaimsToken get claim values from token
func GetClaimsToken(c *fiber.Ctx, key string) (string, error) {
	// Get token value from header Authorization
	authHeader := c.Get("Authorization")
	// fmt.Println("Authorization Header:", authHeader)

	// Remove prefix "Bearer " if exists
	token := authHeader
	// Remove prefix "Bearer " if exist
	if strings.HasPrefix(token, "Bearer ") {
		token = token[len("Bearer "):]
	}
	// fmt.Println("Stripped Token:", token)

	// Check if token is empty
	if token == "" {
		return "", fmt.Errorf("missing token")
	}

	// Verify JWT token
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		secretKey := []byte(os.Getenv("JWT_SECRET"))
		return secretKey, nil
	})

	if err != nil {
		// fmt.Println("Token parse error:", err)
		return "", fmt.Errorf("invalid token")
	}

	if !parsedToken.Valid {
		// fmt.Println("Token is not valid")
		return "", fmt.Errorf("invalid token")
	}

	// Get claims from token
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		// fmt.Println("Failed to cast claims to MapClaims")
		return "", fmt.Errorf("invalid token claims")
	}
	// fmt.Printf("All Claims: %+v\n", claims)

	// Get value from claims
	claimValRaw, ok := claims[key]
	if !ok || claimValRaw == nil {
		// fmt.Println("Claim for key", key, "not found")
		return "", fmt.Errorf("claim '%s' not found in token", key)
	}

	claimedValue := fmt.Sprintf("%v", claimValRaw)
	// fmt.Println("Claimed value for key", key, ":", claimedValue)

	return claimedValue, nil
}

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
