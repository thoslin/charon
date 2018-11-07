// Code generated by mockery v1.0.0. DO NOT EDIT.

package charontest

import charonrpc "github.com/piotrkowalczuk/charon/charonrpc"
import context "context"
import grpc "google.golang.org/grpc"
import mock "github.com/stretchr/testify/mock"

// RefreshTokenManagerClient is an autogenerated mock type for the RefreshTokenManagerClient type
type RefreshTokenManagerClient struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, in, opts
func (_m *RefreshTokenManagerClient) Create(ctx context.Context, in *charonrpc.CreateRefreshTokenRequest, opts ...grpc.CallOption) (*charonrpc.CreateRefreshTokenResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *charonrpc.CreateRefreshTokenResponse
	if rf, ok := ret.Get(0).(func(context.Context, *charonrpc.CreateRefreshTokenRequest, ...grpc.CallOption) *charonrpc.CreateRefreshTokenResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*charonrpc.CreateRefreshTokenResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *charonrpc.CreateRefreshTokenRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, in, opts
func (_m *RefreshTokenManagerClient) List(ctx context.Context, in *charonrpc.ListRefreshTokensRequest, opts ...grpc.CallOption) (*charonrpc.ListRefreshTokensResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *charonrpc.ListRefreshTokensResponse
	if rf, ok := ret.Get(0).(func(context.Context, *charonrpc.ListRefreshTokensRequest, ...grpc.CallOption) *charonrpc.ListRefreshTokensResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*charonrpc.ListRefreshTokensResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *charonrpc.ListRefreshTokensRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Revoke provides a mock function with given fields: ctx, in, opts
func (_m *RefreshTokenManagerClient) Revoke(ctx context.Context, in *charonrpc.RevokeRefreshTokenRequest, opts ...grpc.CallOption) (*charonrpc.RevokeRefreshTokenResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *charonrpc.RevokeRefreshTokenResponse
	if rf, ok := ret.Get(0).(func(context.Context, *charonrpc.RevokeRefreshTokenRequest, ...grpc.CallOption) *charonrpc.RevokeRefreshTokenResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*charonrpc.RevokeRefreshTokenResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *charonrpc.RevokeRefreshTokenRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
