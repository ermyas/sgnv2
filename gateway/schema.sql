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
    dst_chain_id INT NOT NULL,
    src_token_info TEXT NOT NULL DEFAULT '',
    dst_token_info TEXT NOT NULL DEFAULT '',
    src_chain_info TEXT NOT NULL DEFAULT '',
    dst_chain_info TEXT NOT NULL DEFAULT '',
    amt float NOT NULL DEFAULT 0,
    received_amt float NOT NULL DEFAULT 0,
    status INT NOT NULL DEFAULT 1,
    volume float NOT NULL DEFAULT 0
);
CREATE INDEX IF NOT EXISTS tsf_utm_idx ON transfer (update_time);
CREATE INDEX IF NOT EXISTS tsf_ctm_idx ON transfer (create_time);
CREATE INDEX IF NOT EXISTS tsf_sdr_idx ON transfer (usr_addr);
CREATE INDEX IF NOT EXISTS tsf_rcr_idx ON transfer (transfer_id);