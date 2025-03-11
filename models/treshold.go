package models

type Treshold struct {
	Base
	Name   string
	Category string
	TresholdFrom float32
	TresholdTo float32
	Color string
}
