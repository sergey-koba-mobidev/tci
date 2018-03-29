package steps

import (
	"github.com/sergey-koba-mobidev/tci/utils"
)

type StepInterface interface {
	run() ([]byte, error)
	name() string
	setStep(step utils.Step) StepInterface
}