// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	observer "github.com/jjmengze/mygo/pkg/design-pattern/behavioral/observer"
	mock "github.com/stretchr/testify/mock"
)

// Subjecter is an autogenerated mock type for the Subjecter type
type Subjecter struct {
	mock.Mock
}

// Attach provides a mock function with given fields: o
func (_m *Subjecter) Attach(o observer.Observer) {
	_m.Called(o)
}

// UpdateMsg provides a mock function with given fields: msg
func (_m *Subjecter) UpdateMsg(msg string) {
	_m.Called(msg)
}
