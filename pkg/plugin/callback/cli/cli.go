// Copyright (c) 2021 Apptainer a Series of LF Projects LLC
//   For website terms of use, trademark policy, privacy policy and other
//   project policies see https://lfprojects.org/policies
// Copyright (c) 2020, Sylabs Inc. All rights reserved.
// This software is licensed under a 3-clause BSD license. Please consult the
// LICENSE.md file distributed with the URIs of this project regarding your
// rights to use or distribute this software.

package cli

import (
	"github.com/hpcng/singularity/pkg/cmdline"
	"github.com/hpcng/singularity/pkg/runtime/engine/config"
)

// Command callback allows to add/modify commands and/or flags.
// This callback is called in cmd/internal/cli/singularity.go and
// allows plugins to inject/modify commands and/or flags to existing
// singularity commands.
type Command func(*cmdline.CommandManager)

// SingularityEngineConfig callback allows to manipulate Singularity
// runtime engine configuration.
// This callback is called in cmd/internal/cli/actions_linux.go and
// allows plugins to modify/alter runtime engine configuration. This
// is the place to inject custom binds.
type SingularityEngineConfig func(*config.Common)
