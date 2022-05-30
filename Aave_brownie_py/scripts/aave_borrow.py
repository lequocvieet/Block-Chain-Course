from brownie import network, config, interface
from web3 import Web3
from scripts.helpful_script import LOCAL_BLOCKCHAIN_ENVIRONMENTS, get_account
from scripts.get_weth import get_weth


# 0.05
amount = Web3.toWei(0.1, "Ether")


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
        erc20_address,
        amount,
        account.address,
        0,
        {"from": account},
    )
    tx.wait(1)
    print("Deposited!")

    # After deposit your collateral how much money do you want to borrow from aave?
    borrowable_eth, total_debt_eth = get_borrowable_data(lending_pool, account)
    print("Let's borrow!")
    # Dai in term of ETH
    dai_eth_price = get_asset_price(
        config["networks"][network.show_active()]["dai_eth_price_feed"]
    )
    amount_dai_to_borrow = (1 / dai_eth_price) * (borrowable_eth * 0.25)
    # convert borrowable_eth to borrowable_dai * 95% to prevent liquidation
    print(f"We are going to borrow {amount_dai_to_borrow}  Dai")

    # Now we will borrow!
    dai_address = config["networks"][network.show_active()]["dai_token"]
    print(f"Dai_address: {dai_address}")
    borrow_tx = lending_pool.borrow(
        dai_address,
        Web3.toWei(amount_dai_to_borrow, "Ether"),
        1,
        0,
        account.address,
        {"from": account},
    )
    borrow_tx.wait(1)
    print("We borrowed some Dai!")
    get_borrowable_data(lending_pool, account)
    repay_all(amount, lending_pool, account)
    get_borrowable_data(lending_pool, account)
    print(
        "You have Done 4 things "
        + "-Swap ETH to WETH "
        + "-Deposit 0.1 ETH to Aave as Collaterals "
        + "-Borrow Dai From those Collaterals "
        + "-Repay all of money"
    )


def repay_all(amount, lending_pool, account):
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
    repay_tx.wait(1)
    print("Repayed Successfull!")


def get_asset_price(price_feed_address):
    # ABI
    # Address
    dai_eth_price_feed = interface.AggregatorV3Interface(price_feed_address)
    latest_price = dai_eth_price_feed.latestRoundData()[1]
    # [1] make dai_eth_price_feed.latestRoundData() return value at index 1
    convert_latest_price = Web3.fromWei(latest_price, "Ether")
    print(f"The DAI/ETH price is {convert_latest_price}")
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
    tx = erc20.approve(
        spender,
        amount,
        {"from": account},
    )
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
