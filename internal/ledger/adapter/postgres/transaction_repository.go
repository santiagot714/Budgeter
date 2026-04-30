package postgres

import (
	"context"
	"fmt"
	"strings"
	"time"

    "github.com/google/uuid"
	"github.com/shopspring/decimal"
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
    _, err := r.pool.Exec(ctx, "INSERT INTO transactions (id, amount, kind, method, account_id, category_id, occurred_at) VALUES ($1, $2, $3, $4, $5, $6, $7)", transaction.ID(), transaction.Amount().Value(), transaction.Kind(), transaction.Method(), transaction.AccountID(), transaction.CategoryID(), transaction.OccurredAt())
    return err
}
func (r *PostgresTransactionRepository) Update(ctx context.Context, transaction domain.Transaction) error {
    _, err := r.pool.Exec(ctx, "UPDATE transactions SET amount = $1, kind = $2, method = $3, account_id = $4, category_id = $5 WHERE id = $6", transaction.Amount().Value(), transaction.Kind(), transaction.Method(), transaction.AccountID(), transaction.CategoryID(), transaction.ID())
    return err
}
func (r *PostgresTransactionRepository) Delete(ctx context.Context, id uuid.UUID) error {
    _, err := r.pool.Exec(ctx, "UPDATE transactions SET deleted_at = NOW() WHERE id = $1", id)
    return err
}
func (r *PostgresTransactionRepository) Find(ctx context.Context, ids []uuid.UUID) ([]domain.Transaction, error) {
    rows, err := r.pool.Query(ctx, "SELECT id, amount, kind, method, account_id, category_id, occurred_at, created_at, deleted_at FROM transactions WHERE id = ANY($1)", ids)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    var transactions []domain.Transaction
    for rows.Next() {
        var id uuid.UUID
        var rawAmount decimal.Decimal
        var kind domain.TransactionType
        var method domain.TransactionMethod
        var accountID uuid.UUID
        var categoryID *uuid.UUID
        var occurredAt time.Time
        var createdAt time.Time
        var deletedAt *time.Time
        err := rows.Scan(&id, &rawAmount, &kind, &method, &accountID, &categoryID, &occurredAt, &createdAt, &deletedAt)
        if err != nil {
            return nil, err
        }
		amount, err := domain.NewAmount(rawAmount)
		if err != nil {
			return nil, err
		}
        transaction, err := domain.NewTransaction(id, amount, kind, method, accountID, categoryID, occurredAt)
        if err != nil {
            return nil, err
        }
        transactions = append(transactions, transaction)
    }
	if err := rows.Err(); err != nil {
		return nil, err
	}
    return transactions, nil
}
func (r *PostgresTransactionRepository) GetAll(ctx context.Context, filter domain.TransactionFilter) ([]domain.Transaction, error) {
    conditions := []string{}
	values := []any{}
	index := 1
	
	conditions = append(conditions, "deleted_at IS NULL")
	
	if filter.CategoryID != nil {
		conditions = append(conditions, fmt.Sprintf("category_id = $%d", index))
		values = append(values, filter.CategoryID)
		index++
	}
	if filter.StartDate != nil {
		conditions = append(conditions, fmt.Sprintf("occurred_at >= $%d", index))
		values = append(values, filter.StartDate)
		index++
	}
	if filter.EndDate != nil {
		conditions = append(conditions, fmt.Sprintf("occurred_at <= $%d", index))
		values = append(values, filter.EndDate)
		index++
	}
	if filter.TransactionType != nil {
		conditions = append(conditions, fmt.Sprintf("kind = $%d", index))
		values = append(values, filter.TransactionType)
		index++
	}
	if filter.TransactionMethod != nil {
		conditions = append(conditions, fmt.Sprintf("method = $%d", index))
		values = append(values, filter.TransactionMethod)
		index++
	}
	if filter.TransactionAmount != nil {
		conditions = append(conditions, fmt.Sprintf("amount = $%d", index))
		values = append(values, filter.TransactionAmount)
		index++
	}
	if filter.TransactionAmountMin != nil {
		conditions = append(conditions, fmt.Sprintf("amount >= $%d", index))
		values = append(values, filter.TransactionAmountMin)
		index++
	}
	if filter.TransactionAmountMax != nil {
		conditions = append(conditions, fmt.Sprintf("amount <= $%d", index))
		values = append(values, filter.TransactionAmountMax)
		index++
	}
	values = append(values, *filter.Limit)
	index++
	values = append(values, *filter.Offset)
	index++

	query := fmt.Sprintf("SELECT id, amount, kind, method, account_id, category_id, occurred_at, created_at, deleted_at FROM transactions WHERE %s ORDER BY occurred_at DESC LIMIT $%d OFFSET $%d",
    strings.Join(conditions, " AND "), index-2, index-1)
	rows, err := r.pool.Query(ctx, query, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
    var transactions []domain.Transaction
    for rows.Next() {
        var id uuid.UUID
        var rawAmount decimal.Decimal
        var kind domain.TransactionType
        var method domain.TransactionMethod
        var accountID uuid.UUID
        var categoryID *uuid.UUID
        var occurredAt time.Time
        var createdAt time.Time
        var deletedAt *time.Time
        err := rows.Scan(&id, &rawAmount, &kind, &method, &accountID, &categoryID, &occurredAt, &createdAt, &deletedAt)
        if err != nil {
            return nil, err
        }
		amount, err := domain.NewAmount(rawAmount)
		if err != nil {
			return nil, err
		}
        transaction, err := domain.NewTransaction(id, amount, kind, method, accountID, categoryID, occurredAt)
        if err != nil {
            return nil, err
        }
        transactions = append(transactions, transaction)
    }
	if err := rows.Err(); err != nil {
		return nil, err
	}
    return transactions, nil
}