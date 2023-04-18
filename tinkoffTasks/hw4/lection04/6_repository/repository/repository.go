package repository

import (
	"context"
	"tfs-db/6_repository/domain"
	"tfs-db/6_repository/repository/queries"

	"github.com/jackc/pgx/v4/pgxpool"
)

type repo struct {
	*queries.Queries
	pool *pgxpool.Pool // point of adding pool to repo struct???
	Nick string
}

/*
NewRepository has an interface return because of EMBEDDING in repo struct
we have *queries.Queries which has implement of Candles and CandlesByTicker methods.
Struct repo has reference to it, so, we can use these methods and change behavior for
another NewRepositoryFunction (as example)
*/

func (r repo) GetNick() string {
	return r.Nick
}

func NewRepository(pgxPool *pgxpool.Pool) Repository {
	return repo{
		Queries: queries.New(pgxPool),
		pool:    pgxPool, // allows us to change DB everytime we want
		Nick:    "aboba",
	}
}

type Repository interface {
	GetNick() string

	Candles(context.Context) ([]domain.Candle, error)
	CandlesByTicker(ctx context.Context, ticker string) ([]domain.Candle, error)
}
