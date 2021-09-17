# Setup sgn-v2.gateway stack

## Prepare EC2 machine and install dependencies

1. use ubuntu AMI, make sure EBS type is GP3, use eks-vpv-VPC VPC, subnet, security group and ssh key. `sudo apt update; sudo apt upgrade -y; reboot`
2. install go: `sudo snap install go --classic`. also need gcc `sudo apt install build-essential -y`
3. Install CockroachDB:

```sh
curl -sL https://binaries.cockroachdb.com/cockroach-v20.2.7.linux-amd64.tgz | sudo tar -xz --strip 1 -C /usr/local/bin cockroach-v20.2.7.linux-amd64/cockroach
sudo chmod +x /usr/local/bin/cockroach
```

4. Create the cBridge.gateway directory:

```sh
mkdir -p $HOME/cbridge/cockroach
```

## Setup

1. From the $HOME directory, clone the `sgn-v2` repository:

```sh
git clone https://github.com/celer-network/sgn-v2.git
cd sgn-v2/
git checkout main
```

2. Setup and start the CockroachDB system service:

```sh
. crdbsvc.sh
```
Configure the DB schema:
```sh
/usr/local/bin/cockroach sql --insecure < $HOME/sgn-v2/gateway/schema.sql
```
