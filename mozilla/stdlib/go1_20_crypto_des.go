// Code generated by 'mozilla extract crypto/des'. DO NOT EDIT.

//go:build go1.20
// +build go1.20

package stdlib

import (
	"crypto/des"
	"go/constant"
	"go/token"
	"reflect"
)

func init() {
	Symbols["crypto/des/des"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"BlockSize":          reflect.ValueOf(constant.MakeFromLiteral("8", token.INT, 0)),
		"NewCipher":          reflect.ValueOf(des.NewCipher),
		"NewTripleDESCipher": reflect.ValueOf(des.NewTripleDESCipher),

		// type definitions
		"KeySizeError": reflect.ValueOf((*des.KeySizeError)(nil)),
	}
}
