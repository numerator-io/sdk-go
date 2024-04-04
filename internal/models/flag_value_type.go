package models

type Unknown struct {
	Value interface{}
}

type FlagValueType interface {
	string | bool | int64 | float64 | Unknown
}
