from brownie import accounts, config, SimpleStorage, network


def deploy_simple_storage():
    account = get_account()
    #  print(account)
    # account = accounts.add(config["wallets"]["from_key"])
    simple_storage = SimpleStorage.deploy({"from": account})
    print(simple_storage)
    # after deploy let test a view function to interact with this contract
    stored_value = simple_storage.retrieve()
    print(stored_value)

    # let's use another funtions that make state change

    transaction = simple_storage.store(4, {"from": account})
    transaction.wait(1)
    updated_store_number = simple_storage.retrieve()
    print(updated_store_number)


def get_account():
    if network.show_active() == "development":
        return accounts[0]
    else:
        return accounts.add(config["wallets"]["from_key"])


def main():
    deploy_simple_storage()
