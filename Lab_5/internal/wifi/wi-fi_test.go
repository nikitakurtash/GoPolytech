package wifi_test

import (
	"errors"
	myWiFi "example_mock/internal/wifi"
	"github.com/mdlayher/wifi"
	"github.com/stretchr/testify/require"
	"net"
	"testing"
)

var (
	ErrInterfaceConnection = errors.New("невозможно подключиться к интерфейсу Wi-Fi")
	ErrInternal            = errors.New("внутренняя ошибка")
)

var testTable = []struct {
	names       []string
	addrs       []string
	errExpected error
}{
	{names: []string{"WiFi1", "WiFi2"}, addrs: []string{"00:11:22:33:44:55", "66:77:88:99:AA:BB"}, errExpected: nil},
	{names: []string{"WiFi3"}, addrs: []string{"AA:BB:CC:DD:EE:FF"}, errExpected: nil},
	{names: nil, addrs: nil, errExpected: ErrInterfaceConnection},
	{names: nil, addrs: nil, errExpected: ErrInternal},
}

func mockIfaces(addrs []string) []*wifi.Interface {
	var ifaces []*wifi.Interface
	for _, addr := range addrs {
		mac, _ := net.ParseMAC(addr)
		ifaces = append(ifaces, &wifi.Interface{HardwareAddr: mac})
	}
	return ifaces
}

func parseMACs(addrs []string) []net.HardwareAddr {
	var macs []net.HardwareAddr
	for _, addr := range addrs {
		mac, _ := net.ParseMAC(addr)
		macs = append(macs, mac)
	}
	return macs
}

func TestGetNames(t *testing.T) {
	for i, row := range testTable {
		mockWifi := NewWiFi(t)
		wifiService := myWiFi.Service{WiFi: mockWifi}

		var mockIfaces []*wifi.Interface
		for _, name := range row.names {
			mockIfaces = append(mockIfaces, &wifi.Interface{Name: name})
		}

		mockWifi.On("Interfaces").Return(mockIfaces, row.errExpected)

		actualNames, err := wifiService.GetNames()

		if row.errExpected != nil {
			require.ErrorIs(t, err, row.errExpected, "row: %d, expected error: %w, actual error: %w", i, row.errExpected, err)
			mockWifi.AssertExpectations(t)
			continue
		}

		require.NoError(t, err, "row: %d, error must be nil", i)
		require.Equal(t, row.names, actualNames, "row: %d, expected names: %s, actual names: %s", i, row.names, actualNames)
		mockWifi.AssertExpectations(t)
	}
}

func TestGetAddresses(t *testing.T) {
	for i, row := range testTable {
		mockWifi := NewWiFi(t)
		wifiService := myWiFi.Service{WiFi: mockWifi}

		mockIfaces := mockIfaces(row.addrs)
		mockWifi.On("Interfaces").Return(mockIfaces, row.errExpected)

		actualAddrs, err := wifiService.GetAddresses()

		if row.errExpected != nil {
			require.ErrorIs(t, err, row.errExpected, "row: %d, expected error: %w, actual error: %w", i, row.errExpected, err)
			mockWifi.AssertExpectations(t)
			continue
		}

		require.NoError(t, err, "row: %d, error must be nil", i)
		require.Equal(t, parseMACs(row.addrs), actualAddrs, "row: %d, expected addrs: %s, actual addrs: %s", i, parseMACs(row.addrs), actualAddrs)
		mockWifi.AssertExpectations(t)
	}
}
