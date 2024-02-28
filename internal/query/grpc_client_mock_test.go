// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ydb-platform/ydb-go-genproto/Ydb_Query_V1 (interfaces: QueryServiceClient,QueryService_AttachSessionClient,QueryService_ExecuteQueryClient)
//
// Generated by this command:
//
//	mockgen -destination grpc_client_mock_test.go -package query -write_package_comment=false github.com/ydb-platform/ydb-go-genproto/Ydb_Query_V1 QueryServiceClient,QueryService_AttachSessionClient,QueryService_ExecuteQueryClient
package query

import (
	context "context"
	reflect "reflect"

	Ydb_Query_V1 "github.com/ydb-platform/ydb-go-genproto/Ydb_Query_V1"
	Ydb_Operations "github.com/ydb-platform/ydb-go-genproto/protos/Ydb_Operations"
	Ydb_Query "github.com/ydb-platform/ydb-go-genproto/protos/Ydb_Query"
	gomock "go.uber.org/mock/gomock"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
)

// MockQueryServiceClient is a mock of QueryServiceClient interface.
type MockQueryServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockQueryServiceClientMockRecorder
}

// MockQueryServiceClientMockRecorder is the mock recorder for MockQueryServiceClient.
type MockQueryServiceClientMockRecorder struct {
	mock *MockQueryServiceClient
}

// NewMockQueryServiceClient creates a new mock instance.
func NewMockQueryServiceClient(ctrl *gomock.Controller) *MockQueryServiceClient {
	mock := &MockQueryServiceClient{ctrl: ctrl}
	mock.recorder = &MockQueryServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQueryServiceClient) EXPECT() *MockQueryServiceClientMockRecorder {
	return m.recorder
}

