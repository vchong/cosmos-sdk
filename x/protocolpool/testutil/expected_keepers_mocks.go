// Code generated by MockGen. DO NOT EDIT.
// Source: x/protocolpool/types/expected_keepers.go
//
// Generated by this command:
//
//	mockgen -source=x/protocolpool/types/expected_keepers.go -package testutil -destination x/protocolpool/testutil/expected_keepers_mocks.go
//

// Package testutil is a generated GoMock package.
package testutil

import (
	context "context"
	reflect "reflect"

	address "cosmossdk.io/core/address"
	types "github.com/cosmos/cosmos-sdk/types"
	gomock "go.uber.org/mock/gomock"
)

// MockAccountKeeper is a mock of AccountKeeper interface.
type MockAccountKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockAccountKeeperMockRecorder
	isgomock struct{}
}

// MockAccountKeeperMockRecorder is the mock recorder for MockAccountKeeper.
type MockAccountKeeperMockRecorder struct {
	mock *MockAccountKeeper
}

// NewMockAccountKeeper creates a new mock instance.
func NewMockAccountKeeper(ctrl *gomock.Controller) *MockAccountKeeper {
	mock := &MockAccountKeeper{ctrl: ctrl}
	mock.recorder = &MockAccountKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountKeeper) EXPECT() *MockAccountKeeperMockRecorder {
	return m.recorder
}

// AddressCodec mocks base method.
func (m *MockAccountKeeper) AddressCodec() address.Codec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddressCodec")
	ret0, _ := ret[0].(address.Codec)
	return ret0
}

// AddressCodec indicates an expected call of AddressCodec.
func (mr *MockAccountKeeperMockRecorder) AddressCodec() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddressCodec", reflect.TypeOf((*MockAccountKeeper)(nil).AddressCodec))
}

// GetAccount mocks base method.
func (m *MockAccountKeeper) GetAccount(ctx context.Context, addr types.AccAddress) types.AccountI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", ctx, addr)
	ret0, _ := ret[0].(types.AccountI)
	return ret0
}

// GetAccount indicates an expected call of GetAccount.
func (mr *MockAccountKeeperMockRecorder) GetAccount(ctx, addr any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockAccountKeeper)(nil).GetAccount), ctx, addr)
}

// GetModuleAccount mocks base method.
func (m *MockAccountKeeper) GetModuleAccount(ctx context.Context, name string) types.ModuleAccountI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetModuleAccount", ctx, name)
	ret0, _ := ret[0].(types.ModuleAccountI)
	return ret0
}

// GetModuleAccount indicates an expected call of GetModuleAccount.
func (mr *MockAccountKeeperMockRecorder) GetModuleAccount(ctx, name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModuleAccount", reflect.TypeOf((*MockAccountKeeper)(nil).GetModuleAccount), ctx, name)
}

// GetModuleAddress mocks base method.
func (m *MockAccountKeeper) GetModuleAddress(name string) types.AccAddress {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetModuleAddress", name)
	ret0, _ := ret[0].(types.AccAddress)
	return ret0
}

// GetModuleAddress indicates an expected call of GetModuleAddress.
func (mr *MockAccountKeeperMockRecorder) GetModuleAddress(name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModuleAddress", reflect.TypeOf((*MockAccountKeeper)(nil).GetModuleAddress), name)
}

// SetModuleAccount mocks base method.
func (m *MockAccountKeeper) SetModuleAccount(ctx context.Context, macc types.ModuleAccountI) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetModuleAccount", ctx, macc)
}

// SetModuleAccount indicates an expected call of SetModuleAccount.
func (mr *MockAccountKeeperMockRecorder) SetModuleAccount(ctx, macc any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetModuleAccount", reflect.TypeOf((*MockAccountKeeper)(nil).SetModuleAccount), ctx, macc)
}

// MockBankKeeper is a mock of BankKeeper interface.
type MockBankKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockBankKeeperMockRecorder
	isgomock struct{}
}

// MockBankKeeperMockRecorder is the mock recorder for MockBankKeeper.
type MockBankKeeperMockRecorder struct {
	mock *MockBankKeeper
}

// NewMockBankKeeper creates a new mock instance.
func NewMockBankKeeper(ctrl *gomock.Controller) *MockBankKeeper {
	mock := &MockBankKeeper{ctrl: ctrl}
	mock.recorder = &MockBankKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBankKeeper) EXPECT() *MockBankKeeperMockRecorder {
	return m.recorder
}

