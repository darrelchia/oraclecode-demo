package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type citiesResponse struct {
	Cities []string `json:"cities"` // Cities capitalised to export it, otherwise json encoder will ignore it.
}

func eventHandler(res http.ResponseWriter, req *http.Request) {

	mockServer := "https://private-a62371-oraclecodesolution.apiary-mock.com"

	apiary_client := &http.Client{}
	id := strings.TrimPrefix(req.URL.Path, "/events/")

	if id != "" {
		id = fmt.Sprintf("/%s", id)
	}

	endpoint := fmt.Sprintf("%s/events%s", mockServer, id)

	apiary_req, _ := http.NewRequest("GET", endpoint, nil)
	apiary_resp, err := apiary_client.Do(apiary_req)

	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return
	}

	defer apiary_resp.Body.Close()
	apiary_resp_body, _ := ioutil.ReadAll(apiary_resp.Body)

	res.Header().Set("Content-Type", "application/json")
	res.Write([]byte(apiary_resp_body))
}

func cityHandler(res http.ResponseWriter, req *http.Request) {
	cities := citiesResponse{
		Cities: []string{"Los Angeles", "Brasil", "New York", "Chicago", "Boston", "Bogotá", "Bueenos Aires", "Brasil", "Warsaw", "London", "Berlin", "Paris",
			"Hyderabad", "Bengaluru", "Shenzhen"}}
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(res).Encode(cities)
}

func defaultHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/plain; charset=utf-8")
	res.Write([]byte("Welcome to Oracle Code !!!"))
}

func main() {

	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/cities", cityHandler)
	http.HandleFunc("/events/", eventHandler)
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal("Unable to listen on port 5000 : ", err)
	}
}
