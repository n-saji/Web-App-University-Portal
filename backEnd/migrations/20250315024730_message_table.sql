-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS messages(
    id uuid not null,
    account_id uuid not null,
    messages varchar null,
    is_read boolean,
    CONSTRAINT messages_pkey PRIMARY KEY (id),
    CONSTRAINT messages_fkey FOREIGN KEY (account_id) REFERENCES accounts(id)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE IF EXISTS messages DROP CONSTRAINT IF EXISTS messages_pkey;
ALTER TABLE IF EXISTS messages DROP CONSTRAINT IF EXISTS messages_fkey;
DROP TABLE IF EXISTS messages;
-- +goose StatementEnd
