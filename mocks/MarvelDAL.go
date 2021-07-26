// Code generated by mockery v2.8.0. DO NOT EDIT.

package mocks

import (
	model "github.com/fgunawan1995/xendit/model"
	mock "github.com/stretchr/testify/mock"
)

// MarvelDAL is an autogenerated mock type for the MarvelDAL type
type MarvelDAL struct {
	mock.Mock
}

// GetCharacterByID provides a mock function with given fields: id
func (_m *MarvelDAL) GetCharacterByID(id int64) (model.MarvelGetCharactersResponseResult, error) {
	ret := _m.Called(id)

	var r0 model.MarvelGetCharactersResponseResult
	if rf, ok := ret.Get(0).(func(int64) model.MarvelGetCharactersResponseResult); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(model.MarvelGetCharactersResponseResult)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCharacters provides a mock function with given fields: param
func (_m *MarvelDAL) GetCharacters(param model.MarvelGetCharacterRequest) (model.MarvelGetCharactersResponseData, error) {
	ret := _m.Called(param)

	var r0 model.MarvelGetCharactersResponseData
	if rf, ok := ret.Get(0).(func(model.MarvelGetCharacterRequest) model.MarvelGetCharactersResponseData); ok {
		r0 = rf(param)
	} else {
		r0 = ret.Get(0).(model.MarvelGetCharactersResponseData)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.MarvelGetCharacterRequest) error); ok {
		r1 = rf(param)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
