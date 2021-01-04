
package handler

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"

	"github.com/test/layered/service"
)

type SampleHandler interface{
	Sample(c echo.Context) error
	// ここの名前がおかしかった
	// intereface と同じSampleHanderになっていた
	// interface 名と実装する関数の名前が同じだとダメ
}

type sampleHandler struct{
	serviceVal service.SampleInterface
}

func NewSampleHandler(serviceVal service.SampleInterface) SampleHandler{
	return &sampleHandler{serviceVal: serviceVal}
}

func (s *sampleHandler) Sample(c echo.Context) error{
	fmt.Println("handler sampler go")
	return c.JSON(http.StatusOK, s.serviceVal.Method())
}


