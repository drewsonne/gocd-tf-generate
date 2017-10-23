package gocd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func testResourceMaterial(t *testing.T) {
	t.Run("Equality", testMaterialEquality)
	t.Run("AttributeEquality", testMaterialAttributeEquality)
}

func testMaterialEquality(t *testing.T) {
	s1 := Material{
		Type: "git",
		Attributes: MaterialAttributes{
			URL: "https://github.com/gocd/gocd",
		},
	}

	s2 := Material{
		Type: "git",
		Attributes: MaterialAttributes{
			Name: "gocd-src",
			URL:  "https://github.com/gocd/gocd",
		},
	}

	assert.True(t, s1.Equal(&s2))
}

func testMaterialAttributeEquality(t *testing.T) {
	a1 := MaterialAttributes{}
	a2 := MaterialAttributes{}
	assert.True(t, a1.equalGit(&a2))

	a2.URL = "https://github.com/drewsonne/go-gocd"
	assert.False(t, a1.equalGit(&a2))

	a1.URL = "https://github.com/drewsonne/go-gocd"
	a2.Branch = "feature/branch"
	assert.False(t, a1.equalGit(&a2))

	for _, branchCombo := range [][]string{
		{"", "master"},
		{"master", ""},
		{"", ""},
		{"master", "master"},
	} {
		a1.Branch = branchCombo[0]
		a2.Branch = branchCombo[1]
		assert.True(t, a1.equalGit(&a2))
	}
}
