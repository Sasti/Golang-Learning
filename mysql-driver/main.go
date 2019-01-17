package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/go-sql-driver/mysql"
)

func main() {
	dsnConfig := mysql.Config{
		Net:                  "tcp",
		Addr:                 "localhost:13307",
		DBName:               "test_db",
		User:                 "root",
		Passwd:               "root",
		ParseTime:            true,
		AllowNativePasswords: true,
	}

	fmt.Println(dsnConfig.FormatDSN())
	db, err := sql.Open("mysql", dsnConfig.FormatDSN())

	if err != nil {
		log.Fatalln(err)
	}

	pingErr := db.Ping()

	if pingErr != nil {
		log.Fatalln(pingErr)
	}

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

	insertResult, err := db.Exec("INSERT INTO test_table VALUES (NULL, ?, ?, ?);", "testStringValue", time.Now(), rand.Int31())

	if err != nil {
		log.Fatalln(err)
	}

	r, _ = insertResult.RowsAffected()
	fmt.Println("INSERT INTO result: ", r)

	oneRowResult := db.QueryRow("SELECT * from test_table LIMIT 1")

	// Variables to store the intermediate results
	var (
		resultID     int
		resultString string
		resultTime   time.Time
		resultInt    int
	)

	scanErr := oneRowResult.Scan(&resultID, &resultString, &resultTime, &resultInt)

	if scanErr != nil {
		log.Fatalln("Error while scanning the single row query: ", scanErr)
	}

	fmt.Println("Result from single row query: ", resultID, resultString, resultTime, resultInt)

	// Add another row to the table
	if _, err = db.Exec("INSERT INTO test_table VALUES (NULL, ?, ?, ?);", "testString2", time.Now(), rand.Int31()); err != nil {
		log.Fatalln("Error inserting second row: ", err)
	}

	// Execte a select and print all data
	if rows, err := db.Query("SELECT * from test_table;"); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("Printing all found rows to terminal now")
		for rows.Next() {
			scanErr = rows.Scan(&resultID, &resultString, &resultTime, &resultInt)
			if scanErr != nil {
				log.Fatalln(scanErr)
			}
			fmt.Println(resultID, resultString, resultTime, resultInt)
		}
	}
}
