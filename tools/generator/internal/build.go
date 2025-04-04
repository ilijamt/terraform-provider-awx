package internal

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type BuildConfigVersion struct {
	Version string `json:"version" yaml:"version"`
	Build   uint8  `json:"build" yaml:"build"`
	Active  bool   `json:"active,omitempty" yaml:"active,omitempty"`
}

func (c *BuildConfigVersion) Inc() {
	c.Build += 1
}

func (c *BuildConfigVersion) GetBuildVersion() (build string) {
	return fmt.Sprintf("%s-%d", c.Version, c.Build)
}

type BuildConfig []*BuildConfigVersion

func (c *BuildConfig) GetBuildVersion(ver string) (build string, err error) {
	for _, v := range *c {
		if v.Version == ver {
			return v.GetBuildVersion(), nil
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
