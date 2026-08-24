package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	defaultJWTSecret     = "sdkljfoisau89w43ru32krn"
	defaultTokenDuration = 24 * time.Hour
)

type JwtClaims struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`

	jwt.RegisteredClaims
}

type JwtService interface {
	GenerateToken(id, email, name string) (string, error)
	ValidateToken(token string) (*JwtClaims, error)
}

type jwtService struct {
	secretKey     string
	tokenDuration time.Duration
}

func NewJwtService(secretKey string, tokenDuration time.Duration) JwtService {
	if secretKey == "" {
		secretKey = defaultJWTSecret
	}

	if tokenDuration == 0 {
		tokenDuration = defaultTokenDuration
	}

	return &jwtService{
		secretKey:     secretKey,
		tokenDuration: tokenDuration,
	}
}

func (js *jwtService) GenerateToken(id, email, name string) (string, error) {

	claims := JwtClaims{
		ID:    id,
		Name:  name,
		Email: email,

		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(
				time.Now().Add(js.tokenDuration),
			),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	tokenString, err := token.SignedString([]byte(js.secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (js *jwtService) ValidateToken(tokenString string) (*JwtClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(js.secretKey), nil
	})
	if err != nil {
		return nil, fmt.Errorf("invalid token: %v", err)
	}
	if claims, ok := token.Claims.(*JwtClaims); ok && jwt.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}
