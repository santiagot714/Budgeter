package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// ---------------------------------------------------------------------------
// Value Objects
// ---------------------------------------------------------------------------

// Amount represents a monetary value, always positive.
// It is a Value Object: immutable and validated at construction time.
type Amount struct {
	value decimal.Decimal
}

func NewAmount(v decimal.Decimal) (Amount, error) {
	if v.LessThanOrEqual(decimal.Zero) {
		return Amount{}, errors.New("amount must be greater than zero")
	}
	return Amount{value: v}, nil
}

func (a Amount) Value() decimal.Decimal { return a.value }

func (a Amount) String() string { return a.value.String() }

// TransactionType defines whether the transaction is an income or an expense.
type TransactionType string

const (
	TransactionTypeIncome  TransactionType = "income"
	TransactionTypeExpense TransactionType = "expense"
)

func (t TransactionType) IsValid() bool {
	return t == TransactionTypeIncome || t == TransactionTypeExpense
}

// TransactionMethod defines the payment method used.
type TransactionMethod string

const (
	TransactionMethodCash   TransactionMethod = "cash"
	TransactionMethodDebit  TransactionMethod = "debit"
	TransactionMethodCredit TransactionMethod = "credit"
)

func (m TransactionMethod) IsValid() bool {
	return m == TransactionMethodCash ||
		m == TransactionMethodDebit ||
		m == TransactionMethodCredit
}

// ---------------------------------------------------------------------------
// Entity
// ---------------------------------------------------------------------------

// Transaction represents a financial operation in the ledger.
// All fields are private — the only way to create a valid Transaction is through NewTransaction.
type Transaction struct {
	id         uuid.UUID
	amount     Amount
	kind       TransactionType
	method     TransactionMethod
	accountID  uuid.UUID
	categoryID *uuid.UUID
	createdAt  time.Time
	occurredAt time.Time
	deletedAt  *time.Time
}

// NewTransaction is the factory function for Transaction.
// Guarantees that no invalid Transaction can exist in the domain.
//
// The caller is responsible for generating the id (uuid.New()) to ensure
// idempotency on retries: if the operation fails and is retried with the
// same id, the system can detect the duplicate.
func NewTransaction( id uuid.UUID, amount Amount, kind TransactionType, method TransactionMethod, accountID uuid.UUID, categoryID *uuid.UUID, occurredAt time.Time,
) (Transaction, error) {
	if id == uuid.Nil {
		return Transaction{}, errors.New("transaction id is required")
	}
	if accountID == uuid.Nil {
		return Transaction{}, errors.New("account id is required")
	}
	if !kind.IsValid() {
		return Transaction{}, errors.New("invalid transaction type")
	}
	if !method.IsValid() {
		return Transaction{}, errors.New("invalid transaction method")
	}
	if occurredAt.IsZero() {
		return Transaction{}, errors.New("occurredAt is required")
	}
	if categoryID != nil && *categoryID == uuid.Nil {
		return Transaction{}, errors.New("category id must be valid")
	}

	return Transaction{
		id:         id,
		amount:     amount,
		kind:       kind,
		method:     method,
		accountID:  accountID,
		categoryID: categoryID,
		createdAt:  time.Now(),
		occurredAt: occurredAt,
		deletedAt:  nil,
	}, nil
}

// ---------------------------------------------------------------------------
// Getters
// ---------------------------------------------------------------------------

func (t Transaction) ID() uuid.UUID          { return t.id }
func (t Transaction) Amount() Amount          { return t.amount }
func (t Transaction) Kind() TransactionType   { return t.kind }
func (t Transaction) Method() TransactionMethod { return t.method }
func (t Transaction) AccountID() uuid.UUID    { return t.accountID }
func (t Transaction) CategoryID() *uuid.UUID  { return t.categoryID }
func (t Transaction) CreatedAt() time.Time    { return t.createdAt }
func (t Transaction) OccurredAt() time.Time   { return t.occurredAt }
func (t Transaction) DeletedAt() *time.Time   { return t.deletedAt }

// ---------------------------------------------------------------------------
// Setters
// ---------------------------------------------------------------------------

func (t *Transaction) MarkAsDeleted() {
	now := time.Now()
	t.deletedAt = &now
}