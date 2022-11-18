#!/bin/bash

HEIGHT=$1
UPGRADE_NAME="v0.2.0"
BINARY_DIR=".brain"
CHAIN_ID="brain-t1"
DENOM="ubrain"
GAS_PRICE="0.0025"
export DAEMON_NAME="braind"
export DAEMON_HOME="$HOME/$BINARY_DIR"

if ! command -v cosmovisor &>/dev/null; then
  echo "\n\ncosmovisor could not be found"
  exit
fi



# submit upgrade proposal
$DAEMON_NAME tx gov submit-proposal software-upgrade $UPGRADE_NAME --title "Upgrade to $UPGRADE_NAME" --description "Upgrade to $UPGRADE_NAME" --deposit 10000000$DENOM --upgrade-height $HEIGHT --from me --chain-id $CHAIN_ID --keyring-backend test --home $DAEMON_HOME --node tcp://localhost:26657 --yes --gas-prices $GAS_PRICE$DENOM --gas auto --gas-adjustment 1.5 --broadcast-mode block

# vote on the proposal
$DAEMON_NAME tx gov vote 1 yes --from me --chain-id $CHAIN_ID --keyring-backend test --home $DAEMON_HOME --node tcp://localhost:26657 --yes --gas-prices $GAS_PRICE$DENOM --gas auto --gas-adjustment 1.5 --broadcast-mode block


make install
mkdir -p $DAEMON_HOME/cosmovisor/upgrades/$UPGRADE_NAME/bin/
cp `which $DAEMON_NAME` $DAEMON_HOME/cosmovisor/upgrades/$UPGRADE_NAME/bin/