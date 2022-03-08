from brownie import network, config, accounts, MockV3Aggregator
from web3 import Web3

DECIMALS = 8
STARTING_VALUE = 20000000000
LOCAL_BLOCKCHAIN_ENVIRONMENT = ["development", "ganache-local"]
FORKED_LOCAL_ENVIRONMENT = ["mainnet-fork", "mainnet-fork-dev"]


def get_account():
    if (
        network.show_active() in LOCAL_BLOCKCHAIN_ENVIRONMENT
        or network.show_active() in FORKED_LOCAL_ENVIRONMENT
    ):
        return accounts[0]
    else:
        return accounts.add(config["wallets"]["from_key"])


def deploy_mock():
    print(f"Deploying mock...")
    print(f"The active network is {network.show_active()}")  # show network name
    if len(MockV3Aggregator) <= 0:
        MockV3Aggregator.deploy(
            DECIMALS, Web3.toWei(STARTING_VALUE, "ether"), {"from": get_account()}
        )

        print("Mocks Deployed!")
