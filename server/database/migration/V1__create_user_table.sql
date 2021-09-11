CREATE TABLE users (
    id uuid PRIMARY KEY,
    email varchar(225) NOT NULL UNIQUE,
    password varchar(255) NOT NULL,
    created_at timestamptz NOT NULL DEFAULT (now()),
    updated_at timestamptz NULL DEFAULT (null)
);