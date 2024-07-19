package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Country struct {
	ID       int
	Name     string
	Currency string
}

func main() {

	db, err := sql.Open("postgres", "host=localhost  port=5432 user=admin database=students password=1999 sslmode=disable")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	id := 12

	if _, err := db.Exec(`delete from countries where id=$1`, id); err != nil {
		panic(err)
	}

	/*if _, err := db.Exec(`update countries set name=$1,currency=$2 where id=$3`, "malaziya", "malaz", 12); err != nil {
		panic(err)
	}

	fmt.Println("update data")*/
	/*
		countries := []Country{}

		rows, err := db.Query(`select *from countries`)
		if err != nil {
			panic(err)
		}

		for rows.Next() {
			country := Country{}
			if err := rows.Scan(&country.ID, &country.Name, &country.Currency); err != nil {
				panic(err)
			}

			countries = append(countries, country)
		}

		fmt.Println(countries)*/
	/*
		if _, err := db.Exec(`insert into countries(id,name,currency) values ($1,$2,$3)`, c.ID, c.Name, c.Currency); err != nil {
			panic(err)
		}
		fmt.Println(c)
	*/
	/*
	   ****************SELECT  ***************

	   	country := Country{}

	   	row := db.QueryRow(`select id,name ,currency from countries`)
	   	if err = row.Scan(&country.ID, &country.Name, &country.Currency); err != nil {
	   		panic(err)
	   	}

	   	fmt.Println(country)
	*/
}
