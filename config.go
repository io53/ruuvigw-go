package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type conf struct {
	InfluxHost   string   `yaml:"influxHost"`
	InfluxToken  string   `yaml:"influxToken"`
	InfluxBucket string   `yaml:"influxBucket"`
	InfluxOrg    string   `yaml:"influxOrg"`
	PollInterval int64    `yaml:"pollInterval"`
	GatewayIPs   []string `yaml:"gatewayIPs"`
}

func (c *conf) getConf() *conf {
	log.Println("reading config")
	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	log.Println(c)
	return c
}
