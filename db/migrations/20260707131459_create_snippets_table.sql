-- +migrate Up
CREATE TABLE IF NOT EXISTS snippets (
    id SERIAL PRIMARY KEY,
    title varchar(125) NOT NULL,
    content text NOT NULL,
    created TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    expires TIMESTAMP WITH TIME ZONE DEFAULT NOW() + INTERVAL '7 days' NOT NULL
);

-- +migrate Down
DROP TABLE IF EXISTS snippets;
