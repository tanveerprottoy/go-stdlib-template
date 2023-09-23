-- DROP CREATE DB
DROP DATABASE IF EXISTS basic_db;
CREATE DATABASE basic_db;

-- DROP CREATE TABLE
DROP TABLE IF EXISTS users;
CREATE TABLE users (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    `name` varchar NOT NULL,
    `role` varchar NOT NULL,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
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
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    created_at bigint,
    updated_at bigint,
    user_id uuid NOT NULL REFERENCES users(id)
);

-- DROP CREATE TABLE
DROP TABLE IF EXISTS roles;
CREATE TABLE roles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR NOT NULL,
    key VARCHAR NOT NULL,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    created_at BIGINT NOT NULL,
    updated_at BIGINT NOT NULL
);

-- DROP CREATE TABLE
DROP TABLE IF EXISTS actions;
CREATE TABLE actions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR NOT NULL,
    "key" VARCHAR NOT NULL,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    created_at BIGINT NOT NULL,
    updated_at BIGINT NOT NULL
);

CREATE INDEX actions_key_idx ON actions(key);


-- JUNCTION TABLES

-- DROP CREATE TABLE
DROP TABLE IF EXISTS roles_users;
CREATE TABLE roles_users (
    role_id UUID NOT NULL,
    user_id UUID NOT NULL
);

DROP TABLE IF EXISTS roles_actions;
CREATE TABLE roles_actions (
    role_id UUID NOT NULL,
    action_ids JSON NOT NULL
); 
