package steps

import (
	"strings"
	"os/exec"
	"github.com/sergey-koba-mobidev/tci/utils"
)

type CmdStep struct {
	step utils.Step
}

func (c CmdStep) run() ([]byte, error) {
	parts := strings.Fields(c.step.Command)
	head := parts[0]
	parts = parts[1:len(parts)]

	out, err := exec.Command(head, parts...).Output()
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (c CmdStep) name() string {
	return c.step.Command
}

func (c CmdStep) setStep(s utils.Step) StepInterface {
	c.step = s
	return c
}
