from brownie import SimpleStorage, accounts, config


def read_contract():
    simple_storage = SimpleStorage[-1]

    # brownie already know address and ABI

    print(simple_storage.retrieve())


def main():
    read_contract()
