package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"text/template"
)

type People []struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	FullName  string `json:"fullName"`
	Title     string `json:"title"`
	Family    string `json:"family"`
	Image     string `json:"image"`
	ImageURL  string `json:"imageUrl"`
}

type SingleCharacter struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	FullName  string `json:"fullName"`
	Title     string `json:"title"`
	Family    string `json:"family"`
	Image     string `json:"image"`
	ImageURL  string `json:"imageUrl"`
}

func main() {
	launchServer()
}

type Continent []struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func Characters(w http.ResponseWriter) {

	url := "https://thronesapi.com/api/v2/Characters"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	response := People{}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	jsonErr := json.Unmarshal(body, &response)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	indexTpl := template.Must(template.ParseFiles("./index.html"))
	indexTpl.Execute(w, response)
}

func Character(w http.ResponseWriter, r *http.Request) {
	id := strings.ReplaceAll(r.URL.Path, "/character/", "")

	url := "https://thronesapi.com/api/v2/Characters/" + id
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	response := SingleCharacter{}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	jsonErr := json.Unmarshal(body, &response)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	indexTpl := template.Must(template.ParseFiles("./characterdetail.html"))
	indexTpl.Execute(w, response)
}

func Cont(w http.ResponseWriter, r *http.Request) {

	url := "https://thronesapi.com/api/v2/Continents"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	response := Continent{}
	jsonErr := json.Unmarshal(body, &response)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	indexTpl := template.Must(template.ParseFiles("./continents.html"))
	indexTpl.Execute(w, response)
}
func launchServer() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		Characters(w)

	})
	http.HandleFunc("/character/", func(w http.ResponseWriter, r *http.Request) {
		Character(w, r)

	})
	http.HandleFunc("/cont", func(w http.ResponseWriter, r *http.Request) {
		Cont(w, r)

	})

	//cssFolder := http.FileServer(http.Dir("css"))
	//imgFolder := http.FileServer(http.Dir("img"))
	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8080", nil)
}
