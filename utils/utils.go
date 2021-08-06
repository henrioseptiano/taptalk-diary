package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/henrioseptiano/taptalk-diary/entity"
	"github.com/henrioseptiano/taptalk-diary/models"
	"golang.org/x/crypto/bcrypt"
)

func HashedPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func ExtractToken(c *fiber.Ctx) string {
	bearToken := c.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(os.Getenv("SECRET_KEY")), nil
}

func verifyToken(c *fiber.Ctx) (*jwt.Token, error) {
	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, jwtKeyFunc)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func ExtractTokenMetadata(c *fiber.Ctx) (*models.JwtCustomClaims, error) {
	token, err := verifyToken(c)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		username, ok := claims["username"].(string)
		if !ok {
			return nil, err
		}
		userID, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["userID"]), 10, 64)
		if err != nil {
			return nil, err
		}

		return &models.JwtCustomClaims{
			ID:       userID,
			Username: username,
		}, nil
	}
	return nil, err
}

func CreateToken(getUser *entity.MasterUser) (string, error) {
	secretKey := os.Getenv("SECRET_KEY")
	mySecretKey := []byte(secretKey)
	claims := models.UserClaims{
		getUser.ID,
		getUser.Username,
		jwt.StandardClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(mySecretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
