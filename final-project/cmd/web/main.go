package main

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
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
	return conn
}

// connectToDB establishes a connection to the database and returns the connection object.
// It retries the connection if it fails and backs off for 1 second between retries.
func connectToDB() *sql.DB {
	counts := 0

	dsn := os.Getenv("DSN")

	for {
		connection, err := openDB(dsn)
		if err == nil {
			log.Println("connected to database!")
			return connection
		}

		log.Printf("postgres is not yet ready... \n%s", err.Error())

		if counts > 10 {
			return nil
		}

		log.Println("Backing off for 1 second")
		time.Sleep(1 * time.Second)
		counts++
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
