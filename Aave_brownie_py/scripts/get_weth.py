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
    weth = interface.IWeth(config["networks"][network.show_active()]["weth_token"])
    tx = weth.deposit({"from": account, "value": 0.05 * 10**18})
    tx.wait(1)
    print("received 0.05 WETH")
    return tx
