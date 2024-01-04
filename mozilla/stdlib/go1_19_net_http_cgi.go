// Code generated by 'mozilla extract net/http/cgi'. DO NOT EDIT.

//go:build go1.19 && !go1.20
// +build go1.19,!go1.20

package stdlib

import (
	"net/http/cgi"
	"reflect"
)

func init() {
	Symbols["net/http/cgi/cgi"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Request":        reflect.ValueOf(cgi.Request),
		"RequestFromMap": reflect.ValueOf(cgi.RequestFromMap),
		"Serve":          reflect.ValueOf(cgi.Serve),

		// type definitions
		"Handler": reflect.ValueOf((*cgi.Handler)(nil)),
	}
}
