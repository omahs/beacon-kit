// Code generated by mockery v2.42.2. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// IndexDB is an autogenerated mock type for the IndexDB type
type IndexDB struct {
	mock.Mock
}

type IndexDB_Expecter struct {
	mock *mock.Mock
}

func (_m *IndexDB) EXPECT() *IndexDB_Expecter {
	return &IndexDB_Expecter{mock: &_m.Mock}
}

// Delete provides a mock function with given fields: index, key
func (_m *IndexDB) Delete(index uint64, key []byte) error {
	ret := _m.Called(index, key)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64, []byte) error); ok {
		r0 = rf(index, key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IndexDB_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type IndexDB_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - index uint64
//   - key []byte
func (_e *IndexDB_Expecter) Delete(index interface{}, key interface{}) *IndexDB_Delete_Call {
	return &IndexDB_Delete_Call{Call: _e.mock.On("Delete", index, key)}
}

func (_c *IndexDB_Delete_Call) Run(run func(index uint64, key []byte)) *IndexDB_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint64), args[1].([]byte))
	})
	return _c
}

func (_c *IndexDB_Delete_Call) Return(_a0 error) *IndexDB_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IndexDB_Delete_Call) RunAndReturn(run func(uint64, []byte) error) *IndexDB_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteRange provides a mock function with given fields: start, end
func (_m *IndexDB) DeleteRange(start uint64, end uint64) error {
	ret := _m.Called(start, end)

	if len(ret) == 0 {
		panic("no return value specified for DeleteRange")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64, uint64) error); ok {
		r0 = rf(start, end)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IndexDB_DeleteRange_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteRange'
type IndexDB_DeleteRange_Call struct {
	*mock.Call
}

// DeleteRange is a helper method to define mock.On call
//   - start uint64
//   - end uint64
func (_e *IndexDB_Expecter) DeleteRange(start interface{}, end interface{}) *IndexDB_DeleteRange_Call {
	return &IndexDB_DeleteRange_Call{Call: _e.mock.On("DeleteRange", start, end)}
}

func (_c *IndexDB_DeleteRange_Call) Run(run func(start uint64, end uint64)) *IndexDB_DeleteRange_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint64), args[1].(uint64))
	})
	return _c
}

func (_c *IndexDB_DeleteRange_Call) Return(_a0 error) *IndexDB_DeleteRange_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IndexDB_DeleteRange_Call) RunAndReturn(run func(uint64, uint64) error) *IndexDB_DeleteRange_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: index, key
func (_m *IndexDB) Get(index uint64, key []byte) ([]byte, error) {
	ret := _m.Called(index, key)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(uint64, []byte) ([]byte, error)); ok {
		return rf(index, key)
	}
	if rf, ok := ret.Get(0).(func(uint64, []byte) []byte); ok {
		r0 = rf(index, key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(uint64, []byte) error); ok {
		r1 = rf(index, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IndexDB_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type IndexDB_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - index uint64
//   - key []byte
func (_e *IndexDB_Expecter) Get(index interface{}, key interface{}) *IndexDB_Get_Call {
	return &IndexDB_Get_Call{Call: _e.mock.On("Get", index, key)}
}

func (_c *IndexDB_Get_Call) Run(run func(index uint64, key []byte)) *IndexDB_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint64), args[1].([]byte))
	})
	return _c
}

func (_c *IndexDB_Get_Call) Return(_a0 []byte, _a1 error) *IndexDB_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *IndexDB_Get_Call) RunAndReturn(run func(uint64, []byte) ([]byte, error)) *IndexDB_Get_Call {
	_c.Call.Return(run)
	return _c
}

// Has provides a mock function with given fields: index, key
func (_m *IndexDB) Has(index uint64, key []byte) (bool, error) {
	ret := _m.Called(index, key)

	if len(ret) == 0 {
		panic("no return value specified for Has")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(uint64, []byte) (bool, error)); ok {
		return rf(index, key)
	}
	if rf, ok := ret.Get(0).(func(uint64, []byte) bool); ok {
		r0 = rf(index, key)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(uint64, []byte) error); ok {
		r1 = rf(index, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IndexDB_Has_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Has'
type IndexDB_Has_Call struct {
	*mock.Call
}

// Has is a helper method to define mock.On call
//   - index uint64
//   - key []byte
func (_e *IndexDB_Expecter) Has(index interface{}, key interface{}) *IndexDB_Has_Call {
	return &IndexDB_Has_Call{Call: _e.mock.On("Has", index, key)}
}

func (_c *IndexDB_Has_Call) Run(run func(index uint64, key []byte)) *IndexDB_Has_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint64), args[1].([]byte))
	})
	return _c
}

func (_c *IndexDB_Has_Call) Return(_a0 bool, _a1 error) *IndexDB_Has_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *IndexDB_Has_Call) RunAndReturn(run func(uint64, []byte) (bool, error)) *IndexDB_Has_Call {
	_c.Call.Return(run)
	return _c
}

// Set provides a mock function with given fields: index, key, value
func (_m *IndexDB) Set(index uint64, key []byte, value []byte) error {
	ret := _m.Called(index, key, value)

	if len(ret) == 0 {
		panic("no return value specified for Set")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64, []byte, []byte) error); ok {
		r0 = rf(index, key, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IndexDB_Set_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Set'
type IndexDB_Set_Call struct {
	*mock.Call
}

// Set is a helper method to define mock.On call
//   - index uint64
//   - key []byte
//   - value []byte
func (_e *IndexDB_Expecter) Set(index interface{}, key interface{}, value interface{}) *IndexDB_Set_Call {
	return &IndexDB_Set_Call{Call: _e.mock.On("Set", index, key, value)}
}

func (_c *IndexDB_Set_Call) Run(run func(index uint64, key []byte, value []byte)) *IndexDB_Set_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint64), args[1].([]byte), args[2].([]byte))
	})
	return _c
}

func (_c *IndexDB_Set_Call) Return(_a0 error) *IndexDB_Set_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IndexDB_Set_Call) RunAndReturn(run func(uint64, []byte, []byte) error) *IndexDB_Set_Call {
	_c.Call.Return(run)
	return _c
}

// NewIndexDB creates a new instance of IndexDB. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIndexDB(t interface {
	mock.TestingT
	Cleanup(func())
}) *IndexDB {
	mock := &IndexDB{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
