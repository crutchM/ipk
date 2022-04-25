package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"ipk/pkg/data"
	"ipk/pkg/repository"
	"time"
)

const (
	salt      = "fdfsas12dfdsdv4"
	signInKey = "kjngjksdngn"
	tokenTTL  = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId string `json:"userId"`
}

type AuthService struct {
	repo repository.Authorisation
}

func NewAuthService(repo repository.Authorisation) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user data.User) (string, error) {
	user.Password = generatePassword(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username string, password string) (string, error) {
	user, err := s.repo.GetUser(username, generatePassword(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId: user.Id,
	})
	return token.SignedString([]byte(signInKey))
}

func (s AuthService) ParseToken(accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("invalid method")
			}

			return []byte(signInKey), nil
		})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return "", errors.New("token claims are not type *tokenClaims")
	}

	return claims.UserId, nil
}

func generatePassword(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
