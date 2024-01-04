// Code generated by 'mozilla extract runtime/trace'. DO NOT EDIT.

//go:build go1.20
// +build go1.20

package stdlib

import (
	"reflect"
	"runtime/trace"
)

func init() {
	Symbols["runtime/trace/trace"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"IsEnabled":   reflect.ValueOf(trace.IsEnabled),
		"Log":         reflect.ValueOf(trace.Log),
		"Logf":        reflect.ValueOf(trace.Logf),
		"NewTask":     reflect.ValueOf(trace.NewTask),
		"Start":       reflect.ValueOf(trace.Start),
		"StartRegion": reflect.ValueOf(trace.StartRegion),
		"Stop":        reflect.ValueOf(trace.Stop),
		"WithRegion":  reflect.ValueOf(trace.WithRegion),

		// type definitions
		"Region": reflect.ValueOf((*trace.Region)(nil)),
		"Task":   reflect.ValueOf((*trace.Task)(nil)),
	}
}
