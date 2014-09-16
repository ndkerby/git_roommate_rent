package main

import (
    "fmt"
    "net/http"
    "html/template"
)

type Page struct {
    Title string
    Users []string
}

func main() {
    http.HandleFunc("/add", homePage)
    http.HandleFunc("/save", savePage)

    fileServer := http.StripPrefix("/html/", http.FileServer(http.Dir("html")))
    http.Handle("/html/", fileServer)

    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Println(err)
    }
}

func homePage(rw http.ResponseWriter, req *http.Request) {
    var users []string

    for i := 0; i < 100; i++ {
        users = append(users, fmt.Sprintf("natalie %d", i))
    }

    p := Page{
        Title: "My Home",
        Users: users,
    }

    tmpl := make(map[string]*template.Template)
    tmpl["index.html"] = template.Must(template.ParseFiles("html/index.html", "html/layout.html"))
    tmpl["index.html"].ExecuteTemplate(rw, "base", p)
    req.ParseForm()
    totalAmount := req.PostFormValue("totalAmount")
    natalie_pay := req.PostFormValue("natalie_pay")
    fmt.Println(totalAmount)
    fmt.Println(natalie_pay)
    http.Redirect(rw, req, "/", http.StatusFound)
}

func savePage(rw http.ResponseWriter, req *http.Request) {
    req.ParseForm()
    firstName := req.PostFormValue("first_name")
    lastName := req.PostFormValue("last_name")
    fmt.Println(firstName)
    fmt.Println(lastName)

    http.Redirect(rw, req, "/", http.StatusFound)
}
