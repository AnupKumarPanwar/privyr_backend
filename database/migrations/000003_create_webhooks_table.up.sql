CREATE TABLE IF NOT EXISTS webhooks (
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    name varchar(50),
    user_id uuid REFERENCES users(id),
    created_at timestamp WITH time zone DEFAULT NOW (),
    updated_at timestamp WITH time zone DEFAULT NOW (),
    deleted_at timestamp WITH time zone
);

CREATE INDEX IF NOT EXISTS webhooks_created_at_idx ON webhooks (created_at);