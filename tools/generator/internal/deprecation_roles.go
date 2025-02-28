package internal

import (
	"github.com/Masterminds/semver/v3"
)

type Deprecation2430 struct {
	deprecation
}

func (r *Deprecation2430) Check(mc *ModelConfig) (err error) {
	var version *semver.Version
	var constraint *semver.Constraints

	version, _ = semver.NewVersion(mc.ApiVersion)
	constraint, _ = semver.NewConstraint(">= 24.3.0")

	if constraint.Check(version) && mc.HasObjectRoles {
		mc.DeprecatedParts[r.Name()] = true
	}

	return nil
}

var _ Deprecation = (*Deprecation2430)(nil)
