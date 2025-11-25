package service

import (
	"SalimProject/models"
	"SalimProject/pkg/repository"
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"time"
)

const (
	salt        = "adaodkwapd210k1d221"
	signingKeyA = "qweqroqwro123e21edwqdl@@"
	signingKeyR = "wqretgehrgkrm1o3rm3f3p"
	tokenTTL    = 30 * time.Minute
	tokenRTTL   = 7 * 24 * time.Hour
)

type AuthService struct {
	repo repository.Auth
}

func NewAuthService(repo repository.Auth) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user models.User) (string, error) {
	if len(user.Password) < 8 {
		logrus.Infof("User password is shorter than 8 symbols: %s", errors.New("short password"))
		return "", errors.New("short password")
	}
	user.Password = generateHashPassword(user.Password)
	return s.repo.CreateUser(user)
}
func generateHashPassword(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId string `json:"id"`
}

func (s *AuthService) GenerateToken(username, password, email string) (string, string, error) {
	user, err := s.repo.GetUser(username, password, email)
	if err != nil {
		logrus.Fatalf("GeneratetToken/GetUser - Can`t find a user %s:", err.Error())
		return "", "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId: user.Id,
	})
	refToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
		},
		UserId: user.Id,
	})
	rt, err := refToken.SignedString([]byte(signingKeyR))
	at, err := token.SignedString([]byte(signingKeyA))
	return at, rt, err
}
func (s *AuthService) ParseRefToken(tokenR string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenR, &tokenClaims{}, func(token *jwt.Token) (any, error) {
		return []byte(signingKeyR), nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok || !token.Valid {
		return "", errors.New("invalid refresh token")
	}
	return claims.UserId, err
}
func (s *AuthService) ParseToken(accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKeyA), nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return "", errors.New("token claims are not of type *tokenClaims")
	}
	return claims.UserId, nil
}
func (s *AuthService) GenerateAccToken(userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId: userId,
	})
	return token.SignedString([]byte(signingKeyA))
}
