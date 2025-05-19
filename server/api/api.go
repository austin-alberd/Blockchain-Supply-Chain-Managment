package api

import (
	_ "database/sql"
	"fmt"
	"net/http"
	"os"
	"runtime"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

// A struct that holds all of the data of the action
type DataContainer struct {
	ItemName      string // Name of the item effected
	Action        string // What was done to the item ("add, remove, update, etc.")
	DateTimeStamp string // Time and Date stamp of the action
	Comments      string // Comments about the action
}

// A struct that hold the block
type Block struct {
	Hash     string        // The hash of the block
	PrevHash string        // The hash of the previous block
	Data     DataContainer // Data Stored in the Block
}

// Blockchain Helper Functions

/*
	Name:addBlock

Purpose:Adds a block to the blockchain
Returns: true if block is added | false if there is an error
*/
func addBlock() bool {
	return false
}

/*
	Name:validateChain

Purpose:Validates the blockchain
Returns:true if the chain is valid | false if the chain is invalid
*/
func validateChain() bool {
	return false
}

// Inventory System API Endpoint Functions

/*
	Name:newItem

Purpose:HTTP function to add the new item
Returns:Nothing
*/
func newItem(w http.ResponseWriter, req *http.Request) {
	http.Error(w, "Not Implemented", http.StatusNotImplemented)
}

/*
	Name:getItem

Purpose:HTTP function to get an item from the database
Returns:
*/
func getItem(w http.ResponseWriter, req *http.Request) {
	http.Error(w, "Not Implemented", http.StatusNotImplemented)
}

/*
	Name:removeItem

Purpose:HTTP function to remove an item from the database
Returns:
*/
func removeItem(w http.ResponseWriter, req *http.Request) {
	http.Error(w, "Not Implemented", http.StatusNotImplemented)
}

/*
	Name:updateItem

Purpose:HTTP function to update an item in the inventory
Returns:
*/
func updateItem(w http.ResponseWriter, req *http.Request) {
	http.Error(w, "Not Implemented", http.StatusNotImplemented)
}

/*
	Name:ExposeAPI

Purpose: Exposes the API to the user
Returns:
*/

var APIADDR string

func ExposeAPI() {

	if os.Getenv("APIADDR") != "" {
		APIADDR = os.Getenv("APIADDR")
		fmt.Println("✅[SUCCESS]: Loaded connection string from environment variables")
	} else {
		err := godotenv.Load(".env")
		if err != nil {
			fmt.Println("💀 [FATAL ERROR]: Could not load api address from ENV file ", err)
			runtime.Goexit()
		} else {
			APIADDR = os.Getenv("APIADDR")
			fmt.Println("✅[SUCCESS]: Loaded api address from environment variables")
		}
	}

	http.HandleFunc("/api/v1/createItem", newItem)
	http.HandleFunc("/api/v1/getItem", getItem)

	http.HandleFunc("/api/v1/removeItem", removeItem)
	http.HandleFunc("/api/v1/updateItem", updateItem)

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		http.Error(w, "API ROOT PAGE", http.StatusTeapot)
	})

	err := http.ListenAndServe(APIADDR, nil)
	if err != nil {
		fmt.Println("💀[FATAL ERROR]: Could not start server", err)
	}
}
