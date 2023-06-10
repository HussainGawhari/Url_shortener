package dbhelper

import (
	"database/sql"
	"fmt"
	"math/rand"
	"url_shortener/models"

	_ "github.com/go-sql-driver/mysql"
)

// Check for connection if all the things provided are correct
// It will connect to database
func Connect() *sql.DB {

	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "goapi"
	// Open opens a database specified by its database driver name and a driver-specific
	//  data source name, usually consisting of at least a database name and connection information.
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

// with help of this function we are able to add our link to database
func AddLink(request models.Request) error {

	db := Connect()
	defer db.Close()
	// Query executes a query that returns rows, typically a SELECT
	_, err := db.Query(
		"INSERT INTO short (long_url,short_url,create_time) VALUES (?,?, now())",
		request.LongUrl, request.ShortUrl)

	if err != nil {
		fmt.Println("Err", err.Error())
		return err
	}

	return err

}

// Get the link from database
func GetLink(code string) (models.Request, error) {
	// check connection
	db := Connect()
	// declaring varible type of models Request
	var link models.Request
	// 	closing the database connection after the execution of function complete
	defer db.Close()

	results, err := db.Query("SELECT long_url FROM short where short_url=?", code)
	// db.Exec()
	if err != nil {
		fmt.Println("Err", err.Error())
		return models.Request{}, err
	}
	// Next prepares the next result row for reading with the Scan method. It returns true on success, or false
	if results.Next() {
		// Scan copies the columns in the current row into the values pointed at by destenation
		err = results.Scan(&link.LongUrl)
		if err != nil {
			fmt.Println("Err", err.Error())
			fmt.Println("Err", err.Error())
			return models.Request{}, err
		}
	}
	return link, nil
}

// By every calling to this function it will generat one random string of 7 character and return back
func RandomBase62String() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 7)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
