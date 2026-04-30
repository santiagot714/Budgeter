package domain

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"time"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)


func TestNewTransaction(t *testing.T) {
	Amount, err := NewAmount(decimal.NewFromInt(100))
	require.NoError(t, err)
	
    t.Run("valid transaction with category id", func(t *testing.T) {
        categoryID := uuid.New()
        transaction, err := NewTransaction(uuid.New(), Amount, TransactionTypeIncome, TransactionMethodCash, uuid.New(), &categoryID, time.Now())
        assert.NoError(t, err)
        assert.Equal(t, categoryID, *transaction.CategoryID())
    })
    t.Run("valid transaction without category id", func(t *testing.T) {
        transaction, err := NewTransaction(uuid.New(), Amount, TransactionTypeIncome, TransactionMethodCash, uuid.New(), nil, time.Now())
        assert.NoError(t, err)
        assert.Nil(t, transaction.CategoryID())
    })
	t.Run("Invalid transaction with invalid category id", func(t *testing.T) {
		_, err := NewTransaction(uuid.New(), Amount, TransactionTypeIncome, TransactionMethodCash, uuid.New(), &uuid.Nil, time.Now())
		assert.Error(t, err)
	})
	t.Run("Invalid transaction with invalid id", func(t *testing.T) {
		_, err := NewTransaction(uuid.Nil, Amount, TransactionTypeIncome, TransactionMethodCash, uuid.New(), nil, time.Now())
		assert.Error(t, err)
	})
	t.Run("Invalid transaction with invalid account id", func(t *testing.T) {
		_, err := NewTransaction(uuid.New(), Amount, TransactionTypeIncome, TransactionMethodCash, uuid.Nil, nil, time.Now())
		assert.Error(t, err)
	})
	t.Run("Invalid transaction with invalid occurred at", func(t *testing.T) {
		_, err := NewTransaction(uuid.New(), Amount, TransactionTypeIncome, TransactionMethodCash, uuid.New(), nil, time.Time{})
		assert.Error(t, err)
	})
	t.Run("Invalid transaction with invalid transaction type", func(t *testing.T) {
		_, err := NewTransaction(uuid.New(), Amount, TransactionType("invalid"), TransactionMethodCash, uuid.New(), nil, time.Now())
		assert.Error(t, err)
	})
	t.Run("Invalid transaction with invalid transaction method", func(t *testing.T) {
		_, err := NewTransaction(uuid.New(), Amount, TransactionTypeIncome, TransactionMethod("invalid"), uuid.New(), nil, time.Now())
		assert.Error(t, err)
	})
	t.Run("Validate createdAt is within a valid range", func(t *testing.T) {
		timeBefore := time.Now()
		transaction, err := NewTransaction(uuid.New(), Amount, TransactionTypeIncome, TransactionMethodCash, uuid.New(), nil, time.Now())
		assert.NoError(t, err)
		timeAfter := time.Now()
		assert.True(t, transaction.CreatedAt().After(timeBefore))
		assert.True(t, transaction.CreatedAt().Before(timeAfter))
	})

}