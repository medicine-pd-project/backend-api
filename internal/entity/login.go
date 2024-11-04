package entity

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
)

const (
	// JWTKey TODO: вынести в жопу
	JWTKey = "secret"
)

var (
	ErrWrongPassword = errors.New("wrong password")
)

type JWTToken string

type JWTClaims struct {
	OperatorID OperatorID `json:"operator_id"`
	jwt.RegisteredClaims
}
