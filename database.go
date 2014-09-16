package main

import (
    "database/sql"
    "fmt"
    "time"
    "strings"
    "strconv"
    "io/ioutil"
    _ "github.com/go-sql-driver/mysql"
    "net/http"
    "html/template"
)

type Page struct {
    Title string
    Users []string
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    dat, err := ioutil.ReadFile("../env")
    check(err)
    fmt.Print(string(dat))

    var variables map[string]string
    variables = make(map[string]string)

    var rows_list = strings.Split(string(dat), "\n")
    for index := range rows_list {
        var stuff = strings.Split(string(rows_list[index]), "=")
        if len(stuff) == 2{
            variables[stuff[0]] = stuff[1]
        }
    }
    fmt.Print(variables)

    http.HandleFunc("/", homePage)
    http.HandleFunc("/save", savePage)

    fileServer := http.StripPrefix("/html/", http.FileServer(http.Dir("html")))
    http.Handle("/html/", fileServer)

    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Println(err)
    }



    // TODO: Create http template
    db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", variables["database_username"], variables["database_password"],  variables["database_path"], variables["database_name"]))
    defer db.Close()

    if err != nil {
        fmt.Println(err)
    }
    var database_table = variables["database_name"] + "." + variables["database_table"]
    current_day := time.Now().Format("20060102")
    fmt.Printf("Date: %s\n", current_day)

    db.Query(fmt.Sprintf("INSERT INTO variables %s (date, description, category, cost, paid) VALUES (%s, 'Testing', 'Test Category', '10', '0')", database_table, current_day ))

    rows, err := db.Query(fmt.Sprintf("SELECT date, description, category, cost, paid FROM %s", database_table))
    defer rows.Close()
    if err != nil {
                fmt.Println(err)
    }

    var date, description, category, cost, paid string
    var total_cost, total_paid, float_cost, float_paid float64
    total_cost=0
    for rows.Next() {
        rows.Scan(&date, &description, &category, &cost, &paid)
        fmt.Printf("%s, %s, %s, %s, %s\n", date, description, category, cost, paid)
        float_cost, err = strconv.ParseFloat(cost, 64)
        float_paid, err = strconv.ParseFloat(paid, 64)
        total_cost=total_cost + float_cost
        total_paid=total_paid+float_paid
    }
    fmt.Printf("Cost: $%.2f Paid: $%.2f\n",total_cost, total_paid)
}
