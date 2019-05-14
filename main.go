package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

type Welcome struct {
	Name string
	Time string
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "TEST %s", r.URL.Path[1:])
}

func main() {

	router := gin.Default()
	router.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello from %v", "Gin")
	})
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "views/index.html", nil)
	})
	cnn, err := sql.Open("mysql", "docker:docker@tcp(db:3306)/test_db")
	if err != nil {
		log.Fatal(err)
	}
	id := 1
	var name string

	if err := cnn.QueryRow("SELECT name FROM test_tb WHERE id = ? LIMIT 1", id).Scan(&name); err != nil {
		log.Fatal(err)
	}

	fmt.Println(id, name)

	// welcome := Welcome{"Anonymous", time.Now().Format(time.Stamp)}

	// templates := template.Must(template.ParseFiles("templates/welcome.html"))

	// http.Handle("/static/", //final url can be anything
	// 	http.StripPrefix("/static/",
	// 		http.FileServer(http.Dir("static")))) //Go looks in the relative "static" directory first using http.FileServer(), then matches it to a
	// //url of our choice as shown in http.Handle("/static/"). This url is what we need when referencing our css files
	// //once the server begins. Our html code would therefore be <link rel="stylesheet"  href="/static/stylesheet/...">
	// //It is important to note the url in http.Handle can be whatever we like, so long as we are consistent.

	// //This method takes in the URL path "/" and a function that takes in a response writer, and a http request.
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	// 	//Takes the name from the URL query e.g ?name=Martin, will set welcome.Name = Martin.
	// 	if name := r.FormValue("name"); name != "" {
	// 		welcome.Name = name
	// 	}
	// 	//If errors show an internal server error message
	// 	//I also pass the welcome struct to the welcome-template.html file.
	// 	if err := templates.ExecuteTemplate(w, "welcome.html", welcome); err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	}
	// })

	// //Start the web server, set the port to listen to 8080. Without a path it assumes localhost
	// //Print any errors from starting the webserver using fmt
	// fmt.Println("Listening")
	// fmt.Println(http.ListenAndServe(":8080", nil))

	router.Run()

}
