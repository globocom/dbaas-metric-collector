package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
	"strings"
	// "reflect"

	"github.com/otherpirate/dbaas-metric-collector/collector"
	"github.com/otherpirate/dbaas-metric-collector/cron"
	"github.com/otherpirate/dbaas-metric-collector/model"
)

func changeString(date string) string{
	dateList := strings.Split(date, "/")
	dateList[0], dateList[2] = dateList[2], dateList[0]
	return strings.Join(dateList, "-")
}

func getDate(req *http.Request) (time.Time, time.Time){
    queryString := req.URL.Query()
	dateFrom := queryString["from"][0]
	dateTo := queryString["to"][0] 

	to := time.Now()
	from := time.Unix(0, 0)

	queryFromTime := "T00:00:00.000Z"
	queryToTime := "T23:59:59.000Z"

	if (len(dateTo) > 0) {
		dateTo = changeString(dateTo) + queryToTime
		to, _ = time.Parse("2006-01-02T15:04:05.000Z", dateTo)
	} 
	if (len(dateFrom) > 0) {
		dateFrom = changeString(dateFrom) + queryFromTime
		from, _ = time.Parse("2006-01-02T15:04:05.000Z", dateFrom)	
	}

	return from, to	
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("page")))
	http.HandleFunc("/loading", loading)
	http.HandleFunc("/healthcheck", healthcheck)
	http.HandleFunc("/engine_count", engine_count)
	http.HandleFunc("/database_count", database_count)
	http.HandleFunc("/team_count", team_count)
	http.HandleFunc("/project_count", project_count)
	http.HandleFunc("/environment_count", environment_count)

	go cron.DailyLoading()

	fmt.Println("listening...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}

func loading(res http.ResponseWriter, req *http.Request) {
	collector.GetDatabases()
	fmt.Fprintln(res, "Loaded!")
}

func healthcheck(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "WORKING")
}

func database_count(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	dateFrom, dateTo := getDate(req)
	
	connection := model.GetConnection()
	defer connection.Session.Close()
	counters := model.DatabaseCounterGet(connection, dateFrom, dateTo)
	content, err := json.Marshal(counters)
	if err != nil {
		panic(err)
	}
   
	res.Write(content)
}


func team_count(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	dateFrom, dateTo := getDate(req)

	connection := model.GetConnection()
	defer connection.Session.Close()

	counters := model.TeamCounterGetLatest(connection, dateFrom, dateTo)
	content, err := json.Marshal(counters)
	if err != nil {
		panic(err)
	}

	res.Write(content)
}

func project_count(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	dateFrom, dateTo := getDate(req)

	connection := model.GetConnection()
	defer connection.Session.Close()

	counters := model.ProjectCounterGetLatest(connection, dateFrom , dateTo)
	content, err := json.Marshal(counters)
	if err != nil {
		panic(err)
	}

	res.Write(content)
}

func environment_count(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	dateFrom, dateTo := getDate(req)

	connection := model.GetConnection()
	defer connection.Session.Close()

	counters := model.EnvironmentCounterGetLatest(connection, dateFrom, dateTo)
	content, err := json.Marshal(counters)
	if err != nil {
		panic(err)
	}

	res.Write(content)
}

func engine_count(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	dateFrom, dateTo := getDate(req)

	connection := model.GetConnection()
	defer connection.Session.Close()

	counters := model.EngineCounterGet(connection, dateFrom, dateTo)
	content, err := json.Marshal(counters)
	if err != nil {
		panic(err)
	}

	res.Write(content)
}
