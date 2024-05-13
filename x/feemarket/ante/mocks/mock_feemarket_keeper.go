// Code generated by mockery v2.43.0. DO NOT EDIT.

package mocks

import (
	cosmos_sdktypes "github.com/cosmos/cosmos-sdk/types"
	mock "github.com/stretchr/testify/mock"

	types "github.com/skip-mev/feemarket/x/feemarket/types"
)

// FeeMarketKeeper is an autogenerated mock type for the FeeMarketKeeper type
type FeeMarketKeeper struct {
	mock.Mock
}

// GetDenomResolver provides a mock function with given fields:
func (_m *FeeMarketKeeper) GetDenomResolver() types.DenomResolver {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetDenomResolver")
	}

	var r0 types.DenomResolver
	if rf, ok := ret.Get(0).(func() types.DenomResolver); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.DenomResolver)
		}
	}

	return r0
}

// GetMinGasPrice provides a mock function with given fields: ctx
func (_m *FeeMarketKeeper) GetMinGasPrice(ctx cosmos_sdktypes.Context) (cosmos_sdktypes.Coin, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetMinGasPrice")
	}

	var r0 cosmos_sdktypes.Coin
	var r1 error
	if rf, ok := ret.Get(0).(func(cosmos_sdktypes.Context) (cosmos_sdktypes.Coin, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(cosmos_sdktypes.Context) cosmos_sdktypes.Coin); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(cosmos_sdktypes.Coin)
	}

	if rf, ok := ret.Get(1).(func(cosmos_sdktypes.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMinGasPrices provides a mock function with given fields: ctx
func (_m *FeeMarketKeeper) GetMinGasPrices(ctx cosmos_sdktypes.Context) (cosmos_sdktypes.DecCoins, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetMinGasPrices")
	}

	var r0 cosmos_sdktypes.DecCoins
	var r1 error
	if rf, ok := ret.Get(0).(func(cosmos_sdktypes.Context) (cosmos_sdktypes.DecCoins, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(cosmos_sdktypes.Context) cosmos_sdktypes.DecCoins); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(cosmos_sdktypes.DecCoins)
		}
	}

	if rf, ok := ret.Get(1).(func(cosmos_sdktypes.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetParams provides a mock function with given fields: ctx
func (_m *FeeMarketKeeper) GetParams(ctx cosmos_sdktypes.Context) (types.Params, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetParams")
	}

	var r0 types.Params
	var r1 error
	if rf, ok := ret.Get(0).(func(cosmos_sdktypes.Context) (types.Params, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(cosmos_sdktypes.Context) types.Params); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(types.Params)
	}

	if rf, ok := ret.Get(1).(func(cosmos_sdktypes.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetState provides a mock function with given fields: ctx
func (_m *FeeMarketKeeper) GetState(ctx cosmos_sdktypes.Context) (types.State, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetState")
	}

	var r0 types.State
	var r1 error
	if rf, ok := ret.Get(0).(func(cosmos_sdktypes.Context) (types.State, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(cosmos_sdktypes.Context) types.State); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(types.State)
	}

	if rf, ok := ret.Get(1).(func(cosmos_sdktypes.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetParams provides a mock function with given fields: ctx, params
func (_m *FeeMarketKeeper) SetParams(ctx cosmos_sdktypes.Context, params types.Params) error {
	ret := _m.Called(ctx, params)

	if len(ret) == 0 {
		panic("no return value specified for SetParams")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(cosmos_sdktypes.Context, types.Params) error); ok {
		r0 = rf(ctx, params)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetState provides a mock function with given fields: ctx, state
func (_m *FeeMarketKeeper) SetState(ctx cosmos_sdktypes.Context, state types.State) error {
	ret := _m.Called(ctx, state)

	if len(ret) == 0 {
		panic("no return value specified for SetState")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(cosmos_sdktypes.Context, types.State) error); ok {
		r0 = rf(ctx, state)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewFeeMarketKeeper creates a new instance of FeeMarketKeeper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFeeMarketKeeper(t interface {
	mock.TestingT
	Cleanup(func())
}) *FeeMarketKeeper {
	mock := &FeeMarketKeeper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
