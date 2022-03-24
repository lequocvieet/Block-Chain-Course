//SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

import "@chainlink/contracts/src/v0.8/interfaces/AggregatorV3Interface.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@chainlink/contracts/src/v0.8/interfaces/LinkTokenInterface.sol";
import "@chainlink/contracts/src/v0.8/interfaces/VRFCoordinatorV2Interface.sol";
import "@chainlink/contracts/src/v0.8/VRFConsumerBaseV2.sol";

contract Lottery is VRFConsumerBaseV2, Ownable {
    address payable[] public players;

    address payable public recentWinner;

    uint256 public usdEntryFee;

    AggregatorV3Interface internal ethUsdPriceFeed;

    enum LOTTERY_STATE {
        OPEN,
        CLOSED,
        CACULATING_WINNER
    }

    LOTTERY_STATE public lottery_state;

    VRFCoordinatorV2Interface COORDINATOR;

    LinkTokenInterface LINKTOKEN;

    uint64 subscriptionId;

    address link;

    bytes32 keyHash;

    uint32 callbackGasLimit;

    uint16 requestConfirmations;

    uint32 numWords;

    uint256[] public randomWords;

    uint256 public requestId;

    constructor(
        address _priceFeedAddress,
        address _vrfCoordinator,
        address _link,
        uint64 _subscriptionId,
        bytes32 _keyhash,
        uint32 _callbackGasLimit,
        uint16 _requestConfirmations,
        uint32 _numWords
    ) VRFConsumerBaseV2(_vrfCoordinator) {
        usdEntryFee = 50 * (10**18);
        ethUsdPriceFeed = AggregatorV3Interface(_priceFeedAddress);
        COORDINATOR = VRFCoordinatorV2Interface(_vrfCoordinator);
        LINKTOKEN = LinkTokenInterface(_link);
        subscriptionId = _subscriptionId;
        lottery_state = LOTTERY_STATE.CLOSED;
    }

    function enter() public payable {
        //$50 minimum
        require(lottery_state == LOTTERY_STATE.OPEN);
        require(msg.value >= getEntranceFee(), "Not enough ETH");
        players.push(payable(msg.sender));
    }

    function getEntranceFee() public view returns (uint256) {
        (, int256 price, , , ) = ethUsdPriceFeed.latestRoundData();
        uint256 adjustedPrice = uint256(price) * 10**10; // 18 decimals
        uint256 costToEnter = (usdEntryFee * 10**18) / adjustedPrice;
        return costToEnter;
    }

    function startLottery() public onlyOwner {
        require(
            lottery_state == LOTTERY_STATE.CLOSED,
            "Can't start a new lottery yet!"
        );
        lottery_state = LOTTERY_STATE.OPEN;
    }

    function endLottery() public onlyOwner {
        lottery_state = LOTTERY_STATE.CACULATING_WINNER;
        requestId = COORDINATOR.requestRandomWords( //return request Id
            keyHash,
            subscriptionId, // this function will make a request transaction to chain link oracle
            requestConfirmations, // and later chain linh will make a second transactions to return random number
            callbackGasLimit, // in our fullfill function
            numWords
        );
    }

    function fulfillRandomWords(
        uint256 _requestId,
        uint256[] memory _randomWords // automatic fulfill random number in to randomWords by chainlink oracle
    ) internal override {
        require(
            lottery_state == LOTTERY_STATE.CACULATING_WINNER,
            "you aren't there yet!"
        );

        require(_randomWords[0] > 0, "random not found!");
        uint256 indexOfWinner = _randomWords[0] % (players.length);
        recentWinner = players[indexOfWinner];
        recentWinner.transfer(address(this).balance);
        // this contract consumer transfer all balance to the winner

        //Reset
        players = new address payable[](0);
        lottery_state = LOTTERY_STATE.CLOSED;
        randomWords = _randomWords;
    }
}
