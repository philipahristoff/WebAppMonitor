package main

import 
(
    "fmt"
    "database/sql"
    "github.com/go-sql-driver/mysql"
)

func main() 
{
    
    // Open up a database connection.
    // I've set up a database on my local machine.
    // The database is called test.
    db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3000)/test")
    
    // In case there is an error opening the connection.
	if err != nil 
	{
        panic(err.Error())
    }
    
    // Defer the close till after the main function has finished
    // executing. 
    defer db.Close()
    
}