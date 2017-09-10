package gocd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func testResourcePipelineTemplate(t *testing.T) {
	s1 := Stage{Name: "s"}
	s1.CleanWorkingDirectory = false

	s2 := Stage{Name: "s"}
	s2.CleanWorkingDirectory = true

	p := PipelineTemplate{}
	p.AddStage(&s1)

	s := p.GetStage("s")
	assert.False(t, s.CleanWorkingDirectory)

	p.SetStage(&s2)

	s = p.GetStage("s")
	assert.True(t, s.CleanWorkingDirectory)

	s3 := Stage{Name: "s3"}
	p.SetStage(&s3)

	p.SetStage(&s3)

	s = p.GetStage("s3")
	assert.Equal(t, s.Name, "s3")

}
