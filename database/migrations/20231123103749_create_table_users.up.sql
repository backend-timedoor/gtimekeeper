CREATE TABLE users (
	id SERIAL PRIMARY KEY,
    name VARCHAR(250),
    email VARCHAR(250) UNIQUE,
    phone VARCHAR(250) UNIQUE,
    username VARCHAR(250) UNIQUE,
    password VARCHAR(250),
    created_at TIMESTAMP(3),
    updated_at TIMESTAMP(3),
    deleted_at TIMESTAMP(3)
);