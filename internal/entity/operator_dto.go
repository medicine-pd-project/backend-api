package entity

type OperatorDTO struct {
	ID           uint64
	Login        string
	PasswordHash string `db:"password"`
	Name         string
}

func (o *Operator) ToDTO() OperatorDTO {
	return OperatorDTO{
		ID:           uint64(o.id),
		Login:        string(o.login),
		PasswordHash: string(o.password),
		Name:         string(o.name),
	}
}

type LoginOperatorRequest struct {
	Login    OperatorLogin    `json:"login"`
	Password OperatorPassword `json:"password"`
}
