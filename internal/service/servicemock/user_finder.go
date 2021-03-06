// Code generated by mockery v1.0.0. DO NOT EDIT.

package servicemock

import context "context"
import mock "github.com/stretchr/testify/mock"
import model "github.com/piotrkowalczuk/charon/internal/model"

// UserFinder is an autogenerated mock type for the UserFinder type
type UserFinder struct {
	mock.Mock
}

// FindUser provides a mock function with given fields: _a0
func (_m *UserFinder) FindUser(_a0 context.Context) (*model.UserEntity, error) {
	ret := _m.Called(_a0)

	var r0 *model.UserEntity
	if rf, ok := ret.Get(0).(func(context.Context) *model.UserEntity); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.UserEntity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
