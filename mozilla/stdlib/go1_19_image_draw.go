// Code generated by 'mozilla extract image/draw'. DO NOT EDIT.

//go:build go1.19 && !go1.20
// +build go1.19,!go1.20

package stdlib

import (
	"image"
	"image/color"
	"image/draw"
	"reflect"
)

func init() {
	Symbols["image/draw/draw"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Draw":           reflect.ValueOf(draw.Draw),
		"DrawMask":       reflect.ValueOf(draw.DrawMask),
		"FloydSteinberg": reflect.ValueOf(&draw.FloydSteinberg).Elem(),
		"Over":           reflect.ValueOf(draw.Over),
		"Src":            reflect.ValueOf(draw.Src),

		// type definitions
		"Drawer":      reflect.ValueOf((*draw.Drawer)(nil)),
		"Image":       reflect.ValueOf((*draw.Image)(nil)),
		"Op":          reflect.ValueOf((*draw.Op)(nil)),
		"Quantizer":   reflect.ValueOf((*draw.Quantizer)(nil)),
		"RGBA64Image": reflect.ValueOf((*draw.RGBA64Image)(nil)),

		// interface wrapper definitions
		"_Drawer":      reflect.ValueOf((*_image_draw_Drawer)(nil)),
		"_Image":       reflect.ValueOf((*_image_draw_Image)(nil)),
		"_Quantizer":   reflect.ValueOf((*_image_draw_Quantizer)(nil)),
		"_RGBA64Image": reflect.ValueOf((*_image_draw_RGBA64Image)(nil)),
	}
}

// _image_draw_Drawer is an interface wrapper for Drawer type
type _image_draw_Drawer struct {
	IValue interface{}
	WDraw  func(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point)
}

func (W _image_draw_Drawer) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	W.WDraw(dst, r, src, sp)
}

// _image_draw_Image is an interface wrapper for Image type
type _image_draw_Image struct {
	IValue      interface{}
	WAt         func(x int, y int) color.Color
	WBounds     func() image.Rectangle
	WColorModel func() color.Model
	WSet        func(x int, y int, c color.Color)
}

func (W _image_draw_Image) At(x int, y int) color.Color {
	return W.WAt(x, y)
}
func (W _image_draw_Image) Bounds() image.Rectangle {
	return W.WBounds()
}
func (W _image_draw_Image) ColorModel() color.Model {
	return W.WColorModel()
}
func (W _image_draw_Image) Set(x int, y int, c color.Color) {
	W.WSet(x, y, c)
}

// _image_draw_Quantizer is an interface wrapper for Quantizer type
type _image_draw_Quantizer struct {
	IValue    interface{}
	WQuantize func(p color.Palette, m image.Image) color.Palette
}

func (W _image_draw_Quantizer) Quantize(p color.Palette, m image.Image) color.Palette {
	return W.WQuantize(p, m)
}

// _image_draw_RGBA64Image is an interface wrapper for RGBA64Image type
type _image_draw_RGBA64Image struct {
	IValue      interface{}
	WAt         func(x int, y int) color.Color
	WBounds     func() image.Rectangle
	WColorModel func() color.Model
	WRGBA64At   func(x int, y int) color.RGBA64
	WSet        func(x int, y int, c color.Color)
	WSetRGBA64  func(x int, y int, c color.RGBA64)
}

func (W _image_draw_RGBA64Image) At(x int, y int) color.Color {
	return W.WAt(x, y)
}
func (W _image_draw_RGBA64Image) Bounds() image.Rectangle {
	return W.WBounds()
}
func (W _image_draw_RGBA64Image) ColorModel() color.Model {
	return W.WColorModel()
}
func (W _image_draw_RGBA64Image) RGBA64At(x int, y int) color.RGBA64 {
	return W.WRGBA64At(x, y)
}
func (W _image_draw_RGBA64Image) Set(x int, y int, c color.Color) {
	W.WSet(x, y, c)
}
func (W _image_draw_RGBA64Image) SetRGBA64(x int, y int, c color.RGBA64) {
	W.WSetRGBA64(x, y, c)
}
