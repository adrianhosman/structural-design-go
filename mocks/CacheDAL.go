// Code generated by mockery v2.8.0. DO NOT EDIT.

package mocks

import (
	model "github.com/adrianhosman/structural-design-go/model"
	mock "github.com/stretchr/testify/mock"
)

// CacheDAL is an autogenerated mock type for the CacheDAL type
type CacheDAL struct {
	mock.Mock
}

// GetAllMarvelCharacterIDs provides a mock function with given fields:
func (_m *CacheDAL) GetAllMarvelCharacterIDs() ([]int64, error) {
	ret := _m.Called()

	var r0 []int64
	if rf, ok := ret.Get(0).(func() []int64); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int64)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProgressAllMarvelCharacterIDs provides a mock function with given fields:
func (_m *CacheDAL) GetProgressAllMarvelCharacterIDs() (model.MarvelAllCharactersCacheProgress, error) {
	ret := _m.Called()

	var r0 model.MarvelAllCharactersCacheProgress
	if rf, ok := ret.Get(0).(func() model.MarvelAllCharactersCacheProgress); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(model.MarvelAllCharactersCacheProgress)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveAllMarvelCharacterIDs provides a mock function with given fields: characterIDs
func (_m *CacheDAL) SaveAllMarvelCharacterIDs(characterIDs []int64) {
	_m.Called(characterIDs)
}

// SetProgressAllMarvelCharacterIDs provides a mock function with given fields: progress
func (_m *CacheDAL) SetProgressAllMarvelCharacterIDs(progress model.MarvelAllCharactersCacheProgress) {
	_m.Called(progress)
}
