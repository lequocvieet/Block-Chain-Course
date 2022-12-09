// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package online_test

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// OnlineTestWinner is an auto generated low-level Go binding around an user-defined struct.
type OnlineTestWinner struct {
	WinnerAddr common.Address
	SubmitTime *big.Int
}

// OnlineTestMetaData contains all meta data concerning the OnlineTest contract.
var OnlineTestMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"examCode\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"adminAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"winnerAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"submitTime\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structOnlineTest.Winner[]\",\"name\":\"winners\",\"type\":\"tuple[]\"}],\"name\":\"CaculateWinner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"adminAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"adminEntranceFee\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractOwner\",\"type\":\"address\"}],\"name\":\"ChooseAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"examCode\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"adminAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"userSubmitFee\",\"type\":\"uint256\"}],\"name\":\"CreateExam\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"examCode\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"adminAddress\",\"type\":\"address\"}],\"name\":\"EndSubmit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"examCode\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"adminAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"answer\",\"type\":\"string\"}],\"name\":\"SubmitAnswer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"examCode\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"answer\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"submitTime\",\"type\":\"uint256\"}],\"name\":\"SubmitTest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"examFund\",\"type\":\"uint256\"}],\"name\":\"WithDraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"adminAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"adminEntranceFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_examCode\",\"type\":\"uint256\"}],\"name\":\"caculateWinner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_adminAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_adminEntranceFee\",\"type\":\"uint256\"}],\"name\":\"chooseAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_reward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_userSubmitFee\",\"type\":\"uint256\"}],\"name\":\"createExams\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_examCode\",\"type\":\"uint256\"}],\"name\":\"endSubmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"examCode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"examCodes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"examFund\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"exams\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"adminAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"answer\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"userSubmitFee\",\"type\":\"uint256\"},{\"internalType\":\"enumOnlineTest.EXAM_STATE\",\"name\":\"examState\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAccountBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_examCode\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_answer\",\"type\":\"string\"}],\"name\":\"submitAnswer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_examCode\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_answer\",\"type\":\"string\"}],\"name\":\"submitTest\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userExams\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"userAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"examCode\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"answer\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"submitTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6080604052600060055534801561001557600080fd5b50600080546001600160a01b031916331790556119fe806100376000396000f3fe6080604052600436106101025760003560e01c80635ca0b3f311610095578063931486c611610064578063931486c6146102835780639bf2871f146102a3578063ce606ee0146102b9578063d0c9e799146102d9578063e567a4e0146102ef57600080fd5b80635ca0b3f3146101e75780636896fabf146102075780637cb2bb281461021b578063818305931461024b57600080fd5b80632e3f45af116100d15780632e3f45af146101785780633ccfd60b146101985780634198cb3f146101a0578063581c591b146101d157600080fd5b806301fa3bec1461010e57806312065fe0146101305780631506a8581461015257806317029f161461016557600080fd5b3661010957005b600080fd5b34801561011a57600080fd5b5061012e6101293660046114a7565b61030f565b005b34801561013c57600080fd5b50475b6040519081526020015b60405180910390f35b61012e610160366004611562565b610476565b61012e6101733660046114a7565b6106d1565b34801561018457600080fd5b5061013f610193366004611584565b6109a7565b61012e6109c8565b3480156101ac57600080fd5b506101c06101bb366004611584565b610aad565b6040516101499594939291906115f9565b3480156101dd57600080fd5b5061013f60035481565b3480156101f357600080fd5b5061012e610202366004611584565b610b72565b34801561021357600080fd5b50333161013f565b34801561022757600080fd5b5061023b610236366004611584565b61111b565b6040516101499493929190611656565b34801561025757600080fd5b5060015461026b906001600160a01b031681565b6040516001600160a01b039091168152602001610149565b34801561028f57600080fd5b5061012e61029e36600461168e565b6111ea565b3480156102af57600080fd5b5061013f60055481565b3480156102c557600080fd5b5060005461026b906001600160a01b031681565b3480156102e557600080fd5b5061013f60025481565b3480156102fb57600080fd5b5061012e61030a366004611584565b6112b9565b6000828152600960205260409020546001600160a01b031633146103895760405162461bcd60e51b815260206004820152602660248201527f596f75206e65656420746f2062652041646d696e20746f207375626d697420616044820152656e737765722160d01b60648201526084015b60405180910390fd5b600260008381526009602052604090206004015460ff1660038111156103b1576103b16115e3565b146103fe5760405162461bcd60e51b815260206004820152601e60248201527f41646d696e2043616e2774207375626d697420616e73776572207965742100006044820152606401610380565b6000828152600960205260409020600201610419828261174e565b5060008281526009602052604090819020600401805460ff191660031790555133907f9cf87a9c6288a3c37b26ac5bb24282131251f4e24a22c0d141992a7c7d621a199061046a908590859061180e565b60405180910390a25050565b6001546001600160a01b031661048b57600080fd5b60025460000361049a57600080fd5b6001546001600160a01b031633146104f45760405162461bcd60e51b815260206004820181905260248201527f596f75206d7573742062652041646d696e20746f20437265617465204578616d6044820152606401610380565b6002546105019083611845565b3331116105675760405162461bcd60e51b815260206004820152602e60248201527f4e6f7420656e6f75676820457468657220746f20437265617465204578616d2060448201526d30b732102830bc902bb4b73732b960911b6064820152608401610380565b6002543410156105df5760405162461bcd60e51b815260206004820152603860248201527f56616c75652061646d696e2073656e6420746f20636f6e7472616374206d757360448201527f742067726561746572207468616e2061646d696e2066656500000000000000006064820152608401610380565b600580549060006105ef8361185e565b90915550506005805460048054600180820183557f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b90910183905560009283526009602052604080842080546001600160a01b031916331790558454845280842090910186905583548352808320909101805460ff19169055915481529081206003908101839055600254815490929061068a908490611845565b90915550506005546040805191825260208201849052810182905233907fbd6efe366c15cfbe5513653fb28b66c348c6729014b082f21f6a925ab1d8e10f9060600161046a565b60008281526009602052604081206004015460ff1660038111156106f7576106f76115e3565b146107395760405162461bcd60e51b815260206004820152601260248201527143616e6e6f74207375626d6974207965742160701b6044820152606401610380565b600082815260096020526040902060030154333110156107a95760405162461bcd60e51b815260206004820152602560248201527f4e6f7420656e6f75676820457468657220746f207375626d697420796f7572206044820152646578616d2160d81b6064820152608401610380565b6000828152600960205260409020600301543410156108305760405162461bcd60e51b815260206004820152603d60248201527f56616c756520757365722073656e6420746f20636f6e7472616374206d75737460448201527f2067726561746572207468616e2075736572207375626d6974206665650000006064820152608401610380565b610864604051806080016040528060006001600160a01b031681526020016000815260200160608152602001600081525090565b33815260208101838152604082018381524260608401526008805460018101825560009190915283517ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee3600490920291820180546001600160a01b0319166001600160a01b0390921691909117815592517ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee482015590518392917ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee5019061092b908261174e565b5060609190910151600391820155600084815260096020526040812082015482549092919061095b908490611845565b9091555050606081015160405133917f10adfb6a5ccb5230714b769957c062278bad4ef324985b45a0282120fb94f3dd9161099a918791879190611877565b60405180910390a2505050565b600481815481106109b757600080fd5b600091825260209091200154905081565b6000546001600160a01b03163314610a385760405162461bcd60e51b815260206004820152602d60248201527f596f75206d75737420626520636f6e7472616374206f776e657220746f20776960448201526c746864726177206d6f6e65792160981b6064820152608401610380565b600080546040516001600160a01b03909116914780156108fc02929091818181858888f19350505050158015610a72573d6000803e3d6000fd5b5060035460405190815233907f14b43ca4c63c5423006ad978eab8a14386476b52f1d4728070ea20a36f13e83b9060200160405180910390a2565b6009602052600090815260409020805460018201546002830180546001600160a01b03909316939192610adf906116c6565b80601f0160208091040260200160405190810160405280929190818152602001828054610b0b906116c6565b8015610b585780601f10610b2d57610100808354040283529160200191610b58565b820191906000526020600020905b815481529060010190602001808311610b3b57829003601f168201915b50505050600383015460049093015491929160ff16905085565b6000818152600960205260409020546001600160a01b03163314610bd85760405162461bcd60e51b815260206004820152601f60248201527f4f6e6c792061646d696e2063616e20636163756c6174652077696e6e657221006044820152606401610380565b600360008281526009602052604090206004015460ff166003811115610c0057610c006115e3565b14610c575760405162461bcd60e51b815260206004820152602160248201527f41646d696e20646f6573206e6f74207375626d697420616e73776572207965746044820152602160f81b6064820152608401610380565b60005b600854811015610e18578160088281548110610c7857610c786118a0565b90600052602060002090600402016001015403610e06576000828152600960209081526040918290209151610cb19260020191016118b6565b6040516020818303038152906040528051906020012060088281548110610cda57610cda6118a0565b9060005260206000209060040201600201604051602001610cfb91906118b6565b6040516020818303038152906040528051906020012003610e0657604080518082019091526000808252602082015260088281548110610d3d57610d3d6118a0565b60009182526020909120600490910201546001600160a01b031681526008805483908110610d6d57610d6d6118a0565b600091825260208083206004929092029091016003015490830190815260068054600181018255925291517ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f600290920291820180546001600160a01b0319166001600160a01b0390921691909117905590517ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d40909101555b80610e108161185e565b915050610c5a565b506000805b600654610e2c9060019061192c565b811015610f60576000610e40826001611845565b90505b600654811015610f4d5760068181548110610e6057610e606118a0565b90600052602060002090600202016001015460068381548110610e8557610e856118a0565b9060005260206000209060020201600101541115610f3b5760068281548110610eb057610eb06118a0565b906000526020600020906002020160010154925060068181548110610ed757610ed76118a0565b90600052602060002090600202016001015460068381548110610efc57610efc6118a0565b9060005260206000209060020201600101819055508260068281548110610f2557610f256118a0565b9060005260206000209060020201600101819055505b80610f458161185e565b915050610e43565b5080610f588161185e565b915050610e1d565b50600654600a1161100957600082815260096020526040812060010154610f8990600a9061193f565b905060005b600a8110156110025760068181548110610faa57610faa6118a0565b600091825260208220600290910201546040516001600160a01b039091169184156108fc02918591818181858888f19350505050158015610fef573d6000803e3d6000fd5b5080610ffa8161185e565b915050610f8e565b50506110a6565b60065460008381526009602052604081206001015490916110299161193f565b905060005b6006548110156110a3576006818154811061104b5761104b6118a0565b600091825260208220600290910201546040516001600160a01b039091169184156108fc02918591818181858888f19350505050158015611090573d6000803e3d6000fd5b508061109b8161185e565b91505061102e565b50505b336001600160a01b03167f3f88a99e75365aae8f40e52be1f5800cfc13cbfb6787ecbf91df5bfcba87f1858360066040516110e2929190611961565b60405180910390a26000828152600960205260409020600401805460ff1916600117905560078054611116916006916113f3565b505050565b6008818154811061112b57600080fd5b60009182526020909120600490910201805460018201546002830180546001600160a01b039093169450909291611161906116c6565b80601f016020809104026020016040519081016040528092919081815260200182805461118d906116c6565b80156111da5780601f106111af576101008083540402835291602001916111da565b820191906000526020600020905b8154815290600101906020018083116111bd57829003601f168201915b5050505050908060030154905084565b6000546001600160a01b031633146112585760405162461bcd60e51b815260206004820152602b60248201527f546f2063686f6f73652041646d696e20796f75206d75737420626520636f6e7460448201526a72616374206f776e65722160a81b6064820152608401610380565b600180546001600160a01b0319166001600160a01b0384811691821790925560028390556000546040518481529216917f408ed3a4f20bb252e43a352a43ff8c6be1cd281dac72fdac9c96a0668ffca92b9060200160405180910390a35050565b6000818152600960205260409020546001600160a01b0316331461132b5760405162461bcd60e51b815260206004820152602360248201527f596f75206e65656420746f2062652041646d696e20746f20656e64207375626d60448201526269742160e81b6064820152608401610380565b60008181526009602052604081206004015460ff166003811115611351576113516115e3565b1461139e5760405162461bcd60e51b815260206004820152601b60248201527f41646d696e2043616e277420656e64207375626d6974207965742100000000006044820152606401610380565b600081815260096020908152604091829020600401805460ff19166002179055905182815233917fc52616f2c177cdd7f43dd268f7b66a6d48fd7e1c91a8c265359bd0b207cdb00d910160405180910390a250565b82805482825590600052602060002090600202810192821561145b5760005260206000209160020282015b8281111561145b57825482546001600160a01b0319166001600160a01b03909116178255600180840154908301556002928301929091019061141e565b5061146792915061146b565b5090565b5b808211156114675780546001600160a01b03191681556000600182015560020161146c565b634e487b7160e01b600052604160045260246000fd5b600080604083850312156114ba57600080fd5b82359150602083013567ffffffffffffffff808211156114d957600080fd5b818501915085601f8301126114ed57600080fd5b8135818111156114ff576114ff611491565b604051601f8201601f19908116603f0116810190838211818310171561152757611527611491565b8160405282815288602084870101111561154057600080fd5b8260208601602083013760006020848301015280955050505050509250929050565b6000806040838503121561157557600080fd5b50508035926020909101359150565b60006020828403121561159657600080fd5b5035919050565b6000815180845260005b818110156115c3576020818501810151868301820152016115a7565b506000602082860101526020601f19601f83011685010191505092915050565b634e487b7160e01b600052602160045260246000fd5b60018060a01b038616815284602082015260a06040820152600061162060a083018661159d565b90508360608301526004831061164657634e487b7160e01b600052602160045260246000fd5b8260808301529695505050505050565b60018060a01b038516815283602082015260806040820152600061167d608083018561159d565b905082606083015295945050505050565b600080604083850312156116a157600080fd5b82356001600160a01b03811681146116b857600080fd5b946020939093013593505050565b600181811c908216806116da57607f821691505b6020821081036116fa57634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111561111657600081815260208120601f850160051c810160208610156117275750805b601f850160051c820191505b8181101561174657828155600101611733565b505050505050565b815167ffffffffffffffff81111561176857611768611491565b61177c8161177684546116c6565b84611700565b602080601f8311600181146117b157600084156117995750858301515b600019600386901b1c1916600185901b178555611746565b600085815260208120601f198616915b828110156117e0578886015182559484019460019091019084016117c1565b50858210156117fe5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b828152604060208201526000611827604083018461159d565b949350505050565b634e487b7160e01b600052601160045260246000fd5b808201808211156118585761185861182f565b92915050565b6000600182016118705761187061182f565b5060010190565b838152606060208201526000611890606083018561159d565b9050826040830152949350505050565b634e487b7160e01b600052603260045260246000fd5b60008083546118c4816116c6565b600182811680156118dc57600181146118f157611920565b60ff1984168752821515830287019450611920565b8760005260208060002060005b858110156119175781548a8201529084019082016118fe565b50505082870194505b50929695505050505050565b818103818111156118585761185861182f565b60008261195c57634e487b7160e01b600052601260045260246000fd5b500490565b60006040808301858452602082818601528186548084526060870191508760005282600020935060005b818110156119ba5784546001600160a01b0316835260018581015485850152600290950194928601920161198b565b50909897505050505050505056fea26469706673582212207751584683c1ac2cdbdb0801bb00ea71c095f8e4ab9930ab6a8e66983788abf764736f6c63430008110033",
}

