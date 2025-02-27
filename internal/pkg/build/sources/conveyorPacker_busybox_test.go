// Copyright (c) 2021 Apptainer a Series of LF Projects LLC
//   For website terms of use, trademark policy, privacy policy and other
//   project policies see https://lfprojects.org/policies
// Copyright (c) 2018, Sylabs Inc. All rights reserved.
// This software is licensed under a 3-clause BSD license. Please consult the
// LICENSE.md file distributed with the sources of this project regarding your
// rights to use or distribute this software.

package sources_test

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/hpcng/singularity/internal/pkg/build/sources"
	"github.com/hpcng/singularity/internal/pkg/test"
	"github.com/hpcng/singularity/internal/pkg/test/tool/require"
	"github.com/hpcng/singularity/pkg/build/types"
	"github.com/hpcng/singularity/pkg/build/types/parser"
)

const busyBoxDef = "../../../../examples/busybox/Singularity"

func TestBusyBoxConveyor(t *testing.T) {
	// 2021-09-07 - Always skip due to frequent download failures
	t.SkipNow()

	if testing.Short() {
		t.SkipNow()
	}

	// TODO - busybox example has arch hard coded
	require.Arch(t, "amd64")

	test.DropPrivilege(t)
	defer test.ResetPrivilege(t)

	defFile, err := os.Open(busyBoxDef)
	if err != nil {
		t.Fatalf("unable to open file %s: %v\n", busyBoxDef, err)
	}
	defer defFile.Close()

	b, err := types.NewBundle(filepath.Join(os.TempDir(), "sbuild-busybox"), os.TempDir())
	if err != nil {
		return
	}

	b.Recipe, err = parser.ParseDefinitionFile(defFile)
	if err != nil {
		t.Fatalf("failed to parse definition file %s: %v\n", busyBoxDef, err)
	}

	c := &sources.BusyBoxConveyor{}

	err = c.Get(context.Background(), b)
	// clean up tmpfs since assembler isn't called
	defer c.CleanUp()
	if err != nil {
		t.Fatalf("failed to Get from %s: %v\n", busyBoxDef, err)
	}
}

func TestBusyBoxPacker(t *testing.T) {
	// 2021-09-22 - Always skip due to frequent download failures
	t.SkipNow()

	// TODO - busybox example has arch hard coded
	require.Arch(t, "amd64")

	test.DropPrivilege(t)
	defer test.ResetPrivilege(t)

	defFile, err := os.Open(busyBoxDef)
	if err != nil {
		t.Fatalf("unable to open file %s: %v\n", busyBoxDef, err)
	}
	defer defFile.Close()

	b, err := types.NewBundle(filepath.Join(os.TempDir(), "sbuild-busybox"), os.TempDir())
	if err != nil {
		return
	}

	b.Recipe, err = parser.ParseDefinitionFile(defFile)
	if err != nil {
		t.Fatalf("failed to parse definition file %s: %v\n", busyBoxDef, err)
	}

	cp := &sources.BusyBoxConveyorPacker{}

	err = cp.Get(context.Background(), b)
	// clean up tmpfs since assembler isn't called
	defer cp.CleanUp()
	if err != nil {
		t.Fatalf("failed to Get from %s: %v\n", busyBoxDef, err)
	}

	_, err = cp.Pack(context.Background())
	if err != nil {
		t.Fatalf("failed to Pack from %s: %v\n", busyBoxDef, err)
	}
}
