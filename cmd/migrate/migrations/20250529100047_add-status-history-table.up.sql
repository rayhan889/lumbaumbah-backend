CREATE TABLE IF NOT EXISTS status_histories (
    id           varchar(36) PRIMARY KEY DEFAULT uuid_generate_v4(),
	laundry_request_id varchar(36) NOT NULL,
    status varchar(255) NOT NULL,
    updated_at timestamp NOT NULL DEFAULT NOW(),
    updated_by varchar(36) NOT NULL,

    CONSTRAINT fk_laundry_request FOREIGN KEY (laundry_request_id) REFERENCES laundry_requests(id)
);