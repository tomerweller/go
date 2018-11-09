-- +migrate Up

ALTER TABLE history_ledgers ADD failed_transaction_count integer DEFAULT 0 NOT NULL,

-- +migrate Down

ALTER TABLE history_ledgers DROP COLUMN failed_transaction_count;