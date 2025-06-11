CREATE TABLE IF NOT EXISTS notifications (
    id           varchar(36) PRIMARY KEY DEFAULT uuid_generate_v4(),
	user_id varchar(36),
    admin_id varchar(36),
    laundry_request_id varchar(36),
    message text NOT NULL,
    is_read boolean NOT NULL DEFAULT false,
    created_at timestamp NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT fk_admin FOREIGN KEY (admin_id) REFERENCES admins(id),
    CONSTRAINT fk_laundry_request FOREIGN KEY (laundry_request_id) REFERENCES laundry_requests(id)
);