CREATE TABLE IF NOT EXISTS addresses (
    id           varchar(36) PRIMARY KEY DEFAULT uuid_generate_v4(),
	user_id      varchar(36) NOT NULL,
    street_address varchar(255) NOT NULL,
    city varchar(255) NOT NULL,
    state varchar(255) NOT NULL,
    is_default boolean NOT NULL DEFAULT false,
    created_at   timestamp NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id)
);