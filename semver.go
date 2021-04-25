package semver

import "fmt"

type SemVer struct {
	Major int
	Minor int
	Patch int
	Tag   string
	TagN  int
}

func (v SemVer) String() string {
	if v.Tag == "" {
		return fmt.Sprintf("v%d.%d.%d", v.Major, v.Minor, v.Patch)
	} else {
		return fmt.Sprintf("v%d.%d.%d-%s.%d", v.Major, v.Minor, v.Patch, v.Tag, v.TagN)
	}
}
