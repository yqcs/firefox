// Code generated by 'mozilla extract crypto/sha256'. DO NOT EDIT.

//go:build go1.20
// +build go1.20

package stdlib

import (
	"crypto/sha256"
	"go/constant"
	"go/token"
	"reflect"
)

func init() {
	Symbols["crypto/sha256/sha256"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"BlockSize": reflect.ValueOf(constant.MakeFromLiteral("64", token.INT, 0)),
		"New":       reflect.ValueOf(sha256.New),
		"New224":    reflect.ValueOf(sha256.New224),
		"Size":      reflect.ValueOf(constant.MakeFromLiteral("32", token.INT, 0)),
		"Size224":   reflect.ValueOf(constant.MakeFromLiteral("28", token.INT, 0)),
		"Sum224":    reflect.ValueOf(sha256.Sum224),
		"Sum256":    reflect.ValueOf(sha256.Sum256),
	}
}