// OnlineTestABI is the input ABI used to generate the binding from.
// Deprecated: Use OnlineTestMetaData.ABI instead.
var OnlineTestABI = OnlineTestMetaData.ABI

// OnlineTestBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OnlineTestMetaData.Bin instead.
var OnlineTestBin = OnlineTestMetaData.Bin

// DeployOnlineTest deploys a new Ethereum contract, binding an instance of OnlineTest to it.
func DeployOnlineTest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OnlineTest, error) {
	parsed, err := OnlineTestMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OnlineTestBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OnlineTest{OnlineTestCaller: OnlineTestCaller{contract: contract}, OnlineTestTransactor: OnlineTestTransactor{contract: contract}, OnlineTestFilterer: OnlineTestFilterer{contract: contract}}, nil
}

// OnlineTest is an auto generated Go binding around an Ethereum contract.
type OnlineTest struct {
	OnlineTestCaller     // Read-only binding to the contract
	OnlineTestTransactor // Write-only binding to the contract
	OnlineTestFilterer   // Log filterer for contract events
}

// OnlineTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type OnlineTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OnlineTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OnlineTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OnlineTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OnlineTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OnlineTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OnlineTestSession struct {
	Contract     *OnlineTest       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OnlineTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OnlineTestCallerSession struct {
	Contract *OnlineTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// OnlineTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OnlineTestTransactorSession struct {
	Contract     *OnlineTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// OnlineTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type OnlineTestRaw struct {
	Contract *OnlineTest // Generic contract binding to access the raw methods on
}

// OnlineTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OnlineTestCallerRaw struct {
	Contract *OnlineTestCaller // Generic read-only contract binding to access the raw methods on
}

// OnlineTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OnlineTestTransactorRaw struct {
	Contract *OnlineTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOnlineTest creates a new instance of OnlineTest, bound to a specific deployed contract.
func NewOnlineTest(address common.Address, backend bind.ContractBackend) (*OnlineTest, error) {
	contract, err := bindOnlineTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OnlineTest{OnlineTestCaller: OnlineTestCaller{contract: contract}, OnlineTestTransactor: OnlineTestTransactor{contract: contract}, OnlineTestFilterer: OnlineTestFilterer{contract: contract}}, nil
}

// NewOnlineTestCaller creates a new read-only instance of OnlineTest, bound to a specific deployed contract.
func NewOnlineTestCaller(address common.Address, caller bind.ContractCaller) (*OnlineTestCaller, error) {
	contract, err := bindOnlineTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OnlineTestCaller{contract: contract}, nil
}

// NewOnlineTestTransactor creates a new write-only instance of OnlineTest, bound to a specific deployed contract.
func NewOnlineTestTransactor(address common.Address, transactor bind.ContractTransactor) (*OnlineTestTransactor, error) {
	contract, err := bindOnlineTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OnlineTestTransactor{contract: contract}, nil
}

// NewOnlineTestFilterer creates a new log filterer instance of OnlineTest, bound to a specific deployed contract.
func NewOnlineTestFilterer(address common.Address, filterer bind.ContractFilterer) (*OnlineTestFilterer, error) {
	contract, err := bindOnlineTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OnlineTestFilterer{contract: contract}, nil
}

// bindOnlineTest binds a generic wrapper to an already deployed contract.
func bindOnlineTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OnlineTestABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OnlineTest *OnlineTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OnlineTest.Contract.OnlineTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OnlineTest *OnlineTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OnlineTest.Contract.OnlineTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OnlineTest *OnlineTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OnlineTest.Contract.OnlineTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OnlineTest *OnlineTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OnlineTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OnlineTest *OnlineTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OnlineTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OnlineTest *OnlineTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OnlineTest.Contract.contract.Transact(opts, method, params...)
}

