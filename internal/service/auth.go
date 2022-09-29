package service

import (
	"errors"

	"github.com/fpmi-hci/proekt12b-hedgehogs/internal/domain"
	"github.com/fpmi-hci/proekt12b-hedgehogs/internal/repository"
	"github.com/golang-jwt/jwt/v4"

	"time"

	"golang.org/x/crypto/bcrypt"
)

const (
	signingKey = "grk#iwoerjn%324lskdfnHsdj3skldf"
	tokenTTL   = 48 * time.Hour
)

type tokenClaims struct {
	jwt.RegisteredClaims
	UserId int `json:"userId"`
}

type AuthService struct {
	repo repository.Authorization
}

func (s *AuthService) CreateUser(user *domain.User) (domain.User, error) {
	user.PasswordHash = hashPassword(user.PasswordHash)

	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username)
	if err != nil {
		return "", err
	}
	if isValid := checkPasswordHash(password, user.PasswordHash); !isValid {
		err := errors.New("wrong Password")
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		user.ID,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) GetUserByUsername(username string) (*domain.User, error) {
	user, err := s.repo.GetUser(username)
	if err != nil {
		return nil, err
	}

	return user, err
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, err
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func hashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
