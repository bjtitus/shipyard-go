package shipyard

type Application struct {
    BackendPort     string `json:"backend_port"`
    Containers      []string `json:"containers"`
    Description     string `json:"description"`
    DomainName      string `json:"domain_name"`
    ID              int `json:"id"`
    Name            string `json:"name"`
    Protocol        string `json:"protocol"`
    ResourceURI     string `json:"resource_uri"`
    UUID            string `json:"uuid"`
}
