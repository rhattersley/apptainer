// Copyright (c) 2021 Apptainer a Series of LF Projects LLC
//   For website terms of use, trademark policy, privacy policy and other
//   project policies see https://lfprojects.org/policies
// Copyright (c) 2018-2020, Sylabs Inc. All rights reserved.
// This software is licensed under a 3-clause BSD license. Please consult the
// LICENSE.md file distributed with the sources of this project regarding your
// rights to use or distribute this software.

package security

import (
	"fmt"
	"strings"

	"github.com/hpcng/singularity/internal/pkg/security/apparmor"
	"github.com/hpcng/singularity/internal/pkg/security/seccomp"
	"github.com/hpcng/singularity/internal/pkg/security/selinux"
	"github.com/hpcng/singularity/pkg/sylog"
	specs "github.com/opencontainers/runtime-spec/specs-go"
)

// Configure applies security related configuration to current process
func Configure(config *specs.Spec) error {
	if config.Process != nil {
		if config.Process.SelinuxLabel != "" && config.Process.ApparmorProfile != "" {
			return fmt.Errorf("you can't specify both an apparmor profile and a selinux label")
		}
		if config.Process.SelinuxLabel != "" {
			if selinux.Enabled() {
				if err := selinux.SetExecLabel(config.Process.SelinuxLabel); err != nil {
					return err
				}
			} else {
				sylog.Warningf("selinux is not enabled or supported on this system")
			}
		} else if config.Process.ApparmorProfile != "" {
			if apparmor.Enabled() {
				if err := apparmor.LoadProfile(config.Process.ApparmorProfile); err != nil {
					return err
				}
			} else {
				sylog.Warningf("apparmor is not enabled or supported on this system")
			}
		}
	}
	if config.Linux != nil && config.Linux.Seccomp != nil {
		if seccomp.Enabled() {
			if err := seccomp.LoadSeccompConfig(config.Linux.Seccomp, config.Process.NoNewPrivileges, 1); err != nil {
				return err
			}
		} else {
			sylog.Warningf("seccomp requested but not enabled, seccomp library is missing or too old")
		}
	}
	return nil
}

// GetParam iterates over security argument and returns parameters
// for the security feature
func GetParam(security []string, feature string) string {
	for _, param := range security {
		splitted := strings.SplitN(param, ":", 2)
		if splitted[0] == feature {
			if len(splitted) != 2 {
				sylog.Warningf("bad format for parameter %s (format is <security>:<arg>)", param)
			}
			return splitted[1]
		}
	}
	return ""
}
