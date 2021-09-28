## how to start to test?
0. install docker and run sgn multi node from `test/e2e/multinode/cbridge_test.go`
remember to slash line 332~35 in `test/e2e/multinode/e2e_test.go`

1. install cockroach and start it
```sh
cockroach start-single-node --insecure --listen-addr=localhost:26257 --http-addr=127.0.0.1:9099 --store=path=/tmp/relaynode_test_sql_db
```
2. init table for 1st time
```sh
cockroach sql --insecure --host=localhost:26257
```
copy `gateway/dal/schema.sql` and run
3. run `run.sh`
4. input `12341234` when `Enter keyring passphrase:` showed in console


mock data sql record:
```shell
update token set name='test USDT1', icon='test' where chain_id = 883 and symbol='USDT';
update token set name='test USDT2', icon='test' where chain_id = 884 and symbol='USDT';
insert into chain (id, name, icon, tx_url) values(883, 'chain1', 'test', 'http://test');
insert into chain (id, name, icon, tx_url) values(884, 'chain2', 'test2', 'http://test2');
```