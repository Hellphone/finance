-- Table: users

-- DROP TABLE IF EXISTS users;

CREATE TABLE IF NOT EXISTS users
(
    id           BIGSERIAL PRIMARY KEY,
    name         VARCHAR(100)     NOT NULL,
    money_amount DOUBLE PRECISION NOT NULL
)
