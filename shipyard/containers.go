package shipyard

type Container struct {
    Id              int `json:"id"`
    ContainerID     string `json:"container_id"`
    Description     string `json:"description"`
    Protected       bool `json:"protected"`
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
