package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	//"mysqlSample1/utils"

	_ "github.com/go-sql-driver/mysql"
	"nmuk.com/garicsLP/utils"
)

type Longplay struct {
	ID      int
	Box     sql.NullInt64
	Artist  sql.NullString
	Header  sql.NullString
	Lps     sql.NullInt64
	Years   sql.NullString
	Years2  sql.NullString
	Label   sql.NullString
	Country sql.NullString
	Notes   sql.NullString
}

type PageData struct {
	Title     string
	LongPlays []Longplay
}

func main() {
	var err error
	var lp Longplay

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("\tRecovered panic:", r)
		}
	}()
	db, err := sql.Open("mysql", "nick:ghbdtnnt@tcp(localhost:3306)/vinyl")
	utils.PanicErr(err)

	rows, err := db.Query("SELECT id, box, artist, header, lps, years, years2, label, country, notes FROM longplay order by artist")
	// rows, err := db.Query("SELECT id, box, artist, header FROM longplay order by artist")
	utils.PanicErr(err)
	var longPlays []Longplay
	// headers := []string{"ID", "Box", "Artist", "Header", "LPs", "Years", "Years2", "Label", "Country", "Notes"}

	defer rows.Close()

	// i := 0
	for rows.Next() {
		utils.PanicErr(rows.Scan(&lp.ID, &lp.Box, &lp.Artist, &lp.Header, &lp.Lps, &lp.Years, &lp.Years2, &lp.Label, &lp.Country, &lp.Notes))
		longPlays = append(longPlays, lp)
	}
	fmt.Println("LENGTH:", len(longPlays))
	pageData := PageData{"Коллекция пластинок Графа Хортицы", longPlays}
	fmt.Println(pageData.Title)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("table.html"))
		err := tmpl.Execute(w, pageData)
		utils.PanicErr(err)
	})
	fmt.Println("Check the HTML page on localhoct:8080")
	http.ListenAndServe(":8080", nil)

	defer func() {
		if db != nil {
			db.Close()
		}
		fmt.Println("database closed!")
	}()
}
