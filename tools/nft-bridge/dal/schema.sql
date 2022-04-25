CREATE DATABASE nftbr;
CREATE USER nftbr;
GRANT ALL ON DATABASE nftbr TO nftbr;
SET DATABASE TO nftbr;

-- track each nft bridge transfer
-- src_tx can't be unique key if we want to support batch xfer by contract
-- but for one NFT token, all its past status must be done so we only deal with
-- one pending. but dst_tx if set, must be unique as we send onchain tx without batch
-- but this could also change? if we have a contract to call msg bus?

-- status 1: waiting for sgn, 2: waiting for dst tx, 3: done
CREATE TABLE IF NOT EXISTS nftxfer (
    created_at BIGINT NOT NULL, -- epoch second when add row
    src_chid BIGINT NOT NULL,
    dst_chid BIGINT NOT NULL,
    sender TEXT NOT NULL,
    receiver TEXT NOT NULL,
    src_nft TEXT NOT NULL,
    dst_nft TEXT NOT NULL,
    tok_id TEXT NOT NULL, -- string of goutils/big.Int
    src_tx TEXT NOT NULL, -- not unique if a contract calls deposit w/ batch transfers
    dst_tx TEXT NOT NULL, -- empty string in sql statement when first insert
    status INT2 DEFAULT 1 NOT NULL
);
CREATE INDEX IF NOT EXISTS nftxfer_dsttx_idx on nftxfer (dst_tx);
-- query user history
CREATE INDEX IF NOT EXISTS nftxfer_sender_idx on nftxfer (sender);
CREATE INDEX IF NOT EXISTS nftxfer_receiver_idx on nftxfer (receiver);

-- persist block num/index to resume when restart, key is chid-addr
CREATE TABLE IF NOT EXISTS monitor (
    key TEXT PRIMARY KEY NOT NULL,
    blknum BIGINT NOT NULL,
    blkidx INT NOT NULL -- could be -1 when fast forward with no log received
);

-- support which nft tokens a user has
-- chid, nft, user -> list of ids, note list is NOT sorted
-- note No zero addr so mint/burn won't have issue
-- nft and usr are lower case no 0x prefix
CREATE TABLE IF NOT EXISTS usrnfts (
    chid BIGINT NOT NULL,
    nft TEXT NOT NULL,
    usr TEXT NOT NULL,
    tokens TEXT[] NOT NULL,
    UNIQUE (chid, nft, usr)
);
CREATE INDEX IF NOT EXISTS usrnfts_usr on usrnfts (usr);

-- save all erc721 Transfer events since contract deploy so we can cross check
-- note min/burn also emit Transfer event, from or to will be zero addr
-- nft, from and to are lower case no 0x prefix
-- TODO: add unique constrain to handle duplicated event? or anyway we won't allow dup id in usrnfts tokens so it's fine
CREATE TABLE IF NOT EXISTS allevs (
    chid BIGINT NOT NULL,
    nft TEXT NOT NULL,
    from_addr TEXT NOT NULL,
    to_addr TEXT NOT NULL,
    tok_id TEXT NOT NULL
);