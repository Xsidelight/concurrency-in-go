package main

import (
	"database/sql"
	"log"
	"os"
	"time"
)

const webPort = "80"

func main() {
	// connect to DB
	db := initDB()
	db.Ping()


	// create sessions

	// create channels

	// create waitgroup

	// set up the application config

	// set up mail

	// listen for web communication

}

func initDB() *sql.DB {
	conn := connectToDB()
	if conn == nil {
		log.Panic("can't connect to database")
	}
}

func connectToDB() *sql.DB {
	counts := 0

	dsn := os.Getenv("DSN")

	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("postgres is not yet ready...")
		} else {
			log.Println("connected to database!")
			return connection
		}

		if counts > 10 {
			return nil
		}

		log.Println("Backing off for 1 second")
		time.Sleep(1 * time.Second)
		continue

	}
}

func openDB(dsn string) (*sql.DB, error) {
	// Open a database connection using the provided DSN (Data Source Name).
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		// If there's an error opening the connection, return nil for the database connection and the error.
		return nil, err
	}

	// Check the connection by pinging the database.
	err = db.Ping()
	if err != nil {
		// If there's an error pinging the database, close the connection and return nil for the database connection and the error.
		db.Close()
		return nil, err
	}

	// If there are no errors, return the database connection and nil for the error.
	return db, nil
}
