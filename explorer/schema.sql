CREATE DATABASE IF NOT EXISTS explorer;
CREATE USER IF NOT EXISTS explorer;
GRANT ALL ON DATABASE explorer TO explorer;
SET DATABASE TO explorer;

CREATE TABLE IF NOT EXISTS daily_liquidity (
    datetime TIMESTAMPTZ PRIMARY KEY NOT NULL,
    total_liquidity NUMERIC NOT NULL
);

CREATE TABLE IF NOT EXISTS daily_transaction_stat (
    datetime TIMESTAMPTZ PRIMARY KEY NOT NULL,
    transaction_volume NUMERIC NOT NULL,
    transaction_count  INT NOT NULL
);

CREATE TABLE IF NOT EXISTS transaction_stat (
    begin_datetime TIMESTAMPTZ PRIMARY KEY NOT NULL,
    end_datetime TIMESTAMPTZ NOT NULL,
    transaction_volume NUMERIC NOT NULL,
    transaction_count INT NOT NULL
);

CREATE TABLE IF NOT EXISTS wallet (
    addr TEXT PRIMARY KEY NOT NULL,
    create_time TIMESTAMPTZ NOT NULL DEFAULT now()
);