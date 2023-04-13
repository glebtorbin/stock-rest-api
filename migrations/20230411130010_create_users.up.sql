CREATE TABLE users (
    id bigserial not null primary key,
    email varchar not null unique,
    encrypted_password varchar not null,
    PRIMARY KEY(id)
);

CREATE TABLE stocks (
    id bigserial not null primary key,
    title varchar not null,
    avail_sign boolean
    PRIMARY KEY(id)
);

CREATE TABLE products (
    id bigserial not null primary key,
    code bigint not null,
    size varchar not null,
    stock_id int NOT NULL,
    PRIMARY KEY(id)
    CONSTRAINT stocks FOREIGN KEY(stock_id) REFERENCES stocks(id)
);