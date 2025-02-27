// Copyright (c) 2021 Apptainer a Series of LF Projects LLC
//   For website terms of use, trademark policy, privacy policy and other
//   project policies see https://lfprojects.org/policies
// Copyright (c) 2018-2020, Sylabs Inc. All rights reserved.
// This software is licensed under a 3-clause BSD license. Please consult the
// LICENSE.md file distributed with the sources of this project regarding your
// rights to use or distribute this software.

package layout

import (
	"fmt"
	"path/filepath"
	"syscall"

	"github.com/hpcng/singularity/internal/pkg/util/fs/mount"
	"github.com/hpcng/singularity/pkg/sylog"
)

const (
	rootFsDir = "/rootfs"
	finalDir  = "/final"
)

// Session directory layout manager
type Session struct {
	*Manager
	Layer layer
}

// Layer describes a layer interface added on top of session layout
type layer interface {
	Add(*Session, *mount.System) error
	Dir() string
}

// NewSession creates and returns a session directory layout manager
func NewSession(path string, fstype string, size int, system *mount.System, layer layer) (*Session, error) {
	manager := &Manager{VFS: DefaultVFS}
	session := &Session{Manager: manager}

	if err := manager.SetRootPath(path); err != nil {
		return nil, err
	}
	if err := manager.AddDir(rootFsDir); err != nil {
		return nil, err
	}
	if err := manager.AddDir(finalDir); err != nil {
		return nil, err
	}
	options := "mode=1777"
	if size > 0 {
		options = fmt.Sprintf("mode=1777,size=%dm", size)
	}
	err := system.Points.AddFS(mount.SessionTag, path, fstype, syscall.MS_NOSUID, options)
	if err != nil {
		return nil, err
	}
	if err := system.RunAfterTag(mount.SessionTag, session.createLayout); err != nil {
		return nil, err
	}
	if layer != nil {
		if err := layer.Add(session, system); err != nil {
			return nil, fmt.Errorf("failed to init layer: %s", err)
		}
		session.Layer = layer
	}
	return session, nil
}

// Path returns the full path of session directory
func (s *Session) Path() string {
	path, _ := s.GetPath("/")
	return path
}

// FinalPath returns the full path to session final directory
func (s *Session) FinalPath() string {
	if s.Layer != nil {
		path, _ := s.GetPath(finalDir)
		return path
	}
	return s.RootFsPath()
}

// OverrideDir overrides a path in the session directory, it simulates
// a bind mount.
func (s *Session) OverrideDir(path string, realpath string) {
	p := path
	if s.Layer != nil {
		p = filepath.Join(s.Layer.Dir(), path)
	}
	s.overrideDir(p, realpath)
}

// RootFsPath returns the full path to session rootfs directory
func (s *Session) RootFsPath() string {
	path, _ := s.GetPath(rootFsDir)
	return path
}

func (s *Session) createLayout(system *mount.System) error {
	// create directory for registered overridden directory
	for _, tag := range mount.GetTagList() {
		for _, point := range system.Points.GetByTag(tag) {
			if point.Source == "" {
				continue
			}

			// search until we find a parent overridden directory
			overridden := false
			for baseDir := filepath.Dir(point.Destination); baseDir != "/"; {
				if _, err := s.GetOverridePath(baseDir); err == nil {
					overridden = true
					break
				}
				baseDir = filepath.Dir(baseDir)
			}
			if !overridden {
				continue
			}

			dest := point.Destination
			if _, err := s.GetPath(dest); err == nil {
				continue
			}
			flags, _ := mount.ConvertOptions(point.Options)

			// ignore anything which is not a bind mount point
			if flags&syscall.MS_BIND == 0 {
				continue
			}

			// check if the bind source exists
			fi, err := s.VFS.Stat(point.Source)
			if err != nil {
				sylog.Warningf("skipping mount of: %s: %s", point.Source, err)
				continue
			}

			if fi.IsDir() {
				if err := s.AddDir(dest); err != nil {
					return err
				}
			} else {
				if err := s.AddFile(dest, nil); err != nil {
					return err
				}
			}
		}
	}

	return s.Create()
}
