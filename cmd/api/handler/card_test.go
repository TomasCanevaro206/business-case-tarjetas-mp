package handler

import "github.com/stretchr/testify/mock"

type ServiceMock struct {
	mock.Mock
}

func NewServiceMock() *ServiceMock {
	return &ServiceMock{
		Mock: mock.Mock{},
	}
}
