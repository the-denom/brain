# brain

The layer 1 Cosmos blockchain for the Denom

# Info

### [testnet]

- Current binary version: v0.1.0
- Chain ID: `brain-t1`
- Staking denom: `ubrain`
- Denom exponent: `6`
- Wallet prefix: `brain`
- Average block speed: `5s`

# Chain Specs

- Automatic minting disabled
- CosmWasm 1.1 enabled
- IBC enabled
- Staking denom `ubrain` total supply: `1B`

# Install

- [Install Golang v1.18+](https://go.dev/doc/install)
- Set your [$GOBIN and $GOPATH](https://pkg.go.dev/cmd/go#hdr-GOPATH_environment_variable) env
- `git clone https://github.com/cdbo/brain.git`
- `cd brain`
- `make install`

At this point, you have a `braind` binary installed, used to run a node/validator and execute/query the chain.

# Running a local development environment

Read through [init_local.sh](init_local.sh) for more details.

### 1. `./init_local.sh`

### 2. `braind start`

# Running a node

### 1. `braind init <moniker> --chain-id brain-t1`

Moniker is your node name.

### 2. Adjust some configuration parameters

<a name="config"></a>

> The default location for configuration files is `$HOME/.brain/config`

`app.toml`

```bash
minimum-gas-prices = "0.0025ubrain"
```

`config.toml`

```bash
moniker = "moniker_entered_at_step_1"
persistent_peers = <ask on discord>
```

`client.toml`

```bash
chain-id = "brain-t1"
```

### 3. Download the [genesis.json](https://raw.githubusercontent.com/cdbo/brain/master/genesis.json) file to your [config](#config) folder

StateSync is enabled on the public RPC server and can dramatically speed up catch-up time to the latest block.  
It can be enabled by modifying some config parameters before starting the `braind`.  
Here are the changes required to enable StateSync catch-up:

```bash
# config.toml

[statesync]

enable = true
rpc_servers = <ask on discord>
trust_height = <insert previous block height which is a factor of 500>
trust_hash = <insert block hash of that block height>

```

### 4. `braind start`

At this point, your node will start synchronizing with the existing network and catch up on blocks. This might take a while. You can verify the state of your node with the following command: `braind status`, look for the `catching_up` property; once `false`, that means you are in sync with the rest of the chain.

It is recommended to run the this binary as a daemon like systemd. Here is an example of a `/etc/systemd/system/brain.service`:  
_replace $USER with your username and $GOBIN with the path where `braind` is installed._

```bash
[Unit]
Description=Brain Daemon
After=network.target

[Service]
Type=simple
User=$USER
ExecStart=$GOBIN/braind start
Restart=on-abort

[Install]
WantedBy=multi-user.target

[Service]
LimitNOFILE=65535
```

# Running a validator

Once you have a fully sync'd node, you can start signing blocks by becoming a validator.

### 1. Make you have a wallet configured

with `braind keys list`. If you don't, add one: `braind keys add <wallet name>`. If you need funds, hit the faucet on discord with `/request <wallet address>`.

### 2. Execute the **create-validator** transaction:

```bash
braind tx staking create-validator \\
--amount="1000000000ubrain" \\
--pubkey=$(braind tendermint show-validator) \\
--moniker="My Node" \\
--chain-id="brain-t1" \\
--commission-rate="0.05" \\
--commission-max-rate="0.20" \\
--commission-max-change-rate="0.01" \\
--min-self-delegation="1000000" \\
--gas="auto" \\
--gas-prices="0.0025ubrain" \\
--gas-adjustment="1.75" \\
--from="myWalletName" \\
```

- [CosmosHub example](https://hub.cosmos.network/main/validators/validator-setup.html#create-your-validator)
