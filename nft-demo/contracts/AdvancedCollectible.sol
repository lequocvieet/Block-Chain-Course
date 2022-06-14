//An NFT contract
// The token URI can be random in 3 different dogs

// SPDX-License-Identifier: MIT
pragma solidity 0.6.6;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@chainlink/contracts/src/v0.6/VRFConsumerBase.sol";

contract AdvancedCollectible is ERC721, VRFConsumerBase {
    uint256 public tokenCounter;
    bytes32 public keyhash;
    uint256 public fee;
    enum Breed {
        PUG,
        SHIBA_INU,
        ST_BERNARD
    } // three type of dogs need for random rare
    mapping(uint256 => Breed) public tokenIdToBreed; //mapping tokenId to Breed
    mapping(bytes32 => address) public requestIdToSender; //mapping requestId to sender to use it inside the vrf function
    event requestedCollectible(bytes32 indexed requestId, address requester); //the indexed key word just make this event easier to searching
    event breedAssigned(uint256 indexed tokenId, Breed breed);

    constructor(
        address _vrfCoodinator,
        address _linkToken,
        bytes32 _keyhash,
        uint256 _fee
    )
        public
        VRFConsumerBase(_vrfCoodinator, _linkToken)
        ERC721("Doggie", "DOG")
    {
        tokenCounter = 0;
        keyhash = _keyhash;
        fee = _fee;
    }

    function createCollectible() public returns (bytes32) {
        bytes32 requestId = requestRandomness(keyhash, fee);
        requestIdToSender[requestId] = msg.sender;
        // give requestID and return it's address

        emit requestedCollectible(requestId, msg.sender);
    }

    function fulfillRandomness(bytes32 requestId, uint256 randomNumber)
        internal
        override
    {
        //only vrf coodinator can call this function
        Breed breed = Breed(randomNumber % 3);
        uint256 newTokenId = tokenCounter;
        tokenIdToBreed[newTokenId] = breed;
        emit breedAssigned(newTokenId, breed);
        address owner = requestIdToSender[requestId];
        _safeMint(owner, newTokenId);

        tokenCounter = tokenCounter + 1; // increase the TokenId for each nft item
    }

    function setTokenURI(uint256 tokenId, string memory _tokenURI) {
        require(
            _isApprovedOrOwner(_msgSender(), tokenId),
            "ERC721:caller is not owner no approved"
        );
        //both _isApprovedOrOwner and _msgSender are ERC721 openzepllin functions
        _setTokenURI(tokenId, _tokenURI);
    }
}
