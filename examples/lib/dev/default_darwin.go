package dev

import (
	"github.com/modular-id/ble"
	"github.com/modular-id/ble/darwin"
)

// DefaultDevice ...
func DefaultDevice(opts ...ble.Option) (d ble.Device, err error) {
	return darwin.NewDevice(opts...)
}
