// Code generated by mockery v2.22.1. DO NOT EDIT.

package synchronizer

import (
	client "github.com/0xPolygon/cdk-data-availability/client"
	mock "github.com/stretchr/testify/mock"
)

// dataCommitteeClientFactoryMock is an autogenerated mock type for the ClientFactoryInterface type
type dataCommitteeClientFactoryMock struct {
	mock.Mock
}

// New provides a mock function with given fields: url
func (_m *dataCommitteeClientFactoryMock) New(url string) client.ClientInterface {
	ret := _m.Called(url)

	var r0 client.ClientInterface
	if rf, ok := ret.Get(0).(func(string) client.ClientInterface); ok {
		r0 = rf(url)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(client.ClientInterface)
		}
	}

	return r0
}

type mockConstructorTestingTnewDataCommitteeClientFactoryMock interface {
	mock.TestingT
	Cleanup(func())
}

// newDataCommitteeClientFactoryMock creates a new instance of dataCommitteeClientFactoryMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newDataCommitteeClientFactoryMock(t mockConstructorTestingTnewDataCommitteeClientFactoryMock) *dataCommitteeClientFactoryMock {
	mock := &dataCommitteeClientFactoryMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}