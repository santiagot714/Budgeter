package postgres

import (
	"context"
    "github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/santiagot714/Budgeter/internal/ledger/domain"
)

type PostgresTransactionRepository struct {
	pool *pgxpool.Pool
}

func NewTransactionRepository(pool *pgxpool.Pool) domain.TransactionRepository {
    return &PostgresTransactionRepository{pool: pool}
}

func (r *PostgresTransactionRepository) Create(ctx context.Context, transaction domain.Transaction) error {
    panic("not implemented")
}
func (r *PostgresTransactionRepository) Update(ctx context.Context, transaction domain.Transaction) error {
    panic("not implemented")
}
func (r *PostgresTransactionRepository) Delete(ctx context.Context, id uuid.UUID) error {
    panic("not implemented")
}
func (r *PostgresTransactionRepository) Find(ctx context.Context, ids []uuid.UUID) ([]domain.Transaction, error) {
    panic("not implemented")
}
func (r *PostgresTransactionRepository) GetAll(ctx context.Context, filter domain.TransactionFilter) ([]domain.Transaction, error) {
    panic("not implemented")
}