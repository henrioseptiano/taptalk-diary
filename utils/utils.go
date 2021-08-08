package utils

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"

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

func IsEmailValid(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(email)
}

func IsBirthdayValid(birthday string) bool {
	birthdayRegex := regexp.MustCompile(`\d{2}-\d{2}-\d{4}`)
	return birthdayRegex.MatchString(birthday)
}

func IsPasswordValid(password string) bool {
	var (
		hasMinlen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	if len(password) >= 6 && len(password) <= 32 {
		hasMinlen = true
	}
	for _, character := range password {
		switch {
		case unicode.IsUpper(character):
			hasUpper = true
		case unicode.IsLower(character):
			hasLower = true
		case unicode.IsNumber(character):
			hasNumber = true
		case unicode.IsSymbol(character) || unicode.IsPunct(character):
			hasSpecial = true
		}
	}
	return hasMinlen && hasUpper && hasLower && hasNumber && hasSpecial
}

func GetQuarter(currentMonth int64) string {
	if currentMonth >= 1 && currentMonth <= 3 {
		return "1"
	}
	if currentMonth >= 4 && currentMonth <= 6 {
		return "2"
	}
	if currentMonth >= 7 && currentMonth <= 9 {
		return "3"
	}
	if currentMonth >= 10 && currentMonth <= 12 {
		return "4"
	}
	return ""
}

func GetMonthRageFromQuarter(quarter string) []int {
	var monthRange []int
	switch quarter {
	case "1":
		monthRange = append(monthRange, 1, 3)
		break
	case "2":
		monthRange = append(monthRange, 4, 6)
		break
	case "3":
		monthRange = append(monthRange, 7, 9)
		break
	case "4":
		monthRange = append(monthRange, 10, 12)
		break

	}
	return monthRange
}

func NextPage(page, totalPages int) int {
	if page == totalPages {
		return page
	}
	return page + 1
}

func PrevPage(page int) int {
	if page > 1 {
		return page - 1
	}
	return page
}

func SortedBy(sort []string) []string {
	var sorted []string
	for _, value := range sort {
		split := strings.Split(",", value)
		sorted = append(sorted, fmt.Sprintf("%s %s", split[0], split[1]))
	}
	return sorted
}

func Offset(page, limit int) int {
	offset := 0
	if page == 1 {
		offset = 0
	} else {
		offset = (page - 1) * limit
	}
	return offset
}

func TotalPages(count, limit int) int {
	return int(math.Ceil(float64(count) / float64(limit)))
}

func IsNumeric(strings string) bool {
	_, err := strconv.ParseFloat(strings, 64)
	return err == nil
}
