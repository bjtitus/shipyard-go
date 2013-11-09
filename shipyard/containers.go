package shipyard

type Container struct {
    ID              int `json:"id"`
    ContainerID     string `json:"container_id"`
    Description     string `json:"description"`
    Protected       bool `json:"protected"`
    ResourceURI     string `json:"resource_uri"`
    Meta            ContainerInfoMeta
}

type ContainerInfoMeta struct {
    Args        []string
    Config      ContainerConfig
    Created     string
}
type ContainerConfig struct {
    Cmd         []string
    CpuShares   int
    Env         []string
    Hostname    string
    Image       string
    Memory      int
    MemorySwap  int
    PortSpecs   []string
    Privileged  bool

}
