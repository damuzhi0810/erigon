// Code generated by MockGen. DO NOT EDIT.
// Source: ./entity_store.go
//
// Generated by this command:
//
//	mockgen -typed=true -source=./entity_store.go -destination=./entity_store_mock.go -package=heimdall
//

// Package heimdall is a generated GoMock package.
package heimdall

import (
	context "context"
	reflect "reflect"

	snaptype "github.com/erigontech/erigon-lib/downloader/snaptype"
	gomock "go.uber.org/mock/gomock"
)

// MockEntityStore is a mock of EntityStore interface.
type MockEntityStore[TEntity Entity] struct {
	ctrl     *gomock.Controller
	recorder *MockEntityStoreMockRecorder[TEntity]
	isgomock struct{}
}

// MockEntityStoreMockRecorder is the mock recorder for MockEntityStore.
type MockEntityStoreMockRecorder[TEntity Entity] struct {
	mock *MockEntityStore[TEntity]
}

// NewMockEntityStore creates a new mock instance.
func NewMockEntityStore[TEntity Entity](ctrl *gomock.Controller) *MockEntityStore[TEntity] {
	mock := &MockEntityStore[TEntity]{ctrl: ctrl}
	mock.recorder = &MockEntityStoreMockRecorder[TEntity]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEntityStore[TEntity]) EXPECT() *MockEntityStoreMockRecorder[TEntity] {
	return m.recorder
}

// Close mocks base method.
func (m *MockEntityStore[TEntity]) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockEntityStoreMockRecorder[TEntity]) Close() *MockEntityStoreCloseCall[TEntity] {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockEntityStore[TEntity])(nil).Close))
	return &MockEntityStoreCloseCall[TEntity]{Call: call}
}

