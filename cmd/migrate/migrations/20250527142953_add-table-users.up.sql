CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users (
    id           varchar(36) PRIMARY KEY DEFAULT uuid_generate_v4(),
	username     varchar(255) NOT NULL,
	full_name    varchar(255) NOT NULL,
    email        varchar(255) NOT NULL UNIQUE,
    password     varchar(255) NOT NULL,
    phone_number varchar(20),
    created_at   timestamp NOT NULL DEFAULT NOW()
);