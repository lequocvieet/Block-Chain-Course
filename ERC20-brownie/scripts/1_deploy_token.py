from brownie import OurToken
from scripts.helpful_script import get_account
from web3 import Web3

initial_supply = Web3.toWei(1000, "Ether")


def main():
    account = get_account()
    our_token = OurToken.deploy({"from": account})
    print(our_token.name())
