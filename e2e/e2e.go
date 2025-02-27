// Copyright (c) 2021 Apptainer a Series of LF Projects LLC
//   For website terms of use, trademark policy, privacy policy and other
//   project policies see https://lfprojects.org/policies
// Copyright (c) 2019, Sylabs Inc. All rights reserved.
// This software is licensed under a 3-clause BSD license. Please consult the
// LICENSE.md file distributed with the sources of this project regarding your
// rights to use or distribute this software.

package e2e

import (
	"testing"
)

// RunE2ETests is the main func to trigger the test suite.
func RunE2ETests(t *testing.T) {
	t.Log("Running E2E tests for Singularity")
	Run(t)
}
