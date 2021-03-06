// Code generated by mockery v1.0.0. DO NOT EDIT.

package charondmock

import charond "github.com/piotrkowalczuk/charon/pb/rpc/charond/v1"
import context "context"
import grpc "google.golang.org/grpc"
import mock "github.com/stretchr/testify/mock"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"

// UserManagerClient is an autogenerated mock type for the UserManagerClient type
type UserManagerClient struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, in, opts
func (_m *UserManagerClient) Create(ctx context.Context, in *charond.CreateUserRequest, opts ...grpc.CallOption) (*charond.CreateUserResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *charond.CreateUserResponse
	if rf, ok := ret.Get(0).(func(context.Context, *charond.CreateUserRequest, ...grpc.CallOption) *charond.CreateUserResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*charond.CreateUserResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *charond.CreateUserRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, in, opts
func (_m *UserManagerClient) Delete(ctx context.Context, in *charond.DeleteUserRequest, opts ...grpc.CallOption) (*wrappers.BoolValue, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *wrappers.BoolValue
	if rf, ok := ret.Get(0).(func(context.Context, *charond.DeleteUserRequest, ...grpc.CallOption) *wrappers.BoolValue); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*wrappers.BoolValue)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *charond.DeleteUserRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: ctx, in, opts
func (_m *UserManagerClient) Get(ctx context.Context, in *charond.GetUserRequest, opts ...grpc.CallOption) (*charond.GetUserResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *charond.GetUserResponse
	if rf, ok := ret.Get(0).(func(context.Context, *charond.GetUserRequest, ...grpc.CallOption) *charond.GetUserResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*charond.GetUserResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *charond.GetUserRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, in, opts
func (_m *UserManagerClient) List(ctx context.Context, in *charond.ListUsersRequest, opts ...grpc.CallOption) (*charond.ListUsersResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *charond.ListUsersResponse
	if rf, ok := ret.Get(0).(func(context.Context, *charond.ListUsersRequest, ...grpc.CallOption) *charond.ListUsersResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*charond.ListUsersResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *charond.ListUsersRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListGroups provides a mock function with given fields: ctx, in, opts
func (_m *UserManagerClient) ListGroups(ctx context.Context, in *charond.ListUserGroupsRequest, opts ...grpc.CallOption) (*charond.ListUserGroupsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *charond.ListUserGroupsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *charond.ListUserGroupsRequest, ...grpc.CallOption) *charond.ListUserGroupsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*charond.ListUserGroupsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *charond.ListUserGroupsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListPermissions provides a mock function with given fields: ctx, in, opts
func (_m *UserManagerClient) ListPermissions(ctx context.Context, in *charond.ListUserPermissionsRequest, opts ...grpc.CallOption) (*charond.ListUserPermissionsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *charond.ListUserPermissionsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *charond.ListUserPermissionsRequest, ...grpc.CallOption) *charond.ListUserPermissionsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*charond.ListUserPermissionsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *charond.ListUserPermissionsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Modify provides a mock function with given fields: ctx, in, opts
func (_m *UserManagerClient) Modify(ctx context.Context, in *charond.ModifyUserRequest, opts ...grpc.CallOption) (*charond.ModifyUserResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *charond.ModifyUserResponse
	if rf, ok := ret.Get(0).(func(context.Context, *charond.ModifyUserRequest, ...grpc.CallOption) *charond.ModifyUserResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*charond.ModifyUserResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *charond.ModifyUserRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetGroups provides a mock function with given fields: ctx, in, opts
func (_m *UserManagerClient) SetGroups(ctx context.Context, in *charond.SetUserGroupsRequest, opts ...grpc.CallOption) (*charond.SetUserGroupsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *charond.SetUserGroupsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *charond.SetUserGroupsRequest, ...grpc.CallOption) *charond.SetUserGroupsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*charond.SetUserGroupsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *charond.SetUserGroupsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetPermissions provides a mock function with given fields: ctx, in, opts
func (_m *UserManagerClient) SetPermissions(ctx context.Context, in *charond.SetUserPermissionsRequest, opts ...grpc.CallOption) (*charond.SetUserPermissionsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *charond.SetUserPermissionsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *charond.SetUserPermissionsRequest, ...grpc.CallOption) *charond.SetUserPermissionsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*charond.SetUserPermissionsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *charond.SetUserPermissionsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
