package accesors

import _ "github.com/go-sql-driver/mysql"
import "fmt"
import "database/sql"
import "gintest/models"


func TestCreate (test string) int64 {	
	db, err := sql.Open("mysql", MySqlConnection)
    if err != nil {
        fmt.Println(err)
        return -1
    }
    defer db.Close()

    // Create a new user
	
    stmt, err := db.Prepare("INSERT INTO test (test) VALUES (?)")
    if err != nil {
        fmt.Println(err)
        return -1
    }
    res, err := stmt.Exec(test)
    if err != nil {
        fmt.Println(err)
        return -1
    }

	lid, err := res.LastInsertId()
	if err != nil {
        fmt.Println(err)
        return -1
    }
	return lid
}

func TestConsult (id string) models.Test {
	var test models.Test
	test.ID = -1

	db, err := sql.Open("mysql", MySqlConnection)
    if err != nil {
        fmt.Println(err)
        return test
    }
    defer db.Close()

	row := db.QueryRow("SELECT id, test FROM test WHERE id = ?", id)
    
    err = row.Scan(&test.ID, &test.Test)
    if err != nil {
        fmt.Println(err)
        return test
    }
    fmt.Printf("%+v\n", test)
	return test
}

func TestUpdate (test models.Test) bool {
	db, err := sql.Open("mysql", MySqlConnection)
    if err != nil {
        fmt.Println(err)
        return false
    }
    defer db.Close()

	stmt, err := db.Prepare("UPDATE test SET test = ? WHERE id = ?")
	_, err = stmt.Exec(test.Test, test.ID)
    if err != nil {
        fmt.Println(err)
        return false
    }
	return true
}

func TestDelete (id string) bool {
	db, err := sql.Open("mysql", MySqlConnection)
    if err != nil {
        fmt.Println(err)
        return false
    }
    defer db.Close()

	stmt, err := db.Prepare("DELETE FROM test WHERE id = ?")
    if err != nil {
        fmt.Println(err)
        return false
    }
    _, err = stmt.Exec(id)
    if err != nil {
        fmt.Println(err)
        return false
    }
	return true
}