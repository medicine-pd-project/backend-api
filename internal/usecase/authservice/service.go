package authservice

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/medicine-pd-project/backend-common/logger"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"

	"github.com/medicine-pd-project/backend-api/internal/entity"
)

type repo interface {
	GetOperator(ctx context.Context, login entity.OperatorLogin) (entity.Operator, error)
}

type Service struct {
	repo repo
}

func New(repo repo) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Login(
	ctx context.Context,
	log logger.Logger,
	req entity.LoginOperatorRequest,
) (entity.JWTToken, error) {
	operator, err := s.repo.GetOperator(ctx, req.Login)
	if err != nil {
		log.Errorf("failed to get operator by login %v: %v", req.Login, err)

		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(operator.Password()), []byte(req.Password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			log.Errorf("invalid password while logging in operator with name: %v", req.Login)

			return "", entity.ErrWrongPassword
		}

		log.Errorf("failed to compare password: %v", err)

		return "", err
	}

	token, err := s.createToken(operator.ID())
	if err != nil {
		log.Errorf("failed to create token: %v", err)

		return "", err
	}

	return token, nil
}

func (s *Service) createToken(id entity.OperatorID) (entity.JWTToken, error) {
	claims := &entity.JWTClaims{
		OperatorID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(10 * time.Minute)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(entity.JWTKey))
	if err != nil {
		return "", err
	}

	return entity.JWTToken(t), nil
}
