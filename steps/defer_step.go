package steps

import (
	"strings"
	"os/exec"
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

	parts := strings.Fields(c.step.Command)
	head := parts[0]
	parts = parts[1:len(parts)]

	out, err := exec.Command(head, parts...).Output()
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (c DeferStep) name() string {
	return c.step.Command
}

func (c DeferStep) setStep(s utils.Step) StepInterface {
	c.step = s
	return c
}