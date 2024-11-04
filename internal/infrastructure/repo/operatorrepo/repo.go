package operatorrepo

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"

	"github.com/medicine-pd-project/backend-api/internal/entity"
)

type Repo struct {
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *Repo {
	return &Repo{
		pool: pool,
	}
}

func (r *Repo) GetOperator(ctx context.Context, login entity.OperatorLogin) (entity.Operator, error) {
	var dto entity.OperatorDTO

	q := `SELECT id, login, password, name FROM operators WHERE login = $1`

	err := r.pool.QueryRow(ctx, q, login).Scan(&dto.ID, &dto.Login, &dto.PasswordHash, &dto.Name)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.Operator{}, entity.ErrOperatorNotFound
		}

		return entity.Operator{}, err
	}

	return entity.NewOperator(&dto)
}
