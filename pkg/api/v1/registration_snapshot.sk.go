// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"fmt"

	"github.com/solo-io/solo-kit/pkg/utils/hashutils"
	"go.uber.org/zap"
)

type RegistrationSnapshot struct {
	Meshes        MeshesByNamespace
	Meshingresses MeshingressesByNamespace
}

func (s RegistrationSnapshot) Clone() RegistrationSnapshot {
	return RegistrationSnapshot{
		Meshes:        s.Meshes.Clone(),
		Meshingresses: s.Meshingresses.Clone(),
	}
}

func (s RegistrationSnapshot) Hash() uint64 {
	return hashutils.HashAll(
		s.hashMeshes(),
		s.hashMeshingresses(),
	)
}

func (s RegistrationSnapshot) hashMeshes() uint64 {
	return hashutils.HashAll(s.Meshes.List().AsInterfaces()...)
}

func (s RegistrationSnapshot) hashMeshingresses() uint64 {
	return hashutils.HashAll(s.Meshingresses.List().AsInterfaces()...)
}

func (s RegistrationSnapshot) HashFields() []zap.Field {
	var fields []zap.Field
	fields = append(fields, zap.Uint64("meshes", s.hashMeshes()))
	fields = append(fields, zap.Uint64("meshingresses", s.hashMeshingresses()))

	return append(fields, zap.Uint64("snapshotHash", s.Hash()))
}

type RegistrationSnapshotStringer struct {
	Version       uint64
	Meshes        []string
	Meshingresses []string
}

func (ss RegistrationSnapshotStringer) String() string {
	s := fmt.Sprintf("RegistrationSnapshot %v\n", ss.Version)

	s += fmt.Sprintf("  Meshes %v\n", len(ss.Meshes))
	for _, name := range ss.Meshes {
		s += fmt.Sprintf("    %v\n", name)
	}

	s += fmt.Sprintf("  Meshingresses %v\n", len(ss.Meshingresses))
	for _, name := range ss.Meshingresses {
		s += fmt.Sprintf("    %v\n", name)
	}

	return s
}

func (s RegistrationSnapshot) Stringer() RegistrationSnapshotStringer {
	return RegistrationSnapshotStringer{
		Version:       s.Hash(),
		Meshes:        s.Meshes.List().NamespacesDotNames(),
		Meshingresses: s.Meshingresses.List().NamespacesDotNames(),
	}
}
