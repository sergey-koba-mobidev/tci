package steps

import (
	"github.com/sergey-koba-mobidev/tci/utils"
	"fmt"
	"reflect"
	"errors"
)

var stepRegistry = make(map[string]reflect.Type)

func initRegistry() {
	stepRegistry["cmd"] = reflect.TypeOf(CmdStep{})
}

func RunSteps(conf *utils.Conf) error {
	initRegistry()
	for i := 0; i < len(conf.Steps); i++ {
		err := runStep(conf.Steps[i])
		if err != nil {
			return err
		}
	}
	return nil
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
	utils.LogInfo("Running step %s", s.name())

	out, err := s.run()
	if err != nil {
		return err
	}

	fmt.Printf("%s", out)
	utils.LogSuccess("Finished step %s", s.name())

	return nil
}
