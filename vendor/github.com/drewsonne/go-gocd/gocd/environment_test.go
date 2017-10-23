package gocd

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestEnvironment(t *testing.T) {
	setup()
	defer teardown()

	t.Run("List", testEnvironmentList)
}

func testEnvironmentList(t *testing.T) {
	mux.HandleFunc("/api/admin/environments", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "GET", "Unexpected HTTP method")
		assert.Contains(t, r.Header["Accept"], "application/vnd.go.cd.v2+json")

		j, _ := ioutil.ReadFile("test/resources/environment.0.json")
		fmt.Fprint(w, string(j))
	})

	envs, _, err := client.Environments.List(context.Background())
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, envs)

	assert.NotNil(t, envs.Links.Get("Self"))
	assert.Equal(t, "https://ci.example.com/go/api/admin/environments", envs.Links.Get("Self").URL.String())
	assert.NotNil(t, envs.Links.Get("Doc"))
	assert.Equal(t, "https://api.gocd.org/#environment-config", envs.Links.Get("Doc").URL.String())

	assert.NotNil(t, envs.Embedded)
	assert.NotNil(t, envs.Embedded.Environments)
	assert.Len(t, envs.Embedded.Environments, 1)

	env := envs.Embedded.Environments[0]
	assert.NotNil(t, env.Links)
	assert.Equal(t, "https://ci.example.com/go/api/admin/environments/foobar", env.Links.Get("Self").URL.String())
	assert.Equal(t, "https://api.gocd.org/#environment-config", env.Links.Get("Doc").URL.String())
	assert.Equal(t, "https://ci.example.com/go/api/admin/environments/:environment_name", env.Links.Get("Find").URL.String())

	assert.Equal(t, "foobar", env.Name)

	assert.NotNil(t, env.Pipelines)
	assert.Len(t, env.Pipelines, 1)

	p := env.Pipelines[0]
	assert.NotNil(t, p.Links)
	assert.Equal(t, "https://ci.example.com/go/api/admin/pipelines/up42", p.Links.Get("Self").URL.String())
	assert.Equal(t, "https://api.gocd.org/#pipeline-config", p.Links.Get("Doc").URL.String())
	assert.Equal(t, "https://ci.example.com/go/api/admin/pipelines/:pipeline_name", p.Links.Get("Find").URL.String())
	assert.Equal(t, "up42", p.Name)

	assert.NotNil(t, env.Agents)
	assert.Len(t, env.Agents, 1)

	a := env.Agents[0]
	assert.NotNil(t, a.Links)
	assert.Equal(t, "https://ci.example.com/go/api/agents/adb9540a-b954-4571-9d9b-2f330739d4da", a.Links.Get("Self").URL.String())
	assert.Equal(t, "https://api.gocd.org/#agents", a.Links.Get("Doc").URL.String())
	assert.Equal(t, "https://ci.example.com/go/api/agents/:uuid", a.Links.Get("Find").URL.String())
	assert.Equal(t, "12345678-e2f6-4c78-123456789012", a.UUID)

	assert.NotNil(t, env.EnvironmentVariables)
	assert.Len(t, env.EnvironmentVariables, 2)

	ev1 := env.EnvironmentVariables[0]
	assert.Equal(t, "username", ev1.Name)
	assert.False(t, ev1.Secure)
	assert.Equal(t, "admin", ev1.Value)

	ev2 := env.EnvironmentVariables[1]
	assert.Equal(t, "password", ev2.Name)
	assert.True(t, ev2.Secure)
	assert.Equal(t, "LSd1TI0eLa+DjytHjj0qjA==", ev2.EncryptedValue)
}
