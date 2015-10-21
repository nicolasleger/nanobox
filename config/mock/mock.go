// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/nanobox-io/nanobox/config (interfaces: Config)

package mock_config

import (
	gomock "github.com/golang/mock/gomock"
	config "github.com/nanobox-io/nanobox/config"
)

// Mock of Config interface
type MockConfig struct {
	ctrl     *gomock.Controller
	recorder *_MockConfigRecorder
}

// Recorder for MockConfig (not exported)
type _MockConfigRecorder struct {
	mock *MockConfig
}

func NewMockConfig(ctrl *gomock.Controller) *MockConfig {
	mock := &MockConfig{ctrl: ctrl}
	mock.recorder = &_MockConfigRecorder{mock}
	return mock
}

func (_m *MockConfig) EXPECT() *_MockConfigRecorder {
	return _m.recorder
}

func (_m *MockConfig) Debug(_param0 string, _param1 bool) {
	_m.ctrl.Call(_m, "Debug", _param0, _param1)
}

func (_mr *_MockConfigRecorder) Debug(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Debug", arg0, arg1)
}

func (_m *MockConfig) Fatal(_param0 string, _param1 string) {
	_m.ctrl.Call(_m, "Fatal", _param0, _param1)
}

func (_mr *_MockConfigRecorder) Fatal(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Fatal", arg0, arg1)
}

func (_m *MockConfig) Info(_param0 string, _param1 bool) {
	_m.ctrl.Call(_m, "Info", _param0, _param1)
}

func (_mr *_MockConfigRecorder) Info(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Info", arg0, arg1)
}

func (_m *MockConfig) ParseConfig(_param0 string, _param1 interface{}) error {
	ret := _m.ctrl.Call(_m, "ParseConfig", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockConfigRecorder) ParseConfig(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ParseConfig", arg0, arg1)
}

func (_m *MockConfig) ParseNanofile() *config.NanofileConfig {
	ret := _m.ctrl.Call(_m, "ParseNanofile")
	ret0, _ := ret[0].(*config.NanofileConfig)
	return ret0
}

func (_mr *_MockConfigRecorder) ParseNanofile() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ParseNanofile")
}

func (_m *MockConfig) ParseVMfile() *config.VMfileConfig {
	ret := _m.ctrl.Call(_m, "ParseVMfile")
	ret0, _ := ret[0].(*config.VMfileConfig)
	return ret0
}

func (_mr *_MockConfigRecorder) ParseVMfile() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ParseVMfile")
}

func (_m *MockConfig) Root() string {
	ret := _m.ctrl.Call(_m, "Root")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockConfigRecorder) Root() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Root")
}
