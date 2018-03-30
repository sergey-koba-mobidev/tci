package steps

import (
	"strings"
	"github.com/sergey-koba-mobidev/tci/utils"
	"errors"
	"time"
	"strconv"
)

type UntilStep struct {
	step utils.Step
}

func (c UntilStep) run() ([]byte, error) {
	if c.step.Command == "" {
		return nil, errors.New("`command` cannot be blank for until step")
	}
	if c.step.Contains == "" {
		return nil, errors.New("`contains` cannot be blank for until step")
	}
	if c.step.Retries == 0 {
		c.step.Retries = 3
	}
	if c.step.Delay == 0 {
		c.step.Delay = 100
	}

	var out []byte
	retries := 0

	for !strings.Contains(string(out), c.step.Contains) && retries < c.step.Retries {
		retryOut, err := ExecSystemCommand(c.step.Command, c.step.Shell)
		if err != nil {
			return retryOut, err
		}
		out = retryOut
		if !strings.Contains(string(out), c.step.Contains) {
			utils.LogInfo(" ﹂Waiting %d milliseconds before next retry %d", c.step.Delay, retries)
			time.Sleep(time.Duration(c.step.Delay) * time.Millisecond)
		}
		retries++
	}

	if !strings.Contains(string(out), c.step.Contains) {
		return out, errors.New("max retries count " + strconv.Itoa(retries) + " reached")
	}

	return out, nil
}

func (c UntilStep) name() string {
	return "Until `" + c.step.Command + "` contains " + c.step.Contains
}

func (c UntilStep) setStep(s utils.Step) StepInterface {
	c.step = s
	return c
}
