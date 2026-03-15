-- +goose Up
-- +goose StatementBegin
ALTER TABLE sessions ADD COLUMN cache_read_tokens INTEGER NOT NULL DEFAULT 0;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE sessions DROP COLUMN cache_read_tokens;
-- +goose StatementEnd
