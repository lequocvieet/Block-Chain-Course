from distutils.command.config import config
from scripts.helpful_script import get_account, get_contract
from brownie import Lottery, network, config


def deploy_lottery():
    account = get_account()
    lottery = Lottery.deploy(
        get_contract("eth_usd_price_feed").address,
        get_contract("vrf_coordinator").address,
        get_contract("link_token").address,
        config["networks"][network.show_active()]["_subscriptionId"],
        config["networks"][network.show_active()]["_keyhash"],
        config["networks"][network.show_active()]["_callbackGasLimit"],
        config["networks"][network.show_active()]["_requestConfirmations"],
        config["networks"][network.show_active()]["_numWords"],
        {"from": account},
        publish_source=config["networks"][network.show_active()].get("verify", False),
    )
    print("deployed lottery!")


def main():
    deploy_lottery()
