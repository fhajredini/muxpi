// Generated by go-rpcgen. Do not modify.

package stm

import (
	"net/rpc"
	"time"
)

// InterfaceService is generated service for Interface interface.
type InterfaceService struct {
	impl Interface
}

// NewInterfaceService creates a new InterfaceService instance.
func NewInterfaceService(impl Interface) *InterfaceService {
	return &InterfaceService{impl}
}

// RegisterInterfaceService registers impl in server.
func RegisterInterfaceService(server *rpc.Server, impl Interface) error {
	return server.RegisterName("Interface", NewInterfaceService(impl))
}

// InterfacePowerTickRequest is a helper structure for PowerTick method.
type InterfacePowerTickRequest struct {
	D time.Duration
}

// InterfacePowerTickResponse is a helper structure for PowerTick method.
type InterfacePowerTickResponse struct {
}

// PowerTick is RPC implementation of PowerTick calling it.
func (s *InterfaceService) PowerTick(request *InterfacePowerTickRequest, response *InterfacePowerTickResponse) (err error) {
	err = s.impl.PowerTick(request.D)
	return
}

// InterfaceDUTRequest is a helper structure for DUT method.
type InterfaceDUTRequest struct {
}

// InterfaceDUTResponse is a helper structure for DUT method.
type InterfaceDUTResponse struct {
}

// DUT is RPC implementation of DUT calling it.
func (s *InterfaceService) DUT(request *InterfaceDUTRequest, response *InterfaceDUTResponse) (err error) {
	err = s.impl.DUT()
	return
}

// InterfaceTSRequest is a helper structure for TS method.
type InterfaceTSRequest struct {
}

// InterfaceTSResponse is a helper structure for TS method.
type InterfaceTSResponse struct {
}

// TS is RPC implementation of TS calling it.
func (s *InterfaceService) TS(request *InterfaceTSRequest, response *InterfaceTSResponse) (err error) {
	err = s.impl.TS()
	return
}

// InterfaceGetCurrentRequest is a helper structure for GetCurrent method.
type InterfaceGetCurrentRequest struct {
}

// InterfaceGetCurrentResponse is a helper structure for GetCurrent method.
type InterfaceGetCurrentResponse struct {
	Value int
}

// GetCurrent is RPC implementation of GetCurrent calling it.
func (s *InterfaceService) GetCurrent(request *InterfaceGetCurrentRequest, response *InterfaceGetCurrentResponse) (err error) {
	response.Value, err = s.impl.GetCurrent()
	return
}

// InterfaceSetLEDRequest is a helper structure for SetLED method.
type InterfaceSetLEDRequest struct {
	Led     LED
	R, G, B uint8
}

// InterfaceSetLEDResponse is a helper structure for SetLED method.
type InterfaceSetLEDResponse struct {
}

// SetLED is RPC implementation of SetLED calling it.
func (s *InterfaceService) SetLED(request *InterfaceSetLEDRequest, response *InterfaceSetLEDResponse) (err error) {
	err = s.impl.SetLED(request.Led, request.R, request.G, request.B)
	return
}

// InterfaceClearDisplayRequest is a helper structure for ClearDisplay method.
type InterfaceClearDisplayRequest struct {
}

// InterfaceClearDisplayResponse is a helper structure for ClearDisplay method.
type InterfaceClearDisplayResponse struct {
}

// ClearDisplay is RPC implementation of ClearDisplay calling it.
func (s *InterfaceService) ClearDisplay(request *InterfaceClearDisplayRequest, response *InterfaceClearDisplayResponse) (err error) {
	err = s.impl.ClearDisplay()
	return
}

// InterfacePrintTextRequest is a helper structure for PrintText method.
type InterfacePrintTextRequest struct {
	X, Y  uint
	Color Color
	Text  string
}

// InterfacePrintTextResponse is a helper structure for PrintText method.
type InterfacePrintTextResponse struct {
}

// PrintText is RPC implementation of PrintText calling it.
func (s *InterfaceService) PrintText(request *InterfacePrintTextRequest, response *InterfacePrintTextResponse) (err error) {
	err = s.impl.PrintText(request.X, request.Y, request.Color, request.Text)
	return
}

// InterfaceClient is generated client for Interface interface.
type InterfaceClient struct {
	client *rpc.Client
}

// DialInterfaceClient connects to addr and creates a new InterfaceClient instance.
func DialInterfaceClient(addr string) (*InterfaceClient, error) {
	client, err := rpc.Dial("tcp", addr)
	return &InterfaceClient{client}, err
}

// NewInterfaceClient creates a new InterfaceClient instance.
func NewInterfaceClient(client *rpc.Client) *InterfaceClient {
	return &InterfaceClient{client}
}

// Close terminates the connection.
func (_c *InterfaceClient) Close() error {
	return _c.client.Close()
}

// PowerTick is part of implementation of Interface calling corresponding method on RPC server.
func (_c *InterfaceClient) PowerTick(d time.Duration) (err error) {
	_request := &InterfacePowerTickRequest{d}
	_response := &InterfacePowerTickResponse{}
	err = _c.client.Call("Interface.PowerTick", _request, _response)
	return err
}

// DUT is part of implementation of Interface calling corresponding method on RPC server.
func (_c *InterfaceClient) DUT() (err error) {
	_request := &InterfaceDUTRequest{}
	_response := &InterfaceDUTResponse{}
	err = _c.client.Call("Interface.DUT", _request, _response)
	return err
}

// TS is part of implementation of Interface calling corresponding method on RPC server.
func (_c *InterfaceClient) TS() (err error) {
	_request := &InterfaceTSRequest{}
	_response := &InterfaceTSResponse{}
	err = _c.client.Call("Interface.TS", _request, _response)
	return err
}

// GetCurrent is part of implementation of Interface calling corresponding method on RPC server.
func (_c *InterfaceClient) GetCurrent() (value int, err error) {
	_request := &InterfaceGetCurrentRequest{}
	_response := &InterfaceGetCurrentResponse{}
	err = _c.client.Call("Interface.GetCurrent", _request, _response)
	return _response.Value, err
}

// SetLED is part of implementation of Interface calling corresponding method on RPC server.
func (_c *InterfaceClient) SetLED(led LED, r, g, b uint8) (err error) {
	_request := &InterfaceSetLEDRequest{led, r, g, b}
	_response := &InterfaceSetLEDResponse{}
	err = _c.client.Call("Interface.SetLED", _request, _response)
	return err
}

// ClearDisplay is part of implementation of Interface calling corresponding method on RPC server.
func (_c *InterfaceClient) ClearDisplay() (err error) {
	_request := &InterfaceClearDisplayRequest{}
	_response := &InterfaceClearDisplayResponse{}
	err = _c.client.Call("Interface.ClearDisplay", _request, _response)
	return err
}

// PrintText is part of implementation of Interface calling corresponding method on RPC server.
func (_c *InterfaceClient) PrintText(x, y uint, color Color, text string) (err error) {
	_request := &InterfacePrintTextRequest{x, y, color, text}
	_response := &InterfacePrintTextResponse{}
	err = _c.client.Call("Interface.PrintText", _request, _response)
	return err
}