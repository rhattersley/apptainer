// Copyright (c) 2021 Apptainer a Series of LF Projects LLC
//   For website terms of use, trademark policy, privacy policy and other
//   project policies see https://lfprojects.org/policies
// Copyright (c) 2018-2021, Sylabs Inc. All rights reserved.
// This software is licensed under a 3-clause BSD license. Please consult the
// LICENSE.md file distributed with the sources of this project regarding your
// rights to use or distribute this software.

package oci

import (
	"sync"

	"github.com/hpcng/singularity/internal/pkg/cgroups"
	"github.com/hpcng/singularity/internal/pkg/runtime/engine/config/oci"
	"github.com/hpcng/singularity/pkg/ociruntime"
)

// Name of the engine.
const Name = "oci"

// EngineConfig is the config for the OCI engine.
type EngineConfig struct {
	BundlePath    string          `json:"bundlePath"`
	LogPath       string          `json:"logPath"`
	LogFormat     string          `json:"logFormat"`
	PidFile       string          `json:"pidFile"`
	OciConfig     *oci.Config     `json:"ociConfig"`
	MasterPts     int             `json:"masterPts"`
	SlavePts      int             `json:"slavePts"`
	OutputStreams [2]int          `json:"outputStreams"`
	ErrorStreams  [2]int          `json:"errorStreams"`
	InputStreams  [2]int          `json:"inputStreams"`
	SyncSocket    string          `json:"syncSocket"`
	EmptyProcess  bool            `json:"emptyProcess"`
	Exec          bool            `json:"exec"`
	Cgroups       cgroups.Manager `json:"-"`

	sync.Mutex `json:"-"`
	State      ociruntime.State `json:"state"`
}

// NewConfig returns an oci.EngineConfig.
func NewConfig() *EngineConfig {
	ret := &EngineConfig{
		OciConfig: &oci.Config{},
	}

	return ret
}

// SetBundlePath sets the container bundle path.
func (e *EngineConfig) SetBundlePath(path string) {
	e.BundlePath = path
}

// GetBundlePath returns the container bundle path.
func (e *EngineConfig) GetBundlePath() string {
	return e.BundlePath
}

// SetState sets the container state as defined by OCI state specification.
func (e *EngineConfig) SetState(state *ociruntime.State) {
	e.State = *state
}

// GetState returns the container state as defined by OCI state specification.
func (e *EngineConfig) GetState() *ociruntime.State {
	return &e.State
}

// SetLogPath sets the container log path.
func (e *EngineConfig) SetLogPath(path string) {
	e.LogPath = path
}

// GetLogPath returns the container log path.
func (e *EngineConfig) GetLogPath() string {
	return e.LogPath
}

// SetLogFormat sets the container log format.
func (e *EngineConfig) SetLogFormat(format string) {
	e.LogFormat = format
}

// GetLogFormat returns the container log format.
func (e *EngineConfig) GetLogFormat() string {
	return e.LogFormat
}

// SetPidFile sets the pid file path.
func (e *EngineConfig) SetPidFile(path string) {
	e.PidFile = path
}

// GetPidFile gets the pid file path.
func (e *EngineConfig) GetPidFile() string {
	return e.PidFile
}
