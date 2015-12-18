package project

import (
	"errors"
	"github.com/docker/docker/pkg/archive"
	"github.com/seguins/georges-manager/georges/compose"
	"github.com/seguins/georges-manager/georges/consul"
	"github.com/seguins/georges-manager/georges/docker"
	"io"
)

type Project struct {
	Name string
}

const basePath = "/tmp/docker/"

func CreateNewProject(name string, reader io.Reader) (*Project, error) {
	project := &Project{
		Name: name,
	}
	exists, err := project.Exists()
	if err != nil {
		return nil, err
	} else if exists == true {
		return nil, errors.New("Project already exists")
	}

	decompressedArchive, err := archive.DecompressStream(reader)
	if err != nil {
		return nil, err
	}
	defer decompressedArchive.Close()

	project.createDirectory()

	archive.Unpack(decompressedArchive, basePath+name, &archive.TarOptions{})

	return project, nil
}

func (p *Project) Start() error {
	if err := compose.Up(basePath + p.Name); err != nil {
		return err
	}
	address, err := docker.GetExposeAddressByContainerName(p.Name + "_web_1")
	if err != nil {
		return err
	}
	if err := consul.SetKey("sites/"+p.Name, address); err != nil {
		return err
	}
	return nil
}
