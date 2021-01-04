package service

import (
	"fmt"
	"time"
	//"github.com/test/layered/infra"
	"github.com/test/layered/repo"

)

type SampleInterface interface {
	Method() time.Time
}

type sampleService struct {
	repoVal repo.SampleInterface
}

func NewSampleService(repoVal repo.SampleInterface) SampleInterface{
	return &sampleService{repoVal: repoVal}
}

func (s *sampleService) Method() time.Time{
	fmt.Println("SampleInterface")
	return s.repoVal.Method()
}
