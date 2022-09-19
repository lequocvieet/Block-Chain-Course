// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

contract Album {
    struct AlbumInfo {
        int256 price;
        string id;
        string artist;
        string title;
    }
    AlbumInfo public albumInfo;

    function setInfo(
        int256 _price,
        string memory _id,
        string memory _title,
        string memory _artist
    ) public {
        albumInfo.price = _price;
        albumInfo.id = _id;
        albumInfo.title = _title;
        albumInfo.artist = _artist;
    }

    function getInfo() public view returns (AlbumInfo memory) {
        return albumInfo;
    }
}
