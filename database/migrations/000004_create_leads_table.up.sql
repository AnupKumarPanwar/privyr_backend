CREATE TABLE IF NOT EXISTS leads (
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    name varchar(50),
    email varchar(50),
    phone varchar(15),
    other_fields jsonb, 
    webhook_id uuid REFERENCES webhooks(id),
    created_at timestamp WITH time zone DEFAULT NOW (),
    updated_at timestamp WITH time zone DEFAULT NOW (),
    deleted_at timestamp WITH time zone
);

CREATE INDEX IF NOT EXISTS leads_created_at_idx ON leads (created_at);