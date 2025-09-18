-- +goose NO TRANSACTION

-- +goose Up
CREATE UNIQUE INDEX CONCURRENTLY IF NOT EXISTS uniq_feeds_url ON feeds (url);

-- +goose Down
DROP INDEX CONCURRENTLY IF EXISTS idx_feeds_url;

