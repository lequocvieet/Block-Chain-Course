from solcx import compile_standard, install_solc
import json
from web3 import Web3

# Web3.py is a python library for interact with ethereum
import os
from dotenv import load_dotenv

load_dotenv()


with open("./SimpleStorage.sol", "r") as file:
    simple_storage_file = file.read()


# compile our solidity
# Solidity source code
install_solc("0.6.0")

compiled_sol = compile_standard(
    {
        "language": "Solidity",
        "sources": {"SimpleStorage.sol": {"content": simple_storage_file}},
        "settings": {
            "outputSelection": {
                "*": {
                    "*": ["abi", "metadata", "evm.bytecode", "evm.bytecode.sourceMap"]
                }
            }
        },
    },
    solc_version="0.6.0",
)
with open("compiled_code.json", "w") as file:
    json.dump(compiled_sol, file)

    # get bytecode
bytecode = compiled_sol["contracts"]["SimpleStorage.sol"]["SimpleStorage"]["evm"][
    "bytecode"
]["object"]

# get abi
abi = json.loads(
    compiled_sol["contracts"]["SimpleStorage.sol"]["SimpleStorage"]["metadata"]
)["output"]["abi"]


# w3 = Web3(Web3.HTTPProvider(os.getenv("RINKEBY_RPC_URL")))
# chain_id = 4
#
# For connecting to Rinkby
w3 = Web3(
    Web3.HTTPProvider("https://rinkeby.infura.io/v3/b5c49fa17580457e8fbe8457d2ad97a4")
)
chain_id = 4
my_address = "0xf930DcD86eC4570ee16d823Ea823d23BB88B800C"
private_key = os.getenv("PRIVATE_KEY")


# Create the contract in Python
SimpleStorage = w3.eth.contract(abi=abi, bytecode=bytecode)
# Get the latest transaction
nonce = w3.eth.getTransactionCount(my_address)

# 1 Build a transactrion
# 2 Sign a transaction
# 3 Send a transaction

# Submit the transaction that deploys the contract
transaction = SimpleStorage.constructor().buildTransaction(
    {
        "chainId": chain_id,
        "gasPrice": w3.eth.gas_price,
        "from": my_address,
        "nonce": nonce,
    }
)


# Sign the transaction
signed_txn = w3.eth.account.sign_transaction(transaction, private_key=private_key)


# Send this signed transaction

print("Deploying contract!")
tx_hash = w3.eth.send_raw_transaction(signed_txn.rawTransaction)


tx_receipt = w3.eth.wait_for_transaction_receipt(tx_hash)

print("Deploy  Success!")

# Work with contract
# (contract after deploy)

simple_storage = w3.eth.contract(address=tx_receipt.contractAddress, abi=abi)
# tx_receipt.contractAddress is address of contract just deploy inside new transaction


print(simple_storage.functions.retrieve().call())


# create a new transaction to interact like an orange button in Remix
# So this is state change
print("Update contract!")
store_transaction = simple_storage.functions.store(4).buildTransaction(
    {
        "chainId": chain_id,
        "gasPrice": w3.eth.gas_price,
        "from": my_address,
        "nonce": nonce + 1,  # because each transaction have unique nonce
    }
)

singed_store_tx = w3.eth.account.sign_transaction(
    store_transaction, private_key=private_key
)

send_store_tx = w3.eth.send_raw_transaction(singed_store_tx.rawTransaction)
transaction_receipt = w3.eth.wait_for_transaction_receipt(send_store_tx)
print("Update contract success!")
print(simple_storage.functions.retrieve().call())
