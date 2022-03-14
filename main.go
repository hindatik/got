package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type People []struct {
	URL         string   `json:"url"`
	Name        string   `json:"name"`
	Gender      string   `json:"gender"`
	Culture     string   `json:"culture"`
	Born        string   `json:"born"`
	Died        string   `json:"died"`
	Titles      []string `json:"titles"`
	Aliases     []string `json:"aliases"`
	Father      string   `json:"father"`
	Mother      string   `json:"mother"`
	Spouse      string   `json:"spouse"`
	Allegiances []string `json:"allegiances"`
	Books       []string `json:"books"`
	PovBooks    []string `json:"povBooks"`
	TvSeries    []string `json:"tvSeries"`
	PlayedBy    []string `json:"playedBy"`
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		url := "https://anapioficeandfire.com/api/characters/"

		httpClient := http.Client{
			Timeout: time.Second * 8,
		}

		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			log.Fatal(err)
		}
		res, GetErr := httpClient.Do(req)
		if GetErr != nil {
			log.Fatal(GetErr)
		}

		req.Header.Set("User-Agent", "seb go tuto v2")

		if res.Body != nil {
			defer res.Body.Close()
		}

		body, readErr := ioutil.ReadAll(res.Body)
		if readErr != nil {
			log.Fatal(readErr)
		}

		response := People{}
		jsonErr := json.Unmarshal(body, &response)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.Execute(w, response)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		url := "https://anapioficeandfire.com/api/characters/"

		httpClient := http.Client{
			Timeout: time.Second * 8,
		}

		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			log.Fatal(err)
		}
		res, GetErr := httpClient.Do(req)
		if GetErr != nil {
			log.Fatal(GetErr)
		}

		req.Header.Set("User-Agent", "seb go tuto v2")

		if res.Body != nil {
			defer res.Body.Close()
		}

		body, readErr := ioutil.ReadAll(res.Body)
		if readErr != nil {
			log.Fatal(readErr)
		}

		response := People{}
		jsonErr := json.Unmarshal(body, &response)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.Execute(w, response)
	})
	http.ListenAndServe(":8000", nil)

}
