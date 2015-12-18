package docker

import (
	"github.com/samalba/dockerclient"
	"strings"
)

func getDeamon() (*dockerclient.DockerClient, error) {
	return dockerclient.NewDockerClient("unix:///var/run/docker.sock", nil)
}

func GetExposeAddressByContainerName(name string) (string, error) {
	docker, err := getDeamon()
	if err != nil {
		return "", err
	}
	info, err := docker.InspectContainer(name)
	if err != nil {
		return "", err
	}
	adresse := info.NetworkSettings.IPAddress
	for portName := range info.NetworkSettings.Ports {
		adresse += ":" + strings.Split(portName, "/")[0]
	}
	return adresse, nil
}
