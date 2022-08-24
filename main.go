package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type employe struct {
	id       int
	name     string
	address  string
	position string
}

func connectDataBase() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/golang_basic_sql")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println("Connection Success")
	return db, nil
}
func queryRow() {
	db, err := connectDataBase()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var employe = employe{}
	var id = 1
	err = db.QueryRow("SELECT id,name,address,position FROM employe WHERE id = ?", id).
		Scan(&employe.id, &employe.name, &employe.address, &employe.position)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(employe.id, employe.name, employe.address, employe.position)
}
func query() {
	db, err := connectDataBase()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT id,name,address,position FROM employe")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var ress []employe

	for rows.Next() {
		var resp = employe{}
		var err = rows.Scan(&resp.id, &resp.name, &resp.address, &resp.position)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		ress = append(ress, resp)

	}
	for _, response := range ress {
		fmt.Println(response.id, response.name, response.address, response.position)
	}
}
func main() {
	//query()
	queryRow()
}
