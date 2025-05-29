DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'status_type') THEN
        CREATE TYPE status_type AS ENUM ('pending', 'canceled', 'completed', 'processed');
    END IF;
END$$;

CREATE TABLE IF NOT EXISTS laundry_requests (
    id           varchar(36) PRIMARY KEY DEFAULT uuid_generate_v4(),
	user_id      varchar(36) NOT NULL,
    admin_id     varchar(36),
	laundry_type_id varchar(36) NOT NULL,
    address_id   varchar(36) NOT NULL,
    weight       decimal(10, 2) NOT NULL,
    notes        varchar(255) NOT NULL,
    status       varchar(255) NOT NULL,
    completion_date timestamp NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT fk_admin FOREIGN KEY (admin_id) REFERENCES admins(id),
    CONSTRAINT fk_laundry_type FOREIGN KEY (laundry_type_id) REFERENCES laundry_types(id),
    CONSTRAINT fk_address FOREIGN KEY (address_id) REFERENCES addresses(id)
);