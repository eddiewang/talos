// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Package bootloader provides bootloader implementation.
package bootloader

import (
	"os"

	"github.com/siderolabs/talos/internal/app/machined/pkg/runtime/v1alpha1/bootloader/grub"
)

// Bootloader describes a bootloader.
type Bootloader interface {
	// Install installs the bootloader
	Install(bootDisk, arch, cmdline string) error
	// Revert reverts the bootloader entry to the previous state.
	Revert() error
	// PreviousLabel returns the previous bootloader label.
	PreviousLabel() string
}

// Probe checks if any supported bootloaders are installed.
//
// If 'disk' is empty, it will probe all disks.
// Returns nil if it cannot detect any supported bootloader.
func Probe(disk string) (Bootloader, error) {
	grubBootloader, err := grub.Probe(disk)
	if err != nil {
		return nil, err
	}

	if grubBootloader == nil {
		return nil, os.ErrNotExist
	}

	return grubBootloader, nil
}

// New returns a new bootloader.
func New() (Bootloader, error) {
	return grub.NewConfig(), nil
}
