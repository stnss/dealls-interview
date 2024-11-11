// Code generated by mockery v2.46.3. DO NOT EDIT.

package cryptx

import mock "github.com/stretchr/testify/mock"

// MockHelper is an autogenerated mock type for the Helper type
type MockHelper struct {
	mock.Mock
}

type MockHelper_Expecter struct {
	mock *mock.Mock
}

func (_m *MockHelper) EXPECT() *MockHelper_Expecter {
	return &MockHelper_Expecter{mock: &_m.Mock}
}

// BcryptHash provides a mock function with given fields: password
func (_m *MockHelper) BcryptHash(password string) (string, error) {
	ret := _m.Called(password)

	if len(ret) == 0 {
		panic("no return value specified for BcryptHash")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(password)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(password)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockHelper_BcryptHash_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BcryptHash'
type MockHelper_BcryptHash_Call struct {
	*mock.Call
}

// BcryptHash is a helper method to define mock.On call
//   - password string
func (_e *MockHelper_Expecter) BcryptHash(password interface{}) *MockHelper_BcryptHash_Call {
	return &MockHelper_BcryptHash_Call{Call: _e.mock.On("BcryptHash", password)}
}

func (_c *MockHelper_BcryptHash_Call) Run(run func(password string)) *MockHelper_BcryptHash_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockHelper_BcryptHash_Call) Return(_a0 string, _a1 error) *MockHelper_BcryptHash_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockHelper_BcryptHash_Call) RunAndReturn(run func(string) (string, error)) *MockHelper_BcryptHash_Call {
	_c.Call.Return(run)
	return _c
}

// BcryptValidate provides a mock function with given fields: hashedPassword, password
func (_m *MockHelper) BcryptValidate(hashedPassword string, password string) error {
	ret := _m.Called(hashedPassword, password)

	if len(ret) == 0 {
		panic("no return value specified for BcryptValidate")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(hashedPassword, password)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockHelper_BcryptValidate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BcryptValidate'
type MockHelper_BcryptValidate_Call struct {
	*mock.Call
}

// BcryptValidate is a helper method to define mock.On call
//   - hashedPassword string
//   - password string
func (_e *MockHelper_Expecter) BcryptValidate(hashedPassword interface{}, password interface{}) *MockHelper_BcryptValidate_Call {
	return &MockHelper_BcryptValidate_Call{Call: _e.mock.On("BcryptValidate", hashedPassword, password)}
}

func (_c *MockHelper_BcryptValidate_Call) Run(run func(hashedPassword string, password string)) *MockHelper_BcryptValidate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockHelper_BcryptValidate_Call) Return(_a0 error) *MockHelper_BcryptValidate_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockHelper_BcryptValidate_Call) RunAndReturn(run func(string, string) error) *MockHelper_BcryptValidate_Call {
	_c.Call.Return(run)
	return _c
}

// DecryptRSAWithBase64 provides a mock function with given fields: base64PrivateKey, base64Ciphertext
func (_m *MockHelper) DecryptRSAWithBase64(base64PrivateKey string, base64Ciphertext string) (string, error) {
	ret := _m.Called(base64PrivateKey, base64Ciphertext)

	if len(ret) == 0 {
		panic("no return value specified for DecryptRSAWithBase64")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (string, error)); ok {
		return rf(base64PrivateKey, base64Ciphertext)
	}
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(base64PrivateKey, base64Ciphertext)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(base64PrivateKey, base64Ciphertext)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockHelper_DecryptRSAWithBase64_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DecryptRSAWithBase64'
type MockHelper_DecryptRSAWithBase64_Call struct {
	*mock.Call
}

// DecryptRSAWithBase64 is a helper method to define mock.On call
//   - base64PrivateKey string
//   - base64Ciphertext string
func (_e *MockHelper_Expecter) DecryptRSAWithBase64(base64PrivateKey interface{}, base64Ciphertext interface{}) *MockHelper_DecryptRSAWithBase64_Call {
	return &MockHelper_DecryptRSAWithBase64_Call{Call: _e.mock.On("DecryptRSAWithBase64", base64PrivateKey, base64Ciphertext)}
}

func (_c *MockHelper_DecryptRSAWithBase64_Call) Run(run func(base64PrivateKey string, base64Ciphertext string)) *MockHelper_DecryptRSAWithBase64_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockHelper_DecryptRSAWithBase64_Call) Return(_a0 string, _a1 error) *MockHelper_DecryptRSAWithBase64_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockHelper_DecryptRSAWithBase64_Call) RunAndReturn(run func(string, string) (string, error)) *MockHelper_DecryptRSAWithBase64_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockHelper creates a new instance of MockHelper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockHelper(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockHelper {
	mock := &MockHelper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}