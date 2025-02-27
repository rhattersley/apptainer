// Copyright (c) 2021 Apptainer a Series of LF Projects LLC
//   For website terms of use, trademark policy, privacy policy and other
//   project policies see https://lfprojects.org/policies
// Copyright (c) 2019-2020, Sylabs Inc. All rights reserved.
// This software is licensed under a 3-clause BSD license. Please consult the
// LICENSE.md file distributed with the sources of this project regarding your
// rights to use or distribute this software.

package cmdline

import (
	"fmt"
	"strings"

	"github.com/hpcng/singularity/pkg/sylog"
	"github.com/spf13/pflag"
)

func setValue(flag *pflag.Flag, value string) error {
	if err := flag.Value.Set(value); err != nil {
		return fmt.Errorf("unable to set flag %s to value %s: %s", flag.Name, value, err)
	}
	flag.Changed = true
	sylog.Debugf("Updated flag '%s' value to: %s", flag.Name, flag.Value)
	return nil
}

// EnvAppendValue combines command line and environment var into a single argument
func EnvAppendValue(flag *pflag.Flag, value string) error {
	v := strings.TrimSpace(value)
	if v == "" {
		return nil
	}
	return setValue(flag, value)
}

// EnvSetValue set flag value if CLI option/argument is unset and env var is set
func EnvSetValue(flag *pflag.Flag, value string) error {
	v := strings.TrimSpace(value)
	if flag.Changed || v == "" {
		return nil
	}
	// if flag is a string slice, sanitize slice by
	// trimming potential spaces in environment variable
	// value (eg: FOO="val1 , val2,val3")
	if flag.Value.Type() == "stringSlice" {
		vals := strings.Split(v, ",")
		for i, e := range vals {
			vals[i] = strings.TrimSpace(e)
		}
		v = strings.Join(vals, ",")
	}
	return setValue(flag, v)
}

// EnvHandler defines an environment handler type to set flag's values
// from environment variables
type EnvHandler func(*pflag.Flag, string) error
