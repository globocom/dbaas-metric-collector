package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/otherpirate/dbaas-metric-collector/collector"
	"github.com/otherpirate/dbaas-metric-collector/model"
)

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/loading", loading)
	http.HandleFunc("/showing", showing)

	fmt.Println("listening...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "hello, world!")
}

func loading(res http.ResponseWriter, req *http.Request) {
	collector.GetDatabases()
	fmt.Fprintln(res, "Loaded!")
}

func showing(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	connection := model.GetConnection()
	defer connection.Session.Close()

	counters := model.DatabaseCounterLoad(connection)
	content, err := json.Marshal(counters)
	if err != nil {
		panic(err)
	}

	res.Write(content)
}
