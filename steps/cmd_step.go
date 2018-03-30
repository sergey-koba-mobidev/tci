package steps

import (
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

	return ExecSystemCommand(c.step.Command, c.step.Shell)
}

func (c CmdStep) name() string {
	return c.step.Command
}

func (c CmdStep) setStep(s utils.Step) StepInterface {
	c.step = s
	return c
}
