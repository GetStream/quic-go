// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/GetStream/quic-go (interfaces: PacketHandler)
//
// Generated by this command:
//
//	mockgen -typed -build_flags=-tags=gomock -package quic -self_package github.com/GetStream/quic-go -destination mock_packet_handler_test.go github.com/GetStream/quic-go PacketHandler
//
// Package quic is a generated GoMock package.
package quic

import (
	reflect "reflect"

	qerr "github.com/GetStream/quic-go/internal/qerr"
	gomock "go.uber.org/mock/gomock"
)

// MockPacketHandler is a mock of PacketHandler interface.
type MockPacketHandler struct {
	ctrl     *gomock.Controller
	recorder *MockPacketHandlerMockRecorder
}

// MockPacketHandlerMockRecorder is the mock recorder for MockPacketHandler.
type MockPacketHandlerMockRecorder struct {
	mock *MockPacketHandler
}

// NewMockPacketHandler creates a new mock instance.
func NewMockPacketHandler(ctrl *gomock.Controller) *MockPacketHandler {
	mock := &MockPacketHandler{ctrl: ctrl}
	mock.recorder = &MockPacketHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPacketHandler) EXPECT() *MockPacketHandlerMockRecorder {
	return m.recorder
}

// closeWithTransportError mocks base method.
func (m *MockPacketHandler) closeWithTransportError(arg0 qerr.TransportErrorCode) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "closeWithTransportError", arg0)
}

// closeWithTransportError indicates an expected call of closeWithTransportError.
func (mr *MockPacketHandlerMockRecorder) closeWithTransportError(arg0 any) *PacketHandlercloseWithTransportErrorCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "closeWithTransportError", reflect.TypeOf((*MockPacketHandler)(nil).closeWithTransportError), arg0)
	return &PacketHandlercloseWithTransportErrorCall{Call: call}
}

// PacketHandlercloseWithTransportErrorCall wrap *gomock.Call
type PacketHandlercloseWithTransportErrorCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *PacketHandlercloseWithTransportErrorCall) Return() *PacketHandlercloseWithTransportErrorCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *PacketHandlercloseWithTransportErrorCall) Do(f func(qerr.TransportErrorCode)) *PacketHandlercloseWithTransportErrorCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *PacketHandlercloseWithTransportErrorCall) DoAndReturn(f func(qerr.TransportErrorCode)) *PacketHandlercloseWithTransportErrorCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// destroy mocks base method.
func (m *MockPacketHandler) destroy(arg0 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "destroy", arg0)
}

// destroy indicates an expected call of destroy.
func (mr *MockPacketHandlerMockRecorder) destroy(arg0 any) *PacketHandlerdestroyCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "destroy", reflect.TypeOf((*MockPacketHandler)(nil).destroy), arg0)
	return &PacketHandlerdestroyCall{Call: call}
}

// PacketHandlerdestroyCall wrap *gomock.Call
type PacketHandlerdestroyCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *PacketHandlerdestroyCall) Return() *PacketHandlerdestroyCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *PacketHandlerdestroyCall) Do(f func(error)) *PacketHandlerdestroyCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *PacketHandlerdestroyCall) DoAndReturn(f func(error)) *PacketHandlerdestroyCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// handlePacket mocks base method.
func (m *MockPacketHandler) handlePacket(arg0 receivedPacket) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "handlePacket", arg0)
}

// handlePacket indicates an expected call of handlePacket.
func (mr *MockPacketHandlerMockRecorder) handlePacket(arg0 any) *PacketHandlerhandlePacketCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "handlePacket", reflect.TypeOf((*MockPacketHandler)(nil).handlePacket), arg0)
	return &PacketHandlerhandlePacketCall{Call: call}
}

// PacketHandlerhandlePacketCall wrap *gomock.Call
type PacketHandlerhandlePacketCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *PacketHandlerhandlePacketCall) Return() *PacketHandlerhandlePacketCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *PacketHandlerhandlePacketCall) Do(f func(receivedPacket)) *PacketHandlerhandlePacketCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *PacketHandlerhandlePacketCall) DoAndReturn(f func(receivedPacket)) *PacketHandlerhandlePacketCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
