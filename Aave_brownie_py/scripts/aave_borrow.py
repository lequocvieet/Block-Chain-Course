from brownie import network, config, interface
from scripts.helpful_script import LOCAL_BLOCKCHAIN_ENVIRONMENTS, get_account
from scripts.get_weth import get_weth


def main():
    account = get_account()
    erc20_address = config["networks"][network.show_active()]["weth_token"]
    if network.show_active() in ["mainnet-fork"]:
        get_weth()

    # ABI
    # Address
    lending_pool = get_lending_pool()


def get_lending_pool():
    # ABI
    # Address
    lending_pool_addresses_provider = interface.ILendingPoolAddressesProvider(
        config["networks"][network.show_active()]["lending_pool_addresses_provider"]
        # create lending pool provider interface
    )
    # get lending pool address
    lending_pool_address = lending_pool_addresses_provider.getLendingPool()
    # ABI
    # Address-check!
