#!/bin/bash

HEIGHT=$1
KEY_NAME=$2
UPGRADE_NAME="v0.1.0"
BINARY_DIR=".cdnode"
CHAIN_ID="brain-t1"
DENOM="ubrain"
GAS_PRICE="0.0025"
NODE="tcp://localhost:26657"
export DAEMON_NAME="braind"
export DAEMON_HOME="$HOME/$BINARY_DIR"

if ! command -v cosmovisor &>/dev/null; then
  echo "\n\ncosmovisor could not be found"
  exit
fi



# submit upgrade proposal
$DAEMON_NAME tx gov submit-proposal software-upgrade $UPGRADE_NAME --title "Upgrade to $UPGRADE_NAME" --description "Upgrade to $UPGRADE_NAME" --upgrade-info='{"binaries":{"linux/amd64":"https://github.com/cdbo/cdnode/releases/download/'"${UPGRADE_NAME}"'/cdnode_linux_amd64.tar.gz?checksum=sha256:af61e03eb0c3c8b2af43a8dbf61558b2520b2de576f4238ee80a874b21893b71"}}' --deposit 10000000$DENOM --upgrade-height $HEIGHT --from $KEY_NAME --chain-id $CHAIN_ID --keyring-backend test --home $DAEMON_HOME --node $NODE --yes --gas-prices $GAS_PRICE$DENOM --gas auto --gas-adjustment 1.5 --broadcast-mode block

PROPOSAL_ID=$($DAEMON_NAME q gov proposals limit 1 --reverse --output json --home $DAEMON_HOME --node $NODE | jq '.proposals[0].proposal_id | tonumber')

# vote on the proposal
$DAEMON_NAME tx gov vote $PROPOSAL_ID yes --from $KEY_NAME --chain-id $CHAIN_ID --keyring-backend test --home $DAEMON_HOME --node $NODE --yes --gas-prices $GAS_PRICE$DENOM --gas auto --gas-adjustment 1.5 --broadcast-mode block

