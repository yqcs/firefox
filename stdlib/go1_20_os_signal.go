// Code generated by 'mozilla extract os/signal'. DO NOT EDIT.

//go:build go1.20
// +build go1.20

package stdlib

import (
	"os/signal"
	"reflect"
)

func init() {
	Symbols["os/signal/signal"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Ignore":        reflect.ValueOf(signal.Ignore),
		"Ignored":       reflect.ValueOf(signal.Ignored),
		"Notify":        reflect.ValueOf(signal.Notify),
		"NotifyContext": reflect.ValueOf(signal.NotifyContext),
		"Reset":         reflect.ValueOf(signal.Reset),
		"Stop":          reflect.ValueOf(signal.Stop),
	}
}
