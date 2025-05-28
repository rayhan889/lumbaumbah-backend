CREATE TABLE IF NOT EXISTS laundry_types (
    id           varchar(36) PRIMARY KEY DEFAULT uuid_generate_v4(),
	name         varchar(255) NOT NULL,
	description  varchar(255) NOT NULL,
    price       decimal(10, 2) NOT NULL, 
    estimated_days integer NOT NULL,
    is_active boolean NOT NULL DEFAULT false,
    created_at   timestamp NOT NULL DEFAULT NOW()
);