// AttachSession mocks base method.
func (m *MockQueryServiceClient) AttachSession(arg0 context.Context, arg1 *Ydb_Query.AttachSessionRequest, arg2 ...grpc.CallOption) (Ydb_Query_V1.QueryService_AttachSessionClient, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AttachSession", varargs...)
	ret0, _ := ret[0].(Ydb_Query_V1.QueryService_AttachSessionClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AttachSession indicates an expected call of AttachSession.
func (mr *MockQueryServiceClientMockRecorder) AttachSession(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AttachSession", reflect.TypeOf((*MockQueryServiceClient)(nil).AttachSession), varargs...)
}

// BeginTransaction mocks base method.
func (m *MockQueryServiceClient) BeginTransaction(arg0 context.Context, arg1 *Ydb_Query.BeginTransactionRequest, arg2 ...grpc.CallOption) (*Ydb_Query.BeginTransactionResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BeginTransaction", varargs...)
	ret0, _ := ret[0].(*Ydb_Query.BeginTransactionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BeginTransaction indicates an expected call of BeginTransaction.
func (mr *MockQueryServiceClientMockRecorder) BeginTransaction(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginTransaction", reflect.TypeOf((*MockQueryServiceClient)(nil).BeginTransaction), varargs...)
}

// CommitTransaction mocks base method.
func (m *MockQueryServiceClient) CommitTransaction(arg0 context.Context, arg1 *Ydb_Query.CommitTransactionRequest, arg2 ...grpc.CallOption) (*Ydb_Query.CommitTransactionResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CommitTransaction", varargs...)
	ret0, _ := ret[0].(*Ydb_Query.CommitTransactionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CommitTransaction indicates an expected call of CommitTransaction.
func (mr *MockQueryServiceClientMockRecorder) CommitTransaction(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommitTransaction", reflect.TypeOf((*MockQueryServiceClient)(nil).CommitTransaction), varargs...)
}

// CreateSession mocks base method.
func (m *MockQueryServiceClient) CreateSession(arg0 context.Context, arg1 *Ydb_Query.CreateSessionRequest, arg2 ...grpc.CallOption) (*Ydb_Query.CreateSessionResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateSession", varargs...)
	ret0, _ := ret[0].(*Ydb_Query.CreateSessionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSession indicates an expected call of CreateSession.
func (mr *MockQueryServiceClientMockRecorder) CreateSession(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSession", reflect.TypeOf((*MockQueryServiceClient)(nil).CreateSession), varargs...)
}

// DeleteSession mocks base method.
func (m *MockQueryServiceClient) DeleteSession(arg0 context.Context, arg1 *Ydb_Query.DeleteSessionRequest, arg2 ...grpc.CallOption) (*Ydb_Query.DeleteSessionResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteSession", varargs...)
	ret0, _ := ret[0].(*Ydb_Query.DeleteSessionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteSession indicates an expected call of DeleteSession.
func (mr *MockQueryServiceClientMockRecorder) DeleteSession(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSession", reflect.TypeOf((*MockQueryServiceClient)(nil).DeleteSession), varargs...)
}

// ExecuteQuery mocks base method.
func (m *MockQueryServiceClient) ExecuteQuery(arg0 context.Context, arg1 *Ydb_Query.ExecuteQueryRequest, arg2 ...grpc.CallOption) (Ydb_Query_V1.QueryService_ExecuteQueryClient, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ExecuteQuery", varargs...)
	ret0, _ := ret[0].(Ydb_Query_V1.QueryService_ExecuteQueryClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExecuteQuery indicates an expected call of ExecuteQuery.
func (mr *MockQueryServiceClientMockRecorder) ExecuteQuery(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecuteQuery", reflect.TypeOf((*MockQueryServiceClient)(nil).ExecuteQuery), varargs...)
}

// ExecuteScript mocks base method.
func (m *MockQueryServiceClient) ExecuteScript(arg0 context.Context, arg1 *Ydb_Query.ExecuteScriptRequest, arg2 ...grpc.CallOption) (*Ydb_Operations.Operation, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ExecuteScript", varargs...)
	ret0, _ := ret[0].(*Ydb_Operations.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExecuteScript indicates an expected call of ExecuteScript.
func (mr *MockQueryServiceClientMockRecorder) ExecuteScript(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecuteScript", reflect.TypeOf((*MockQueryServiceClient)(nil).ExecuteScript), varargs...)
}

// FetchScriptResults mocks base method.
func (m *MockQueryServiceClient) FetchScriptResults(arg0 context.Context, arg1 *Ydb_Query.FetchScriptResultsRequest, arg2 ...grpc.CallOption) (*Ydb_Query.FetchScriptResultsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FetchScriptResults", varargs...)
	ret0, _ := ret[0].(*Ydb_Query.FetchScriptResultsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchScriptResults indicates an expected call of FetchScriptResults.
func (mr *MockQueryServiceClientMockRecorder) FetchScriptResults(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchScriptResults", reflect.TypeOf((*MockQueryServiceClient)(nil).FetchScriptResults), varargs...)
}

// RollbackTransaction mocks base method.
func (m *MockQueryServiceClient) RollbackTransaction(arg0 context.Context, arg1 *Ydb_Query.RollbackTransactionRequest, arg2 ...grpc.CallOption) (*Ydb_Query.RollbackTransactionResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RollbackTransaction", varargs...)
	ret0, _ := ret[0].(*Ydb_Query.RollbackTransactionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RollbackTransaction indicates an expected call of RollbackTransaction.
func (mr *MockQueryServiceClientMockRecorder) RollbackTransaction(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RollbackTransaction", reflect.TypeOf((*MockQueryServiceClient)(nil).RollbackTransaction), varargs...)
}

// MockQueryService_AttachSessionClient is a mock of QueryService_AttachSessionClient interface.
type MockQueryService_AttachSessionClient struct {
	ctrl     *gomock.Controller
	recorder *MockQueryService_AttachSessionClientMockRecorder
}

// MockQueryService_AttachSessionClientMockRecorder is the mock recorder for MockQueryService_AttachSessionClient.
type MockQueryService_AttachSessionClientMockRecorder struct {
	mock *MockQueryService_AttachSessionClient
}

// NewMockQueryService_AttachSessionClient creates a new mock instance.
func NewMockQueryService_AttachSessionClient(ctrl *gomock.Controller) *MockQueryService_AttachSessionClient {
	mock := &MockQueryService_AttachSessionClient{ctrl: ctrl}
	mock.recorder = &MockQueryService_AttachSessionClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQueryService_AttachSessionClient) EXPECT() *MockQueryService_AttachSessionClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method.
func (m *MockQueryService_AttachSessionClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend.
func (mr *MockQueryService_AttachSessionClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockQueryService_AttachSessionClient)(nil).CloseSend))
}

// Context mocks base method.
func (m *MockQueryService_AttachSessionClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockQueryService_AttachSessionClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockQueryService_AttachSessionClient)(nil).Context))
}

// Header mocks base method.
func (m *MockQueryService_AttachSessionClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header.
func (mr *MockQueryService_AttachSessionClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockQueryService_AttachSessionClient)(nil).Header))
}

// Recv mocks base method.
func (m *MockQueryService_AttachSessionClient) Recv() (*Ydb_Query.SessionState, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*Ydb_Query.SessionState)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockQueryService_AttachSessionClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockQueryService_AttachSessionClient)(nil).Recv))
}

// RecvMsg mocks base method.
func (m *MockQueryService_AttachSessionClient) RecvMsg(arg0 any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockQueryService_AttachSessionClientMockRecorder) RecvMsg(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockQueryService_AttachSessionClient)(nil).RecvMsg), arg0)
}

// SendMsg mocks base method.
func (m *MockQueryService_AttachSessionClient) SendMsg(arg0 any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockQueryService_AttachSessionClientMockRecorder) SendMsg(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockQueryService_AttachSessionClient)(nil).SendMsg), arg0)
}

// Trailer mocks base method.
func (m *MockQueryService_AttachSessionClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer.
func (mr *MockQueryService_AttachSessionClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockQueryService_AttachSessionClient)(nil).Trailer))
}

