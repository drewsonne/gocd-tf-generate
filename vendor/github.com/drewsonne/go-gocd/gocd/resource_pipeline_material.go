package gocd

import (
	"fmt"
)

// Equal is true if the two materials are logically equivalent. Not neccesarily literally equal.
func (m Material) Equal(a *Material) bool {
	if m.Type != a.Type {
		return false
	}
	switch m.Type {
	case "git":
		return m.Attributes.equalGit(&a.Attributes)
	default:
		panic(fmt.Errorf("Material comparison not implemented for '%s'", m.Type))
	}
}

func (a1 MaterialAttributes) equalGit(a2 *MaterialAttributes) bool {
	urlsEqual := a1.URL == a2.URL
	branchesEqual := a1.Branch == a2.Branch ||
		a1.Branch == "" && a2.Branch == "master" ||
		a1.Branch == "master" && a2.Branch == ""

	if !urlsEqual {
		return false
	}
	return branchesEqual
}
