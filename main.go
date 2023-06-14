package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {
	connStr := "postgres://postgres:P@ssw0rd@localhost/pabuccucom?sslmode=disable"
	conn, _ := sql.Open("postgres", connStr)
	result := conn.Ping()

	if result != nil {
		fmt.Println(result.Error())
	} else {
		fmt.Println("Connected!!")
	}

	rows, err := conn.Query("select * from marka")
	if err != nil {
		fmt.Println(err.Error())
	}

	defer rows.Close()
	defer conn.Close()

	var markalar Markalar
	for rows.Next() {
		var marka Marka

		rowError := rows.Scan(&marka.name, &marka.id, &marka.is_active)
		if rowError != nil {
			fmt.Println(rowError.Error())
		}

		markalar = append(markalar, marka)
	}

	for _, mrk := range markalar {
		fmt.Println(mrk)
	}

	testMarka := Marka{id: 11, name: "Nike", is_active: true}
	testMarka.updateMarkaName("New Nike")
	fmt.Println(testMarka)
}

func (currentMarka *Marka) updateMarkaName(newName string) {
	currentMarka.name = newName
}

type Marka struct {
	id        int
	name      string
	is_active bool
}

type Markalar []Marka
