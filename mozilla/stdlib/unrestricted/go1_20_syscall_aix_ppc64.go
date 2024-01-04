// Code generated by 'mozilla extract syscall'. DO NOT EDIT.

//go:build go1.20
// +build go1.20

package unrestricted

import (
	"reflect"
	"syscall"
)

func init() {
	Symbols["syscall/syscall"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Exec":             reflect.ValueOf(syscall.Exec),
		"Exit":             reflect.ValueOf(syscall.Exit),
		"ForkExec":         reflect.ValueOf(syscall.ForkExec),
		"Kill":             reflect.ValueOf(syscall.Kill),
		"PtraceAttach":     reflect.ValueOf(syscall.PtraceAttach),
		"PtraceCont":       reflect.ValueOf(syscall.PtraceCont),
		"PtraceDetach":     reflect.ValueOf(syscall.PtraceDetach),
		"PtracePeekData":   reflect.ValueOf(syscall.PtracePeekData),
		"PtracePeekText":   reflect.ValueOf(syscall.PtracePeekText),
		"PtracePokeData":   reflect.ValueOf(syscall.PtracePokeData),
		"PtracePokeText":   reflect.ValueOf(syscall.PtracePokeText),
		"PtraceSingleStep": reflect.ValueOf(syscall.PtraceSingleStep),
		"RawSyscall":       reflect.ValueOf(syscall.RawSyscall),
		"RawSyscall6":      reflect.ValueOf(syscall.RawSyscall6),
		"Reboot":           reflect.ValueOf(syscall.Reboot),
		"Shutdown":         reflect.ValueOf(syscall.Shutdown),
		"StartProcess":     reflect.ValueOf(syscall.StartProcess),
		"Syscall":          reflect.ValueOf(syscall.Syscall),
		"Syscall6":         reflect.ValueOf(syscall.Syscall6),
	}
}
