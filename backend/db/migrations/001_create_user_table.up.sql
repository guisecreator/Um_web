-- Миграция 001_create_user_table.up.sql

CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       name TEXT NOT NULL,
                       email TEXT NOT NULL,
                       password TEXT NOT NULL,
                       role TEXT NOT NULL,
                       createAt TIMESTAMP,
                       updateAt TIMESTAMP,
                       deletedAt TIMESTAMP
);