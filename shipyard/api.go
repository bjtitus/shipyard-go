package shipyard

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Meta struct {
	limit      int
	next       int
	offset     int
	previous   int
	totalCount int `json:"total_count,string"`
}

type UserMeta struct {
	ApiKey   string `json:"api_key"`
	Email    string
	Status   string
	Username string
}

type AppMeta struct {
	Meta
	Objects []Application `json:"objects"`
}

type ContainerMeta struct {
	Meta
	Objects []Container `json:"objects"`
}

type HostMeta struct {
	Meta
	Objects []Host `json:"objects"`
}

type API struct {
	username string
	apiKey   string
	url      string
	version  string
}

func NewAPI(username string, apiKey string, url string, version string) *API {
	api := &API{
		username: username,
		apiKey:   apiKey,
		url:      url,
		version:  version,
	}
	return api
}

func (api *API) makeRequest(path string, method string) ([]byte, int) {
	client := &http.Client{}
	url := fmt.Sprintf("%v/api/v%v%v", api.url, api.version, path)
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Panic(err)
	}
	// set headers
	req.Header.Set("Authorization", fmt.Sprintf("ApiKey %v:%s", api.username, api.apiKey))
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}
	return body, resp.StatusCode
}

func (api *API) GetInfo() string {
	info := "Info"
	return info
}

func (api *API) Login(baseUrl string, username string, password string) (UserMeta, error) {
	loginUrl := fmt.Sprintf("%v/api/login", baseUrl)
	resp, err := http.PostForm(loginUrl, url.Values{"username": {username}, "password": {password}})
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}
	var data UserMeta
	e := json.Unmarshal(body, &data)
	return data, e
}

func (api *API) GetApplications() (AppMeta, error) {
	resp, _ := api.makeRequest("/applications/", "GET")
	var data AppMeta
	err := json.Unmarshal(resp, &data)
	return data, err
}

func (api *API) GetContainers() (ContainerMeta, error) {
	resp, _ := api.makeRequest("/containers/", "GET")
	var data ContainerMeta
	err := json.Unmarshal(resp, &data)
	return data, err
}

func (api *API) GetHosts() (HostMeta, error) {
	resp, _ := api.makeRequest("/hosts/", "GET")
	var data HostMeta
	err := json.Unmarshal(resp, &data)
	return data, err
}
