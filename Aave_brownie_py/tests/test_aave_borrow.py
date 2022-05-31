from scripts.aave_borrow import (
    get_asset_price,
    get_lending_pool,
    approve_erc20,
    get_account,
    repay_all,
    get_borrowable_data,
)
from scripts.get_weth import get_weth
from brownie import config, network
from web3 import Web3


def test_get_asset_price():
    # Arrange / Act
    asset_price = get_asset_price(
        config["networks"][network.show_active()]["dai_eth_price_feed"]
    )
    # Assert
    assert asset_price > 0


def test_get_lending_pool():
    # Arrange / Act
    lending_pool = get_lending_pool()
    # Assert
    assert lending_pool is not None


def test_repay_all():
    # Arrange
    account = get_account()
    lending_pool = get_lending_pool()
    amount = Web3.toWei(0.1, "Ether")
    erc20_address = config["networks"][network.show_active()]["weth_token"]
    get_weth()
    approve_erc20(amount, lending_pool.address, erc20_address, account)
    lending_pool.deposit(
        erc20_address,
        amount,
        account.address,
        0,
        {"from": account},
    )
    borrowable_eth, total_debt_eth = get_borrowable_data(lending_pool, account)
    print("Let's borrow!")
    # Dai in term of ETH
    dai_eth_price = get_asset_price(
        config["networks"][network.show_active()]["dai_eth_price_feed"]
    )
    amount_dai_to_borrow = (1 / dai_eth_price) * (borrowable_eth * 0.95)
    # convert borrowable_eth to borrowable_dai * 95% to prevent liquidation
    print(f"We are going to borrow {amount_dai_to_borrow}  Dai")

    # Now we will borrow!
    dai_address = config["networks"][network.show_active()]["dai_token"]
    print(f"Dai_address: {dai_address}")
    lending_pool.borrow(
        dai_address,
        Web3.toWei(amount_dai_to_borrow, "Ether"),
        1,
        0,
        account.address,
        {"from": account},
    )
    print("Borrow Successfull")
    # Act
    approve_erc20(
        Web3.toWei(amount, "Ether"),
        lending_pool,
        config["networks"][network.show_active()]["dai_token"],
        account,
    )
    repay_tx = lending_pool.repay(
        config["networks"][network.show_active()]["dai_token"],
        amount,
        1,
        account.address,
        {"from": account},
    )
    print("Repay successfull")
    # Assert
    assert repay_tx is not True


def test_approve_erc20():
    # Arrange
    account = get_account()
    lending_pool = get_lending_pool()
    amount = 1000000000000000000  # 1
    erc20_address = config["networks"][network.show_active()]["weth_token"]
    # Act
    tx = approve_erc20(amount, lending_pool.address, erc20_address, account)
    # Assert
    assert tx is not True
