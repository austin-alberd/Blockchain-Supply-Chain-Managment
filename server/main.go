package main

import (
	"database/sql"
	"fmt"
	"main/api"
	"os"
	"runtime"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	//Get the DB connection string
	var connectionString string

	if os.Getenv("DBCONNSTRING") != "" {
		connectionString = os.Getenv("DBCONNSTRING")
		fmt.Println("✅[SUCCESS]: Loaded connection string from environment variables")
	} else {
		err := godotenv.Load(".env")
		if err != nil {
			fmt.Println("💀 [FATAL ERROR]: Could not load connection string from ENV file ", err)
			runtime.Goexit()
		} else {
			connectionString = os.Getenv("DBCONNSTRING")
			fmt.Println("✅[SUCCESS]: Loaded connection string from environment variables")
		}
	}

	// Connect to the database
	db, err := sql.Open("sqlite3", connectionString)

	// Can not connect to the database. UhOh
	if err != nil {
		fmt.Println("💀 [FATAL ERROR]: Could not connect to db", err)
		runtime.Goexit()
	}

	// SQL Statement to create the database needed to store the blockchain data.
	sqlStatement := `
		CREATE TABLE IF NOT EXISTS blockchain (
			hash LONGTEXT NOT NULL PRIMARY KEY,
			prevHash LONGTEXT,
			data LONGTEXT NOT NULL
		)
	`

	// Execute the statement and handle any errors that result from it.
	_, err = db.Exec(sqlStatement)
	if err != nil {
		fmt.Println("💀[FATAL ERROR]: Could not create blockchain table \n", err)
		runtime.Goexit()
	} else {
		fmt.Println("✅[SUCCESS]: Blockchain Table is Ready")
	}

	// SQL Statement to create the inventory table.
	sqlStatement = `
		CREATE TABLE IF NOT EXISTS inventory (
			itemName LONGTEXT NOT NULL PRIMARY KEY,
			stockNumber LONGTEXT,
			itemDescription LONGTEXT,
			itemsQty MEDIUMINT
		)
	`

	// Execute the statement and handle any errors that result from it.
	_, err = db.Exec(sqlStatement)
	if err != nil {
		fmt.Println("💀[FATAL ERROR]: Could not create inventory table \n", err)
		runtime.Goexit()
	} else {
		fmt.Println("✅[SUCCESS]: Inventory Table is Ready")
	}

	// Success Message
	fmt.Println("✅[SUCCESS]: Completed Initial Setup \n")

	// Expose the API Endpoint
	api.ExposeAPI()
}
