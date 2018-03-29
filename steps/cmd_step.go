package steps

import (
	"strings"
	"os/exec"
	"github.com/sergey-koba-mobidev/tci/utils"
	"errors"
)

type CmdStep struct {
	step utils.Step
}

func (c CmdStep) run() ([]byte, error) {
	if c.step.Command == "" {
		return nil, errors.New("`command` cannot be blank for step")
	}

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
