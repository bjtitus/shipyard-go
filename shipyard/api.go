package shipyard

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "log"
)

type API struct {
    username    string
    apiKey      string
    url         string
    version     string
    Containers  *ContainerManager
}
func NewAPI(username string, apiKey string, url string, version string) *API {
    containerMgr, _ := NewContainerManager()
    api := &API{
        username: username,
        apiKey: apiKey,
        url: url,
        version: version,
        Containers: containerMgr,
    }
    return api
}

func (api *API) makeRequest(path string, method string) (string, int) {
    client := &http.Client{}
    url := fmt.Sprintf("%v/api/v%v%v", api.url, api.version, path)
    req, err := http.NewRequest(method, url, nil)
    if err != nil { log.Panic(err) }
    // set headers
    req.Header.Set("Authorization", fmt.Sprintf("ApiKey %v:%s", api.username, api.apiKey))
    resp, err := client.Do(req)
    if err != nil { log.Panic(err) }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err!= nil { log.Panic(err) }
    return string(body), resp.StatusCode
}

func (api *API) GetInfo() string {
    //version := fmt.Sprintf("Shipyard API %v (%v)", api.version, api.url)
    cnt, _ := api.makeRequest("/", "GET")
    return cnt
}
