package shipyard

type Container struct {
    containerId string
    description string
    host        string
    resourceUri string
}

type ContainerManager struct {
}

func NewContainerManager() (*ContainerManager, error) {
    mgr := &ContainerManager{}
    return mgr, nil
}

func (mgr *ContainerManager) ContainerInfo() (string, error) {
    return "Container Info 12345", nil
}
