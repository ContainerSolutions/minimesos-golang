package minimesos

import "fmt"
// import "github.com/fsouza/go-dockerclient"

type MesosCluster struct { }

func (mc MesosCluster) Create() {
  // client, _ := docker.NewClientFromEnv()

  // Add random ID
  // opts := docker.CreateContainerOptions{Name: "minimesos-api-123"}
  //c, err := client.CreateContainer(opts)
  //if err != nil || c == nil {
  //  fmt.Println("Could not create MesosCluster ", err)
  //  return
  //}
  fmt.Println("Created Mesos cluster. API running at port 9999")
}

func (mc MesosCluster) Start() {
  fmt.Println("Started Mesos cluster.")
}

func (mc MesosCluster) Destroy() {
  fmt.Println("Destroyed Mesos cluster.")
}
