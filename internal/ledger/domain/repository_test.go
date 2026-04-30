package domain

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"time"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

func TestNewFilter(t *testing.T) {
	Amount, err := NewAmount(decimal.NewFromInt(100))
	require.NoError(t, err)
	AmountMin, err := NewAmount(decimal.NewFromInt(50))
	require.NoError(t, err)
	AmountMax, err := NewAmount(decimal.NewFromInt(150))
	require.NoError(t, err)
	startDate := time.Now()
	endDate := startDate.Add(1*time.Hour)
	offset := 0
	limit := 10
	transactionType := TransactionTypeIncome
	transactionMethod := TransactionMethodCash
	t.Run("invalid filter with no fields set", func(t *testing.T) {
		_, err := NewFilter(nil, nil, nil, nil, nil, nil, nil, nil, &offset, &limit)
		assert.Error(t, err)
	})
	t.Run("invalid filter with end date before start date", func(t *testing.T) {
		_, err := NewFilter(nil, &endDate, &startDate,nil, nil, nil, nil, nil, &offset, &limit)
		assert.Error(t, err)
	})
	t.Run("invalid filter with transaction amount and transaction amount min/max set", func(t *testing.T) {
		_, err := NewFilter(nil, nil, nil, nil, nil, &Amount, &AmountMin, &AmountMax, &offset, &limit)
		assert.Error(t, err)
	})
	t.Run("invalid filter with limit less than 1", func(t *testing.T) {
		invalidLimit := 0
		_, err := NewFilter(nil, &startDate, &endDate, nil, nil, nil, nil, nil, &offset, &invalidLimit)
		assert.Error(t, err)
	})	
	t.Run("invalid filter with offset less than 0", func(t *testing.T) {
		invalidOffset := -1
		_, err := NewFilter(nil, &startDate, &endDate, nil, nil, nil, nil, nil, &invalidOffset, &limit)
		assert.Error(t, err)
	})
	t.Run("valid filter with amount range", func(t *testing.T) {
		categoryID := uuid.New()
		filter, err := NewFilter(&categoryID, &startDate, &endDate, &transactionType, &transactionMethod, nil, &AmountMin, &AmountMax, &offset, &limit)
		assert.NoError(t, err)
		assert.Equal(t, categoryID, *filter.CategoryID)
		assert.Equal(t, startDate, *filter.StartDate)
		assert.Equal(t, endDate, *filter.EndDate)
		assert.Equal(t, transactionType, *filter.TransactionType)
		assert.Equal(t, transactionMethod, *filter.TransactionMethod)
		assert.Equal(t, AmountMin, *filter.TransactionAmountMin)
		assert.Equal(t, AmountMax, *filter.TransactionAmountMax)
		assert.Equal(t, 0, *filter.Offset)
		assert.Equal(t, 10, *filter.Limit)
	})
	t.Run("Valid filter with amount", func(t *testing.T) {
		categoryID := uuid.New()
		filter, err := NewFilter(&categoryID, &startDate, &endDate, &transactionType, &transactionMethod, &Amount, nil, nil, &offset, &limit)
		assert.NoError(t, err)
		assert.Equal(t, categoryID, *filter.CategoryID)
		assert.Equal(t, startDate, *filter.StartDate)
		assert.Equal(t, endDate, *filter.EndDate)
		assert.Equal(t, transactionType, *filter.TransactionType)
		assert.Equal(t, transactionMethod, *filter.TransactionMethod)
		assert.Equal(t, Amount, *filter.TransactionAmount)
		assert.Equal(t, 0, *filter.Offset)
		assert.Equal(t, 10, *filter.Limit)
	})
}