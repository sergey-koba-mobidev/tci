package commands

import (
	"github.com/sergey-koba-mobidev/tci/utils"
	"os"
)

func Deploy(filename string) error {
	utils.LogInfo("Reading " + filename + " ...")

	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return err;
	}
	conf := &utils.Conf{}
	conf, err = conf.GetConf(filename)
	if err != nil {
		return err
	}

	utils.LogInfo("Found %d step(s).", len(conf.Steps))
	err = RunSteps(conf)
	if err != nil {
		return err
	}

	return nil
}
