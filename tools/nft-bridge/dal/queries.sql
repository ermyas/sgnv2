-- NOTE: sqlc gen query.sql.go sorted by func name, so all func should have table name prefix
-- comment must be "name: TableXXX :yyy" see https://docs.sqlc.dev/en/latest/reference/query-annotations.html

-- name: MonGet :one
SELECT * FROM monitor WHERE key = $1;

-- name: MonSet :exec
INSERT INTO monitor (key, blknum, blkidx) VALUES ($1, $2, $3) ON CONFLICT (key) DO UPDATE
SET blknum = excluded.blknum, blkidx = excluded.blkidx;

-- name: NftAddSend :exec
INSERT INTO nftxfer (created_at, src_chid, dst_chid, sender, receiver, src_nft, dst_nft, tok_id, src_tx, dst_tx)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, '');

-- name: NftGetByDstInfo :one
SELECT src_tx FROM nftxfer WHERE src_chid = $1 AND dst_chid = $2 AND receiver = $3 AND dst_nft = $4 AND tok_id = $5 AND status = $6;

-- name: NftGetBySender :many
-- user's history, support pagination
SELECT * FROM nftxfer WHERE sender = $1 AND created_at < $2 ORDER BY created_at desc LIMIT $3;

-- name: NftSetDstTx :exec
-- also set status to 2 wait for dst tx
UPDATE nftxfer SET status = 2, dst_tx = $6 WHERE src_chid = $1 AND dst_chid = $2 AND receiver = $3 AND dst_nft = $4 AND tok_id = $5 AND status = 1;

-- name: NftSetDoneByDstTx :exec
UPDATE nftxfer SET status = 3 WHERE dst_tx = $1;