// BlockedAddr mocks base method.
func (m *MockBankKeeper) BlockedAddr(addr types.AccAddress) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockedAddr", addr)
	ret0, _ := ret[0].(bool)
	return ret0
}

// BlockedAddr indicates an expected call of BlockedAddr.
func (mr *MockBankKeeperMockRecorder) BlockedAddr(addr any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockedAddr", reflect.TypeOf((*MockBankKeeper)(nil).BlockedAddr), addr)
}

// GetAllBalances mocks base method.
func (m *MockBankKeeper) GetAllBalances(ctx context.Context, addr types.AccAddress) types.Coins {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllBalances", ctx, addr)
	ret0, _ := ret[0].(types.Coins)
	return ret0
}

// GetAllBalances indicates an expected call of GetAllBalances.
func (mr *MockBankKeeperMockRecorder) GetAllBalances(ctx, addr any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllBalances", reflect.TypeOf((*MockBankKeeper)(nil).GetAllBalances), ctx, addr)
}

// GetBalance mocks base method.
func (m *MockBankKeeper) GetBalance(ctx context.Context, addr types.AccAddress, denom string) types.Coin {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBalance", ctx, addr, denom)
	ret0, _ := ret[0].(types.Coin)
	return ret0
}

// GetBalance indicates an expected call of GetBalance.
func (mr *MockBankKeeperMockRecorder) GetBalance(ctx, addr, denom any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBalance", reflect.TypeOf((*MockBankKeeper)(nil).GetBalance), ctx, addr, denom)
}

// SendCoinsFromAccountToModule mocks base method.
func (m *MockBankKeeper) SendCoinsFromAccountToModule(ctx context.Context, senderAddr types.AccAddress, recipientModule string, amt types.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCoinsFromAccountToModule", ctx, senderAddr, recipientModule, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCoinsFromAccountToModule indicates an expected call of SendCoinsFromAccountToModule.
func (mr *MockBankKeeperMockRecorder) SendCoinsFromAccountToModule(ctx, senderAddr, recipientModule, amt any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCoinsFromAccountToModule", reflect.TypeOf((*MockBankKeeper)(nil).SendCoinsFromAccountToModule), ctx, senderAddr, recipientModule, amt)
}

// SendCoinsFromModuleToAccount mocks base method.
func (m *MockBankKeeper) SendCoinsFromModuleToAccount(ctx context.Context, senderModule string, recipientAddr types.AccAddress, amt types.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCoinsFromModuleToAccount", ctx, senderModule, recipientAddr, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCoinsFromModuleToAccount indicates an expected call of SendCoinsFromModuleToAccount.
func (mr *MockBankKeeperMockRecorder) SendCoinsFromModuleToAccount(ctx, senderModule, recipientAddr, amt any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCoinsFromModuleToAccount", reflect.TypeOf((*MockBankKeeper)(nil).SendCoinsFromModuleToAccount), ctx, senderModule, recipientAddr, amt)
}

// SendCoinsFromModuleToModule mocks base method.
func (m *MockBankKeeper) SendCoinsFromModuleToModule(ctx context.Context, senderModule, recipientModule string, amt types.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCoinsFromModuleToModule", ctx, senderModule, recipientModule, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCoinsFromModuleToModule indicates an expected call of SendCoinsFromModuleToModule.
func (mr *MockBankKeeperMockRecorder) SendCoinsFromModuleToModule(ctx, senderModule, recipientModule, amt any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCoinsFromModuleToModule", reflect.TypeOf((*MockBankKeeper)(nil).SendCoinsFromModuleToModule), ctx, senderModule, recipientModule, amt)
}

// SpendableCoins mocks base method.
func (m *MockBankKeeper) SpendableCoins(ctx context.Context, addr types.AccAddress) types.Coins {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SpendableCoins", ctx, addr)
	ret0, _ := ret[0].(types.Coins)
	return ret0
}

// SpendableCoins indicates an expected call of SpendableCoins.
func (mr *MockBankKeeperMockRecorder) SpendableCoins(ctx, addr any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SpendableCoins", reflect.TypeOf((*MockBankKeeper)(nil).SpendableCoins), ctx, addr)
}
