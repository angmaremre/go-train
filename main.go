package main

import (
	"database/sql"
	_interface "flo.com.tr/interfaces"
	"flo.com.tr/types"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {
	var logObject _interface.ElkLogMessage
	logObject.Message = "Will be logged to ELK"
	logObject.Log()

	var logObject2 _interface.DbLogMessage
	logObject2.Message = "Will be logged to DB"
	logObject2.Log()

	var lo _interface.LogMessage
	lo = &logObject2
	lo.Log()

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
