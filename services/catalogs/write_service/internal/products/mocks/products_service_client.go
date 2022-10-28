// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	products_service "github.com/mehdihadeli/store-golang-microservice-sample/services/catalogs/write_service/internal/products/contracts/proto/service_clients"
)

// ProductsServiceClient is an autogenerated mock type for the ProductsServiceClient type
type ProductsServiceClient struct {
	mock.Mock
}

// CreateProduct provides a mock function with given fields: ctx, in, opts
func (_m *ProductsServiceClient) CreateProduct(ctx context.Context, in *products_service.CreateProductReq, opts ...grpc.CallOption) (*products_service.CreateProductRes, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *products_service.CreateProductRes
	if rf, ok := ret.Get(0).(func(context.Context, *products_service.CreateProductReq, ...grpc.CallOption) *products_service.CreateProductRes); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*products_service.CreateProductRes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *products_service.CreateProductReq, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProductById provides a mock function with given fields: ctx, in, opts
func (_m *ProductsServiceClient) GetProductById(ctx context.Context, in *products_service.GetProductByIdReq, opts ...grpc.CallOption) (*products_service.GetProductByIdRes, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *products_service.GetProductByIdRes
	if rf, ok := ret.Get(0).(func(context.Context, *products_service.GetProductByIdReq, ...grpc.CallOption) *products_service.GetProductByIdRes); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*products_service.GetProductByIdRes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *products_service.GetProductByIdReq, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateProduct provides a mock function with given fields: ctx, in, opts
func (_m *ProductsServiceClient) UpdateProduct(ctx context.Context, in *products_service.UpdateProductReq, opts ...grpc.CallOption) (*products_service.UpdateProductRes, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *products_service.UpdateProductRes
	if rf, ok := ret.Get(0).(func(context.Context, *products_service.UpdateProductReq, ...grpc.CallOption) *products_service.UpdateProductRes); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*products_service.UpdateProductRes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *products_service.UpdateProductReq, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewProductsServiceClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewProductsServiceClient creates a new instance of ProductsServiceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewProductsServiceClient(t mockConstructorTestingTNewProductsServiceClient) *ProductsServiceClient {
	mock := &ProductsServiceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
