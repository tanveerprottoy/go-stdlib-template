-- DROP CREATE DB
DROP DATABASE IF EXISTS basic_db;
CREATE DATABASE basic_db;

-- DROP CREATE TABLE
DROP TABLE IF EXISTS users;
CREATE TABLE users (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    `name` VARCHAR NOT NULL,
    `role` VARCHAR NOT NULL,
    created_at BIGINT,
    updated_at BIGINT
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
    `name` VARCHAR NOT NULL,
    created_at BIGINT,
    updated_at BIGINT
);