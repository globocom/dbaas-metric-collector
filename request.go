package main

import (
	"fmt"
	"time"
	"encoding/json"
	"net/http"
	"io/ioutil"

	"current/settings"
)


type DatabaseAPI struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Environment string `json:"environment"`
	Project string `json:"project"`
	Team string `json:"team"`
	Created_At string `json:"created_at"`
	Quarantine_At string `json:"quarantine_dt"`
}

type LinkAPI struct {
	Next string `json:"next"`
}

type DatabaseListAPI struct {
	Link LinkAPI `json:"_links"`
	Databases []DatabaseAPI `json:"database"`
}

func GetDatabases() {
	url := settings.DBAAS_ENDPOINT + "/api/database/"
    for {
		body, err := GetJson(url)
    	if (err != nil) {
        	panic(err)
    	}

    	database_list := ParseResponse(body)
		for _, database := range database_list.Databases {
			fmt.Println(database);	
		}
    	if database_list.Link.Next == "" {
    		break
    	}
    	url = database_list.Link.Next
    }
}

func GetJson(url string) ([]byte, error) {

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	
	request, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, err
    }

	request.SetBasicAuth(settings.DBAAS_USER, settings.DBAAS_PASSWORD)
	response, err := client.Do(request)
    if err != nil {
        return nil, err
    }

    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        return nil, err
    }
    
    return []byte(body), err
}

func ParseResponse(body []byte) (*DatabaseListAPI) {
    var database_list = new(DatabaseListAPI)
    err := json.Unmarshal(body, &database_list)
    if(err != nil) {
        fmt.Println("Error in parser", err)
    }
    return database_list
}