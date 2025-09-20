-- +goose Up
CREATE TABLE posts(
	id UUID PRIMARY KEY NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	title TEXT,
	url TEXT UNIQUE,
	description TEXT,
	published_at TIMESTAMP,
	feed_id UUID references feeds (id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE posts;
