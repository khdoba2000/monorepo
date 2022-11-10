CREATE TABLE IF NOT EXISTS staff_auth (
    id uuid PRIMARY KEY,
    phone_number VARCHAR(16) UNIQUE
    username VARCHAR(64) UNIQUE
    name VARCHAR(64)
    role VARCHAR(16) NOT NULL
    password VARCHAR(64) NOT NULL
    is_active BOOLEAN DEFAULT TRUE
    branch_id uuid DEFAULT NULL
    create_date TIMESTAMPTZ NOT NULL DEFAULT(NOW())
    update_date TIMESTAMPTZ NOT NULL
)