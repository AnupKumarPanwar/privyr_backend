CREATE TABLE IF NOT EXISTS users (
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    email varchar(50),
    created_at timestamp WITH time zone DEFAULT NOW (),
    updated_at timestamp WITH time zone DEFAULT NOW (),
    deleted_at timestamp WITH time zone
);

CREATE INDEX IF NOT EXISTS users_created_at_idx ON users (created_at);