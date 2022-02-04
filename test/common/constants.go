package common

import (
	"time"

	"github.com/celer-network/sgn-v2/eth"
)

const (
	// outPathPrefix is the path prefix for all output from e2e (incl. chain data, binaries etc)
	// the code will append epoch second to this and create the folder
	// the folder will be deleted after test ends successfully
	OutRootDirPrefix = "/tmp/celer_e2e_"
	EnvDir           = "../../env"
	LocalGeth        = "http://127.0.0.1:8545"
	LocalGeth2       = "http://127.0.0.1:8547"

	SgnChainID    = "sgn-localnet-1000"
	SgnPassphrase = "12341234"
	SgnGasPrice   = ""
	SgnValAcct    = "sgn15h2geedmud70gvpajdwpcaxfs4qcrw4z92zlqe"
	SgnNodeURI    = "tcp://localhost:26657"

	SgnBlockInterval = 1
	DefaultTimeout   = 60 * time.Second
	waitMinedTimeout = 180 * time.Second
	BlockDelay       = 0
	PollingInterval  = time.Second
	DisputeTimeout   = 100

	RetryPeriod = 500 * time.Millisecond
	RetryLimit  = 100
)

var (
	SgnHomes = [...]string{
		"../../../docker-volumes/node0/sgnd",
		"../../../docker-volumes/node1/sgnd",
		"../../../docker-volumes/node2/sgnd",
		"../../../docker-volumes/node3/sgnd",
	}

	// validators
	ValSgnAddrStrs = [...]string{
		"sgn15h2geedmud70gvpajdwpcaxfs4qcrw4z92zlqe",
		"sgn1egtta7su5jxjahtw56pe07qerz4lwvrlttac6y",
		"sgn19q9usqmjcmx8vynynfl5tj5n2k22gc5f6wjvd7",
		"sgn1rjr9uaewus3qh4vs4sqdkdvepwyxq8ql84udfh",
	}
	ValConsensusAddrs = [...]string{
		"sgnvalcons150sptaghyax9g3zwcg9hx8rshug9grdtzdfran",
		"sgnvalcons1lp7hgu47jlyuwg5ed6ep5zhmk2s4r32q8emln0",
		"sgnvalcons1cvrh4k08y73shw8ug5a3uex6huq6j9tm2sf2vs",
		"sgnvalcons1hpw9w3rfuq4mteu6kjme9tye05ea0x7x7hpy8l",
	}
	ValEthKs = [...]string{
		"../../keys/vethks0.json",
		"../../keys/vethks1.json",
		"../../keys/vethks2.json",
		"../../keys/vethks3.json",
	}
	ValEthAddrs = [...]eth.Addr{
		eth.Hex2Addr("00078b31fa8b29a76bce074b5ea0d515a6aeaee7"),
		eth.Hex2Addr("0015f5863ddc59ab6610d7b6d73b2eacd43e6b7e"),
		eth.Hex2Addr("00290a43e5b2b151d530845b2d5a818240bc7c70"),
		eth.Hex2Addr("003ea363bccfd7d14285a34a6b1deb862df0bc84"),
	}
	ValSignerKs = [...]string{
		"../../keys/vsigner0.json",
		"../../keys/vsigner1.json",
		"../../keys/vsigner2.json",
		"../../keys/vsigner3.json",
	}
	ValSignerAddrs = [...]eth.Addr{
		eth.Hex2Addr("00a99dc08476bf4e0f8d68f32fcaa991b7836464"),
		eth.Hex2Addr("00bee3477b0d08217642a3b53704a2f716571070"),
		eth.Hex2Addr("00cee31b12d213987db5ea478aec02ad6f2ba3b6"),
		eth.Hex2Addr("00de0c0fb32979b269686785bcf79d948d9a2d0e"),
	}

	// delegators
	DelEthKs = [...]string{
		"../../keys/dethks0.json",
		"../../keys/dethks1.json",
		"../../keys/dethks2.json",
		"../../keys/dethks3.json",
	}
	DelEthAddrs = [...]eth.Addr{
		eth.Hex2Addr("d0f2596d700c9bd4d605c938e586ec67b01c7364"),
		eth.Hex2Addr("d199de50946314ca94b8e967a18d9c1ce5cc9251"),
		eth.Hex2Addr("d290938754df5eecf95f05ebd801c50a43c3231f"),
		eth.Hex2Addr("d3f716da96b893d4bcefa489f65e4b3e9a3dd3e6"),
	}

	// clients
	ClientEthKs = [...]string{
		"../../keys/cethks0.json",
		"../../keys/cethks1.json",
		"../../keys/cethks2.json",
		"../../keys/cethks3.json",
	}
	ClientEthAddrs = [...]eth.Addr{
		eth.Hex2Addr("c06fdd796e140aee53de5111607e8ded93ebdca3"),
		eth.Hex2Addr("c1699e89639adda8f39faefc0fc294ee5c3b462d"),
		eth.Hex2Addr("c22c304660d5f1d2a7a459ceefc0c2cb30f5cfe4"),
		eth.Hex2Addr("c310195e2844791cfd11d158bd13e5649bae832d"),
	}

	// executor
	ExecutorEthKs   = "../../keys/eethks.json"
	ExecutorEthAddr = eth.Hex2Addr("5047fcf61b30685a0aabd569668cb46a4898404e")

	// used by local manual tests
	SgnNodeURIs = [...]string{
		"tcp://localhost:26657",
		"tcp://localhost:26660",
		"tcp://localhost:26662",
		"tcp://localhost:26664",
	}

	Addrs = []eth.Addr{
		ValEthAddrs[0],
		ValEthAddrs[1],
		ValEthAddrs[2],
		ValEthAddrs[3],
		ValSignerAddrs[0],
		ValSignerAddrs[1],
		ValSignerAddrs[2],
		ValSignerAddrs[3],
		DelEthAddrs[0],
		DelEthAddrs[1],
		DelEthAddrs[2],
		DelEthAddrs[3],
		ClientEthAddrs[0],
		ClientEthAddrs[1],
		ClientEthAddrs[2],
		ClientEthAddrs[3],
		ExecutorEthAddr,
	}

	Addrs2 = []eth.Addr{
		ValSignerAddrs[0],
		ValSignerAddrs[1],
		ValSignerAddrs[2],
		ValSignerAddrs[3],
		ClientEthAddrs[0],
		ClientEthAddrs[1],
		ClientEthAddrs[2],
		ClientEthAddrs[3],
		ExecutorEthAddr,
	}

	ValDelAddrs = []eth.Addr{
		ValEthAddrs[0],
		ValEthAddrs[1],
		ValEthAddrs[2],
		ValEthAddrs[3],
		DelEthAddrs[0],
		DelEthAddrs[1],
		DelEthAddrs[2],
		DelEthAddrs[3],
	}
)
