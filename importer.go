package main

import (
	"bufio"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
	"strings"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	inputFile, err := os.Open(os.Args[1])
	checkErr(err)

	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		line := scanner.Text()

		lastIndex := strings.LastIndex(line, "mst")

		if lastIndex > 0 {
			name := line[lastIndex+3 : len(line)]
			name = strings.TrimSpace(name)

			fields := strings.Fields(line[:lastIndex])
			key := fields[len(fields)-1]
			key = strings.TrimSuffix(key, ".")

			fmt.Println(key, name)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(scanner.Err())
	}

	db, err := sql.Open("postgres", "user=chg dbname=budget host=localhost sslmode=disable")
	checkErr(err)

	rows, err := db.Query("SELECT * FROM test")
	checkErr(err)
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		fmt.Println(id, name)
	}

	defer db.Close()

}
