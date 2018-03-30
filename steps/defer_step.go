package steps

import (
	"github.com/sergey-koba-mobidev/tci/utils"
	"errors"
)

type DeferStep struct {
	step utils.Step
}

func (c DeferStep) run() ([]byte, error) {
	if c.step.Command == "" {
		return nil, errors.New("`command` cannot be blank for defer step")
	}

	return ExecSystemCommand(c.step.Command, c.step.Shell)
}

func (c DeferStep) name() string {
	return c.step.Command
}

func (c DeferStep) setStep(s utils.Step) StepInterface {
	c.step = s
	return c
}