from distutils.command.config import config
from scripts.helpful_script import get_account
from brownie import interface, config, network


def main():
    get_weth()


def get_weth():

    """
    Mints WETHS  by deposit ETH
    """
    # ABI
    # Address

    account = get_account()
    weth = interface.IWeth(config["networks"][network.show_active()]["wet_token"])
    tx = weth.deposit({"from": account, "value": 0.1 * 10**18})
    print("received 0.1 WETH")
    pass
