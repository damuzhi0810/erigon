// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/erigontech/erigon-lib/gointerfaces/remoteproto (interfaces: KVClient)
//
// Generated by this command:
//
//	mockgen -typed=true -destination=./kv_client_mock.go -package=remoteproto . KVClient
//

// Package remoteproto is a generated GoMock package.
package remoteproto

import (
	context "context"
	reflect "reflect"

	typesproto "github.com/erigontech/erigon-lib/gointerfaces/typesproto"
	gomock "go.uber.org/mock/gomock"
	grpc "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// MockKVClient is a mock of KVClient interface.
type MockKVClient struct {
	ctrl     *gomock.Controller
	recorder *MockKVClientMockRecorder
	isgomock struct{}
}

// MockKVClientMockRecorder is the mock recorder for MockKVClient.
type MockKVClientMockRecorder struct {
	mock *MockKVClient
}

// NewMockKVClient creates a new mock instance.
func NewMockKVClient(ctrl *gomock.Controller) *MockKVClient {
	mock := &MockKVClient{ctrl: ctrl}
	mock.recorder = &MockKVClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKVClient) EXPECT() *MockKVClientMockRecorder {
	return m.recorder
}

// GetLatest mocks base method.
func (m *MockKVClient) GetLatest(ctx context.Context, in *GetLatestReq, opts ...grpc.CallOption) (*GetLatestReply, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetLatest", varargs...)
	ret0, _ := ret[0].(*GetLatestReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLatest indicates an expected call of GetLatest.
func (mr *MockKVClientMockRecorder) GetLatest(ctx, in any, opts ...any) *MockKVClientGetLatestCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLatest", reflect.TypeOf((*MockKVClient)(nil).GetLatest), varargs...)
	return &MockKVClientGetLatestCall{Call: call}
}

// MockKVClientGetLatestCall wrap *gomock.Call
type MockKVClientGetLatestCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockKVClientGetLatestCall) Return(arg0 *GetLatestReply, arg1 error) *MockKVClientGetLatestCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockKVClientGetLatestCall) Do(f func(context.Context, *GetLatestReq, ...grpc.CallOption) (*GetLatestReply, error)) *MockKVClientGetLatestCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockKVClientGetLatestCall) DoAndReturn(f func(context.Context, *GetLatestReq, ...grpc.CallOption) (*GetLatestReply, error)) *MockKVClientGetLatestCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// HistoryRange mocks base method.
func (m *MockKVClient) HistoryRange(ctx context.Context, in *HistoryRangeReq, opts ...grpc.CallOption) (*Pairs, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "HistoryRange", varargs...)
	ret0, _ := ret[0].(*Pairs)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HistoryRange indicates an expected call of HistoryRange.
func (mr *MockKVClientMockRecorder) HistoryRange(ctx, in any, opts ...any) *MockKVClientHistoryRangeCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HistoryRange", reflect.TypeOf((*MockKVClient)(nil).HistoryRange), varargs...)
	return &MockKVClientHistoryRangeCall{Call: call}
}

// MockKVClientHistoryRangeCall wrap *gomock.Call
type MockKVClientHistoryRangeCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockKVClientHistoryRangeCall) Return(arg0 *Pairs, arg1 error) *MockKVClientHistoryRangeCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockKVClientHistoryRangeCall) Do(f func(context.Context, *HistoryRangeReq, ...grpc.CallOption) (*Pairs, error)) *MockKVClientHistoryRangeCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockKVClientHistoryRangeCall) DoAndReturn(f func(context.Context, *HistoryRangeReq, ...grpc.CallOption) (*Pairs, error)) *MockKVClientHistoryRangeCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// HistorySeek mocks base method.
func (m *MockKVClient) HistorySeek(ctx context.Context, in *HistorySeekReq, opts ...grpc.CallOption) (*HistorySeekReply, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "HistorySeek", varargs...)
	ret0, _ := ret[0].(*HistorySeekReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HistorySeek indicates an expected call of HistorySeek.
func (mr *MockKVClientMockRecorder) HistorySeek(ctx, in any, opts ...any) *MockKVClientHistorySeekCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HistorySeek", reflect.TypeOf((*MockKVClient)(nil).HistorySeek), varargs...)
	return &MockKVClientHistorySeekCall{Call: call}
}

