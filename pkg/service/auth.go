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

//рандомные значения(не менять, все посыпется)(последнее-время действия одного токена)
const (
	salt      = "fdfsas12dfdsdv4"
	signInKey = "kjngjksdngn"
	tokenTTL  = 12 * time.Hour
)

//расширение базовых claims jwt либы, нужно чтобы генерить токен на основе id пользователя
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

//просто хэшируем пароль и вызываем метод репозитория
func (s *AuthService) CreateUser(user data.User) (string, error) {
	if user.Post == 3 {
		return s.repo.CreateUser(user)
	} else {
		user.Password = generatePassword(user.Password)
	}
	return s.repo.CreateUser(user)

}

//при авторизации нам нужен токен, поэтому снала получаем пользователя, а потом уже генерим токен посредством либы(вся докумка есть на гите либы, ссылка в импорте)
func (s *AuthService) GenerateToken(username string, password string) (string, error) {
	var user data.User
	var err error
	user, err = s.repo.GetUser(username, generatePassword(password))
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

//проверяем валидность токена, получаем id
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

func (s AuthService) GetAll() []data.User {
	return s.repo.GetAll()
}

func generatePassword(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
