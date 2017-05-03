package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/otherpirate/dbaas-metric-collector/collector"
)

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/loading", loading)
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
	fmt.Fprintln(res, "Loading!")
	collector.GetDatabases()
	fmt.Fprintln(res, "Loaded!")
}
