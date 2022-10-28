// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	metadata "github.com/mehdihadeli/store-golang-microservice-sample/pkg/core/metadata"
	mock "github.com/stretchr/testify/mock"

	types "github.com/mehdihadeli/store-golang-microservice-sample/pkg/messaging/types"
)

// Producer is an autogenerated mock type for the Producer type
type Producer struct {
	mock.Mock
}

// AddMessageProducedHandler provides a mock function with given fields: _a0
func (_m *Producer) AddMessageProducedHandler(_a0 func(types.IMessage)) {
	_m.Called(_a0)
}

// PublishMessage provides a mock function with given fields: ctx, message, meta
func (_m *Producer) PublishMessage(ctx context.Context, message types.IMessage, meta metadata.Metadata) error {
	ret := _m.Called(ctx, message, meta)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, types.IMessage, metadata.Metadata) error); ok {
		r0 = rf(ctx, message, meta)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PublishMessageWithTopicName provides a mock function with given fields: ctx, message, meta, topicOrExchangeName
func (_m *Producer) PublishMessageWithTopicName(ctx context.Context, message types.IMessage, meta metadata.Metadata, topicOrExchangeName string) error {
	ret := _m.Called(ctx, message, meta, topicOrExchangeName)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, types.IMessage, metadata.Metadata, string) error); ok {
		r0 = rf(ctx, message, meta, topicOrExchangeName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewProducer interface {
	mock.TestingT
	Cleanup(func())
}

// NewProducer creates a new instance of Producer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewProducer(t mockConstructorTestingTNewProducer) *Producer {
	mock := &Producer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
