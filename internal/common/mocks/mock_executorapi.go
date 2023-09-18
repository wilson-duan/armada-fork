// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/armadaproject/armada/pkg/executorapi (interfaces: ExecutorApiClient,ExecutorApi_LeaseJobRunsClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	executorapi "github.com/armadaproject/armada/pkg/executorapi"
	types "github.com/gogo/protobuf/types"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
)

// MockExecutorApiClient is a mock of ExecutorApiClient interface.
type MockExecutorApiClient struct {
	ctrl     *gomock.Controller
	recorder *MockExecutorApiClientMockRecorder
}

// MockExecutorApiClientMockRecorder is the mock recorder for MockExecutorApiClient.
type MockExecutorApiClientMockRecorder struct {
	mock *MockExecutorApiClient
}

// NewMockExecutorApiClient creates a new mock instance.
func NewMockExecutorApiClient(ctrl *gomock.Controller) *MockExecutorApiClient {
	mock := &MockExecutorApiClient{ctrl: ctrl}
	mock.recorder = &MockExecutorApiClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExecutorApiClient) EXPECT() *MockExecutorApiClientMockRecorder {
	return m.recorder
}

// LeaseJobRuns mocks base method.
func (m *MockExecutorApiClient) LeaseJobRuns(arg0 context.Context, arg1 ...grpc.CallOption) (executorapi.ExecutorApi_LeaseJobRunsClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "LeaseJobRuns", varargs...)
	ret0, _ := ret[0].(executorapi.ExecutorApi_LeaseJobRunsClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LeaseJobRuns indicates an expected call of LeaseJobRuns.
func (mr *MockExecutorApiClientMockRecorder) LeaseJobRuns(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LeaseJobRuns", reflect.TypeOf((*MockExecutorApiClient)(nil).LeaseJobRuns), varargs...)
}

// ReportEvents mocks base method.
func (m *MockExecutorApiClient) ReportEvents(arg0 context.Context, arg1 *executorapi.EventList, arg2 ...grpc.CallOption) (*types.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ReportEvents", varargs...)
	ret0, _ := ret[0].(*types.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReportEvents indicates an expected call of ReportEvents.
func (mr *MockExecutorApiClientMockRecorder) ReportEvents(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReportEvents", reflect.TypeOf((*MockExecutorApiClient)(nil).ReportEvents), varargs...)
}

// MockExecutorApi_LeaseJobRunsClient is a mock of ExecutorApi_LeaseJobRunsClient interface.
type MockExecutorApi_LeaseJobRunsClient struct {
	ctrl     *gomock.Controller
	recorder *MockExecutorApi_LeaseJobRunsClientMockRecorder
}

// MockExecutorApi_LeaseJobRunsClientMockRecorder is the mock recorder for MockExecutorApi_LeaseJobRunsClient.
type MockExecutorApi_LeaseJobRunsClientMockRecorder struct {
	mock *MockExecutorApi_LeaseJobRunsClient
}

// NewMockExecutorApi_LeaseJobRunsClient creates a new mock instance.
func NewMockExecutorApi_LeaseJobRunsClient(ctrl *gomock.Controller) *MockExecutorApi_LeaseJobRunsClient {
	mock := &MockExecutorApi_LeaseJobRunsClient{ctrl: ctrl}
	mock.recorder = &MockExecutorApi_LeaseJobRunsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExecutorApi_LeaseJobRunsClient) EXPECT() *MockExecutorApi_LeaseJobRunsClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method.
func (m *MockExecutorApi_LeaseJobRunsClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend.
func (mr *MockExecutorApi_LeaseJobRunsClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockExecutorApi_LeaseJobRunsClient)(nil).CloseSend))
}

// Context mocks base method.
func (m *MockExecutorApi_LeaseJobRunsClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockExecutorApi_LeaseJobRunsClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockExecutorApi_LeaseJobRunsClient)(nil).Context))
}

// Header mocks base method.
func (m *MockExecutorApi_LeaseJobRunsClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header.
func (mr *MockExecutorApi_LeaseJobRunsClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockExecutorApi_LeaseJobRunsClient)(nil).Header))
}

// Recv mocks base method.
func (m *MockExecutorApi_LeaseJobRunsClient) Recv() (*executorapi.LeaseStreamMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*executorapi.LeaseStreamMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockExecutorApi_LeaseJobRunsClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockExecutorApi_LeaseJobRunsClient)(nil).Recv))
}

// RecvMsg mocks base method.
func (m *MockExecutorApi_LeaseJobRunsClient) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockExecutorApi_LeaseJobRunsClientMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockExecutorApi_LeaseJobRunsClient)(nil).RecvMsg), arg0)
}

// Send mocks base method.
func (m *MockExecutorApi_LeaseJobRunsClient) Send(arg0 *executorapi.LeaseRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockExecutorApi_LeaseJobRunsClientMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockExecutorApi_LeaseJobRunsClient)(nil).Send), arg0)
}

// SendMsg mocks base method.
func (m *MockExecutorApi_LeaseJobRunsClient) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockExecutorApi_LeaseJobRunsClientMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockExecutorApi_LeaseJobRunsClient)(nil).SendMsg), arg0)
}

// Trailer mocks base method.
func (m *MockExecutorApi_LeaseJobRunsClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer.
func (mr *MockExecutorApi_LeaseJobRunsClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockExecutorApi_LeaseJobRunsClient)(nil).Trailer))
}
