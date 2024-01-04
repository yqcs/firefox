// Code generated by 'mozilla extract hash/maphash'. DO NOT EDIT.

//go:build go1.19 && !go1.20
// +build go1.19,!go1.20

package stdlib

import (
	"hash/maphash"
	"reflect"
)

func init() {
	Symbols["hash/maphash/maphash"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Bytes":    reflect.ValueOf(maphash.Bytes),
		"MakeSeed": reflect.ValueOf(maphash.MakeSeed),
		"String":   reflect.ValueOf(maphash.String),

		// type definitions
		"Hash": reflect.ValueOf((*maphash.Hash)(nil)),
		"Seed": reflect.ValueOf((*maphash.Seed)(nil)),
	}
}
