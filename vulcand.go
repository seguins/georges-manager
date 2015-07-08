package main

import (
  "fmt"
  "encoding/json"
  "github.com/coreos/go-etcd/etcd"
)

type VulcandManager struct {
  etcdHosts []string
  projectId string
}

func (vulcand VulcandManager) initBackend() {
  client := etcd.NewClient(vulcand.etcdHosts)
  defer client.Close();

  backend := BackendVulcand{"http"}
  jsonBackend, err := json.Marshal(backend)
  if err != nil {
    return
  }

  etcdKey := fmt.Sprintf("/vulcand/backends/%s/backend", vulcand.projectId)

  client.Set(etcdKey, string(jsonBackend), 0)
}

func (vulcand VulcandManager) SetServer(url string, numberServer int) {
  client := etcd.NewClient(vulcand.etcdHosts)
  defer client.Close();

  vulcand.initBackend()

  server := ServerVulcand{url}
  jsonServer, err := json.Marshal(server)
	if err != nil {
		return
	}

  etcdKey := fmt.Sprintf("/vulcand/backends/%s/servers/srv%d", vulcand.projectId, numberServer)

  client.Set(etcdKey, string(jsonServer), 0)
}

func (vulcand VulcandManager) SetFrontend(route string) {
  client := etcd.NewClient(vulcand.etcdHosts)
  defer client.Close();

  frontend := FrontendVulcand{"http", vulcand.projectId, route}

  jsonFrontend, err := json.Marshal(frontend)
  if err != nil {
    return
  }

  etcdKey := fmt.Sprintf("/vulcand/frontends/%s/frontend", vulcand.projectId)

  client.Set(etcdKey, string(jsonFrontend), 0)
}