// MockQueryService_ExecuteQueryClient is a mock of QueryService_ExecuteQueryClient interface.
type MockQueryService_ExecuteQueryClient struct {
	ctrl     *gomock.Controller
	recorder *MockQueryService_ExecuteQueryClientMockRecorder
}

// MockQueryService_ExecuteQueryClientMockRecorder is the mock recorder for MockQueryService_ExecuteQueryClient.
type MockQueryService_ExecuteQueryClientMockRecorder struct {
	mock *MockQueryService_ExecuteQueryClient
}

// NewMockQueryService_ExecuteQueryClient creates a new mock instance.
func NewMockQueryService_ExecuteQueryClient(ctrl *gomock.Controller) *MockQueryService_ExecuteQueryClient {
	mock := &MockQueryService_ExecuteQueryClient{ctrl: ctrl}
	mock.recorder = &MockQueryService_ExecuteQueryClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQueryService_ExecuteQueryClient) EXPECT() *MockQueryService_ExecuteQueryClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method.
func (m *MockQueryService_ExecuteQueryClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend.
func (mr *MockQueryService_ExecuteQueryClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockQueryService_ExecuteQueryClient)(nil).CloseSend))
}

// Context mocks base method.
func (m *MockQueryService_ExecuteQueryClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockQueryService_ExecuteQueryClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockQueryService_ExecuteQueryClient)(nil).Context))
}

// Header mocks base method.
func (m *MockQueryService_ExecuteQueryClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header.
func (mr *MockQueryService_ExecuteQueryClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockQueryService_ExecuteQueryClient)(nil).Header))
}

// Recv mocks base method.
func (m *MockQueryService_ExecuteQueryClient) Recv() (*Ydb_Query.ExecuteQueryResponsePart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*Ydb_Query.ExecuteQueryResponsePart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockQueryService_ExecuteQueryClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockQueryService_ExecuteQueryClient)(nil).Recv))
}

// RecvMsg mocks base method.
func (m *MockQueryService_ExecuteQueryClient) RecvMsg(arg0 any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockQueryService_ExecuteQueryClientMockRecorder) RecvMsg(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockQueryService_ExecuteQueryClient)(nil).RecvMsg), arg0)
}

// SendMsg mocks base method.
func (m *MockQueryService_ExecuteQueryClient) SendMsg(arg0 any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockQueryService_ExecuteQueryClientMockRecorder) SendMsg(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockQueryService_ExecuteQueryClient)(nil).SendMsg), arg0)
}

// Trailer mocks base method.
func (m *MockQueryService_ExecuteQueryClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer.
func (mr *MockQueryService_ExecuteQueryClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockQueryService_ExecuteQueryClient)(nil).Trailer))
}
