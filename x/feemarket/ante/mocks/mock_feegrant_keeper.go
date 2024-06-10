// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	types "github.com/cosmos/cosmos-sdk/types"
	mock "github.com/stretchr/testify/mock"
)

// FeeGrantKeeper is an autogenerated mock type for the FeeGrantKeeper type
type FeeGrantKeeper struct {
	mock.Mock
}

// UseGrantedFees provides a mock function with given fields: ctx, granter, grantee, fee, msgs
func (_m *FeeGrantKeeper) UseGrantedFees(ctx types.Context, granter types.AccAddress, grantee types.AccAddress, fee types.Coins, msgs []types.Msg) error {
	ret := _m.Called(ctx, granter, grantee, fee, msgs)

	if len(ret) == 0 {
		panic("no return value specified for UseGrantedFees")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, types.AccAddress, types.AccAddress, types.Coins, []types.Msg) error); ok {
		r0 = rf(ctx, granter, grantee, fee, msgs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewFeeGrantKeeper creates a new instance of FeeGrantKeeper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFeeGrantKeeper(t interface {
	mock.TestingT
	Cleanup(func())
},
) *FeeGrantKeeper {
	mock := &FeeGrantKeeper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
