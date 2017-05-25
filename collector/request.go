package collector

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/otherpirate/dbaas-metric-collector/settings"
)

type Database struct {
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	Environment   string `json:"environment"`
	Project       string `json:"project"`
	Team          string `json:"team"`
	Created_At    string `json:"created_at"`
	Quarantine_At string `json:"quarantine_dt"`
}

type LinkAPI struct {
	Next string `json:"next"`
	Self string `json:"self"`
}

type DatabaseListAPI struct {
	Link      LinkAPI    `json:"_links"`
	Databases []Database `json:"database"`
}

type SubModelAPI struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func GetDatabases() {
	url := settings.DBAAS_ENDPOINT + "/api/database/"
	databases := []Database{}
	teams := make(map[string]string)
	projects := make(map[string]string)
	environments := make(map[string]string)

	for {
		body, err := GetJson(url)
		if err != nil {
			panic(err)
		}

		database_page := ParseResponseDatabases(body)
		for _, database := range database_page.Databases {
			_, ok := teams[database.Team]
			if !(ok) {
				teams[database.Team] = SubModelName(database.Team)
			}
			database.Team = teams[database.Team]

			if database.Project == "" {
				projects[database.Project] = "No project"
			} else {
				_, ok = projects[database.Project]
				if !(ok) {
					projects[database.Project] = SubModelName(database.Project)
				}
			}
			database.Project = projects[database.Project]

			_, ok = environments[database.Environment]
			if !(ok) {
				environments[database.Environment] = SubModelName(database.Environment)
			}
			database.Environment = environments[database.Environment]

			databases = append(databases, database)
		}

		if database_page.Link.Next == "" {
			break
		}
		url = database_page.Link.Next
	}

	Extractor(databases)
}

func SubModelName(url string) string {
	body, err := GetJson(url)
	if err != nil {
		panic(err)
	}
	return ParseResponseSubModel(body).Name
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

func ParseResponseDatabases(body []byte) *DatabaseListAPI {
	api_obj := new(DatabaseListAPI)
	err := json.Unmarshal(body, &api_obj)
	if err != nil {
		fmt.Println("Error in parser", err)
	}
	return api_obj
}

func ParseResponseSubModel(body []byte) *SubModelAPI {
	api_obj := new(SubModelAPI)
	err := json.Unmarshal(body, &api_obj)
	if err != nil {
		fmt.Println("Error in parser", err)
	}
	return api_obj
}
