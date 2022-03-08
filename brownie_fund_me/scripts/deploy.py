from brownie import FundMe, MockV3Aggregator, network, config
from scripts.helpful_script import (
    get_account,
    deploy_mock,
    LOCAL_BLOCKCHAIN_ENVIRONMENT,
)


def deploy_fund_me():
    account = get_account()

    # if we are on the persistent network like rinkeby use the associated address
    # otherwise deploy mocks

    if network.show_active() not in LOCAL_BLOCKCHAIN_ENVIRONMENT:
        # if not a development network then get address of eth price from network rinkeby

        price_feed_address = config["networks"][network.show_active()][
            "eth_usd_price_feed"
        ]
        # if is development network deploy mock aggregator contract
    else:
        deploy_mock()
        price_feed_address = MockV3Aggregator[-1].address

    fund_me = FundMe.deploy(
        price_feed_address,
        {"from": account},
        publish_source=config["networks"][network.show_active()].get("verify"),
    )
    print(f"contract deployed to {fund_me.address}")
    print(network.show_active())
    return fund_me


def main():
    deploy_fund_me()
