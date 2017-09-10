package gocd

import "testing"

func TestResource(t *testing.T) {
	t.Run("Pipeline", testResourcePipeline)
	t.Run("PipelineTemplate", testResourcePipelineTemplate)
	t.Run("PipelineGroups", testResourcePipelineGroups)
}
