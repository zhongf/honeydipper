// Copyright 2019 Honey Science Corporation
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, you can obtain one at http://mozilla.org/MPL/2.0/.

package driver

// Meta holds the meta information about the driver itself.
type Meta struct {
	Name        string
	Type        string
	Executable  string
	Arguments   []string
	HandlerData map[string]interface{}
}

// DriverStates represents driver states.
const (
	DriverLoading = iota
	DriverReloading
	DriverAlive
	DriverFailed
)
