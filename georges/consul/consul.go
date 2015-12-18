package consul

import (
	log "github.com/Sirupsen/logrus"
	consulapi "github.com/hashicorp/consul/api"
)

func getConsul() (*consulapi.Client, error) {
	config := consulapi.DefaultConfig()
	config.Address = "consul:8500"
	consul, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return consul, nil
}

func SetKey(key string, value string) error {
	consul, err := getConsul()
	if err != nil {
		return err
	}
	d := &consulapi.KVPair{Key: key, Value: []byte(value)}
	_, err = consul.KV().Put(d, nil)
	return err
}
