// Code generated by mockery v1.0.0
package charontest

import charonrpc "github.com/piotrkowalczuk/charon/charonrpc"
import context "context"
import grpc "google.golang.org/grpc"
import mock "github.com/stretchr/testify/mock"

// PermissionManagerClient is an autogenerated mock type for the PermissionManagerClient type
type PermissionManagerClient struct {
	mock.Mock
}

// Get provides a mock function with given fields: ctx, in, opts
func (_m *PermissionManagerClient) Get(ctx context.Context, in *charonrpc.GetPermissionRequest, opts ...grpc.CallOption) (*charonrpc.GetPermissionResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *charonrpc.GetPermissionResponse
	if rf, ok := ret.Get(0).(func(context.Context, *charonrpc.GetPermissionRequest, ...grpc.CallOption) *charonrpc.GetPermissionResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*charonrpc.GetPermissionResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *charonrpc.GetPermissionRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, in, opts
func (_m *PermissionManagerClient) List(ctx context.Context, in *charonrpc.ListPermissionsRequest, opts ...grpc.CallOption) (*charonrpc.ListPermissionsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *charonrpc.ListPermissionsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *charonrpc.ListPermissionsRequest, ...grpc.CallOption) *charonrpc.ListPermissionsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*charonrpc.ListPermissionsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *charonrpc.ListPermissionsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: ctx, in, opts
func (_m *PermissionManagerClient) Register(ctx context.Context, in *charonrpc.RegisterPermissionsRequest, opts ...grpc.CallOption) (*charonrpc.RegisterPermissionsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *charonrpc.RegisterPermissionsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *charonrpc.RegisterPermissionsRequest, ...grpc.CallOption) *charonrpc.RegisterPermissionsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*charonrpc.RegisterPermissionsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *charonrpc.RegisterPermissionsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}