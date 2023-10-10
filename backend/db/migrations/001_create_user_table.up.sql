-- Миграция 001_create_user_table.up.sql

CREATE TABLE users (
										 id SERIAL PRIMARY KEY,
										 email TEXT NOT NULL,
										 password TEXT NOT NULL,
										 role TEXT NOT NULL,
										 create_at TIMESTAMP,
										 update_at TIMESTAMP,
										 deleted_at TIMESTAMP
);
