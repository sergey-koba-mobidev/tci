package utils

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type Step struct {
	Mode string `yaml:"mode"`
	Command string `yaml:"command"`
}

type Conf struct {
	Steps []Step `yaml:"steps"`
}

func (c *Conf) GetConf(filename string) (*Conf, error) {

	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}