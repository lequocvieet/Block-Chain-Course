pragma solidity ^0.6.0; // declare a solidity version

contract SimpleStorage {
    // create a Contract class with some variables
    uint256 favoriteNumber;

    struct People {
        //  create a type People inside contract class
        uint256 favoriteNumber;
        string name;
    }

    People[] public people;

    mapping(string => uint256) public name2FavoriteNumber; // create a method convert from name 2 favorite number

    //People public person=People({favoriteNumber: 2,name:"viet"});//create a person and assign  value

    function store(uint256 _favoriteNumber) public returns (uint256) {
        favoriteNumber = _favoriteNumber;
        return favoriteNumber;
    }

    //view.pure
    function retrieve() public view returns (uint256) {
        return favoriteNumber;
    }

    function addPerson(string memory _name, uint256 _favoriteNumber) public {
        people.push(People(_favoriteNumber, _name));
        name2FavoriteNumber[_name] = _favoriteNumber;
    }
}