// AdminAddr is a free data retrieval call binding the contract method 0x81830593.
//
// Solidity: function adminAddr() view returns(address)
func (_OnlineTest *OnlineTestCaller) AdminAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OnlineTest.contract.Call(opts, &out, "adminAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AdminAddr is a free data retrieval call binding the contract method 0x81830593.
//
// Solidity: function adminAddr() view returns(address)
func (_OnlineTest *OnlineTestSession) AdminAddr() (common.Address, error) {
	return _OnlineTest.Contract.AdminAddr(&_OnlineTest.CallOpts)
}

// AdminAddr is a free data retrieval call binding the contract method 0x81830593.
//
// Solidity: function adminAddr() view returns(address)
func (_OnlineTest *OnlineTestCallerSession) AdminAddr() (common.Address, error) {
	return _OnlineTest.Contract.AdminAddr(&_OnlineTest.CallOpts)
}

// AdminEntranceFee is a free data retrieval call binding the contract method 0xd0c9e799.
//
// Solidity: function adminEntranceFee() view returns(uint256)
func (_OnlineTest *OnlineTestCaller) AdminEntranceFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OnlineTest.contract.Call(opts, &out, "adminEntranceFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminEntranceFee is a free data retrieval call binding the contract method 0xd0c9e799.
//
// Solidity: function adminEntranceFee() view returns(uint256)
func (_OnlineTest *OnlineTestSession) AdminEntranceFee() (*big.Int, error) {
	return _OnlineTest.Contract.AdminEntranceFee(&_OnlineTest.CallOpts)
}

// AdminEntranceFee is a free data retrieval call binding the contract method 0xd0c9e799.
//
// Solidity: function adminEntranceFee() view returns(uint256)
func (_OnlineTest *OnlineTestCallerSession) AdminEntranceFee() (*big.Int, error) {
	return _OnlineTest.Contract.AdminEntranceFee(&_OnlineTest.CallOpts)
}

// ContractOwner is a free data retrieval call binding the contract method 0xce606ee0.
//
// Solidity: function contractOwner() view returns(address)
func (_OnlineTest *OnlineTestCaller) ContractOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OnlineTest.contract.Call(opts, &out, "contractOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ContractOwner is a free data retrieval call binding the contract method 0xce606ee0.
//
// Solidity: function contractOwner() view returns(address)
func (_OnlineTest *OnlineTestSession) ContractOwner() (common.Address, error) {
	return _OnlineTest.Contract.ContractOwner(&_OnlineTest.CallOpts)
}

// ContractOwner is a free data retrieval call binding the contract method 0xce606ee0.
//
// Solidity: function contractOwner() view returns(address)
func (_OnlineTest *OnlineTestCallerSession) ContractOwner() (common.Address, error) {
	return _OnlineTest.Contract.ContractOwner(&_OnlineTest.CallOpts)
}

// ExamCode is a free data retrieval call binding the contract method 0x9bf2871f.
//
// Solidity: function examCode() view returns(uint256)
func (_OnlineTest *OnlineTestCaller) ExamCode(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OnlineTest.contract.Call(opts, &out, "examCode")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExamCode is a free data retrieval call binding the contract method 0x9bf2871f.
//
// Solidity: function examCode() view returns(uint256)
func (_OnlineTest *OnlineTestSession) ExamCode() (*big.Int, error) {
	return _OnlineTest.Contract.ExamCode(&_OnlineTest.CallOpts)
}

// ExamCode is a free data retrieval call binding the contract method 0x9bf2871f.
//
// Solidity: function examCode() view returns(uint256)
func (_OnlineTest *OnlineTestCallerSession) ExamCode() (*big.Int, error) {
	return _OnlineTest.Contract.ExamCode(&_OnlineTest.CallOpts)
}

// ExamCodes is a free data retrieval call binding the contract method 0x2e3f45af.
//
// Solidity: function examCodes(uint256 ) view returns(uint256)
func (_OnlineTest *OnlineTestCaller) ExamCodes(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _OnlineTest.contract.Call(opts, &out, "examCodes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExamCodes is a free data retrieval call binding the contract method 0x2e3f45af.
//
// Solidity: function examCodes(uint256 ) view returns(uint256)
func (_OnlineTest *OnlineTestSession) ExamCodes(arg0 *big.Int) (*big.Int, error) {
	return _OnlineTest.Contract.ExamCodes(&_OnlineTest.CallOpts, arg0)
}

// ExamCodes is a free data retrieval call binding the contract method 0x2e3f45af.
//
// Solidity: function examCodes(uint256 ) view returns(uint256)
func (_OnlineTest *OnlineTestCallerSession) ExamCodes(arg0 *big.Int) (*big.Int, error) {
	return _OnlineTest.Contract.ExamCodes(&_OnlineTest.CallOpts, arg0)
}

// ExamFund is a free data retrieval call binding the contract method 0x581c591b.
//
// Solidity: function examFund() view returns(uint256)
func (_OnlineTest *OnlineTestCaller) ExamFund(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OnlineTest.contract.Call(opts, &out, "examFund")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExamFund is a free data retrieval call binding the contract method 0x581c591b.
//
// Solidity: function examFund() view returns(uint256)
func (_OnlineTest *OnlineTestSession) ExamFund() (*big.Int, error) {
	return _OnlineTest.Contract.ExamFund(&_OnlineTest.CallOpts)
}

// ExamFund is a free data retrieval call binding the contract method 0x581c591b.
//
// Solidity: function examFund() view returns(uint256)
func (_OnlineTest *OnlineTestCallerSession) ExamFund() (*big.Int, error) {
	return _OnlineTest.Contract.ExamFund(&_OnlineTest.CallOpts)
}

// Exams is a free data retrieval call binding the contract method 0x4198cb3f.
//
// Solidity: function exams(uint256 ) view returns(address adminAddr, uint256 reward, string answer, uint256 userSubmitFee, uint8 examState)
func (_OnlineTest *OnlineTestCaller) Exams(opts *bind.CallOpts, arg0 *big.Int) (struct {
	AdminAddr     common.Address
	Reward        *big.Int
	Answer        string
	UserSubmitFee *big.Int
	ExamState     uint8
}, error) {
	var out []interface{}
	err := _OnlineTest.contract.Call(opts, &out, "exams", arg0)

	outstruct := new(struct {
		AdminAddr     common.Address
		Reward        *big.Int
		Answer        string
		UserSubmitFee *big.Int
		ExamState     uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AdminAddr = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Reward = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Answer = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.UserSubmitFee = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ExamState = *abi.ConvertType(out[4], new(uint8)).(*uint8)

	return *outstruct, err

}

// Exams is a free data retrieval call binding the contract method 0x4198cb3f.
//
// Solidity: function exams(uint256 ) view returns(address adminAddr, uint256 reward, string answer, uint256 userSubmitFee, uint8 examState)
func (_OnlineTest *OnlineTestSession) Exams(arg0 *big.Int) (struct {
	AdminAddr     common.Address
	Reward        *big.Int
	Answer        string
	UserSubmitFee *big.Int
	ExamState     uint8
}, error) {
	return _OnlineTest.Contract.Exams(&_OnlineTest.CallOpts, arg0)
}

// Exams is a free data retrieval call binding the contract method 0x4198cb3f.
//
// Solidity: function exams(uint256 ) view returns(address adminAddr, uint256 reward, string answer, uint256 userSubmitFee, uint8 examState)
func (_OnlineTest *OnlineTestCallerSession) Exams(arg0 *big.Int) (struct {
	AdminAddr     common.Address
	Reward        *big.Int
	Answer        string
	UserSubmitFee *big.Int
	ExamState     uint8
}, error) {
	return _OnlineTest.Contract.Exams(&_OnlineTest.CallOpts, arg0)
}

// GetAccountBalance is a free data retrieval call binding the contract method 0x6896fabf.
//
// Solidity: function getAccountBalance() view returns(uint256)
func (_OnlineTest *OnlineTestCaller) GetAccountBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OnlineTest.contract.Call(opts, &out, "getAccountBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAccountBalance is a free data retrieval call binding the contract method 0x6896fabf.
//
// Solidity: function getAccountBalance() view returns(uint256)
func (_OnlineTest *OnlineTestSession) GetAccountBalance() (*big.Int, error) {
	return _OnlineTest.Contract.GetAccountBalance(&_OnlineTest.CallOpts)
}

// GetAccountBalance is a free data retrieval call binding the contract method 0x6896fabf.
//
// Solidity: function getAccountBalance() view returns(uint256)
func (_OnlineTest *OnlineTestCallerSession) GetAccountBalance() (*big.Int, error) {
	return _OnlineTest.Contract.GetAccountBalance(&_OnlineTest.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_OnlineTest *OnlineTestCaller) GetBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OnlineTest.contract.Call(opts, &out, "getBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_OnlineTest *OnlineTestSession) GetBalance() (*big.Int, error) {
	return _OnlineTest.Contract.GetBalance(&_OnlineTest.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_OnlineTest *OnlineTestCallerSession) GetBalance() (*big.Int, error) {
	return _OnlineTest.Contract.GetBalance(&_OnlineTest.CallOpts)
}

// UserExams is a free data retrieval call binding the contract method 0x7cb2bb28.
//
// Solidity: function userExams(uint256 ) view returns(address userAddr, uint256 examCode, string answer, uint256 submitTime)
func (_OnlineTest *OnlineTestCaller) UserExams(opts *bind.CallOpts, arg0 *big.Int) (struct {
	UserAddr   common.Address
	ExamCode   *big.Int
	Answer     string
	SubmitTime *big.Int
}, error) {
	var out []interface{}
	err := _OnlineTest.contract.Call(opts, &out, "userExams", arg0)

	outstruct := new(struct {
		UserAddr   common.Address
		ExamCode   *big.Int
		Answer     string
		SubmitTime *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.UserAddr = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.ExamCode = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Answer = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.SubmitTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserExams is a free data retrieval call binding the contract method 0x7cb2bb28.
//
// Solidity: function userExams(uint256 ) view returns(address userAddr, uint256 examCode, string answer, uint256 submitTime)
func (_OnlineTest *OnlineTestSession) UserExams(arg0 *big.Int) (struct {
	UserAddr   common.Address
	ExamCode   *big.Int
	Answer     string
	SubmitTime *big.Int
}, error) {
	return _OnlineTest.Contract.UserExams(&_OnlineTest.CallOpts, arg0)
}

// UserExams is a free data retrieval call binding the contract method 0x7cb2bb28.
//
// Solidity: function userExams(uint256 ) view returns(address userAddr, uint256 examCode, string answer, uint256 submitTime)
func (_OnlineTest *OnlineTestCallerSession) UserExams(arg0 *big.Int) (struct {
	UserAddr   common.Address
	ExamCode   *big.Int
	Answer     string
	SubmitTime *big.Int
}, error) {
	return _OnlineTest.Contract.UserExams(&_OnlineTest.CallOpts, arg0)
}

// CaculateWinner is a paid mutator transaction binding the contract method 0x5ca0b3f3.
//
// Solidity: function caculateWinner(uint256 _examCode) returns()
func (_OnlineTest *OnlineTestTransactor) CaculateWinner(opts *bind.TransactOpts, _examCode *big.Int) (*types.Transaction, error) {
	return _OnlineTest.contract.Transact(opts, "caculateWinner", _examCode)
}

// CaculateWinner is a paid mutator transaction binding the contract method 0x5ca0b3f3.
//
// Solidity: function caculateWinner(uint256 _examCode) returns()
func (_OnlineTest *OnlineTestSession) CaculateWinner(_examCode *big.Int) (*types.Transaction, error) {
	return _OnlineTest.Contract.CaculateWinner(&_OnlineTest.TransactOpts, _examCode)
}

// CaculateWinner is a paid mutator transaction binding the contract method 0x5ca0b3f3.
//
// Solidity: function caculateWinner(uint256 _examCode) returns()
func (_OnlineTest *OnlineTestTransactorSession) CaculateWinner(_examCode *big.Int) (*types.Transaction, error) {
	return _OnlineTest.Contract.CaculateWinner(&_OnlineTest.TransactOpts, _examCode)
}

// ChooseAdmin is a paid mutator transaction binding the contract method 0x931486c6.
//
// Solidity: function chooseAdmin(address _adminAddr, uint256 _adminEntranceFee) returns()
func (_OnlineTest *OnlineTestTransactor) ChooseAdmin(opts *bind.TransactOpts, _adminAddr common.Address, _adminEntranceFee *big.Int) (*types.Transaction, error) {
	return _OnlineTest.contract.Transact(opts, "chooseAdmin", _adminAddr, _adminEntranceFee)
}

// ChooseAdmin is a paid mutator transaction binding the contract method 0x931486c6.
//
// Solidity: function chooseAdmin(address _adminAddr, uint256 _adminEntranceFee) returns()
func (_OnlineTest *OnlineTestSession) ChooseAdmin(_adminAddr common.Address, _adminEntranceFee *big.Int) (*types.Transaction, error) {
	return _OnlineTest.Contract.ChooseAdmin(&_OnlineTest.TransactOpts, _adminAddr, _adminEntranceFee)
}

// ChooseAdmin is a paid mutator transaction binding the contract method 0x931486c6.
//
// Solidity: function chooseAdmin(address _adminAddr, uint256 _adminEntranceFee) returns()
func (_OnlineTest *OnlineTestTransactorSession) ChooseAdmin(_adminAddr common.Address, _adminEntranceFee *big.Int) (*types.Transaction, error) {
	return _OnlineTest.Contract.ChooseAdmin(&_OnlineTest.TransactOpts, _adminAddr, _adminEntranceFee)
}

// CreateExams is a paid mutator transaction binding the contract method 0x1506a858.
//
// Solidity: function createExams(uint256 _reward, uint256 _userSubmitFee) payable returns()
func (_OnlineTest *OnlineTestTransactor) CreateExams(opts *bind.TransactOpts, _reward *big.Int, _userSubmitFee *big.Int) (*types.Transaction, error) {
	return _OnlineTest.contract.Transact(opts, "createExams", _reward, _userSubmitFee)
}

// CreateExams is a paid mutator transaction binding the contract method 0x1506a858.
//
// Solidity: function createExams(uint256 _reward, uint256 _userSubmitFee) payable returns()
func (_OnlineTest *OnlineTestSession) CreateExams(_reward *big.Int, _userSubmitFee *big.Int) (*types.Transaction, error) {
	return _OnlineTest.Contract.CreateExams(&_OnlineTest.TransactOpts, _reward, _userSubmitFee)
}

// CreateExams is a paid mutator transaction binding the contract method 0x1506a858.
//
// Solidity: function createExams(uint256 _reward, uint256 _userSubmitFee) payable returns()
func (_OnlineTest *OnlineTestTransactorSession) CreateExams(_reward *big.Int, _userSubmitFee *big.Int) (*types.Transaction, error) {
	return _OnlineTest.Contract.CreateExams(&_OnlineTest.TransactOpts, _reward, _userSubmitFee)
}

// EndSubmit is a paid mutator transaction binding the contract method 0xe567a4e0.
//
// Solidity: function endSubmit(uint256 _examCode) returns()
func (_OnlineTest *OnlineTestTransactor) EndSubmit(opts *bind.TransactOpts, _examCode *big.Int) (*types.Transaction, error) {
	return _OnlineTest.contract.Transact(opts, "endSubmit", _examCode)
}

// EndSubmit is a paid mutator transaction binding the contract method 0xe567a4e0.
//
// Solidity: function endSubmit(uint256 _examCode) returns()
func (_OnlineTest *OnlineTestSession) EndSubmit(_examCode *big.Int) (*types.Transaction, error) {
	return _OnlineTest.Contract.EndSubmit(&_OnlineTest.TransactOpts, _examCode)
}

// EndSubmit is a paid mutator transaction binding the contract method 0xe567a4e0.
//
// Solidity: function endSubmit(uint256 _examCode) returns()
func (_OnlineTest *OnlineTestTransactorSession) EndSubmit(_examCode *big.Int) (*types.Transaction, error) {
	return _OnlineTest.Contract.EndSubmit(&_OnlineTest.TransactOpts, _examCode)
}

// SubmitAnswer is a paid mutator transaction binding the contract method 0x01fa3bec.
//
// Solidity: function submitAnswer(uint256 _examCode, string _answer) returns()
func (_OnlineTest *OnlineTestTransactor) SubmitAnswer(opts *bind.TransactOpts, _examCode *big.Int, _answer string) (*types.Transaction, error) {
	return _OnlineTest.contract.Transact(opts, "submitAnswer", _examCode, _answer)
}

// SubmitAnswer is a paid mutator transaction binding the contract method 0x01fa3bec.
//
// Solidity: function submitAnswer(uint256 _examCode, string _answer) returns()
func (_OnlineTest *OnlineTestSession) SubmitAnswer(_examCode *big.Int, _answer string) (*types.Transaction, error) {
	return _OnlineTest.Contract.SubmitAnswer(&_OnlineTest.TransactOpts, _examCode, _answer)
}

// SubmitAnswer is a paid mutator transaction binding the contract method 0x01fa3bec.
//
// Solidity: function submitAnswer(uint256 _examCode, string _answer) returns()
func (_OnlineTest *OnlineTestTransactorSession) SubmitAnswer(_examCode *big.Int, _answer string) (*types.Transaction, error) {
	return _OnlineTest.Contract.SubmitAnswer(&_OnlineTest.TransactOpts, _examCode, _answer)
}

// SubmitTest is a paid mutator transaction binding the contract method 0x17029f16.
//
// Solidity: function submitTest(uint256 _examCode, string _answer) payable returns()
func (_OnlineTest *OnlineTestTransactor) SubmitTest(opts *bind.TransactOpts, _examCode *big.Int, _answer string) (*types.Transaction, error) {
	return _OnlineTest.contract.Transact(opts, "submitTest", _examCode, _answer)
}

// SubmitTest is a paid mutator transaction binding the contract method 0x17029f16.
//
// Solidity: function submitTest(uint256 _examCode, string _answer) payable returns()
func (_OnlineTest *OnlineTestSession) SubmitTest(_examCode *big.Int, _answer string) (*types.Transaction, error) {
	return _OnlineTest.Contract.SubmitTest(&_OnlineTest.TransactOpts, _examCode, _answer)
}

// SubmitTest is a paid mutator transaction binding the contract method 0x17029f16.
//
// Solidity: function submitTest(uint256 _examCode, string _answer) payable returns()
func (_OnlineTest *OnlineTestTransactorSession) SubmitTest(_examCode *big.Int, _answer string) (*types.Transaction, error) {
	return _OnlineTest.Contract.SubmitTest(&_OnlineTest.TransactOpts, _examCode, _answer)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() payable returns()
func (_OnlineTest *OnlineTestTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OnlineTest.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() payable returns()
func (_OnlineTest *OnlineTestSession) Withdraw() (*types.Transaction, error) {
	return _OnlineTest.Contract.Withdraw(&_OnlineTest.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() payable returns()
func (_OnlineTest *OnlineTestTransactorSession) Withdraw() (*types.Transaction, error) {
	return _OnlineTest.Contract.Withdraw(&_OnlineTest.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_OnlineTest *OnlineTestTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OnlineTest.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_OnlineTest *OnlineTestSession) Receive() (*types.Transaction, error) {
	return _OnlineTest.Contract.Receive(&_OnlineTest.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_OnlineTest *OnlineTestTransactorSession) Receive() (*types.Transaction, error) {
	return _OnlineTest.Contract.Receive(&_OnlineTest.TransactOpts)
}

// OnlineTestCaculateWinnerIterator is returned from FilterCaculateWinner and is used to iterate over the raw logs and unpacked data for CaculateWinner events raised by the OnlineTest contract.
type OnlineTestCaculateWinnerIterator struct {
	Event *OnlineTestCaculateWinner // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OnlineTestCaculateWinnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnlineTestCaculateWinner)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OnlineTestCaculateWinner)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OnlineTestCaculateWinnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OnlineTestCaculateWinnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OnlineTestCaculateWinner represents a CaculateWinner event raised by the OnlineTest contract.
type OnlineTestCaculateWinner struct {
	ExamCode     *big.Int
	AdminAddress common.Address
	Winners      []OnlineTestWinner
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCaculateWinner is a free log retrieval operation binding the contract event 0x3f88a99e75365aae8f40e52be1f5800cfc13cbfb6787ecbf91df5bfcba87f185.
//
// Solidity: event CaculateWinner(uint256 examCode, address indexed adminAddress, (address,uint256)[] winners)
func (_OnlineTest *OnlineTestFilterer) FilterCaculateWinner(opts *bind.FilterOpts, adminAddress []common.Address) (*OnlineTestCaculateWinnerIterator, error) {

	var adminAddressRule []interface{}
	for _, adminAddressItem := range adminAddress {
		adminAddressRule = append(adminAddressRule, adminAddressItem)
	}

	logs, sub, err := _OnlineTest.contract.FilterLogs(opts, "CaculateWinner", adminAddressRule)
	if err != nil {
		return nil, err
	}
	return &OnlineTestCaculateWinnerIterator{contract: _OnlineTest.contract, event: "CaculateWinner", logs: logs, sub: sub}, nil
}

// WatchCaculateWinner is a free log subscription operation binding the contract event 0x3f88a99e75365aae8f40e52be1f5800cfc13cbfb6787ecbf91df5bfcba87f185.
//
// Solidity: event CaculateWinner(uint256 examCode, address indexed adminAddress, (address,uint256)[] winners)
func (_OnlineTest *OnlineTestFilterer) WatchCaculateWinner(opts *bind.WatchOpts, sink chan<- *OnlineTestCaculateWinner, adminAddress []common.Address) (event.Subscription, error) {

	var adminAddressRule []interface{}
	for _, adminAddressItem := range adminAddress {
		adminAddressRule = append(adminAddressRule, adminAddressItem)
	}

	logs, sub, err := _OnlineTest.contract.WatchLogs(opts, "CaculateWinner", adminAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OnlineTestCaculateWinner)
				if err := _OnlineTest.contract.UnpackLog(event, "CaculateWinner", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCaculateWinner is a log parse operation binding the contract event 0x3f88a99e75365aae8f40e52be1f5800cfc13cbfb6787ecbf91df5bfcba87f185.
//
// Solidity: event CaculateWinner(uint256 examCode, address indexed adminAddress, (address,uint256)[] winners)
func (_OnlineTest *OnlineTestFilterer) ParseCaculateWinner(log types.Log) (*OnlineTestCaculateWinner, error) {
	event := new(OnlineTestCaculateWinner)
	if err := _OnlineTest.contract.UnpackLog(event, "CaculateWinner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OnlineTestChooseAdminIterator is returned from FilterChooseAdmin and is used to iterate over the raw logs and unpacked data for ChooseAdmin events raised by the OnlineTest contract.
type OnlineTestChooseAdminIterator struct {
	Event *OnlineTestChooseAdmin // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OnlineTestChooseAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnlineTestChooseAdmin)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OnlineTestChooseAdmin)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OnlineTestChooseAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OnlineTestChooseAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OnlineTestChooseAdmin represents a ChooseAdmin event raised by the OnlineTest contract.
type OnlineTestChooseAdmin struct {
	AdminAddress     common.Address
	AdminEntranceFee *big.Int
	ContractOwner    common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterChooseAdmin is a free log retrieval operation binding the contract event 0x408ed3a4f20bb252e43a352a43ff8c6be1cd281dac72fdac9c96a0668ffca92b.
//
// Solidity: event ChooseAdmin(address indexed adminAddress, uint256 adminEntranceFee, address indexed contractOwner)
func (_OnlineTest *OnlineTestFilterer) FilterChooseAdmin(opts *bind.FilterOpts, adminAddress []common.Address, contractOwner []common.Address) (*OnlineTestChooseAdminIterator, error) {

	var adminAddressRule []interface{}
	for _, adminAddressItem := range adminAddress {
		adminAddressRule = append(adminAddressRule, adminAddressItem)
	}

	var contractOwnerRule []interface{}
	for _, contractOwnerItem := range contractOwner {
		contractOwnerRule = append(contractOwnerRule, contractOwnerItem)
	}

	logs, sub, err := _OnlineTest.contract.FilterLogs(opts, "ChooseAdmin", adminAddressRule, contractOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OnlineTestChooseAdminIterator{contract: _OnlineTest.contract, event: "ChooseAdmin", logs: logs, sub: sub}, nil
}

// WatchChooseAdmin is a free log subscription operation binding the contract event 0x408ed3a4f20bb252e43a352a43ff8c6be1cd281dac72fdac9c96a0668ffca92b.
//
// Solidity: event ChooseAdmin(address indexed adminAddress, uint256 adminEntranceFee, address indexed contractOwner)
func (_OnlineTest *OnlineTestFilterer) WatchChooseAdmin(opts *bind.WatchOpts, sink chan<- *OnlineTestChooseAdmin, adminAddress []common.Address, contractOwner []common.Address) (event.Subscription, error) {

	var adminAddressRule []interface{}
	for _, adminAddressItem := range adminAddress {
		adminAddressRule = append(adminAddressRule, adminAddressItem)
	}

	var contractOwnerRule []interface{}
	for _, contractOwnerItem := range contractOwner {
		contractOwnerRule = append(contractOwnerRule, contractOwnerItem)
	}

	logs, sub, err := _OnlineTest.contract.WatchLogs(opts, "ChooseAdmin", adminAddressRule, contractOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OnlineTestChooseAdmin)
				if err := _OnlineTest.contract.UnpackLog(event, "ChooseAdmin", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseChooseAdmin is a log parse operation binding the contract event 0x408ed3a4f20bb252e43a352a43ff8c6be1cd281dac72fdac9c96a0668ffca92b.
//
// Solidity: event ChooseAdmin(address indexed adminAddress, uint256 adminEntranceFee, address indexed contractOwner)
func (_OnlineTest *OnlineTestFilterer) ParseChooseAdmin(log types.Log) (*OnlineTestChooseAdmin, error) {
	event := new(OnlineTestChooseAdmin)
	if err := _OnlineTest.contract.UnpackLog(event, "ChooseAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OnlineTestCreateExamIterator is returned from FilterCreateExam and is used to iterate over the raw logs and unpacked data for CreateExam events raised by the OnlineTest contract.
type OnlineTestCreateExamIterator struct {
	Event *OnlineTestCreateExam // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OnlineTestCreateExamIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnlineTestCreateExam)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OnlineTestCreateExam)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OnlineTestCreateExamIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OnlineTestCreateExamIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OnlineTestCreateExam represents a CreateExam event raised by the OnlineTest contract.
type OnlineTestCreateExam struct {
	ExamCode      *big.Int
	AdminAddress  common.Address
	Reward        *big.Int
	UserSubmitFee *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCreateExam is a free log retrieval operation binding the contract event 0xbd6efe366c15cfbe5513653fb28b66c348c6729014b082f21f6a925ab1d8e10f.
//
// Solidity: event CreateExam(uint256 examCode, address indexed adminAddress, uint256 reward, uint256 userSubmitFee)
func (_OnlineTest *OnlineTestFilterer) FilterCreateExam(opts *bind.FilterOpts, adminAddress []common.Address) (*OnlineTestCreateExamIterator, error) {

	var adminAddressRule []interface{}
	for _, adminAddressItem := range adminAddress {
		adminAddressRule = append(adminAddressRule, adminAddressItem)
	}

	logs, sub, err := _OnlineTest.contract.FilterLogs(opts, "CreateExam", adminAddressRule)
	if err != nil {
		return nil, err
	}
	return &OnlineTestCreateExamIterator{contract: _OnlineTest.contract, event: "CreateExam", logs: logs, sub: sub}, nil
}

// WatchCreateExam is a free log subscription operation binding the contract event 0xbd6efe366c15cfbe5513653fb28b66c348c6729014b082f21f6a925ab1d8e10f.
//
// Solidity: event CreateExam(uint256 examCode, address indexed adminAddress, uint256 reward, uint256 userSubmitFee)
func (_OnlineTest *OnlineTestFilterer) WatchCreateExam(opts *bind.WatchOpts, sink chan<- *OnlineTestCreateExam, adminAddress []common.Address) (event.Subscription, error) {

	var adminAddressRule []interface{}
	for _, adminAddressItem := range adminAddress {
		adminAddressRule = append(adminAddressRule, adminAddressItem)
	}

	logs, sub, err := _OnlineTest.contract.WatchLogs(opts, "CreateExam", adminAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OnlineTestCreateExam)
				if err := _OnlineTest.contract.UnpackLog(event, "CreateExam", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCreateExam is a log parse operation binding the contract event 0xbd6efe366c15cfbe5513653fb28b66c348c6729014b082f21f6a925ab1d8e10f.
//
// Solidity: event CreateExam(uint256 examCode, address indexed adminAddress, uint256 reward, uint256 userSubmitFee)
func (_OnlineTest *OnlineTestFilterer) ParseCreateExam(log types.Log) (*OnlineTestCreateExam, error) {
	event := new(OnlineTestCreateExam)
	if err := _OnlineTest.contract.UnpackLog(event, "CreateExam", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OnlineTestEndSubmitIterator is returned from FilterEndSubmit and is used to iterate over the raw logs and unpacked data for EndSubmit events raised by the OnlineTest contract.
type OnlineTestEndSubmitIterator struct {
	Event *OnlineTestEndSubmit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OnlineTestEndSubmitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnlineTestEndSubmit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OnlineTestEndSubmit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OnlineTestEndSubmitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OnlineTestEndSubmitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OnlineTestEndSubmit represents a EndSubmit event raised by the OnlineTest contract.
type OnlineTestEndSubmit struct {
	ExamCode     *big.Int
	AdminAddress common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterEndSubmit is a free log retrieval operation binding the contract event 0xc52616f2c177cdd7f43dd268f7b66a6d48fd7e1c91a8c265359bd0b207cdb00d.
//
// Solidity: event EndSubmit(uint256 examCode, address indexed adminAddress)
func (_OnlineTest *OnlineTestFilterer) FilterEndSubmit(opts *bind.FilterOpts, adminAddress []common.Address) (*OnlineTestEndSubmitIterator, error) {

	var adminAddressRule []interface{}
	for _, adminAddressItem := range adminAddress {
		adminAddressRule = append(adminAddressRule, adminAddressItem)
	}

	logs, sub, err := _OnlineTest.contract.FilterLogs(opts, "EndSubmit", adminAddressRule)
	if err != nil {
		return nil, err
	}
	return &OnlineTestEndSubmitIterator{contract: _OnlineTest.contract, event: "EndSubmit", logs: logs, sub: sub}, nil
}

// WatchEndSubmit is a free log subscription operation binding the contract event 0xc52616f2c177cdd7f43dd268f7b66a6d48fd7e1c91a8c265359bd0b207cdb00d.
//
// Solidity: event EndSubmit(uint256 examCode, address indexed adminAddress)
func (_OnlineTest *OnlineTestFilterer) WatchEndSubmit(opts *bind.WatchOpts, sink chan<- *OnlineTestEndSubmit, adminAddress []common.Address) (event.Subscription, error) {

	var adminAddressRule []interface{}
	for _, adminAddressItem := range adminAddress {
		adminAddressRule = append(adminAddressRule, adminAddressItem)
	}

	logs, sub, err := _OnlineTest.contract.WatchLogs(opts, "EndSubmit", adminAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OnlineTestEndSubmit)
				if err := _OnlineTest.contract.UnpackLog(event, "EndSubmit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEndSubmit is a log parse operation binding the contract event 0xc52616f2c177cdd7f43dd268f7b66a6d48fd7e1c91a8c265359bd0b207cdb00d.
//
// Solidity: event EndSubmit(uint256 examCode, address indexed adminAddress)
func (_OnlineTest *OnlineTestFilterer) ParseEndSubmit(log types.Log) (*OnlineTestEndSubmit, error) {
	event := new(OnlineTestEndSubmit)
	if err := _OnlineTest.contract.UnpackLog(event, "EndSubmit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OnlineTestSubmitAnswerIterator is returned from FilterSubmitAnswer and is used to iterate over the raw logs and unpacked data for SubmitAnswer events raised by the OnlineTest contract.
type OnlineTestSubmitAnswerIterator struct {
	Event *OnlineTestSubmitAnswer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OnlineTestSubmitAnswerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnlineTestSubmitAnswer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OnlineTestSubmitAnswer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OnlineTestSubmitAnswerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OnlineTestSubmitAnswerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OnlineTestSubmitAnswer represents a SubmitAnswer event raised by the OnlineTest contract.
type OnlineTestSubmitAnswer struct {
	ExamCode     *big.Int
	AdminAddress common.Address
	Answer       string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSubmitAnswer is a free log retrieval operation binding the contract event 0x9cf87a9c6288a3c37b26ac5bb24282131251f4e24a22c0d141992a7c7d621a19.
//
// Solidity: event SubmitAnswer(uint256 examCode, address indexed adminAddress, string answer)
func (_OnlineTest *OnlineTestFilterer) FilterSubmitAnswer(opts *bind.FilterOpts, adminAddress []common.Address) (*OnlineTestSubmitAnswerIterator, error) {

	var adminAddressRule []interface{}
	for _, adminAddressItem := range adminAddress {
		adminAddressRule = append(adminAddressRule, adminAddressItem)
	}

	logs, sub, err := _OnlineTest.contract.FilterLogs(opts, "SubmitAnswer", adminAddressRule)
	if err != nil {
		return nil, err
	}
	return &OnlineTestSubmitAnswerIterator{contract: _OnlineTest.contract, event: "SubmitAnswer", logs: logs, sub: sub}, nil
}

// WatchSubmitAnswer is a free log subscription operation binding the contract event 0x9cf87a9c6288a3c37b26ac5bb24282131251f4e24a22c0d141992a7c7d621a19.
//
// Solidity: event SubmitAnswer(uint256 examCode, address indexed adminAddress, string answer)
func (_OnlineTest *OnlineTestFilterer) WatchSubmitAnswer(opts *bind.WatchOpts, sink chan<- *OnlineTestSubmitAnswer, adminAddress []common.Address) (event.Subscription, error) {

	var adminAddressRule []interface{}
	for _, adminAddressItem := range adminAddress {
		adminAddressRule = append(adminAddressRule, adminAddressItem)
	}

	logs, sub, err := _OnlineTest.contract.WatchLogs(opts, "SubmitAnswer", adminAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OnlineTestSubmitAnswer)
				if err := _OnlineTest.contract.UnpackLog(event, "SubmitAnswer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSubmitAnswer is a log parse operation binding the contract event 0x9cf87a9c6288a3c37b26ac5bb24282131251f4e24a22c0d141992a7c7d621a19.
//
// Solidity: event SubmitAnswer(uint256 examCode, address indexed adminAddress, string answer)
func (_OnlineTest *OnlineTestFilterer) ParseSubmitAnswer(log types.Log) (*OnlineTestSubmitAnswer, error) {
	event := new(OnlineTestSubmitAnswer)
	if err := _OnlineTest.contract.UnpackLog(event, "SubmitAnswer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OnlineTestSubmitTestIterator is returned from FilterSubmitTest and is used to iterate over the raw logs and unpacked data for SubmitTest events raised by the OnlineTest contract.
type OnlineTestSubmitTestIterator struct {
	Event *OnlineTestSubmitTest // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OnlineTestSubmitTestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnlineTestSubmitTest)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OnlineTestSubmitTest)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OnlineTestSubmitTestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OnlineTestSubmitTestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OnlineTestSubmitTest represents a SubmitTest event raised by the OnlineTest contract.
type OnlineTestSubmitTest struct {
	ExamCode    *big.Int
	UserAddress common.Address
	Answer      string
	SubmitTime  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSubmitTest is a free log retrieval operation binding the contract event 0x10adfb6a5ccb5230714b769957c062278bad4ef324985b45a0282120fb94f3dd.
//
// Solidity: event SubmitTest(uint256 examCode, address indexed userAddress, string answer, uint256 submitTime)
func (_OnlineTest *OnlineTestFilterer) FilterSubmitTest(opts *bind.FilterOpts, userAddress []common.Address) (*OnlineTestSubmitTestIterator, error) {

	var userAddressRule []interface{}
	for _, userAddressItem := range userAddress {
		userAddressRule = append(userAddressRule, userAddressItem)
	}

	logs, sub, err := _OnlineTest.contract.FilterLogs(opts, "SubmitTest", userAddressRule)
	if err != nil {
		return nil, err
	}
	return &OnlineTestSubmitTestIterator{contract: _OnlineTest.contract, event: "SubmitTest", logs: logs, sub: sub}, nil
}

// WatchSubmitTest is a free log subscription operation binding the contract event 0x10adfb6a5ccb5230714b769957c062278bad4ef324985b45a0282120fb94f3dd.
//
// Solidity: event SubmitTest(uint256 examCode, address indexed userAddress, string answer, uint256 submitTime)
func (_OnlineTest *OnlineTestFilterer) WatchSubmitTest(opts *bind.WatchOpts, sink chan<- *OnlineTestSubmitTest, userAddress []common.Address) (event.Subscription, error) {

	var userAddressRule []interface{}
	for _, userAddressItem := range userAddress {
		userAddressRule = append(userAddressRule, userAddressItem)
	}

	logs, sub, err := _OnlineTest.contract.WatchLogs(opts, "SubmitTest", userAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OnlineTestSubmitTest)
				if err := _OnlineTest.contract.UnpackLog(event, "SubmitTest", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSubmitTest is a log parse operation binding the contract event 0x10adfb6a5ccb5230714b769957c062278bad4ef324985b45a0282120fb94f3dd.
//
// Solidity: event SubmitTest(uint256 examCode, address indexed userAddress, string answer, uint256 submitTime)
func (_OnlineTest *OnlineTestFilterer) ParseSubmitTest(log types.Log) (*OnlineTestSubmitTest, error) {
	event := new(OnlineTestSubmitTest)
	if err := _OnlineTest.contract.UnpackLog(event, "SubmitTest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OnlineTestWithDrawIterator is returned from FilterWithDraw and is used to iterate over the raw logs and unpacked data for WithDraw events raised by the OnlineTest contract.
type OnlineTestWithDrawIterator struct {
	Event *OnlineTestWithDraw // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OnlineTestWithDrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnlineTestWithDraw)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OnlineTestWithDraw)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OnlineTestWithDrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OnlineTestWithDrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OnlineTestWithDraw represents a WithDraw event raised by the OnlineTest contract.
type OnlineTestWithDraw struct {
	ContractOwner common.Address
	ExamFund      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterWithDraw is a free log retrieval operation binding the contract event 0x14b43ca4c63c5423006ad978eab8a14386476b52f1d4728070ea20a36f13e83b.
//
// Solidity: event WithDraw(address indexed contractOwner, uint256 examFund)
func (_OnlineTest *OnlineTestFilterer) FilterWithDraw(opts *bind.FilterOpts, contractOwner []common.Address) (*OnlineTestWithDrawIterator, error) {

	var contractOwnerRule []interface{}
	for _, contractOwnerItem := range contractOwner {
		contractOwnerRule = append(contractOwnerRule, contractOwnerItem)
	}

	logs, sub, err := _OnlineTest.contract.FilterLogs(opts, "WithDraw", contractOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OnlineTestWithDrawIterator{contract: _OnlineTest.contract, event: "WithDraw", logs: logs, sub: sub}, nil
}

// WatchWithDraw is a free log subscription operation binding the contract event 0x14b43ca4c63c5423006ad978eab8a14386476b52f1d4728070ea20a36f13e83b.
//
// Solidity: event WithDraw(address indexed contractOwner, uint256 examFund)
func (_OnlineTest *OnlineTestFilterer) WatchWithDraw(opts *bind.WatchOpts, sink chan<- *OnlineTestWithDraw, contractOwner []common.Address) (event.Subscription, error) {

	var contractOwnerRule []interface{}
	for _, contractOwnerItem := range contractOwner {
		contractOwnerRule = append(contractOwnerRule, contractOwnerItem)
	}

	logs, sub, err := _OnlineTest.contract.WatchLogs(opts, "WithDraw", contractOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OnlineTestWithDraw)
				if err := _OnlineTest.contract.UnpackLog(event, "WithDraw", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithDraw is a log parse operation binding the contract event 0x14b43ca4c63c5423006ad978eab8a14386476b52f1d4728070ea20a36f13e83b.
//
// Solidity: event WithDraw(address indexed contractOwner, uint256 examFund)
func (_OnlineTest *OnlineTestFilterer) ParseWithDraw(log types.Log) (*OnlineTestWithDraw, error) {
	event := new(OnlineTestWithDraw)
	if err := _OnlineTest.contract.UnpackLog(event, "WithDraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
