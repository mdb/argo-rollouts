// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	v1alpha1 "github.com/argoproj/argo-rollouts/pkg/apis/rollouts/v1alpha1"
)

// TrafficRoutingReconciler is an autogenerated mock type for the TrafficRoutingReconciler type
type TrafficRoutingReconciler struct {
	mock.Mock
}

// SetHeaderRouting provides a mock function with given fields: headerRouting
func (_m *TrafficRoutingReconciler) SetHeaderRouting(headerRouting *v1alpha1.SetHeaderRouting) error {
	ret := _m.Called(headerRouting)

	var r0 error
	if rf, ok := ret.Get(0).(func(*v1alpha1.SetHeaderRouting) error); ok {
		r0 = rf(headerRouting)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetWeight provides a mock function with given fields: desiredWeight, additionalDestinations
func (_m *TrafficRoutingReconciler) SetWeight(desiredWeight int32, additionalDestinations ...v1alpha1.WeightDestination) error {
	_va := make([]interface{}, len(additionalDestinations))
	for _i := range additionalDestinations {
		_va[_i] = additionalDestinations[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, desiredWeight)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(int32, ...v1alpha1.WeightDestination) error); ok {
		r0 = rf(desiredWeight, additionalDestinations...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Type provides a mock function with given fields:
func (_m *TrafficRoutingReconciler) Type() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// UpdateHash provides a mock function with given fields: canaryHash, stableHash, additionalDestinations
func (_m *TrafficRoutingReconciler) UpdateHash(canaryHash string, stableHash string, additionalDestinations ...v1alpha1.WeightDestination) error {
	_va := make([]interface{}, len(additionalDestinations))
	for _i := range additionalDestinations {
		_va[_i] = additionalDestinations[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, canaryHash, stableHash)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, ...v1alpha1.WeightDestination) error); ok {
		r0 = rf(canaryHash, stableHash, additionalDestinations...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// VerifyWeight provides a mock function with given fields: desiredWeight, additionalDestinations
func (_m *TrafficRoutingReconciler) VerifyWeight(desiredWeight int32, additionalDestinations ...v1alpha1.WeightDestination) (*bool, error) {
	_va := make([]interface{}, len(additionalDestinations))
	for _i := range additionalDestinations {
		_va[_i] = additionalDestinations[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, desiredWeight)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *bool
	if rf, ok := ret.Get(0).(func(int32, ...v1alpha1.WeightDestination) *bool); ok {
		r0 = rf(desiredWeight, additionalDestinations...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bool)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int32, ...v1alpha1.WeightDestination) error); ok {
		r1 = rf(desiredWeight, additionalDestinations...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
