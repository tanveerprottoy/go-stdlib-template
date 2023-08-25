-- DROP CREATE DB
DROP DATABASE IF EXISTS basic_db;
CREATE DATABASE basic_db;

-- DROP CREATE TABLE
DROP TABLE IF EXISTS users;
CREATE TABLE users (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    `name` varchar NOT NULL,
    `role` varchar NOT NULL,
    created_at bigint,
    updated_at bigint
);

/*
email VARCHAR NOT NULL,
age INT,
phone VARCHAR NOT NULL,
*/

-- DROP CREATE TABLE
DROP TABLE IF EXISTS contents;
CREATE TABLE contents (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    `name` varchar NOT NULL,
    created_at bigint,
    updated_at bigint,
    user_id uuid NOT NULL REFERENCES users(id)
);