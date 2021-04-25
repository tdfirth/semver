package semver_test

import (
	"testing"

	"github.com/tdfirth/semver"
)

func TestVersionString(t *testing.T) {
	v := semver.SemVer{
		Major: 1,
		Minor: 2,
		Patch: 3,
	}

	if v.String() != "v1.2.3" {
		t.Errorf("Expected version to be 'v1.2.3' but found %s", v.String())
	}
}

func TestVersionStringWithTag(t *testing.T) {
	v := semver.SemVer{
		Major: 1,
		Minor: 2,
		Patch: 3,
		Tag:   "alpha",
		TagN:  1,
	}

	if v.String() != "v1.2.3-alpha.1" {
		t.Errorf("Expected version to be 'v1.2.3-alpha.1' but found %s", v.String())
	}
}