// MockKVClientHistorySeekCall wrap *gomock.Call
type MockKVClientHistorySeekCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockKVClientHistorySeekCall) Return(arg0 *HistorySeekReply, arg1 error) *MockKVClientHistorySeekCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockKVClientHistorySeekCall) Do(f func(context.Context, *HistorySeekReq, ...grpc.CallOption) (*HistorySeekReply, error)) *MockKVClientHistorySeekCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockKVClientHistorySeekCall) DoAndReturn(f func(context.Context, *HistorySeekReq, ...grpc.CallOption) (*HistorySeekReply, error)) *MockKVClientHistorySeekCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// IndexRange mocks base method.
func (m *MockKVClient) IndexRange(ctx context.Context, in *IndexRangeReq, opts ...grpc.CallOption) (*IndexRangeReply, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "IndexRange", varargs...)
	ret0, _ := ret[0].(*IndexRangeReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IndexRange indicates an expected call of IndexRange.
func (mr *MockKVClientMockRecorder) IndexRange(ctx, in any, opts ...any) *MockKVClientIndexRangeCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IndexRange", reflect.TypeOf((*MockKVClient)(nil).IndexRange), varargs...)
	return &MockKVClientIndexRangeCall{Call: call}
}

// MockKVClientIndexRangeCall wrap *gomock.Call
type MockKVClientIndexRangeCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockKVClientIndexRangeCall) Return(arg0 *IndexRangeReply, arg1 error) *MockKVClientIndexRangeCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockKVClientIndexRangeCall) Do(f func(context.Context, *IndexRangeReq, ...grpc.CallOption) (*IndexRangeReply, error)) *MockKVClientIndexRangeCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockKVClientIndexRangeCall) DoAndReturn(f func(context.Context, *IndexRangeReq, ...grpc.CallOption) (*IndexRangeReply, error)) *MockKVClientIndexRangeCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Range mocks base method.
func (m *MockKVClient) Range(ctx context.Context, in *RangeReq, opts ...grpc.CallOption) (*Pairs, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Range", varargs...)
	ret0, _ := ret[0].(*Pairs)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Range indicates an expected call of Range.
func (mr *MockKVClientMockRecorder) Range(ctx, in any, opts ...any) *MockKVClientRangeCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Range", reflect.TypeOf((*MockKVClient)(nil).Range), varargs...)
	return &MockKVClientRangeCall{Call: call}
}

// MockKVClientRangeCall wrap *gomock.Call
type MockKVClientRangeCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockKVClientRangeCall) Return(arg0 *Pairs, arg1 error) *MockKVClientRangeCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockKVClientRangeCall) Do(f func(context.Context, *RangeReq, ...grpc.CallOption) (*Pairs, error)) *MockKVClientRangeCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockKVClientRangeCall) DoAndReturn(f func(context.Context, *RangeReq, ...grpc.CallOption) (*Pairs, error)) *MockKVClientRangeCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// RangeAsOf mocks base method.
func (m *MockKVClient) RangeAsOf(ctx context.Context, in *RangeAsOfReq, opts ...grpc.CallOption) (*Pairs, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RangeAsOf", varargs...)
	ret0, _ := ret[0].(*Pairs)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RangeAsOf indicates an expected call of RangeAsOf.
func (mr *MockKVClientMockRecorder) RangeAsOf(ctx, in any, opts ...any) *MockKVClientRangeAsOfCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RangeAsOf", reflect.TypeOf((*MockKVClient)(nil).RangeAsOf), varargs...)
	return &MockKVClientRangeAsOfCall{Call: call}
}

