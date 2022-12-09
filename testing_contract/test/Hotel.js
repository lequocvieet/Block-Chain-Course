const { expect } = require("chai");

const { loadFixture } = require("@nomicfoundation/hardhat-network-helpers");
const { ethers } = require("hardhat");
const { int } = require("hardhat/internal/core/params/argumentTypes");

describe.only("Test Hotel contract", function () {


  async function deployRoomNFTFixture() {

    const RoomNFT = await ethers.getContractFactory("RoomNFT");
    const [owner, addr1, addr2] = await ethers.getSigners();
    const roomNFT = await RoomNFT.deploy();
    await roomNFT.deployed();
    console.log("roomNFT address 1", roomNFT.address);
    return roomNFT;
  }

  async function deployHotelFixture() {
    // Contracts are deployed using the first signer/account by default
    const [owner, landLord, user1, user2] = await ethers.getSigners();
    const roomNFT = await loadFixture(deployRoomNFTFixture);
    console.log("roomNFT Address", roomNFT.address)

    const Hotel = await ethers.getContractFactory("Hotel");
    const hotel = await Hotel.deploy(roomNFT.address, landLord.address);
    await hotel.deployed();
    console.log("hotel address", hotel.address)
    return { hotel, owner, landLord, user1 };

  }

  async function listenCreateListRoomEvent() {
    const { hotel, landLord } = await loadFixture(deployHotelFixture);
    await hotel.connect(landLord).createListRoom();
    const filter = hotel.filters.CreateListRoom(
      null,
      landLord.address,
      null,
    )
    const results = await hotel.queryFilter(filter)
    console.log("results", results)
    const listRooms = await Promise.all(
      results.map(async (event) => {
        event = event.args
        let listRoom = {
          listRoomId: event.listRoomId,
          owner: event.owner.toString(),
          timestamp: event.timestamp,
        }
        console.log(listRoom)
        return listRoom

      }),
    )
    return listRooms

  }


  async function listenBookRoomEvent(roomId, price, numberOfDates) {
    const { hotel, user1 } = await loadFixture(deployHotelFixture);
    var totalPrice = ethers.utils.parseEther(price.toString());
    await hotel.connect(user1).bookRoom(roomId, price, numberOfDates, { value: totalPrice });
    const filter = hotel.filters.BookRoom(
      null,
      null,
      null,
      user1.address,   //filter by booker address
      null,
    )
    const results = await hotel.queryFilter(filter)
    const bookedRooms = await Promise.all(
      results.map(async (event) => {
        event = event.args
        let bookedRoom = {
          roomId: event.roomId,
          numberOfDates: event.numberOfdates,
          startTokenId: event.startTokenId,
          booker: event.booker.toString(),
          timestamp: event.timestamp,
        }
        return bookedRoom

      }),
    )
    return bookedRooms

  }

  async function isApprovedForAll(roomNFT, user1, hotel) {

    if (!(await roomNFT.isApprovedForAll(user1, hotel.address))) {
      await (await roomNFT.setApprovalForAll(hotel.address, true)).wait();
    }
  }

  async function listenCancelBookRoomEvent(roomId, price, numberOfDates, tokenIds, cancelPrice) {
    const { hotel, user1 } = await loadFixture(deployHotelFixture);
    var totalPrice = ethers.utils.parseEther(price.toString());
    await hotel.connect(user1).bookRoom(roomId, price, numberOfDates, { value: totalPrice });

    //approve for hotel to burn nft

    const roomNFT = await loadFixture(deployRoomNFTFixture);
    console.log("pass here")
    await roomNFT.setApprovalForAll(hotel.address, true);

    await hotel.connect(user1).cancelBookRoom(tokenIds, cancelPrice);

    const filter = hotel.filters.cancelBookRoom(
      null,
      user1.address,   //filter by cancler address
      null,
    )

    const results = await hotel.queryFilter(filter)
    const canceledRooms = await Promise.all(
      results.map(async (event) => {
        event = event.args
        let canceledRoom = {
          tokenIds: event.tokenIds,
          canceler: event.canceler.toString(),
          timestamp: event.timestamp,
        }
        return canceledRoom

      }),
    )
    return canceledRooms

  }


  async function listenCheckinEvent(roomId, price, numberOfDates, tokenIds) {
    const { hotel, user1 } = await loadFixture(deployHotelFixture);
    var totalPrice = ethers.utils.parseEther(price.toString());
    await hotel.connect(user1).bookRoom(roomId, price, numberOfDates, { value: totalPrice });

    //approve for hotel to tranfer nft

    const roomNFT = await loadFixture(deployRoomNFTFixture);
    console.log("pass here")
    await roomNFT.setApprovalForAll(hotel.address, true);

    await hotel.connect(user1).checkIn(tokenIds);

    const filter = hotel.filters.CheckIn(
      null,
      user1.address,   //filter by owner who call check in 
      null,
    )

    const results = await hotel.queryFilter(filter)
    const checkInEvents = await Promise.all(
      results.map(async (event) => {
        event = event.args
        let checkInEvent = {
          tokenIds: event.tokenIds,
          checker: event.checker.toString(),
          timestamp: event.timestamp,
        }
        return checkInEvent

      }),
    )
    return checkInEvents

  }





  describe("Test case", function () {

    it("Should create new list room and owner is landlord", async function () {

      const { hotel, landLord } = await loadFixture(deployHotelFixture);
      await hotel.connect(landLord).createListRoom();
      const element = await hotel.listRooms(0);
      expect(element.owner).to.equal(landLord.address);
    });


    it("Should delete list room ", async function () {

      const { hotel, landLord } = await loadFixture(deployHotelFixture);
      await hotel.connect(landLord).createListRoom();
      const listRoomId = await hotel.listRoomCount();
      await hotel.connect(landLord).deleteListRoom(listRoomId);
      const element = await hotel.listRooms(0);//after delete 
      expect(element.owner).to.equal("0x0000000000000000000000000000000000000000");
    });

    it("Should book room ", async function () {
      const roomId = 1;
      price = 5; //5 ether
      const numberOfDates = 5;
      const bookedRooms = await listenBookRoomEvent(roomId, price, numberOfDates)
      //console.log("bookedRooms",bookedRooms[0])
      expect(bookedRooms[0].numberOfDates).to.equal(numberOfDates)

    });


    // it("Should cancel book room ", async function () {
    //   const roomId = 1;
    //   const price = 5; //5 ether                       // parameter for book
    //   const numberOfDates = 5;

    //   const tokenIds = [0, 1, 2];//burn 0,1,2      //parameter for cancel
    //   const cancelPrice = 2;// 2 ether


    //   const canceledRooms = await listenCancelBookRoomEvent(roomId, price, numberOfDates, tokenIds, cancelPrice)
    //   var burnedTokenId = [];
    //   for (let index = 0; index < canceledRooms.length; index++) {
    //     const tokenId = canceledRooms[index].tokenIds;
    //     burnedTokenId.push(tokenId);

    //   }
    //   expect(burnedTokenId).to.equal(tokenIds)

    // }).timeout(100000);


    // it("Should check in ", async function () {
    //   const roomId = 1;
    //   const price = 5; //5 ether                       // parameter for book
    //   const numberOfDates = 5;


    //   const tokenIds = [0, 1, 2]// checkin for 3 days in a row

    //   const checkInEvents = await listenCheckinEvent(roomId, price, numberOfDates,tokenIds)
    //   var checkInTokenIds = [];
    //   for (let index = 0; index < checkInEvents.length; index++) {
    //     const tokenId = checkInEvents[index].tokenIds;
    //     checkInTokenIds.push(tokenId);

    //   }
    //   expect(checkInTokenIds).to.equal(tokenIds)

    // });


    it("Land lord can request payment ", async function () {
      const { hotel, landLord,user1 } = await loadFixture(deployHotelFixture);
      const roomId = 1;
      const price = 5; //5 ether                       // parameter for book
      const numberOfDates = 5;
      const totalPrice=ethers.utils.parseEther(price.toString());
      await hotel.connect(user1).bookRoom(roomId, price, numberOfDates, { value: totalPrice });
      
      var balanceBefore=await landLord.getBalance();
      balanceBefore=ethers.utils.formatEther(balanceBefore);
      console.log(balanceBefore);
      await hotel.connect(landLord).requestPayment(price);
      var balanceAfter=await landLord.getBalance();
      balanceAfter=ethers.utils.formatEther(balanceAfter);


      expect(balanceBefore+price).to.equal(balanceAfter);
    });

  });
});