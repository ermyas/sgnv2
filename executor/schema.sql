CREATE DATABASE IF NOT EXISTS executor;

SET
  DATABASE TO executor;

CREATE TABLE IF NOT EXISTS execution_context (
  -- computed message id, hashed differently from transfer or routing info
  id BYTES PRIMARY KEY,
  -- proto serialized bytes of the proto sgn.message.v1.ExecutionContext, 
  exec_ctx BYTES,
  -- see dal.ExecutionStatus
  status INT DEFAULT 0
);

-- v1.6.3-dev.1 migration
alter table execution_context alter column status set default 1; 

CREATE TABLE IF NOT EXISTS monitor_block (
  event TEXT,
  block_num INT,
  block_idx INT,
  restart BOOLEAN,
  PRIMARY KEY (event)
);