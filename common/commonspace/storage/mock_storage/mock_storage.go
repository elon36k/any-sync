// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/anytypeio/go-anytype-infrastructure-experiments/common/commonspace/storage (interfaces: SpaceStorageProvider,SpaceStorage)

// Package mock_storage is a generated GoMock package.
package mock_storage

import (
	reflect "reflect"

	app "github.com/anytypeio/go-anytype-infrastructure-experiments/app"
	spacesyncproto "github.com/anytypeio/go-anytype-infrastructure-experiments/common/commonspace/spacesyncproto"
	storage "github.com/anytypeio/go-anytype-infrastructure-experiments/common/commonspace/storage"
	storage0 "github.com/anytypeio/go-anytype-infrastructure-experiments/pkg/acl/storage"
	gomock "github.com/golang/mock/gomock"
)

// MockSpaceStorageProvider is a mock of SpaceStorageProvider interface.
type MockSpaceStorageProvider struct {
	ctrl     *gomock.Controller
	recorder *MockSpaceStorageProviderMockRecorder
}

// MockSpaceStorageProviderMockRecorder is the mock recorder for MockSpaceStorageProvider.
type MockSpaceStorageProviderMockRecorder struct {
	mock *MockSpaceStorageProvider
}

// NewMockSpaceStorageProvider creates a new mock instance.
func NewMockSpaceStorageProvider(ctrl *gomock.Controller) *MockSpaceStorageProvider {
	mock := &MockSpaceStorageProvider{ctrl: ctrl}
	mock.recorder = &MockSpaceStorageProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpaceStorageProvider) EXPECT() *MockSpaceStorageProviderMockRecorder {
	return m.recorder
}

// CreateSpaceStorage mocks base method.
func (m *MockSpaceStorageProvider) CreateSpaceStorage(arg0 storage.SpaceStorageCreatePayload) (storage.SpaceStorage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSpaceStorage", arg0)
	ret0, _ := ret[0].(storage.SpaceStorage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSpaceStorage indicates an expected call of CreateSpaceStorage.
func (mr *MockSpaceStorageProviderMockRecorder) CreateSpaceStorage(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSpaceStorage", reflect.TypeOf((*MockSpaceStorageProvider)(nil).CreateSpaceStorage), arg0)
}

// Init mocks base method.
func (m *MockSpaceStorageProvider) Init(arg0 *app.App) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init.
func (mr *MockSpaceStorageProviderMockRecorder) Init(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockSpaceStorageProvider)(nil).Init), arg0)
}

// Name mocks base method.
func (m *MockSpaceStorageProvider) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockSpaceStorageProviderMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockSpaceStorageProvider)(nil).Name))
}

// SpaceStorage mocks base method.
func (m *MockSpaceStorageProvider) SpaceStorage(arg0 string) (storage.SpaceStorage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SpaceStorage", arg0)
	ret0, _ := ret[0].(storage.SpaceStorage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SpaceStorage indicates an expected call of SpaceStorage.
func (mr *MockSpaceStorageProviderMockRecorder) SpaceStorage(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SpaceStorage", reflect.TypeOf((*MockSpaceStorageProvider)(nil).SpaceStorage), arg0)
}

// MockSpaceStorage is a mock of SpaceStorage interface.
type MockSpaceStorage struct {
	ctrl     *gomock.Controller
	recorder *MockSpaceStorageMockRecorder
}

// MockSpaceStorageMockRecorder is the mock recorder for MockSpaceStorage.
type MockSpaceStorageMockRecorder struct {
	mock *MockSpaceStorage
}

// NewMockSpaceStorage creates a new mock instance.
func NewMockSpaceStorage(ctrl *gomock.Controller) *MockSpaceStorage {
	mock := &MockSpaceStorage{ctrl: ctrl}
	mock.recorder = &MockSpaceStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpaceStorage) EXPECT() *MockSpaceStorageMockRecorder {
	return m.recorder
}

// ACLStorage mocks base method.
func (m *MockSpaceStorage) ACLStorage() (storage0.ListStorage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ACLStorage")
	ret0, _ := ret[0].(storage0.ListStorage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ACLStorage indicates an expected call of ACLStorage.
func (mr *MockSpaceStorageMockRecorder) ACLStorage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ACLStorage", reflect.TypeOf((*MockSpaceStorage)(nil).ACLStorage))
}

// CreateTreeStorage mocks base method.
func (m *MockSpaceStorage) CreateTreeStorage(arg0 storage0.TreeStorageCreatePayload) (storage0.TreeStorage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTreeStorage", arg0)
	ret0, _ := ret[0].(storage0.TreeStorage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTreeStorage indicates an expected call of CreateTreeStorage.
func (mr *MockSpaceStorageMockRecorder) CreateTreeStorage(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTreeStorage", reflect.TypeOf((*MockSpaceStorage)(nil).CreateTreeStorage), arg0)
}

// SpaceHeader mocks base method.
func (m *MockSpaceStorage) SpaceHeader() (*spacesyncproto.RawSpaceHeaderWithId, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SpaceHeader")
	ret0, _ := ret[0].(*spacesyncproto.RawSpaceHeaderWithId)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SpaceHeader indicates an expected call of SpaceHeader.
func (mr *MockSpaceStorageMockRecorder) SpaceHeader() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SpaceHeader", reflect.TypeOf((*MockSpaceStorage)(nil).SpaceHeader))
}

// StoredIds mocks base method.
func (m *MockSpaceStorage) StoredIds() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoredIds")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StoredIds indicates an expected call of StoredIds.
func (mr *MockSpaceStorageMockRecorder) StoredIds() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoredIds", reflect.TypeOf((*MockSpaceStorage)(nil).StoredIds))
}

// TreeStorage mocks base method.
func (m *MockSpaceStorage) TreeStorage(arg0 string) (storage0.TreeStorage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TreeStorage", arg0)
	ret0, _ := ret[0].(storage0.TreeStorage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TreeStorage indicates an expected call of TreeStorage.
func (mr *MockSpaceStorageMockRecorder) TreeStorage(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TreeStorage", reflect.TypeOf((*MockSpaceStorage)(nil).TreeStorage), arg0)
}
