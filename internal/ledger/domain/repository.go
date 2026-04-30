package domain

import (
	"errors"
	"time"
	"github.com/google/uuid"
	"context"
)

type TransactionRepository interface {
	Create(ctx context.Context, transaction Transaction) error
	Update(ctx context.Context, transaction Transaction) error
	Delete(ctx context.Context, id uuid.UUID) error
	Find(ctx context.Context, ids []uuid.UUID) ([]Transaction, error)
	GetAll(ctx context.Context, filter TransactionFilter) ([]Transaction, error)
}

// TransactionFilter is the filter options for the GetAll method
type TransactionFilter struct {
	CategoryID *uuid.UUID
	StartDate *time.Time
	EndDate *time.Time
	TransactionType *TransactionType
	TransactionMethod *TransactionMethod
	TransactionAmount *Amount
	TransactionAmountMin *Amount
	TransactionAmountMax *Amount
	Offset *int
	Limit *int
}

// NewFilter checks if the filter is valid
func NewFilter(categoryID *uuid.UUID, startDate *time.Time, endDate *time.Time, transactionType *TransactionType, transactionMethod *TransactionMethod, transactionAmount *Amount, transactionAmountMin *Amount, transactionAmountMax *Amount, offset *int, limit *int) (TransactionFilter, error) {
	// Return error if there are no fields set
	if categoryID == nil && startDate == nil && endDate == nil && transactionType == nil && transactionMethod == nil && transactionAmount == nil && transactionAmountMin == nil && transactionAmountMax == nil {
		return TransactionFilter{}, errors.New("no filter options set")
	}

	// If StartDate and EndDate are set, EndDate must be after StartDate
	if startDate != nil && endDate != nil && endDate.Before(*startDate) {
		return TransactionFilter{}, errors.New("end date must be after start date")
	}

	// TransactionAmount and TransactionAmountMin/Max are mutually exclusive
	if transactionAmount != nil && (transactionAmountMin != nil || transactionAmountMax != nil) {
		return TransactionFilter{}, errors.New("transaction amount and transaction amount min/max are mutually exclusive")
	}

	// If limit is set, it must be greater than 0
	if limit != nil && *limit <= 0 {
		return TransactionFilter{}, errors.New("limit must be greater than 0")
	}

	// If offset is set, it must be greater than or equal to 0
	if offset != nil && *offset < 0 {
		return TransactionFilter{}, errors.New("offset must be greater than or equal to 0")
	}

	return TransactionFilter{
		CategoryID: categoryID,
		StartDate: startDate,
		EndDate: endDate,
		TransactionType: transactionType,
		TransactionMethod: transactionMethod,
		TransactionAmount: transactionAmount,
		TransactionAmountMin: transactionAmountMin,
		TransactionAmountMax: transactionAmountMax,
		Offset: offset,
		Limit: limit,
	}, nil
}