// MockKVClientRangeAsOfCall wrap *gomock.Call
type MockKVClientRangeAsOfCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockKVClientRangeAsOfCall) Return(arg0 *Pairs, arg1 error) *MockKVClientRangeAsOfCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockKVClientRangeAsOfCall) Do(f func(context.Context, *RangeAsOfReq, ...grpc.CallOption) (*Pairs, error)) *MockKVClientRangeAsOfCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockKVClientRangeAsOfCall) DoAndReturn(f func(context.Context, *RangeAsOfReq, ...grpc.CallOption) (*Pairs, error)) *MockKVClientRangeAsOfCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Snapshots mocks base method.
func (m *MockKVClient) Snapshots(ctx context.Context, in *SnapshotsRequest, opts ...grpc.CallOption) (*SnapshotsReply, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Snapshots", varargs...)
	ret0, _ := ret[0].(*SnapshotsReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Snapshots indicates an expected call of Snapshots.
func (mr *MockKVClientMockRecorder) Snapshots(ctx, in any, opts ...any) *MockKVClientSnapshotsCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Snapshots", reflect.TypeOf((*MockKVClient)(nil).Snapshots), varargs...)
	return &MockKVClientSnapshotsCall{Call: call}
}

// MockKVClientSnapshotsCall wrap *gomock.Call
type MockKVClientSnapshotsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockKVClientSnapshotsCall) Return(arg0 *SnapshotsReply, arg1 error) *MockKVClientSnapshotsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockKVClientSnapshotsCall) Do(f func(context.Context, *SnapshotsRequest, ...grpc.CallOption) (*SnapshotsReply, error)) *MockKVClientSnapshotsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockKVClientSnapshotsCall) DoAndReturn(f func(context.Context, *SnapshotsRequest, ...grpc.CallOption) (*SnapshotsReply, error)) *MockKVClientSnapshotsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// StateChanges mocks base method.
func (m *MockKVClient) StateChanges(ctx context.Context, in *StateChangeRequest, opts ...grpc.CallOption) (KV_StateChangesClient, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StateChanges", varargs...)
	ret0, _ := ret[0].(KV_StateChangesClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateChanges indicates an expected call of StateChanges.
func (mr *MockKVClientMockRecorder) StateChanges(ctx, in any, opts ...any) *MockKVClientStateChangesCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateChanges", reflect.TypeOf((*MockKVClient)(nil).StateChanges), varargs...)
	return &MockKVClientStateChangesCall{Call: call}
}

// MockKVClientStateChangesCall wrap *gomock.Call
type MockKVClientStateChangesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockKVClientStateChangesCall) Return(arg0 KV_StateChangesClient, arg1 error) *MockKVClientStateChangesCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockKVClientStateChangesCall) Do(f func(context.Context, *StateChangeRequest, ...grpc.CallOption) (KV_StateChangesClient, error)) *MockKVClientStateChangesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockKVClientStateChangesCall) DoAndReturn(f func(context.Context, *StateChangeRequest, ...grpc.CallOption) (KV_StateChangesClient, error)) *MockKVClientStateChangesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Tx mocks base method.
func (m *MockKVClient) Tx(ctx context.Context, opts ...grpc.CallOption) (KV_TxClient, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Tx", varargs...)
	ret0, _ := ret[0].(KV_TxClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Tx indicates an expected call of Tx.
func (mr *MockKVClientMockRecorder) Tx(ctx any, opts ...any) *MockKVClientTxCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx}, opts...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tx", reflect.TypeOf((*MockKVClient)(nil).Tx), varargs...)
	return &MockKVClientTxCall{Call: call}
}

// MockKVClientTxCall wrap *gomock.Call
type MockKVClientTxCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockKVClientTxCall) Return(arg0 KV_TxClient, arg1 error) *MockKVClientTxCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockKVClientTxCall) Do(f func(context.Context, ...grpc.CallOption) (KV_TxClient, error)) *MockKVClientTxCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockKVClientTxCall) DoAndReturn(f func(context.Context, ...grpc.CallOption) (KV_TxClient, error)) *MockKVClientTxCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Version mocks base method.
func (m *MockKVClient) Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*typesproto.VersionReply, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Version", varargs...)
	ret0, _ := ret[0].(*typesproto.VersionReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Version indicates an expected call of Version.
func (mr *MockKVClientMockRecorder) Version(ctx, in any, opts ...any) *MockKVClientVersionCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Version", reflect.TypeOf((*MockKVClient)(nil).Version), varargs...)
	return &MockKVClientVersionCall{Call: call}
}

// MockKVClientVersionCall wrap *gomock.Call
type MockKVClientVersionCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockKVClientVersionCall) Return(arg0 *typesproto.VersionReply, arg1 error) *MockKVClientVersionCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockKVClientVersionCall) Do(f func(context.Context, *emptypb.Empty, ...grpc.CallOption) (*typesproto.VersionReply, error)) *MockKVClientVersionCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockKVClientVersionCall) DoAndReturn(f func(context.Context, *emptypb.Empty, ...grpc.CallOption) (*typesproto.VersionReply, error)) *MockKVClientVersionCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
