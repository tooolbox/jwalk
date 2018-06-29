// Code generated by go run make_value.go. DO NOT EDIT.
package jwalk

import (
	"github.com/mailru/easyjson/jlexer"
)

type Value interface {
	Bytes() []byte
	String() string
	Int8() int8
	Int16() int16
	Int32() int32
	Int64() int64
	Int() int
	Uint() uint
	Uint8() uint8
	Uint16() uint16
	Uint32() uint32
	Uint64() uint64
	Float32() float32
	Float64() float64
	Interface() interface{}
}

type value struct {
	raw []byte
}

func (v value) MarshalJSON() ([]byte, error) {
	return v.raw, nil
}

func (v value) Bytes() []byte {
	l := &jlexer.Lexer{Data: v.raw}
	return l.Bytes()
}

func (v value) String() string {
	l := &jlexer.Lexer{Data: v.raw}
	return l.String()
}

func (v value) Int8() int8 {
	l := &jlexer.Lexer{Data: v.raw}
	return l.Int8()
}

func (v value) Int16() int16 {
	l := &jlexer.Lexer{Data: v.raw}
	return l.Int16()
}

func (v value) Int32() int32 {
	l := &jlexer.Lexer{Data: v.raw}
	return l.Int32()
}

func (v value) Int64() int64 {
	l := &jlexer.Lexer{Data: v.raw}
	return l.Int64()
}

func (v value) Int() int {
	l := &jlexer.Lexer{Data: v.raw}
	return l.Int()
}

func (v value) Uint() uint {
	l := &jlexer.Lexer{Data: v.raw}
	return l.Uint()
}

func (v value) Uint8() uint8 {
	l := &jlexer.Lexer{Data: v.raw}
	return l.Uint8()
}

func (v value) Uint16() uint16 {
	l := &jlexer.Lexer{Data: v.raw}
	return l.Uint16()
}

func (v value) Uint32() uint32 {
	l := &jlexer.Lexer{Data: v.raw}
	return l.Uint32()
}

func (v value) Uint64() uint64 {
	l := &jlexer.Lexer{Data: v.raw}
	return l.Uint64()
}

func (v value) Float32() float32 {
	l := &jlexer.Lexer{Data: v.raw}
	return l.Float32()
}

func (v value) Float64() float64 {
	l := &jlexer.Lexer{Data: v.raw}
	return l.Float64()
}

func (v value) Interface() interface{} {
	l := &jlexer.Lexer{Data: v.raw}
	return l.Interface()
}