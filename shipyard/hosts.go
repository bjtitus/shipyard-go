package shipyard

type Host struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Hostname    string `json:"hostname"`
	Port        int    `json:"port"`
	ResourceURI string `json:"resource_uri"`
	Enabled     bool   `json:"enabled"`
}
