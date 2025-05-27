CREATE TABLE IF NOT EXISTS admins (
    id           varchar(36) PRIMARY KEY DEFAULT uuid_generate_v4(),
	username     varchar(255) NOT NULL,
    email        varchar(255) NOT NULL UNIQUE,
    password     varchar(255) NOT NULL,
    created_at   timestamp NOT NULL DEFAULT NOW()
);