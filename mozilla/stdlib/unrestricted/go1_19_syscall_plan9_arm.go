// Code generated by 'mozilla extract syscall'. DO NOT EDIT.

//go:build go1.19 && !go1.20
// +build go1.19,!go1.20

package unrestricted

import (
	"reflect"
	"syscall"
)

func init() {
	Symbols["syscall/syscall"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Exec":         reflect.ValueOf(syscall.Exec),
		"Exit":         reflect.ValueOf(syscall.Exit),
		"ForkExec":     reflect.ValueOf(syscall.ForkExec),
		"RawSyscall":   reflect.ValueOf(syscall.RawSyscall),
		"RawSyscall6":  reflect.ValueOf(syscall.RawSyscall6),
		"StartProcess": reflect.ValueOf(syscall.StartProcess),
		"Syscall":      reflect.ValueOf(syscall.Syscall),
		"Syscall6":     reflect.ValueOf(syscall.Syscall6),
	}
}
