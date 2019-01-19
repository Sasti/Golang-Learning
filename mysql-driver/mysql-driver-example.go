package main

// mysql-driver demonstrates the usage of the sql interface by using the MySQL driver from github.com/go-sql-driver/mysql

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/go-sql-driver/mysql"
)

// main connects to a MySQL database. Addas a table test_table, adds some data to the table and reads that data again.
func main() {
	// mysql.Config is struct that can be used to create a DNS connection string
	dsnConfig := new(mysql.Config)
	dsnConfig.Net = "tcp"
	dsnConfig.Addr = "localhost:13307"
	dsnConfig.DBName = "test_db"
	dsnConfig.User = "root"
	dsnConfig.Passwd = "root"
	// We need to enable the ParseTime option, otherwise the mysql driver won't parse the datetime fields
	// into a time.Time struct. If this option is false the driver needs a []byte as Scan destination for
	// the datetime field.
	dsnConfig.ParseTime = true
	// Enable classic authentication to mysql with a password. MySQL 8.0 uses something different.
	dsnConfig.AllowNativePasswords = true

	// sql.Open is an initialization for the driver and connection interface. It doesn't attempt to connect to the DB
	fmt.Println(dsnConfig.FormatDSN())
	db, err := sql.Open("mysql", dsnConfig.FormatDSN())

	// Check if the driver called through sql.Open had a problem with our DSN
	if err != nil {
		log.Fatalln(err)
	}

	// Now we try to connect to the database. db.Ping will initialize the connection.
	// The driver will initialize a first connection if necessary. The connection would be initialized as
	// well if we would use a Query here insted.
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatalln(pingErr)
	}

	// db.Exec can be used to executed querys without resulting rows. The createResult will contain the
	// affected rows information.
	createResult, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS test_table (
			ID int NOT NULL AUTO_INCREMENT,
			textColumn varchar(255) NOT NULL,
			dateColumn DATETIME,
			intColumn int,
			PRIMARY KEY (ID)
		);`)
	if err != nil {
		log.Fatalln(err)
	}
	r, _ := createResult.RowsAffected()
	fmt.Println("CREATE result: ", r)

	// For the INSERT INTO statement we use the question mark syntax from mysql to inject the values.
	// This syntax is specific to MySQL. For PostgreSQL you would need to use the $1, $2, ... syntax.
	// The MySQL driver will actually send the query and the parameters as seperated parts and the
	// MySQL database will assemble the query. This behaviour can be changed through DSN parameters.
	// The MySQL driver will then inject the paramters into the query byitself.
	// See also: https://github.com/go-sql-driver/mysql#interpolateparams
	insertResult, err := db.Exec("INSERT INTO test_table VALUES (NULL, ?, ?, ?);", "testStringValue", time.Now(), rand.Int31())
	if err != nil {
		log.Fatalln(err)
	}
	r, _ = insertResult.RowsAffected()
	fmt.Println("INSERT INTO result: ", r)

	// Variables to store the intermediate results. Yes you need to actually provide a target where the
	// values you like to read are going to be stored ;)
	var (
		resultID     int
		resultString string
		resultTime   time.Time
		resultInt    int
	)

	// With db.QueryRow we can select a single row. It will just return the first row in the result set and ignore the rest.
	// After the query we "Scan" the results into our buffer variables. Please take note on the pointer syntax!
	oneRowResult := db.QueryRow("SELECT * from test_table ORDER BY dateColumn ASC LIMIT 1")
	scanErr := oneRowResult.Scan(&resultID, &resultString, &resultTime, &resultInt)
	if scanErr != nil {
		log.Fatalln("Error while scanning the single row query: ", scanErr)
	}
	fmt.Println("Result from single row query: ", resultID, resultString, resultTime, resultInt)

	// Add another row to the table. Nothing special to see here. The query result will only be visible in the
	// if/else block.
	if _, err = db.Exec("INSERT INTO test_table VALUES (NULL, ?, ?, ?);", "testString2", time.Now(), rand.Int31()); err != nil {
		log.Fatalln("Error inserting second row: ", err)
	}

	// Execute a select and print all data. If you want to store the values in a map or a slice you need to implement
	// that by yourself.
	if rows, err := db.Query("SELECT * from test_table;"); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("Printing all found rows to terminal")
		for rows.Next() {
			scanErr = rows.Scan(&resultID, &resultString, &resultTime, &resultInt)
			if scanErr != nil {
				log.Fatalln(scanErr)
			}
			fmt.Println(resultID, resultString, resultTime, resultInt)
		}
	}
}
