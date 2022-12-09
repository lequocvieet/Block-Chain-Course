package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"fmt"
	"log"
	"net/rpc"
	//"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	//"github.com/ethereum/go-ethereum/ethclient"
	//"github.com/ethereum/go-ethereum/accounts/keystore"
	/*"github.com/ethereum/go-ethereum/accounts"
    
    "github.com/ethereum/go-ethereum/common"
	"context"*/

)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	deploy()
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)
	router.Run("localhost:8080")
}

const key = `UTC--2022-09-20T08-38-43.848836669Z--2b8a1413bff36ec8b8345db291b014a47fe8ea58`

func deploy(){
	/*infura := "https://rinkeby.infura.io/v3/0c84ae418d3042a490f3f8f46bd823c4"
	cl, err := ethclient.Dial(infura)
	ctx := context.Background()
	addr := common.HexToAddress("0xf930DcD86eC4570ee16d823Ea823d23BB88B800C")
	balance, err := cl.BalanceAt(ctx, addr, nil)
	fmt.Println(balance)
	fmt.Println(err)

    
	// Deploy a new contract
	addr, tx, ctr, err := albumcontract.DeployAlbum(transactOpts, backend)
	// Check if the contract was deployed successfully
	_, err = bind.WaitDeployed(ctx, backend, tx)
	*/
    
    // Open Keystore
	//ks := keystore.NewKeyStore("/home/vieet/.ethereum/keystore", keystore.StandardScryptN, keystore.StandardScryptP)
	// Create an account
	//acc, _ := ks.NewAccount("password")
	// List all accounts
	//accs := ks.Accounts()
	//fmt.Println(accs)

	// Unlock an account
	//ks.Unlock(accs[1], "password")
	// Create a TransactOpts object
	//ksOpts, err := bind.NewKeyStoreTransactor(ks, accs[1])

	// Create an IPC based RPC connection to a remote node and an authorized transactor
	conn, err := rpc.NewIPCClient("/home/go-ethereum/goerli/geth.ipc")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	auth, err := bind.NewTransactor(strings.NewReader(key), "<<strong_password>>")
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}
	// Deploy the contract passing the newly created `auth` and `conn` vars
	address, tx, _, err:= DeployAlbum(auth, conn)
	//tx:= new(big.Int)
	//instance := "Album contract in Go!"
	//err := "Storage contract in Go!", 0, "Go!"
	if err != nil {
		log.Fatalf("Failed to deploy new storage contract: %v", err)
	}
	fmt.Printf("Contract pending deploy: 0x%x\n", address)
	fmt.Printf("Transaction waiting to be mined: 0x%x\n\n", tx.Hash())

	time.Sleep(250 * time.Millisecond) // Allow it to be processed by the local node :P

	// function call on `instance`. Retrieves pending name
	/*name, err := instance.Name(&bind.CallOpts{Pending: true})
	if err != nil {
		log.Fatalf("Failed to retrieve pending name: %v", err)
	}
	fmt.Println("Pending name:", name)*/
    


}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
