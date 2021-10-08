CREATE DATABASE IF NOT EXISTS gateway;
SET DATABASE TO gateway;

CREATE TABLE IF NOT EXISTS transfer (
    transfer_id TEXT PRIMARY KEY NOT NULL,
    dst_transfer_id TEXT NOT NULL DEFAULT '',
    usr_addr TEXT NOT NULL,
    update_time TIMESTAMPTZ NOT NULL DEFAULT now(),
    create_time TIMESTAMPTZ NOT NULL DEFAULT now(),
    token_symbol TEXT NOT NULL,
    src_chain_id INT NOT NULL,
    src_tx_hash TEXT NOT NULL DEFAULT '',
    dst_chain_id INT NOT NULL,
    dst_tx_hash TEXT NOT NULL DEFAULT '',
    amt TEXT NOT NULL DEFAULT '0',
    received_amt TEXT NOT NULL DEFAULT '0',
    status INT NOT NULL DEFAULT 1,
    volume FLOAT NOT NULL DEFAULT 0,
    refund_tx TEXT NOT NULL DEFAULT '',
    refund_seq_num INT NOT NULL DEFAULT 0
);
CREATE INDEX IF NOT EXISTS tsf_utm_idx ON transfer (update_time);
CREATE INDEX IF NOT EXISTS tsf_ctm_idx ON transfer (create_time);
CREATE INDEX IF NOT EXISTS tsf_addr_idx ON transfer (usr_addr);
CREATE INDEX IF NOT EXISTS tsf_tid_idx ON transfer (transfer_id);
CREATE INDEX IF NOT EXISTS tsf_dtid_idx ON transfer (dst_transfer_id);
CREATE INDEX IF NOT EXISTS tsf_sqn_idx ON transfer (refund_seq_num);

CREATE TABLE IF NOT EXISTS lp (
    usr_addr TEXT NOT NULL,
    chain_id INT NOT NULL,
    token_symbol TEXT NOT NULL,
    token_addr TEXT NOT NULL,
    amt TEXT NOT NULL DEFAULT '0',
    tx_hash TEXT NOT NULL DEFAULT '',
    update_time TIMESTAMPTZ NOT NULL DEFAULT now(),
    create_time TIMESTAMPTZ NOT NULL DEFAULT now(),
    status INT NOT NULL DEFAULT 1,
    lp_type INT NOT NULL DEFAULT 1,
    seq_num INT NOT NULL DEFAULT 0,
    PRIMARY KEY (usr_addr, chain_id, seq_num, lp_type)
    );
CREATE INDEX IF NOT EXISTS lp_utm_idx ON lp (update_time);
CREATE INDEX IF NOT EXISTS lp_ctm_idx ON lp (create_time);
CREATE INDEX IF NOT EXISTS lp_addr_idx ON lp (usr_addr);


CREATE TABLE IF NOT EXISTS token (
    symbol TEXT NOT NULL,
    chain_id INT NOT NULL,
    update_time TIMESTAMPTZ NOT NULL DEFAULT now(),
    decimal INT NOT NULL,
    address TEXT NOT NULL,
    name TEXT NOT NULL DEFAULT '',
    icon TEXT NOT NULL DEFAULT '',
    contract TEXT NOT NULL,
    PRIMARY KEY (symbol, chain_id)
);
CREATE INDEX IF NOT EXISTS tk_addr_idx ON token (address, chain_id);

CREATE TABLE IF NOT EXISTS chain (
     id INT PRIMARY KEY NOT NULL,
     name TEXT NOT NULL DEFAULT '',
     icon TEXT NOT NULL DEFAULT '',
     tx_url TEXT NOT NULL DEFAULT ''
);


CREATE TABLE IF NOT EXISTS usr_slippage (
     addr TEXT PRIMARY KEY NOT NULL,
     slippage INT NOT NULL DEFAULT 0
);