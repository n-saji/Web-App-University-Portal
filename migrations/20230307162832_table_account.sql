-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS account(
    id uuid not null,
    name varchar null,
    info jsonb,
    CONSTRAINT account_pkey PRIMARY KEY (id)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE IF EXISTS account DROP CONSTRAINT IF EXISTS account_pkey;
DROP TABLE IF EXISTS account;
-- +goose StatementEnd
