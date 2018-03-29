package utils

import (
	"github.com/fatih/color"
)

func LogInfo(msg string, a ...interface{}) {
	if len(a) > 0 {
		color.Yellow(msg, a)
	} else {
		color.Yellow(msg)
	}
}

func LogError(msg string, a ...interface{}) {
	if len(a) > 0 {
		color.Red(msg, a)
	} else {
		color.Red(msg)
	}
}

func LogSuccess(msg string, a ...interface{}) {
	if len(a) > 0 {
		color.Green(msg, a)
	} else {
		color.Green(msg)
	}
}