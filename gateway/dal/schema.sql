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
    fee_perc INT NOT NULL DEFAULT 0,
    status INT NOT NULL DEFAULT 1,
    volume FLOAT NOT NULL DEFAULT 0,
    refund_tx TEXT NOT NULL DEFAULT '',
    refund_seq_num INT NOT NULL DEFAULT 0,
    refund_id TEXT,
    UNIQUE (refund_id)
);
CREATE INDEX IF NOT EXISTS tsf_utm_idx ON transfer (update_time);
CREATE INDEX IF NOT EXISTS tsf_ctm_idx ON transfer (create_time);
CREATE INDEX IF NOT EXISTS tsf_addr_idx ON transfer (usr_addr);
CREATE INDEX IF NOT EXISTS tsf_tid_idx ON transfer (transfer_id);
CREATE INDEX IF NOT EXISTS tsf_dtid_idx ON transfer (dst_transfer_id);
CREATE INDEX IF NOT EXISTS tsf_sqn_idx ON transfer (refund_seq_num);
CREATE INDEX IF NOT EXISTS tsf_shs_idx ON transfer (src_tx_hash);
ALTER TABLE IF EXISTS transfer ADD COLUMN IF NOT EXISTS refund_id TEXT UNIQUE;

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
    volume FLOAT NOT NULL DEFAULT 0,
    lp_type INT NOT NULL DEFAULT 1,
    seq_num INT NOT NULL DEFAULT 0,
    withdraw_method_type INT NOT NULL DEFAULT 1,
    withdraw_id TEXT,
    PRIMARY KEY (usr_addr, chain_id, seq_num, lp_type),
    UNIQUE (withdraw_id),
    UNIQUE (usr_addr, chain_id, tx_hash, lp_type)
);
    
CREATE INDEX IF NOT EXISTS lp_utm_idx ON lp (update_time);
CREATE INDEX IF NOT EXISTS lp_ctm_idx ON lp (create_time);
CREATE INDEX IF NOT EXISTS lp_addr_idx ON lp (usr_addr);
ALTER TABLE IF EXISTS lp ADD COLUMN IF NOT EXISTS withdraw_id TEXT UNIQUE;

CREATE TABLE IF NOT EXISTS reward_token (
    symbol TEXT NOT NULL,
    chain_id INT NOT NULL,
    update_time TIMESTAMPTZ NOT NULL DEFAULT now(),
    decimal INT NOT NULL,
    address TEXT NOT NULL,
    PRIMARY KEY (symbol, chain_id)
    );
CREATE INDEX IF NOT EXISTS rwtk_addr_idx ON reward_token (address, chain_id);

CREATE TABLE IF NOT EXISTS token (
    symbol TEXT NOT NULL,
    chain_id INT NOT NULL,
    update_time TIMESTAMPTZ NOT NULL DEFAULT now(),
    decimal INT NOT NULL,
    address TEXT NOT NULL,
    name TEXT NOT NULL DEFAULT '',
    icon TEXT NOT NULL DEFAULT '',
    disabled BOOL NOT NULL DEFAULT false,
    PRIMARY KEY (symbol, chain_id)
);
CREATE INDEX IF NOT EXISTS tk_addr_idx ON token (address, chain_id);

CREATE TABLE IF NOT EXISTS chain (
     id INT PRIMARY KEY NOT NULL,
     name TEXT NOT NULL DEFAULT '',
     icon TEXT NOT NULL DEFAULT '',
     tx_url TEXT NOT NULL DEFAULT '',
     block_delay INT NOT NULL DEFAULT 0,
     gas_token_symbol TEXT NOT NULL DEFAULT '',
     explore_url TEXT NOT NULL DEFAULT '',
     rpc_url TEXT NOT NULL DEFAULT '',
     contract TEXT NOT NULL DEFAULT ''
);

CREATE TABLE IF NOT EXISTS claim_withdraw_reward_log (
     usr_addr TEXT NOT NULL,
     create_time TIMESTAMPTZ NOT NULL DEFAULT now(),
     PRIMARY KEY (usr_addr, create_time)
);

CREATE TABLE IF NOT EXISTS admin_addr (
     addr TEXT PRIMARY KEY NOT NULL
);

CREATE TABLE IF NOT EXISTS apy (
    create_time TIMESTAMPTZ PRIMARY KEY NOT NULL DEFAULT now(),
    apy TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS delayed_op (
    id TEXT NOT NULL,
    type INT DEFAULT 0, -- dal.DelayedOpType
    tx_hash TEXT,
    PRIMARY KEY (id)
);