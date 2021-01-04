
package handler

import (
	"github.com/test/layered/service"
)


type Handler interface {
	SampleHandler
}

type handler struct {
	SampleHandler
}

func NewHandler(serviceVal service.SampleInterface) Handler{
	return &handler{
		// type Handler interface を記述したら個々の中身は必須になる
		SampleHandler: NewSampleHandler(serviceVal),
	}
}
