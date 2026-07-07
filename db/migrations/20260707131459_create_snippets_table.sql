-- +migrate Up
CREATE TABLE IF NOT EXISTS snippets (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- +migrate Down
DROP TABLE IF EXISTS snippets;
