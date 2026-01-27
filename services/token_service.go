package services

import (
	"fmt"
	"os"
	"strings"

	fiber "github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt/v5"
)

// GetUserID get user_id from token
func GetUserID(c *fiber.Ctx) (string, error) {
	// Get user_id from claims
	userID, _ := GetClaimsToken(c, "sub")
	// println("+" + userID)
	return userID, nil
}

// GetBranchID get branch_id from token
func GetBranchID(c *fiber.Ctx) (string, error) {
	// Get branch_id from claims
	branchID, _ := GetClaimsToken(c, "branch_id")

	return branchID, nil
}

// GetUserRole get user_role from token
func GetUserRole(c *fiber.Ctx) (string, error) {
	// Get user_role from claims
	userRole, _ := GetClaimsToken(c, "user_role")

	return userRole, nil
}

// GetDefaultMember get default_member from token
func GetDefaultMember(c *fiber.Ctx) (string, error) {
	defaultMember, _ := GetClaimsToken(c, "default_member")

	return defaultMember, nil
}

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
