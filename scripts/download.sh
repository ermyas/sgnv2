GETH_VER="geth-linux-amd64-1.10.8-26675454"

download_geth() {
	curl -sL https://gethstore.blob.core.windows.net/builds/$GETH_VER.tar.gz | sudo tar -xz --strip 1 -C /usr/local/bin $GETH_VER/geth
	sudo chmod +x /usr/local/bin/geth
}
