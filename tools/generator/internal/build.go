package internal

import (
	"fmt"
	"os"

	"github.com/Masterminds/semver/v3"
	"gopkg.in/yaml.v3"
)

const buildsPerPatch = 100

type BuildConfigVersion struct {
	Version string `json:"version" yaml:"version"`
	Build   uint8  `json:"build" yaml:"build"`
	Active  bool   `json:"active,omitempty" yaml:"active,omitempty"`
}

func (c *BuildConfigVersion) Inc() {
	c.Build += 1
}

// Folds the build into the patch component (24.6.1 build 2 -> 24.6.102).
// "24.6.1-2" would be a semver prerelease, which Terraform skips when
// resolving range constraints. See README "Provider versioning".
func (c *BuildConfigVersion) GetBuildVersion() (build string, err error) {
	v, err := semver.NewVersion(c.Version)
	if err != nil {
		return "", fmt.Errorf("invalid version %q: %w", c.Version, err)
	}
	if uint64(c.Build) >= buildsPerPatch {
		return "", fmt.Errorf("build %d of version %s must be below %d, it would collide with patch %d build 0",
			c.Build, c.Version, buildsPerPatch, v.Patch()+1)
	}
	return fmt.Sprintf("%d.%d.%d", v.Major(), v.Minor(), v.Patch()*buildsPerPatch+uint64(c.Build)), nil
}

type BuildConfig []*BuildConfigVersion

func (c *BuildConfig) GetBuildVersion(ver string) (build string, err error) {
	for _, v := range *c {
		if v.Version == ver {
			return v.GetBuildVersion()
		}
	}
	return "", fmt.Errorf("version not found")
}

func (c *BuildConfig) Load(filename string) error {
	var payload, err = os.ReadFile(filename)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(payload, &c)
	if err != nil {
		return err
	}
	return nil
}

func (c *BuildConfig) Save(filename string) error {
	var payload, _ = yaml.Marshal(c)
	return os.WriteFile(filename, payload, 0655)
}
