-- Transactions table
CREATE TABLE transactions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    amount NUMERIC(15,2) NOT NULL CHECK (amount > 0),
    kind VARCHAR(255) NOT NULL CHECK (kind IN ('income', 'expense')),
    method VARCHAR(255) NOT NULL CHECK (method IN ('cash', 'debit', 'credit')),
    account_id UUID NOT NULL REFERENCES accounts(id),
    category_id UUID,
    occurred_at TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    deleted_at TIMESTAMPTZ
);

-- Partial index for active transactions
CREATE INDEX idx_transactions_active
ON transactions(account_id)
WHERE deleted_at IS NULL;
