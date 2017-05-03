package collector

import (
	"fmt"
	"time"
	"encoding/json"
	"net/http"
	"io/ioutil"

	"github.com/otherpirate/dbaas-metric-collector/settings"
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

type EnvironmentAPI struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Link LinkAPI `json:"_links"`
}

type LinkAPI struct {
	Next string `json:"next"`
	Self string `json:"self"`
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
	api_obj := new(DatabaseListAPI)
    err := json.Unmarshal(body, &api_obj)
    if(err != nil) {
        fmt.Println("Error in parser", err)
    }
    return api_obj
}