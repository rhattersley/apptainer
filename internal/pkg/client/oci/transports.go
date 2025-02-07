// Copyright (c) 2021 Apptainer a Series of LF Projects LLC
//   For website terms of use, trademark policy, privacy policy and other
//   project policies see https://lfprojects.org/policies
// Copyright (c) 2018, Sylabs Inc. All rights reserved.
// This software is licensed under a 3-clause BSD license. Please consult the
// LICENSE.md file distributed with the sources of this project regarding your
// rights to use or distribute this software.

package oci

import (
	"github.com/containers/image/v5/transports"
)

// IsSupported returns whether or not the transport given is supported. To fit within a switch/case
// statement, this function will return transport if it is supported
func IsSupported(transport string) string {
	for _, t := range transports.ListNames() {
		if transport == t {
			return transport
		}
	}

	return ""
}
