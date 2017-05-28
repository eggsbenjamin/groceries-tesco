package mocks

import domain "github.com/eggsbenjamin/groceries-tesco/domain"
import mock "github.com/stretchr/testify/mock"

// ProductGetter is an autogenerated mock type for the ProductGetter type
type ProductGetter struct {
	mock.Mock
}

// Get provides a mock function with given fields: barcode
func (_m *ProductGetter) Get(barcode string) ([]*domain.Product, error) {
	ret := _m.Called(barcode)

	var r0 []*domain.Product
	if rf, ok := ret.Get(0).(func(string) []*domain.Product); ok {
		r0 = rf(barcode)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(barcode)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
