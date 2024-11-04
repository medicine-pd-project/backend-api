package entity

import "github.com/pkg/errors"

var (
	ErrInvalidOperatorID       = errors.New("invalid operator id")
	ErrInvalidOperatorLogin    = errors.New("invalid operator login")
	ErrInvalidOperatorPassword = errors.New("invalid operator password")
	ErrInvalidOperatorName     = errors.New("invalid operator name")

	ErrOperatorNotFound = errors.New("operator not found")
)

type OperatorID uint64
type OperatorLogin string
type OperatorPassword string
type OperatorName string

type Operator struct {
	id       OperatorID
	login    OperatorLogin
	password OperatorPassword
	name     OperatorName
}

func NewOperator(dto *OperatorDTO) (Operator, error) {
	id, err := NewOperatorID(dto.ID)
	if err != nil {
		return Operator{}, err
	}

	login, err := NewOperatorLogin(dto.Login)
	if err != nil {
		return Operator{}, err
	}

	password, err := NewOperatorPassword(dto.PasswordHash)
	if err != nil {
		return Operator{}, err
	}

	name, err := NewOperatorName(dto.Name)
	if err != nil {
		return Operator{}, err
	}

	return Operator{
		id:       id,
		login:    login,
		password: password,
		name:     name,
	}, nil
}

func (o *Operator) ID() OperatorID {
	return o.id
}

func (o *Operator) Login() OperatorLogin {
	return o.login
}

func (o *Operator) Password() OperatorPassword {
	return o.password
}

func (o *Operator) Name() OperatorName {
	return o.name
}

func NewOperatorID(id uint64) (OperatorID, error) {
	if id == 0 {
		return 0, ErrInvalidOperatorID
	}

	return OperatorID(id), nil
}

func NewOperatorLogin(login string) (OperatorLogin, error) {
	if login == "" {
		return "", ErrInvalidOperatorLogin
	}

	return OperatorLogin(login), nil
}

func NewOperatorPassword(password string) (OperatorPassword, error) {
	if password == "" {
		return "", ErrInvalidOperatorPassword
	}

	return OperatorPassword(password), nil
}

func NewOperatorName(name string) (OperatorName, error) {
	if name == "" {
		return "", ErrInvalidOperatorName
	}

	return OperatorName(name), nil
}
