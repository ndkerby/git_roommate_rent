package main

import (
        "database/sql"
        "fmt"
        "time"
        _ "github.com/go-sql-driver/mysql"
)

func main() {
        db, err := sql.Open("mysql", "<username>:<password>@<path>/<database_name>")
        defer db.Close()

        if err != nil {
                fmt.Println(err)
        }

        current_day := time.Now().Format("20060102")
        fmt.Printf("Date: %s\n", current_day)

        db.Query(fmt.Sprintf("INSERT INTO <database_name>.<table_name> (date, description, category, cost, paid) VALUES (%s, 'Testing', 'Test Category', '10', '0')", current_day ))

        rows, err := db.Query("SELECT date, description, category, cost, paid FROM <database_name>.<table_name>")
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
        fmt.Printf("Cost: %.Write failed: Broken pipecost, total_paid)

}
