-- +goose Up
ALTER TABLE accounts ADD CONSTRAINT unique_provider_account UNIQUE (provider_id, account_id);

-- +goose Down
ALTER TABLE accounts DROP CONSTRAINT unique_provider_account;