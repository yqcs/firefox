// Code generated by 'mozilla extract runtime/pprof'. DO NOT EDIT.

//go:build go1.19 && !go1.20
// +build go1.19,!go1.20

package stdlib

import (
	"reflect"
	"runtime/pprof"
)

func init() {
	Symbols["runtime/pprof/pprof"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Do":                 reflect.ValueOf(pprof.Do),
		"ForLabels":          reflect.ValueOf(pprof.ForLabels),
		"Label":              reflect.ValueOf(pprof.Label),
		"Labels":             reflect.ValueOf(pprof.Labels),
		"Lookup":             reflect.ValueOf(pprof.Lookup),
		"NewProfile":         reflect.ValueOf(pprof.NewProfile),
		"Profiles":           reflect.ValueOf(pprof.Profiles),
		"SetGoroutineLabels": reflect.ValueOf(pprof.SetGoroutineLabels),
		"StartCPUProfile":    reflect.ValueOf(pprof.StartCPUProfile),
		"StopCPUProfile":     reflect.ValueOf(pprof.StopCPUProfile),
		"WithLabels":         reflect.ValueOf(pprof.WithLabels),
		"WriteHeapProfile":   reflect.ValueOf(pprof.WriteHeapProfile),

		// type definitions
		"LabelSet": reflect.ValueOf((*pprof.LabelSet)(nil)),
		"Profile":  reflect.ValueOf((*pprof.Profile)(nil)),
	}
}
