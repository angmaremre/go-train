package main

import (
	"database/sql"
	"flo.com.tr/types"
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

	var markalar types.Markalar
	for rows.Next() {
		var marka types.Marka

		rowError := rows.Scan(&marka.Name, &marka.Id, &marka.Is_active)
		if rowError != nil {
			fmt.Println(rowError.Error())
		}

		markalar = append(types.Markalar{}, marka)
	}

	for _, mrk := range markalar {
		fmt.Println(mrk)
	}

	testMarka := types.Marka{Id: 11, Name: "Nike", Is_active: true}
	// receiver function, strcutın bir metotuymuşçasına
	testMarka.UpdateMarkaName("New Nike")
	fmt.Println(testMarka)
}
