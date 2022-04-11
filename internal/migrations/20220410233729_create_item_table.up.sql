CREATE TABLE IF NOT EXISTS item
(
    item_id bigserial NOT NULL PRIMARY KEY,
    name varchar(100) NOT NULL UNIQUE,
    description text NOT NULL,
    price numeric(15,4) NOT NULL,
    currency varchar(3) NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by varchar(100) NOT NULL,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by varchar(100) NOT NULL
);