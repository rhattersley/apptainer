// Copyright (c) 2021 Apptainer a Series of LF Projects LLC
//   For website terms of use, trademark policy, privacy policy and other
//   project policies see https://lfprojects.org/policies
// Copyright (c) 2018-2021, Sylabs Inc. All rights reserved.
// This software is licensed under a 3-clause BSD license. Please consult the
// LICENSE.md file distributed with the sources of this project regarding your
// rights to use or distribute this software.

//go:build !apparmor
// +build !apparmor

package apparmor

import "errors"

// Enabled returns whether AppArmor is enabled.
func Enabled() bool {
	return false
}

// LoadProfile loads the specified AppArmor profile.
func LoadProfile(profile string) error {
	return errors.New("can't load AppArmor profile: not enabled at compilation time")
}
