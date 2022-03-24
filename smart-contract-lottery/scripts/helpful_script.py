from brownie import (
    network,
    config,
    accounts,
    MockV3Aggregator,
    Contract,
    VRFCoordinatorV2Mock,
    LinkToken,
)


LOCAL_BLOCKCHAIN_ENVIRONMENT = ["development", "ganache-local"]
FORKED_LOCAL_ENVIRONMENT = ["mainnet-fork", "mainnet-fork-dev"]


def get_account(index=None, id=None):  # 3 method to get account
    # 1 accounts[0]
    # 2 accounts.add(".env")
    # 3 accounts.load("id")

    if index:
        return accounts[index]
    if id:
        return accounts.load(id)

    if (
        network.show_active() in LOCAL_BLOCKCHAIN_ENVIRONMENT
        or network.show_active() in FORKED_LOCAL_ENVIRONMENT
    ):
        return accounts[0]

    return accounts.add(config["wallets"]["from_key"])


contract_to_mock = {
    "eth_usd_price_feed": MockV3Aggregator,
    "vrf_coordinator": VRFCoordinatorV2Mock,
    "link_token": LinkToken,
}

# mapping contract name to contract type


def get_contract(contract_name):
    """
    This function will grab the contract address from the brownie
    config if defined,otherwise it will deploy a mock version
    of that contract and return that mock contract

        Arguments:
            contract_name: string
        Return:
            brownie.network.contract.project contract: The most recently deployed
            version of this contract

    """

    contract_type = contract_to_mock[contract_name]
    if network.show_active() in LOCAL_BLOCKCHAIN_ENVIRONMENT:
        if len(contract_type) <= 0:
            # contract_type== MockV3Aggregator
            deploy_mocks()
        contract = contract_type[-1]  # MockV3Aggregator[-1]

    else:
        contract_address = config["networks"][network.show_active()][contract_name]

        # To intereact with a contract we will create an instances need 2 agurment:
        # Contract address
        # Abi

        contract = Contract.from_abi(
            contract_type._name, contract_address, contract_type.abi
        )
    return contract


DECIMALS = 8
INITIAL_VALUE = 200000000000


def deploy_mocks(decimal=DECIMALS, initial_value=INITIAL_VALUE):
    account = get_account()
    MockV3Aggregator.deploy(decimal, initial_value, {"from": account})
    LinkToken.deploy({"from": account})
    VRFCoordinatorV2Mock.deploy(10, 10, {"from": account})
    print("deployed!")
