-- Buat tabel users
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    full_name VARCHAR(255) NOT NULL,
    nik VARCHAR(16) UNIQUE NOT NULL,
    phone_number VARCHAR(15) UNIQUE NOT NULL,
    account_number VARCHAR(20) UNIQUE NOT NULL,
    balance INT DEFAULT 0 NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Index untuk optimasi pencarian berdasarkan account_number
CREATE INDEX idx_users_account_number ON users(account_number);

-- Buat tabel transactions
CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    account_number VARCHAR(20) NOT NULL,
    transaction_type VARCHAR(10) CHECK (transaction_type IN ('deposit', 'withdraw')),
    amount INT NOT NULL CHECK (amount > 0),
    transaction_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (account_number) REFERENCES users(account_number) ON DELETE CASCADE
);

-- Index untuk optimasi pencarian transaksi berdasarkan account_number
CREATE INDEX idx_transactions_account_number ON transactions(account_number);
