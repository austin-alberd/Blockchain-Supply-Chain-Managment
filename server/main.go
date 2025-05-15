package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var connectionString = "./prod.db"

type DataContainer struct {
	StockNumber     string // The identifier number of the item
	ItemName        string // The name of the item
	ItemDescription string // The item's description
	ItemsInStock    int    // How many are in stock
}

type Block struct {
	Hash     string        // The hash of the block
	PrevHash string        // The hash of the previous block
	Data     DataContainer // Data Stored in the Block
}

func main() {

	// Set up the database.
	db, err := sql.Open("sqlite3", connectionString)

	if err != nil {
		fmt.Println("💀 [FATAL ERROR]: Could not connect to db", err)
	}

	defer db.Close()
	sqlStatement := `
		CREATE TABLE IF NOT EXISTS blockchain (
			hash LONGTEXT NOT NULL PRIMARY KEY,
			prevHash LONGTEXT,
			data LONGTEXT NOT NULL
		)
	`

	_, err = db.Exec(sqlStatement)
	if err != nil {
		fmt.Println("💀[FATAL ERROR]: Could not create blockchain table \n", err)
	} else {
		fmt.Println("✅[SUCCESS]: Created blockchain table")
	}

	sqlStatement = `
		CREATE TABLE IF NOT EXISTS inventory (
			itemName LONGTEXT NOT NULL PRIMARY KEY,
			stockNumber LONGTEXT,
			itemDescription LONGTEXT,
			itemsQty MEDIUMINT
		)
	`

	_, err = db.Exec(sqlStatement)
	if err != nil {
		fmt.Println("💀[FATAL ERROR]: Could not create inventory table \n", err)
	} else {
		fmt.Println("✅[SUCCESS]: Created inventory table")
	}

	fmt.Println("✅[SUCCESS]: Completed Initial Setup \n")

	//Variable initializations
	var userMenuChoice int // A holder for the user's choice in the menu

	for {
		// Print out the menu
		fmt.Println("1) Expose API")
		fmt.Println("2) Validate Block Chain")
		fmt.Println("3) Manual Insert Block")
		fmt.Println("4) Exit")
		fmt.Println("CHOICE (1-4):")
		fmt.Scan(&userMenuChoice)

		switch userMenuChoice {
		case 1:
			fmt.Println("\nExpose API")
		case 2:
			fmt.Println("\nValidate Block Chain")
		case 3:
			fmt.Println("\nManual Insert Block")
		case 4:
			break
		}

		if userMenuChoice == 4 {
			// Its messy I know. Womp Womp
			break
		}
	}

}
