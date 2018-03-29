package commands

import (
	"github.com/sergey-koba-mobidev/tci/utils"
	"fmt"
	"os/exec"
	"strings"
)

func RunSteps(conf *utils.Conf) error {
	for i := 0; i < len(conf.Steps); i++ {
		err := runStep(conf.Steps[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func runStep(step utils.Step) error {
	utils.LogInfo("Running step %s", step.Command)
	parts := strings.Fields(step.Command)
	head := parts[0]
	parts = parts[1:len(parts)]

	out, err := exec.Command(head,parts...).Output()
	if err != nil {
		return err
	}

	fmt.Printf("%s", out)
	utils.LogSuccess("Finished step %s", step.Command)

	return nil
}