// MockEntityStoreCloseCall wrap *gomock.Call
type MockEntityStoreCloseCall[TEntity Entity] struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockEntityStoreCloseCall[TEntity]) Return() *MockEntityStoreCloseCall[TEntity] {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockEntityStoreCloseCall[TEntity]) Do(f func()) *MockEntityStoreCloseCall[TEntity] {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockEntityStoreCloseCall[TEntity]) DoAndReturn(f func()) *MockEntityStoreCloseCall[TEntity] {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Entity mocks base method.
func (m *MockEntityStore[TEntity]) Entity(ctx context.Context, id uint64) (TEntity, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Entity", ctx, id)
	ret0, _ := ret[0].(TEntity)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Entity indicates an expected call of Entity.
func (mr *MockEntityStoreMockRecorder[TEntity]) Entity(ctx, id any) *MockEntityStoreEntityCall[TEntity] {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Entity", reflect.TypeOf((*MockEntityStore[TEntity])(nil).Entity), ctx, id)
	return &MockEntityStoreEntityCall[TEntity]{Call: call}
}

// MockEntityStoreEntityCall wrap *gomock.Call
type MockEntityStoreEntityCall[TEntity Entity] struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockEntityStoreEntityCall[TEntity]) Return(arg0 TEntity, arg1 bool, arg2 error) *MockEntityStoreEntityCall[TEntity] {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockEntityStoreEntityCall[TEntity]) Do(f func(context.Context, uint64) (TEntity, bool, error)) *MockEntityStoreEntityCall[TEntity] {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockEntityStoreEntityCall[TEntity]) DoAndReturn(f func(context.Context, uint64) (TEntity, bool, error)) *MockEntityStoreEntityCall[TEntity] {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// LastEntity mocks base method.
func (m *MockEntityStore[TEntity]) LastEntity(ctx context.Context) (TEntity, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LastEntity", ctx)
	ret0, _ := ret[0].(TEntity)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// LastEntity indicates an expected call of LastEntity.
func (mr *MockEntityStoreMockRecorder[TEntity]) LastEntity(ctx any) *MockEntityStoreLastEntityCall[TEntity] {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastEntity", reflect.TypeOf((*MockEntityStore[TEntity])(nil).LastEntity), ctx)
	return &MockEntityStoreLastEntityCall[TEntity]{Call: call}
}

// MockEntityStoreLastEntityCall wrap *gomock.Call
type MockEntityStoreLastEntityCall[TEntity Entity] struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockEntityStoreLastEntityCall[TEntity]) Return(arg0 TEntity, arg1 bool, arg2 error) *MockEntityStoreLastEntityCall[TEntity] {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockEntityStoreLastEntityCall[TEntity]) Do(f func(context.Context) (TEntity, bool, error)) *MockEntityStoreLastEntityCall[TEntity] {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockEntityStoreLastEntityCall[TEntity]) DoAndReturn(f func(context.Context) (TEntity, bool, error)) *MockEntityStoreLastEntityCall[TEntity] {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// LastEntityId mocks base method.
func (m *MockEntityStore[TEntity]) LastEntityId(ctx context.Context) (uint64, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LastEntityId", ctx)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// LastEntityId indicates an expected call of LastEntityId.
func (mr *MockEntityStoreMockRecorder[TEntity]) LastEntityId(ctx any) *MockEntityStoreLastEntityIdCall[TEntity] {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastEntityId", reflect.TypeOf((*MockEntityStore[TEntity])(nil).LastEntityId), ctx)
	return &MockEntityStoreLastEntityIdCall[TEntity]{Call: call}
}

// MockEntityStoreLastEntityIdCall wrap *gomock.Call
type MockEntityStoreLastEntityIdCall[TEntity Entity] struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockEntityStoreLastEntityIdCall[TEntity]) Return(arg0 uint64, arg1 bool, arg2 error) *MockEntityStoreLastEntityIdCall[TEntity] {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockEntityStoreLastEntityIdCall[TEntity]) Do(f func(context.Context) (uint64, bool, error)) *MockEntityStoreLastEntityIdCall[TEntity] {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockEntityStoreLastEntityIdCall[TEntity]) DoAndReturn(f func(context.Context) (uint64, bool, error)) *MockEntityStoreLastEntityIdCall[TEntity] {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// LastFrozenEntityId mocks base method.
func (m *MockEntityStore[TEntity]) LastFrozenEntityId() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LastFrozenEntityId")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// LastFrozenEntityId indicates an expected call of LastFrozenEntityId.
func (mr *MockEntityStoreMockRecorder[TEntity]) LastFrozenEntityId() *MockEntityStoreLastFrozenEntityIdCall[TEntity] {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastFrozenEntityId", reflect.TypeOf((*MockEntityStore[TEntity])(nil).LastFrozenEntityId))
	return &MockEntityStoreLastFrozenEntityIdCall[TEntity]{Call: call}
}

// MockEntityStoreLastFrozenEntityIdCall wrap *gomock.Call
type MockEntityStoreLastFrozenEntityIdCall[TEntity Entity] struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockEntityStoreLastFrozenEntityIdCall[TEntity]) Return(arg0 uint64) *MockEntityStoreLastFrozenEntityIdCall[TEntity] {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockEntityStoreLastFrozenEntityIdCall[TEntity]) Do(f func() uint64) *MockEntityStoreLastFrozenEntityIdCall[TEntity] {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockEntityStoreLastFrozenEntityIdCall[TEntity]) DoAndReturn(f func() uint64) *MockEntityStoreLastFrozenEntityIdCall[TEntity] {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Prepare mocks base method.
func (m *MockEntityStore[TEntity]) Prepare(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Prepare", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Prepare indicates an expected call of Prepare.
func (mr *MockEntityStoreMockRecorder[TEntity]) Prepare(ctx any) *MockEntityStorePrepareCall[TEntity] {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Prepare", reflect.TypeOf((*MockEntityStore[TEntity])(nil).Prepare), ctx)
	return &MockEntityStorePrepareCall[TEntity]{Call: call}
}

// MockEntityStorePrepareCall wrap *gomock.Call
type MockEntityStorePrepareCall[TEntity Entity] struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockEntityStorePrepareCall[TEntity]) Return(arg0 error) *MockEntityStorePrepareCall[TEntity] {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockEntityStorePrepareCall[TEntity]) Do(f func(context.Context) error) *MockEntityStorePrepareCall[TEntity] {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockEntityStorePrepareCall[TEntity]) DoAndReturn(f func(context.Context) error) *MockEntityStorePrepareCall[TEntity] {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// PutEntity mocks base method.
func (m *MockEntityStore[TEntity]) PutEntity(ctx context.Context, id uint64, entity TEntity) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutEntity", ctx, id, entity)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutEntity indicates an expected call of PutEntity.
func (mr *MockEntityStoreMockRecorder[TEntity]) PutEntity(ctx, id, entity any) *MockEntityStorePutEntityCall[TEntity] {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutEntity", reflect.TypeOf((*MockEntityStore[TEntity])(nil).PutEntity), ctx, id, entity)
	return &MockEntityStorePutEntityCall[TEntity]{Call: call}
}

// MockEntityStorePutEntityCall wrap *gomock.Call
type MockEntityStorePutEntityCall[TEntity Entity] struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockEntityStorePutEntityCall[TEntity]) Return(arg0 error) *MockEntityStorePutEntityCall[TEntity] {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockEntityStorePutEntityCall[TEntity]) Do(f func(context.Context, uint64, TEntity) error) *MockEntityStorePutEntityCall[TEntity] {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockEntityStorePutEntityCall[TEntity]) DoAndReturn(f func(context.Context, uint64, TEntity) error) *MockEntityStorePutEntityCall[TEntity] {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// RangeFromBlockNum mocks base method.
func (m *MockEntityStore[TEntity]) RangeFromBlockNum(ctx context.Context, startBlockNum uint64) ([]TEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RangeFromBlockNum", ctx, startBlockNum)
	ret0, _ := ret[0].([]TEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RangeFromBlockNum indicates an expected call of RangeFromBlockNum.
func (mr *MockEntityStoreMockRecorder[TEntity]) RangeFromBlockNum(ctx, startBlockNum any) *MockEntityStoreRangeFromBlockNumCall[TEntity] {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RangeFromBlockNum", reflect.TypeOf((*MockEntityStore[TEntity])(nil).RangeFromBlockNum), ctx, startBlockNum)
	return &MockEntityStoreRangeFromBlockNumCall[TEntity]{Call: call}
}

// MockEntityStoreRangeFromBlockNumCall wrap *gomock.Call
type MockEntityStoreRangeFromBlockNumCall[TEntity Entity] struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockEntityStoreRangeFromBlockNumCall[TEntity]) Return(arg0 []TEntity, arg1 error) *MockEntityStoreRangeFromBlockNumCall[TEntity] {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockEntityStoreRangeFromBlockNumCall[TEntity]) Do(f func(context.Context, uint64) ([]TEntity, error)) *MockEntityStoreRangeFromBlockNumCall[TEntity] {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockEntityStoreRangeFromBlockNumCall[TEntity]) DoAndReturn(f func(context.Context, uint64) ([]TEntity, error)) *MockEntityStoreRangeFromBlockNumCall[TEntity] {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SnapType mocks base method.
func (m *MockEntityStore[TEntity]) SnapType() snaptype.Type {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SnapType")
	ret0, _ := ret[0].(snaptype.Type)
	return ret0
}

// SnapType indicates an expected call of SnapType.
func (mr *MockEntityStoreMockRecorder[TEntity]) SnapType() *MockEntityStoreSnapTypeCall[TEntity] {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SnapType", reflect.TypeOf((*MockEntityStore[TEntity])(nil).SnapType))
	return &MockEntityStoreSnapTypeCall[TEntity]{Call: call}
}

// MockEntityStoreSnapTypeCall wrap *gomock.Call
type MockEntityStoreSnapTypeCall[TEntity Entity] struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockEntityStoreSnapTypeCall[TEntity]) Return(arg0 snaptype.Type) *MockEntityStoreSnapTypeCall[TEntity] {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockEntityStoreSnapTypeCall[TEntity]) Do(f func() snaptype.Type) *MockEntityStoreSnapTypeCall[TEntity] {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockEntityStoreSnapTypeCall[TEntity]) DoAndReturn(f func() snaptype.Type) *MockEntityStoreSnapTypeCall[TEntity] {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// EntityIdFromBlockNum mocks base method.
func (m *MockEntityStore[TEntity]) EntityIdFromBlockNum(ctx context.Context, blockNum uint64) (uint64, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EntityIdFromBlockNum", ctx, blockNum)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// EntityIdFromBlockNum indicates an expected call of EntityIdFromBlockNum.
func (mr *MockEntityStoreMockRecorder[TEntity]) EntityIdFromBlockNum(ctx any, blockNum any) *MockEntityStoreEntityIdFromBlockNumCall[TEntity] {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EntityIdFromBlockNum", reflect.TypeOf((*MockEntityStore[TEntity])(nil).EntityIdFromBlockNum), ctx, blockNum)
	return &MockEntityStoreEntityIdFromBlockNumCall[TEntity]{Call: call}
}

// MockEntityStoreEntityIdFromBlockNumCall wrap *gomock.Call
type MockEntityStoreEntityIdFromBlockNumCall[TEntity Entity] struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockEntityStoreEntityIdFromBlockNumCall[TEntity]) Return(arg0 uint64, arg1 bool, arg2 error) *MockEntityStoreEntityIdFromBlockNumCall[TEntity] {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockEntityStoreEntityIdFromBlockNumCall[TEntity]) Do(f func(context.Context, uint64) (uint64, bool, error)) *MockEntityStoreEntityIdFromBlockNumCall[TEntity] {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockEntityStoreEntityIdFromBlockNumCall[TEntity]) DoAndReturn(f func(context.Context, uint64) (uint64, bool, error)) *MockEntityStoreEntityIdFromBlockNumCall[TEntity] {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// EntityIdFromBlockNum mocks base method.
func (m *MockEntityStore[TEntity]) DeleteToBlockNum(ctx context.Context, blockNum uint64, limit int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteToBlockNum", ctx, blockNum, limit)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteToBlockNum indicates an expected call of DeleteToBlockNum.
func (mr *MockEntityStoreMockRecorder[TEntity]) DeleteToBlockNum(ctx any, blockNum any, limit any) *MockEntityStoreDeleteToBlockNumCall[TEntity] {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteToBlockNum", reflect.TypeOf((*MockEntityStore[TEntity])(nil).DeleteToBlockNum), ctx, blockNum, limit)
	return &MockEntityStoreDeleteToBlockNumCall[TEntity]{Call: call}
}

// MockEntityStoreDeleteToBlockNumCall wrap *gomock.Call
type MockEntityStoreDeleteToBlockNumCall[TEntity Entity] struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockEntityStoreDeleteToBlockNumCall[TEntity]) Return(arg0 int, arg1 error) *MockEntityStoreDeleteToBlockNumCall[TEntity] {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockEntityStoreDeleteToBlockNumCall[TEntity]) Do(f func(context.Context, uint64, int) (int, error)) *MockEntityStoreDeleteToBlockNumCall[TEntity] {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockEntityStoreDeleteToBlockNumCall[TEntity]) DoAndReturn(f func(context.Context, uint64, int) (int, error)) *MockEntityStoreDeleteToBlockNumCall[TEntity] {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// EntityIdFromBlockNum mocks base method.
func (m *MockEntityStore[TEntity]) DeleteFromBlockNum(ctx context.Context, blockNum uint64) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFromBlockNum", ctx, blockNum)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteFromBlockNum indicates an expected call of DeleteFromBlockNum.
func (mr *MockEntityStoreMockRecorder[TEntity]) DeleteFromBlockNum(ctx any, blockNum any) *MockEntityStoreDeleteFromBlockNumCall[TEntity] {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFromBlockNum", reflect.TypeOf((*MockEntityStore[TEntity])(nil).DeleteFromBlockNum), ctx, blockNum)
	return &MockEntityStoreDeleteFromBlockNumCall[TEntity]{Call: call}
}

// MockEntityStoreDeleteToBlockNumCall wrap *gomock.Call
type MockEntityStoreDeleteFromBlockNumCall[TEntity Entity] struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockEntityStoreDeleteFromBlockNumCall[TEntity]) Return(arg0 int, arg1 error) *MockEntityStoreDeleteFromBlockNumCall[TEntity] {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockEntityStoreDeleteFromBlockNumCall[TEntity]) Do(f func(context.Context, uint64) (int, error)) *MockEntityStoreDeleteFromBlockNumCall[TEntity] {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockEntityStoreDeleteFromBlockNumCall[TEntity]) DoAndReturn(f func(context.Context, uint64) (int, error)) *MockEntityStoreDeleteFromBlockNumCall[TEntity] {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
