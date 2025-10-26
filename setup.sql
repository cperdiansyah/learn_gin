-- CREATE DATABASE gorest;
-- USE DATABASE gorest
-- \c gorest;

CREATE TABLE users (
   id SERIAL PRIMARY KEY,
   first_name VARCHAR(100) NOT NULL,
   last_name VARCHAR(100) NOT NULL,
   email VARCHAR(255) UNIQUE NOT NULL,
   phone_number VARCHAR(20)
);