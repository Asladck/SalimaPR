package service

import (
	"SalimProject/models"
	"SalimProject/pkg/dto"
	"SalimProject/pkg/repository"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const (
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

func (s *AuthService) CreateUser(input dto.SignUpInput) (string, error) {
	if len(input.Password) < 8 {
		return "", errors.New("password must be at least 8 characters")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	user := &models.User{
		Id:           uuid.New().String(),
		Email:        input.Email,
		Name:         input.Name,
		Username:     input.Username,
		PasswordHash: string(hash),
	}
	return s.repo.CreateUser(user)
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId string `json:"id"`
}

func (s *AuthService) SignIn(input dto.SignInInput) (string, string, error) {
	user, err := s.repo.GetUser(input.Username, input.Email)
	if err != nil {
		return "", "", errors.New("user not found")
	}
	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(input.Password),
	); err != nil {
		return "", "", errors.New("invalid password")
	}

	access, refresh, err := s.GenerateToken(user.Id)
	if err != nil {
		return "", "", err
	}

	return access, refresh, nil
}

func (s *AuthService) GenerateToken(userId string) (string, string, error) {

	access := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  userId,
		"exp": time.Now().Add(tokenTTL).Unix(),
		"iat": time.Now().Unix(),
	})

	refresh := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  userId,
		"exp": time.Now().Add(tokenRTTL).Unix(), // исправлено!
	})

	at, err := access.SignedString([]byte(signingKeyA))
	if err != nil {
		return "", "", err
	}

	rt, err := refresh.SignedString([]byte(signingKeyR))
	if err != nil {
		return "", "", err
	}

	return at, rt, err
}
func (s *AuthService) ParseRefresh(refresh string) (string, error) {
	token, err := jwt.Parse(refresh, func(t *jwt.Token) (interface{}, error) {
		return []byte(signingKeyR), nil
	})
	if err != nil {
		return "", err
	}

	claims := token.Claims.(jwt.MapClaims)
	return claims["id"].(string), nil
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
