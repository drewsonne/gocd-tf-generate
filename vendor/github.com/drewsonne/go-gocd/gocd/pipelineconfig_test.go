package gocd

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestPipelineConfig(t *testing.T) {
	setup()
	defer teardown()
	t.Run("Create", testPipelineConfigCreate)
	t.Run("Update", testPipelineConfigUpdate)
	t.Run("Delete", testPipelineConfigDelete)
	t.Run("Get", testPipelineConfigGet)
}

func testPipelineConfigGet(t *testing.T) {
	mux.HandleFunc("/api/admin/pipelines/test-pipeline0", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "GET", "Unexpected HTTP method")
		assert.Contains(t, r.Header["Accept"], "application/vnd.go.cd.v4+json")

		j, _ := ioutil.ReadFile("test/resources/pipelineconfig.0.json")
		fmt.Fprint(w, string(j))
	})

	pc, _, err := client.PipelineConfigs.Get(context.Background(), "test-pipeline0")
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, pc)

	assert.NotNil(t, pc.Links.Self)
	assert.Equal(t, "https://ci.example.com/go/api/admin/pipelines/new_pipeline", pc.Links.Self.String())
	assert.NotNil(t, pc.Links.Doc)
	assert.Equal(t, "https://api.gocd.org/#pipeline-config", pc.Links.Doc.String())
	assert.NotNil(t, pc.Links.Find)
	assert.Equal(t, "https://ci.example.com/go/api/admin/pipelines/:name", pc.Links.Find.String())

	assert.Equal(t, "${COUNT}", pc.LabelTemplate)
	assert.True(t, pc.EnablePipelineLocking)
	assert.Equal(t, "new_pipeline", pc.Name)
	assert.Empty(t, pc.Template)

	assert.NotNil(t, pc.Origin)
	assert.Equal(t, "local", pc.Origin.Type)
	assert.Equal(t, "cruise-config.xml", pc.Origin.File)

	assert.Len(t, pc.Parameters, 0)
	assert.Len(t, pc.EnvironmentVariables, 0)

	assert.NotNil(t, pc.Materials)
	assert.Len(t, pc.Materials, 1)
	m := pc.Materials[0]
	assert.Equal(t, "git", m.Type)
	assert.NotNil(t, m.Attributes)
	assert.Equal(t, "git@github.com:sample_repo/example.git", m.Attributes.URL)
	assert.Equal(t, "dest", m.Attributes.Destination)
	assert.Nil(t, m.Attributes.Filter)
	assert.False(t, m.Attributes.InvertFilter)
	assert.Empty(t, m.Attributes.Name)
	assert.True(t, m.Attributes.AutoUpdate)
	assert.Equal(t, "master", m.Attributes.Branch)
	assert.Empty(t, m.Attributes.SubmoduleFolder)
	assert.True(t, m.Attributes.ShallowClone)

	assert.NotNil(t, pc.Stages)
	assert.Len(t, pc.Stages, 1)

	s := pc.Stages[0]
	assert.Equal(t, "defaultStage", s.Name)
	assert.True(t, s.FetchMaterials)
	assert.False(t, s.CleanWorkingDirectory)
	assert.False(t, s.NeverCleanupArtifacts)
	assert.NotNil(t, s.Approval)
	assert.Equal(t, "success", s.Approval.Type)
	assert.NotNil(t, s.Approval.Authorization)
	assert.Len(t, s.Approval.Authorization.Roles, 0)
	assert.Len(t, s.Approval.Authorization.Users, 0)
	assert.Len(t, s.EnvironmentVariables, 0)

	assert.NotNil(t, s.Jobs)
	assert.Len(t, s.Jobs, 1)

	j := s.Jobs[0]
	assert.Equal(t, "defaultJob", j.Name)
	assert.Empty(t, j.RunInstanceCount)
	assert.Equal(t, 0, j.Timeout)
	assert.Len(t, j.EnvironmentVariables, 0)
	assert.Len(t, j.Resources, 0)

	assert.NotNil(t, j.Tasks)
	assert.Len(t, j.Tasks, 1)

	tsk := j.Tasks[0]
	assert.Equal(t, "exec", tsk.Type)
	assert.NotNil(t, tsk.Attributes)
	assert.Len(t, tsk.Attributes.RunIf, 1)
	assert.Equal(t, "passed", tsk.Attributes.RunIf[0])
	assert.Equal(t, "ls", tsk.Attributes.Command)
	assert.Empty(t, tsk.Attributes.WorkingDirectory)

	assert.Len(t, j.Tabs, 0)
	assert.Len(t, j.Artifacts, 0)
	assert.Nil(t, j.Properties)

	// @TODO implement timer and trackingtool
	//assert.Empty(t, pc.TrackingTool)
	//assert.Empty(t, pc.Timer)
}

func testPipelineConfigDelete(t *testing.T) {

	mux.HandleFunc("/api/admin/pipelines/test-pipeline", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "DELETE", "Unexpected HTTP method")
		assert.Equal(t, r.Header.Get("Accept"), apiV4)

		fmt.Fprint(w, `{
  "message": "Pipeline 'test-pipeline' was deleted successfully."
}`)
	})
	message, resp, err := client.PipelineConfigs.Delete(context.Background(), "test-pipeline")
	if err != nil {
		assert.Error(t, err)
	}
	assert.NotNil(t, resp)
	assert.Equal(t, "Pipeline 'test-pipeline' was deleted successfully.", message)
}

func testPipelineConfigCreate(t *testing.T) {
	mux.HandleFunc("/api/admin/pipelines", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "POST", "Unexpected HTTP method")
		//b, err := ioutil.ReadAll(r.Body)
		//if err != nil {
		//	t.Error(err)
		//}
		//assert.Equal(
		//	t,
		//	"{\n  \"group\": \"test-group\",\n  \"pipeline\": {\n    \"name\": \"\",\n    \"stages\": null\n  }\n}\n",
		//	string(b))
		j, _ := ioutil.ReadFile("test/resources/pipelineconfig.0.json")
		fmt.Fprint(w, string(j))
	})

	p := Pipeline{}
	pgs, _, err := client.PipelineConfigs.Create(context.Background(), "test-group", &p)
	if err != nil {
		t.Error(t, err)
	}

	assert.NotNil(t, pgs)
}

func testPipelineConfigUpdate(t *testing.T) {
	mux.HandleFunc("/api/admin/pipelines/test-name", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "PUT", r.Method, "Unexpected HTTP method")
		//b, err := ioutil.ReadAll(r.Body)
		//if err != nil {
		//	t.Error(err)
		//}
		//assert.Equal(
		//	t,
		//	"{\n  \"pipeline\": {\n    \"name\": \"\",\n    \"stages\": null,\n    \"version\": \"test-version\"\n  }\n}\n",
		//	string(b))
		j, _ := ioutil.ReadFile("test/resources/pipelineconfig.0.json")
		fmt.Fprint(w, string(j))
	})

	p := Pipeline{
		Version: "test-version",
	}
	pcs, _, err := client.PipelineConfigs.Update(context.Background(), "test-name", &p)
	if err != nil {
		t.Error(t, err)
	}

	assert.NotNil(t, pcs)
}
