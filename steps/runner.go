package steps

import (
	"github.com/sergey-koba-mobidev/tci/utils"
	"fmt"
	"reflect"
	"errors"
	"strings"
)

var stepRegistry = make(map[string]reflect.Type)
var runSteps []utils.Step
var deferSteps []utils.Step

func initRegistry() {
	stepRegistry["cmd"] = reflect.TypeOf(CmdStep{})
	stepRegistry["until"] = reflect.TypeOf(UntilStep{})
	stepRegistry["defer"] = reflect.TypeOf(DeferStep{})
}

func RunSteps(conf *utils.Conf) error {
	var err error = nil
	initRegistry()

	for i := 0; i < len(conf.Steps); i++ {
		if conf.Steps[i].Mode == "defer" {
			deferSteps = append(deferSteps, conf.Steps[i])
		} else {
			runSteps = append(runSteps, conf.Steps[i])
		}
	}

	// Run steps
	for i := 0; i < len(runSteps); i++ {
		err = runStep(runSteps[i])
		if err != nil {
			utils.LogError(" ﹂Step failed with error %s", err.Error())
			break
		}
	}

	// Run defer steps
	for i := 0; i < len(deferSteps); i++ {
		deferErr := runStep(deferSteps[i])
		if deferErr != nil {
			utils.LogError(" ﹂Step failed with error %s", deferErr.Error())
		}
	}

	return err
}

func runStep(step utils.Step) error {
	if step.Mode == "" {
		step.Mode = "cmd"
	}

	if _, ok := stepRegistry[step.Mode]; !ok {
		return errors.New("not supported step mode " + step.Mode)
	}

	s := reflect.New(stepRegistry[step.Mode]).Elem().Interface().(StepInterface)
	s = s.setStep(step)
	utils.LogInfo("\r\nRunning step %s", s.name())

	out, err := s.run()
	if err != nil {
		return err
	}

	printStepOutput(out)
	utils.LogSuccess("\r ﹂Finished step %s", s.name())

	return nil
}

func printStepOutput(out []byte) {
	var lines []string = strings.Split(string(out), "\n")
	for index, line := range lines {
		line = "  " + line
		if index != len(lines)-1 {
			line += "\n"
		}
		fmt.Printf("%s", line)
	}
}
