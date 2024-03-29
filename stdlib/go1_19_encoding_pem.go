// Code generated by 'mozilla extract encoding/pem'. DO NOT EDIT.

//go:build go1.19 && !go1.20
// +build go1.19,!go1.20

package stdlib

import (
	"encoding/pem"
	"reflect"
)

func init() {
	Symbols["encoding/pem/pem"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Decode":         reflect.ValueOf(pem.Decode),
		"Encode":         reflect.ValueOf(pem.Encode),
		"EncodeToMemory": reflect.ValueOf(pem.EncodeToMemory),

		// type definitions
		"Block": reflect.ValueOf((*pem.Block)(nil)),
	}
}
