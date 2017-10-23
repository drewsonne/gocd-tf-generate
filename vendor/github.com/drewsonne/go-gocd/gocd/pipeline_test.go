package gocd

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestPipelineService(t *testing.T) {
	setup()
	defer teardown()

	t.Run("Get", testPipelineServiceGet)
	t.Run("Create", testPipelineServiceCreate)
	t.Run("GetHistory", testPipelineServiceGetHistory)
	t.Run("GetStatus", testPipelineServiceGetStatus)
	t.Run("Pause", testPipelineServicePause)
	t.Run("Unpause", testPipelineServiceUnpause)
	t.Run("ReleaseLock", testPipelineServiceReleaseLock)
	t.Run("PaginationStub", testPipelineServicePaginationStub)
	t.Run("StageContainer", testPipelineStageContainer)
}

func testPipelineStageContainer(t *testing.T) {

	p := &Pipeline{
		Name:   "mock-name",
		Stages: []*Stage{{Name: "1"}, {Name: "2"}},
	}

	i := StageContainer(p)

	assert.Equal(t, "mock-name", i.GetName())
	assert.Len(t, i.GetStages(), 2)

	i.AddStage(&Stage{Name: "3"})
	assert.Len(t, i.GetStages(), 3)

	s1 := i.GetStage("1")
	assert.Equal(t, s1.Name, "1")

	sn := i.GetStage("hello")
	assert.Nil(t, sn)

	i.SetStages([]*Stage{})
	assert.Len(t, i.GetStages(), 0)
}

func testPipelineServicePaginationStub(t *testing.T) {
	pgs := PipelinesService{}

	assert.Equal(t, "a/b/c/4",
		pgs.buildPaginatedStub("a/%s/c", "b", 4))

	assert.Equal(t, "a/b/c",
		pgs.buildPaginatedStub("a/%s/c", "b", 0))

}

func testPipelineServiceGetStatus(t *testing.T) {
	mux.HandleFunc("/api/pipelines/test-pipeline/status", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "GET", "Unexpected HTTP method")
		b, _ := ioutil.ReadFile("test/resources/pipeline.2.json")
		fmt.Fprint(w, string(b))
	})

	ps, _, err := client.Pipelines.GetStatus(context.Background(), "test-pipeline", 0)
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, ps)
	assert.False(t, ps.Locked)
	assert.True(t, ps.Paused)
	assert.False(t, ps.Schedulable)
}

func testPipelineServiceCreate(t *testing.T) {
	mux.HandleFunc("/api/admin/pipelines", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "POST", "Unexpected HTTP method")
		// @TODO Renable and fix diff
		//expectedBody, _ := ioutil.ReadFile("test/request/pipeline.0.json")
		//actual, _ := ioutil.ReadAll(r.Body)
		//assert.Equal(t, string(expectedBody), string(actual))

		j, _ := ioutil.ReadFile("test/resources/pipeline.0.json")
		fmt.Fprint(w, string(j))
	})

	p := Pipeline{
		LabelTemplate:         "${COUNT}",
		EnablePipelineLocking: true,
		Name: "new_pipeline",
		Materials: []Material{
			{
				Type: "git", Attributes: MaterialAttributes{
					URL:          "git@github.com:sample_repo/example.git",
					Destination:  "dest",
					InvertFilter: false,
					AutoUpdate:   true,
					Branch:       "master",
					ShallowClone: true,
				},
			},
		},
		Stages: []*Stage{
			{
				Name:           "defaultStage",
				FetchMaterials: true,
				Approval: &Approval{
					Type: "success",
					Authorization: &Authorization{
						Roles: []string{},
						Users: []string{},
					},
				},
				Jobs: []*Job{
					{
						Name: "defaultJob",
						Tasks: []*Task{
							{
								Type: "exec",
								Attributes: TaskAttributes{
									RunIf:   []string{"passed"},
									Command: "ls",
								},
							},
						},
					},
				},
			},
		},
		Version: "mock-version",
	}
	pr, _, err := client.PipelineConfigs.Create(context.Background(), "first", &p)
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, pr)
	assert.Equal(t, "test-pipeline", pr.Name)

	assert.Len(t, pr.Stages, 1)

	stage := pr.Stages[0]
	assert.Equal(t, "stage1", stage.Name)
}

func testPipelineServiceGet(t *testing.T) {
	mux.HandleFunc("/api/admin/pipelines/test-pipeline/instance", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "GET", "Unexpected HTTP method")
		j, _ := ioutil.ReadFile("test/resources/pipeline.0.json")
		fmt.Fprint(w, string(j))
	})

	p, _, err := client.Pipelines.GetInstance(context.Background(), "test-pipeline", 0)
	if err != nil {
		t.Error(err)
	}
	assert.NotNil(t, p)
	assert.Equal(t, p.Name, "test-pipeline")

	assert.Len(t, p.Stages, 1)

	s := p.Stages[0]
	assert.Equal(t, "stage1", s.Name)
	assert.Equal(t, false, s.FetchMaterials)
	assert.Equal(t, false, s.CleanWorkingDirectory)
	assert.Equal(t, false, s.NeverCleanupArtifacts)

	assert.Len(t, s.EnvironmentVariables, 0)
	assert.Nil(t, s.Approval)
}

func testPipelineServiceGetHistory(t *testing.T) {
	mux.HandleFunc("/api/pipelines/test-pipeline/history", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "GET", "Unexpected HTTP method")
		j, _ := ioutil.ReadFile("test/resources/pipeline.1.json")
		fmt.Fprint(w, string(j))
	})
	ph, _, err := client.Pipelines.GetHistory(context.Background(), "test-pipeline", 0)
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, ph)
	assert.Len(t, ph.Pipelines, 2)

	h1 := ph.Pipelines[0]
	assert.True(t, h1.CanRun)
	assert.Equal(t, h1.Name, "pipeline1")
	assert.Equal(t, h1.NaturalOrder, 11)
	assert.Equal(t, h1.Comment, "")
	assert.Len(t, h1.Stages, 1)

	h1s := h1.Stages[0]
	assert.Equal(t, h1s.Name, "stage1")

	h2 := ph.Pipelines[1]
	assert.True(t, h2.CanRun)
	assert.Equal(t, h2.Name, "pipeline1")
	assert.Equal(t, h2.NaturalOrder, 10)
	assert.Equal(t, h2.Comment, "")
	assert.Len(t, h2.Stages, 1)

	h2s := h2.Stages[0]
	assert.Equal(t, h2s.Name, "stage1")

}
