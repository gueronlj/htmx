package main

import (
	"fmt"
	"net/http"
	"log"
	"html/template"
	"time"
)

type Fruit struct {
	Name string
	Color string
}

func viewFruits(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	//data is a map with keys of type strings and values of type []Fruit (defined by Fruit struct)
	data := map[string][]Fruit{
		"Stuff" : {
			{Name: "Apple", Color: "Red"},
			{Name: "Banana", Color: "Yellow"},
			{Name: "Orange", Color: "Orange"},
			{Name: "Pear", Color: "Green"},
		},
	}
	tmpl.Execute(w, data)
}

func addFruit(w http.ResponseWriter, r *http.Request) {
	//sleep 1 second to simulate a slow server
	time.Sleep(1 * time.Second)

	name := r.FormValue("name")
	color := r.FormValue("color")
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.ExecuteTemplate(w, "fruit-ele", Fruit{Name: name, Color: color})
}
 
func main() { 
	fmt.Println("Running Go server on port 8080")

	http.HandleFunc("/", viewFruits)
	http.HandleFunc("/add-fruit", addFruit)
	
	// prints error and exits program
	log.Fatal(http.ListenAndServe(":8080", nil))
}