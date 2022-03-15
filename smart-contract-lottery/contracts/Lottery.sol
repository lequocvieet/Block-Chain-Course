//SPDX-License-Identifier: MIT
pragma solidity ^0.6.6;

import "@chainlink/contracts/src/v0.6/interfaces/AggregatorV3Interface.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@chainlink/contracts/src/v0.8/interfaces/LinkTokenInterface.sol";
import "@chainlink/contracts/src/v0.8/interfaces/VRFCoordinatorV2Interface.sol";
import "@chainlink/contracts/src/v0.8/VRFConsumerBaseV2.sol";

contract Lottery is VRFConsumerBaseV2, Ownable {
    address payable[] public players;

    address public recentWinner;

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

    uint256[] public _randomWords;

    uint256 public requestId;

    constructor(
        address _priceFeedAddress,
        address _vrfCoordinator,
        address _link,
        uint64 _subscriptionId,
        bytes32 _keyhash,
        uint32 _callbackGasLimit,
        uint16 _requestConfirmations,
        uint32 numWords,
        uint256[] _randomWords
    ) public VRFConsumerBaseV2(_vrfCoordinator) {
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
        players.push(msg.sender);
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
        requestId = COORDINATOR.requestRandomWords(
            keyHash,
            _subscriptionId,
            requestConfirmations,
            callbackGasLimit,
            numWords
        );
    }

    function fulfillRandomWords(
        uint256 _requestId,
        uint256[] memory randomWords
    ) internal override {
        require(
            lottery_state == LOTTERY_STATE.CACULATING_WINNER,
            "you aren't there yet!"
        );
        _randomWords = randomWords;
        require(_randomWords > 0, "random not found!");
        uint256 indexOfWinner = _randomWords % players.length();
        recentWinner = players[indexOfWinner];
        recentWinner.transfer(address(this).balance);
        // this contract consumer transfer all balance to the winner
    }
}
