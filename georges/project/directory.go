package project

import (
	log "github.com/Sirupsen/logrus"
	"os"
)

func (p *Project) Exists() (bool, error) {
	stat, err := os.Stat(basePath + p.Name)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		log.Warnf("Error during the stat of project directory %s", err.Error())
		return false, err
	}
	return stat.IsDir(), nil
}

func (p *Project) createDirectory() {
	os.Mkdir(basePath+p.Name, 0775)
}
