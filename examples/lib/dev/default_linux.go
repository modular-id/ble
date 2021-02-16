package dev

import (
	"github.com/modular-id/ble"
	"github.com/modular-id/ble/linux"
)

// DefaultDevice ...
func DefaultDevice(opts ...ble.Option) (d ble.Device, err error) {
	return linux.NewDevice(opts...)
}
