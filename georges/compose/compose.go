package compose

import (
	log "github.com/Sirupsen/logrus"
	"os"
	"os/exec"
)

func Up(path string) error {
	if err := os.Chdir(path); err != nil {
		log.Errorf("docker-compose chdir : %s", err.Error())
		return err
	}

	cmd := exec.Command("docker-compose", "up", "-d")
	if err := cmd.Start(); err != nil {
		log.Errorf("docker-compose up count not start : %s", err.Error())
		return err
	}
	if err := cmd.Wait(); err != nil {
		log.Errorf("docker-compose up execution error : %s", err.Error())
		return err
	}
	return nil
}
