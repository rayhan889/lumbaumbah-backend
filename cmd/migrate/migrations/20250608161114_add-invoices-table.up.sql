DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'invoice_type') THEN
        CREATE TYPE invoice_type AS ENUM ('pending', 'paid', 'failed');
    END IF;
END$$;

CREATE TABLE IF NOT EXISTS invoices (
    id           varchar(36) PRIMARY KEY DEFAULT uuid_generate_v4(),
	user_id     varchar(36) NOT NULL,
    admin_id   varchar(36),
    amount     decimal(10, 2) NOT NULL,
    payment_method varchar(50) NOT NULL,
    status invoice_type NOT NULL DEFAULT 'pending',
    issued_at timestamp NOT NULL DEFAULT NOW(),
    paid_at timestamp,

    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT fk_admin FOREIGN KEY (admin_id) REFERENCES admins(id)
);

ALTER TABLE invoices
    ADD laundry_request_id varchar(36) NOT NULL,
    ADD CONSTRAINT fk_laundry_request FOREIGN KEY (laundry_request_id) REFERENCES laundry_requests(id);