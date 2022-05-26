from brownie import network, config, interface
from web3 import Web3
from scripts.helpful_script import LOCAL_BLOCKCHAIN_ENVIRONMENTS, get_account
from scripts.get_weth import get_weth


# 0.05
amount = Web3.toWei(0.05, "Ether")


def main():
    account = get_account()
    erc20_address = config["networks"][network.show_active()]["weth_token"]
    if network.show_active() in ["mainnet-fork"]:
        get_weth()

    # ABI
    # Address
    lending_pool = get_lending_pool()

    # Approve sending out ERC20 Token
    approve_erc20(amount, lending_pool.address, erc20_address, account)
    print("Depositing...")
    tx = lending_pool.deposit(
        erc20_address, amount, account.address, 0, {"from": account}
    )
    tx.wait(1)
    print("Deposited!")

    # After deposit your collateral how much money do you want to borrow from aave?
    borrowable_eth, total_debt_eth = get_borrowable_data(lending_pool, account)
    print("Let's borrow!")
    # USDT in term of ETH
    usdt_eth_price = get_asset_price(
        config["networks"][network.show_active()]["usdt_eth_price_feed"]
    )
    amount_usdt_to_borrow = (1 / usdt_eth_price) * (borrowable_eth * 0.95)
    # convert borrowable_eth to borrowable_usdt * 95% to prevent liquidation
    print(f"We are going to borrow {amount_usdt_to_borrow}  USDT")

    # Now we will borrow!


def get_asset_price(price_feed_address):
    # ABI
    # Address
    usdt_eth_price_feed = interface.AggregatorV3Interface(price_feed_address)
    latest_price = usdt_eth_price_feed.latestRoundData()[1]
    # [1] make usdt_eth_price_feed.latestRoundData() return value at index 1
    convert_latest_price = Web3.fromWei(latest_price, "Ether")
    print(f"The USDT/ETH price is {convert_latest_price}")
    return float(convert_latest_price)


def get_borrowable_data(lending_pool, account):
    (
        total_collateral_eth,
        total_debt_eth,
        available_borrow_eth,
        current_liquidation_threshold,
        ltv,
        health_factor,
    ) = lending_pool.getUserAccountData(account.address)
    available_borrow_eth = Web3.fromWei(available_borrow_eth, "Ether")
    total_collateral_eth = Web3.fromWei(total_collateral_eth, "Ether")
    total_debt_eth = Web3.fromWei(total_debt_eth, "Ether")
    print(f"you have {total_collateral_eth} Ether as collateral")
    print(f"you can borrow {available_borrow_eth} Ether from aave")
    print(f"you have {total_debt_eth} Ether as debt")
    return (float(available_borrow_eth), float(total_debt_eth))


def approve_erc20(amount, spender, erc20_address, account):
    # ABI
    # Address
    print("Approving ERC 20 token...")
    erc20 = interface.IERC20(erc20_address)
    tx = erc20.approve(spender, amount, {"from": account})
    tx.wait(1)
    print("Approved!")
    return tx


def get_lending_pool():
    # ABI
    # Address
    lending_pool_addresses_provider = interface.ILendingPoolAddressesProvider(
        config["networks"][network.show_active()]["lending_pool_addresses_provider"]
        # create lending pool provider interface(contract)
    )
    # get lending pool address
    lending_pool_address = lending_pool_addresses_provider.getLendingPool()
    # ABI
    # Address-check!
    lending_pool = interface.ILendingPool(lending_pool_address)
    # create leding_pool interface(contract)

    print(lending_pool_addresses_provider)
    print(lending_pool_address)
    print(lending_pool)

    return lending_pool
