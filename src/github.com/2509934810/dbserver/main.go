package dbserver

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Person struct {
	name string
	age  int8
}

type PersonMan struct {
	personman map[string]Person
}

var person = make([]PersonMan, 100)

func (p *Person) createUser() {
	db, err := sql.Open("mysql", "root:lei123@/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	stmt, err := db.Prepare("insert into lei values(?, ?)")
	if err != nil {
		panic(err.Error())
	}
	stmt.Exec(p.name, p.age)
	defer stmt.Close()
}

func updateUser() {

}

func deleteUser() {

}
