

# gocd
`import "github.com/drewsonne/go-gocd/gocd"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Subdirectories](#pkg-subdirectories)

## <a name="pkg-overview">Overview</a>
Package gocd provides a client for using the GoCD Server API.

Usage:


	import "github.com/drewsonne/go-gocd/gocd"

Construct a new GoCD client and supply the URL to your GoCD server and if required, username and password. Then use the
various services on the client to access different parts of the GoCD API.
For example:


	package main
	import (
		"github.com/drewsonne/go-gocd/gocd"
		"context"
		"fmt"
	)
	
	func main() {
		cfg := gocd.Configuration{
			Server: "<a href="https://my_gocd/go/">https://my_gocd/go/</a>",
			Username: "ApiUser",
			Password: "MySecretPassword",
		}
	
		c := cfg.Client()
	
		// list all agents in use by the GoCD Server
		var a []*gocd.Agent
		var err error
		var r *gocd.APIResponse
		if a, r, err = c.Agents.List(context.Background()); err != nil {
			if r.HTTP.StatusCode == 404 {
				fmt.Println("Couldn't find agent")
			} else {
				panic(err)
			}
		}
	
		fmt.Println(a)
	}

If you wish to use your own http client, you can use the following idiom


	package main
	
	import (
		"github.com/drewsonne/go-gocd/gocd"
		"net/http"
		"context"
	)
	
	func main() {
		client := gocd.NewClient(
			&gocd.Configuration{},
			&http.Client{},
		)
		client.Login(context.Background())
	}

The services of a client divide the API into logical chunks and correspond to
the structure of the GoCD API documentation at
<a href="https://api.gocd.org/current/">https://api.gocd.org/current/</a>.




## <a name="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [func CheckResponse(response *http.Response) error](#CheckResponse)
* [func ConfigFilePath() (configPath string, err error)](#ConfigFilePath)
* [func LoadConfigByName(name string, cfg *Configuration) (err error)](#LoadConfigByName)
* [func LoadConfigFromFile() (cfgs map[string]*Configuration, err error)](#LoadConfigFromFile)
* [func SetupLogging(log *logrus.Logger)](#SetupLogging)
* [type APIClientRequest](#APIClientRequest)
* [type APIRequest](#APIRequest)
* [type APIResponse](#APIResponse)
* [type Agent](#Agent)
  * [func (a *Agent) GetLinks() *HALLinks](#Agent.GetLinks)
  * [func (a *Agent) RemoveLinks()](#Agent.RemoveLinks)
* [type AgentBulkOperationUpdate](#AgentBulkOperationUpdate)
* [type AgentBulkOperationsUpdate](#AgentBulkOperationsUpdate)
* [type AgentBulkUpdate](#AgentBulkUpdate)
* [type AgentsResponse](#AgentsResponse)
* [type AgentsService](#AgentsService)
  * [func (s *AgentsService) BulkUpdate(ctx context.Context, agents AgentBulkUpdate) (string, *APIResponse, error)](#AgentsService.BulkUpdate)
  * [func (s *AgentsService) Delete(ctx context.Context, uuid string) (string, *APIResponse, error)](#AgentsService.Delete)
  * [func (s *AgentsService) Get(ctx context.Context, uuid string) (*Agent, *APIResponse, error)](#AgentsService.Get)
  * [func (s *AgentsService) JobRunHistory(ctx context.Context, uuid string) ([]*Job, *APIResponse, error)](#AgentsService.JobRunHistory)
  * [func (s *AgentsService) List(ctx context.Context) ([]*Agent, *APIResponse, error)](#AgentsService.List)
  * [func (s *AgentsService) Update(ctx context.Context, uuid string, agent *Agent) (*Agent, *APIResponse, error)](#AgentsService.Update)
* [type Approval](#Approval)
  * [func (a *Approval) Clean()](#Approval.Clean)
* [type Artifact](#Artifact)
* [type Auth](#Auth)
* [type Authorization](#Authorization)
* [type BuildCause](#BuildCause)
* [type BuildDetails](#BuildDetails)
* [type CipherText](#CipherText)
* [type Client](#Client)
  * [func NewClient(cfg *Configuration, httpClient *http.Client) *Client](#NewClient)
  * [func (c *Client) Do(ctx context.Context, req *APIRequest, v interface{}, responseType string) (*APIResponse, error)](#Client.Do)
  * [func (c *Client) Lock()](#Client.Lock)
  * [func (c *Client) Login(ctx context.Context) error](#Client.Login)
  * [func (c *Client) NewRequest(method, urlStr string, body interface{}, apiVersion string) (req *APIRequest, err error)](#Client.NewRequest)
  * [func (c *Client) Unlock()](#Client.Unlock)
* [type ConfigApproval](#ConfigApproval)
* [type ConfigArtifact](#ConfigArtifact)
* [type ConfigAuthConfig](#ConfigAuthConfig)
* [type ConfigElastic](#ConfigElastic)
* [type ConfigElasticProfile](#ConfigElasticProfile)
* [type ConfigEnvironmentVariable](#ConfigEnvironmentVariable)
* [type ConfigFilter](#ConfigFilter)
* [type ConfigJob](#ConfigJob)
* [type ConfigMaterialRepository](#ConfigMaterialRepository)
* [type ConfigPackage](#ConfigPackage)
* [type ConfigParam](#ConfigParam)
* [type ConfigPipeline](#ConfigPipeline)
* [type ConfigPipelineGroup](#ConfigPipelineGroup)
* [type ConfigPluginConfiguration](#ConfigPluginConfiguration)
* [type ConfigProperty](#ConfigProperty)
* [type ConfigRepository](#ConfigRepository)
* [type ConfigRepositoryGit](#ConfigRepositoryGit)
* [type ConfigRole](#ConfigRole)
* [type ConfigSCM](#ConfigSCM)
* [type ConfigSecurity](#ConfigSecurity)
* [type ConfigServer](#ConfigServer)
* [type ConfigStage](#ConfigStage)
* [type ConfigTask](#ConfigTask)
* [type ConfigTaskRunIf](#ConfigTaskRunIf)
* [type ConfigTasks](#ConfigTasks)
* [type ConfigXML](#ConfigXML)
* [type Configuration](#Configuration)
  * [func (c *Configuration) Client() *Client](#Configuration.Client)
  * [func (c *Configuration) HasAuth() bool](#Configuration.HasAuth)
* [type ConfigurationService](#ConfigurationService)
  * [func (cs *ConfigurationService) Get(ctx context.Context) (*ConfigXML, *APIResponse, error)](#ConfigurationService.Get)
  * [func (cs *ConfigurationService) GetVersion(ctx context.Context) (*Version, *APIResponse, error)](#ConfigurationService.GetVersion)
* [type EmbeddedEnvironments](#EmbeddedEnvironments)
* [type EncryptionService](#EncryptionService)
  * [func (es *EncryptionService) Encrypt(ctx context.Context, plaintext string) (*CipherText, *APIResponse, error)](#EncryptionService.Encrypt)
* [type Environment](#Environment)
  * [func (env *Environment) GetLinks() *HALLinks](#Environment.GetLinks)
  * [func (env *Environment) GetVersion() (version string)](#Environment.GetVersion)
  * [func (env *Environment) RemoveLinks()](#Environment.RemoveLinks)
  * [func (env *Environment) SetVersion(version string)](#Environment.SetVersion)
* [type EnvironmentPatchRequest](#EnvironmentPatchRequest)
* [type EnvironmentVariable](#EnvironmentVariable)
* [type EnvironmentVariablesAction](#EnvironmentVariablesAction)
* [type EnvironmentsResponse](#EnvironmentsResponse)
  * [func (er *EnvironmentsResponse) GetLinks() *HALLinks](#EnvironmentsResponse.GetLinks)
  * [func (er *EnvironmentsResponse) RemoveLinks()](#EnvironmentsResponse.RemoveLinks)
* [type EnvironmentsService](#EnvironmentsService)
  * [func (es *EnvironmentsService) Create(ctx context.Context, name string) (*Environment, *APIResponse, error)](#EnvironmentsService.Create)
  * [func (es *EnvironmentsService) Delete(ctx context.Context, name string) (string, *APIResponse, error)](#EnvironmentsService.Delete)
  * [func (es *EnvironmentsService) Get(ctx context.Context, name string) (*Environment, *APIResponse, error)](#EnvironmentsService.Get)
  * [func (es *EnvironmentsService) List(ctx context.Context) (*EnvironmentsResponse, *APIResponse, error)](#EnvironmentsService.List)
  * [func (es *EnvironmentsService) Patch(ctx context.Context, name string, patch *EnvironmentPatchRequest) (*Environment, *APIResponse, error)](#EnvironmentsService.Patch)
* [type GitRepositoryMaterial](#GitRepositoryMaterial)
* [type HALContainer](#HALContainer)
* [type HALLink](#HALLink)
* [type HALLinks](#HALLinks)
  * [func (al *HALLinks) Add(link *HALLink)](#HALLinks.Add)
  * [func (al HALLinks) Get(name string) *HALLink](#HALLinks.Get)
  * [func (al HALLinks) GetOk(name string) (*HALLink, bool)](#HALLinks.GetOk)
  * [func (al HALLinks) Keys() []string](#HALLinks.Keys)
  * [func (al HALLinks) MarshallJSON() ([]byte, error)](#HALLinks.MarshallJSON)
  * [func (al *HALLinks) UnmarshalJSON(j []byte) (e error)](#HALLinks.UnmarshalJSON)
* [type Job](#Job)
  * [func (j *Job) JSONString() (string, error)](#Job.JSONString)
  * [func (j *Job) Validate() error](#Job.Validate)
* [type JobProperty](#JobProperty)
* [type JobRunHistoryResponse](#JobRunHistoryResponse)
* [type JobSchedule](#JobSchedule)
* [type JobScheduleEnvVar](#JobScheduleEnvVar)
* [type JobScheduleLink](#JobScheduleLink)
* [type JobScheduleResponse](#JobScheduleResponse)
* [type JobStateTransition](#JobStateTransition)
* [type JobsService](#JobsService)
  * [func (js *JobsService) ListScheduled(ctx context.Context) ([]*JobSchedule, *APIResponse, error)](#JobsService.ListScheduled)
* [type MailHost](#MailHost)
* [type Material](#Material)
  * [func (m Material) Equal(a *Material) (isEqual bool, err error)](#Material.Equal)
  * [func (m *Material) UnmarshalJSON(b []byte) error](#Material.UnmarshalJSON)
* [type MaterialAttribute](#MaterialAttribute)
* [type MaterialAttributesDependency](#MaterialAttributesDependency)
* [type MaterialAttributesGit](#MaterialAttributesGit)
* [type MaterialAttributesHg](#MaterialAttributesHg)
* [type MaterialAttributesP4](#MaterialAttributesP4)
* [type MaterialAttributesPackage](#MaterialAttributesPackage)
* [type MaterialAttributesPlugin](#MaterialAttributesPlugin)
* [type MaterialAttributesSvn](#MaterialAttributesSvn)
* [type MaterialAttributesTfs](#MaterialAttributesTfs)
* [type MaterialFilter](#MaterialFilter)
* [type MaterialRevision](#MaterialRevision)
* [type Modification](#Modification)
* [type PaginationResponse](#PaginationResponse)
* [type Parameter](#Parameter)
* [type PasswordFilePath](#PasswordFilePath)
* [type PatchStringAction](#PatchStringAction)
* [type Pipeline](#Pipeline)
  * [func (p *Pipeline) AddStage(stage *Stage)](#Pipeline.AddStage)
  * [func (p *Pipeline) GetLinks() *HALLinks](#Pipeline.GetLinks)
  * [func (p *Pipeline) GetName() string](#Pipeline.GetName)
  * [func (p *Pipeline) GetStage(stageName string) *Stage](#Pipeline.GetStage)
  * [func (p *Pipeline) GetStages() []*Stage](#Pipeline.GetStages)
  * [func (p *Pipeline) GetVersion() (version string)](#Pipeline.GetVersion)
  * [func (p *Pipeline) RemoveLinks()](#Pipeline.RemoveLinks)
  * [func (p *Pipeline) SetStage(newStage *Stage)](#Pipeline.SetStage)
  * [func (p *Pipeline) SetStages(stages []*Stage)](#Pipeline.SetStages)
  * [func (p *Pipeline) SetVersion(version string)](#Pipeline.SetVersion)
* [type PipelineConfigOrigin](#PipelineConfigOrigin)
* [type PipelineConfigRequest](#PipelineConfigRequest)
  * [func (pr *PipelineConfigRequest) GetVersion() (version string)](#PipelineConfigRequest.GetVersion)
  * [func (pr *PipelineConfigRequest) SetVersion(version string)](#PipelineConfigRequest.SetVersion)
* [type PipelineConfigsService](#PipelineConfigsService)
  * [func (pcs *PipelineConfigsService) Create(ctx context.Context, group string, p *Pipeline) (*Pipeline, *APIResponse, error)](#PipelineConfigsService.Create)
  * [func (pcs *PipelineConfigsService) Delete(ctx context.Context, name string) (string, *APIResponse, error)](#PipelineConfigsService.Delete)
  * [func (pcs *PipelineConfigsService) Get(ctx context.Context, name string) (*Pipeline, *APIResponse, error)](#PipelineConfigsService.Get)
  * [func (pcs *PipelineConfigsService) Update(ctx context.Context, name string, p *Pipeline) (*Pipeline, *APIResponse, error)](#PipelineConfigsService.Update)
* [type PipelineGroup](#PipelineGroup)
* [type PipelineGroups](#PipelineGroups)
  * [func (pg *PipelineGroups) GetGroupByPipeline(pipeline *Pipeline) *PipelineGroup](#PipelineGroups.GetGroupByPipeline)
  * [func (pg *PipelineGroups) GetGroupByPipelineName(pipelineName string) *PipelineGroup](#PipelineGroups.GetGroupByPipelineName)
* [type PipelineGroupsService](#PipelineGroupsService)
  * [func (pgs *PipelineGroupsService) List(ctx context.Context, name string) (*PipelineGroups, *APIResponse, error)](#PipelineGroupsService.List)
* [type PipelineHistory](#PipelineHistory)
* [type PipelineInstance](#PipelineInstance)
* [type PipelineMaterial](#PipelineMaterial)
* [type PipelineRequest](#PipelineRequest)
* [type PipelineStatus](#PipelineStatus)
* [type PipelineTemplate](#PipelineTemplate)
  * [func (pt *PipelineTemplate) AddStage(stage *Stage)](#PipelineTemplate.AddStage)
  * [func (pt PipelineTemplate) GetName() string](#PipelineTemplate.GetName)
  * [func (pt PipelineTemplate) GetStage(stageName string) *Stage](#PipelineTemplate.GetStage)
  * [func (pt PipelineTemplate) GetStages() []*Stage](#PipelineTemplate.GetStages)
  * [func (pt PipelineTemplate) GetVersion() (version string)](#PipelineTemplate.GetVersion)
  * [func (pt PipelineTemplate) Pipelines() []*Pipeline](#PipelineTemplate.Pipelines)
  * [func (pt *PipelineTemplate) RemoveLinks()](#PipelineTemplate.RemoveLinks)
  * [func (pt *PipelineTemplate) SetStage(newStage *Stage)](#PipelineTemplate.SetStage)
  * [func (pt *PipelineTemplate) SetStages(stages []*Stage)](#PipelineTemplate.SetStages)
  * [func (pt *PipelineTemplate) SetVersion(version string)](#PipelineTemplate.SetVersion)
* [type PipelineTemplateRequest](#PipelineTemplateRequest)
  * [func (pt PipelineTemplateRequest) GetVersion() (version string)](#PipelineTemplateRequest.GetVersion)
  * [func (pt *PipelineTemplateRequest) SetVersion(version string)](#PipelineTemplateRequest.SetVersion)
* [type PipelineTemplateResponse](#PipelineTemplateResponse)
* [type PipelineTemplatesResponse](#PipelineTemplatesResponse)
* [type PipelineTemplatesService](#PipelineTemplatesService)
  * [func (pts *PipelineTemplatesService) Create(ctx context.Context, name string, st []*Stage) (*PipelineTemplate, *APIResponse, error)](#PipelineTemplatesService.Create)
  * [func (pts *PipelineTemplatesService) Delete(ctx context.Context, name string) (string, *APIResponse, error)](#PipelineTemplatesService.Delete)
  * [func (pts *PipelineTemplatesService) Get(ctx context.Context, name string) (*PipelineTemplate, *APIResponse, error)](#PipelineTemplatesService.Get)
  * [func (pts *PipelineTemplatesService) List(ctx context.Context) ([]*PipelineTemplate, *APIResponse, error)](#PipelineTemplatesService.List)
  * [func (pts *PipelineTemplatesService) Update(ctx context.Context, name string, template *PipelineTemplate) (*PipelineTemplate, *APIResponse, error)](#PipelineTemplatesService.Update)
* [type PipelinesService](#PipelinesService)
  * [func (pgs *PipelinesService) GetHistory(ctx context.Context, name string, offset int) (*PipelineHistory, *APIResponse, error)](#PipelinesService.GetHistory)
  * [func (pgs *PipelinesService) GetInstance(ctx context.Context, name string, offset int) (*PipelineInstance, *APIResponse, error)](#PipelinesService.GetInstance)
  * [func (pgs *PipelinesService) GetStatus(ctx context.Context, name string, offset int) (*PipelineStatus, *APIResponse, error)](#PipelinesService.GetStatus)
  * [func (pgs *PipelinesService) Pause(ctx context.Context, name string) (bool, *APIResponse, error)](#PipelinesService.Pause)
  * [func (pgs *PipelinesService) ReleaseLock(ctx context.Context, name string) (bool, *APIResponse, error)](#PipelinesService.ReleaseLock)
  * [func (pgs *PipelinesService) Unpause(ctx context.Context, name string) (bool, *APIResponse, error)](#PipelinesService.Unpause)
* [type PluggableInstanceSettings](#PluggableInstanceSettings)
* [type Plugin](#Plugin)
* [type PluginConfiguration](#PluginConfiguration)
* [type PluginConfigurationKVPair](#PluginConfigurationKVPair)
* [type PluginConfigurationMetadata](#PluginConfigurationMetadata)
* [type PluginView](#PluginView)
* [type PluginsResponse](#PluginsResponse)
* [type PluginsService](#PluginsService)
  * [func (ps *PluginsService) Get(ctx context.Context, name string) (*Plugin, *APIResponse, error)](#PluginsService.Get)
  * [func (ps *PluginsService) List(ctx context.Context) (*PluginsResponse, *APIResponse, error)](#PluginsService.List)
* [type Properties](#Properties)
  * [func NewPropertiesFrame(frame [][]string) *Properties](#NewPropertiesFrame)
  * [func (pr *Properties) AddRow(r []string)](#Properties.AddRow)
  * [func (pr Properties) Get(row int, column string) string](#Properties.Get)
  * [func (pr *Properties) MarshalJSON() ([]byte, error)](#Properties.MarshalJSON)
  * [func (pr Properties) MarshallCSV() (string, error)](#Properties.MarshallCSV)
  * [func (pr *Properties) SetRow(row int, r []string)](#Properties.SetRow)
  * [func (pr *Properties) UnmarshallCSV(raw string) error](#Properties.UnmarshallCSV)
  * [func (pr *Properties) Write(p []byte) (n int, err error)](#Properties.Write)
* [type PropertiesService](#PropertiesService)
  * [func (ps *PropertiesService) Create(ctx context.Context, name string, value string, pr *PropertyRequest) (bool, *APIResponse, error)](#PropertiesService.Create)
  * [func (ps *PropertiesService) Get(ctx context.Context, name string, pr *PropertyRequest) (*Properties, *APIResponse, error)](#PropertiesService.Get)
  * [func (ps *PropertiesService) List(ctx context.Context, pr *PropertyRequest) (*Properties, *APIResponse, error)](#PropertiesService.List)
  * [func (ps *PropertiesService) ListHistorical(ctx context.Context, pr *PropertyRequest) (*Properties, *APIResponse, error)](#PropertiesService.ListHistorical)
* [type PropertyCreateResponse](#PropertyCreateResponse)
* [type PropertyRequest](#PropertyRequest)
* [type Stage](#Stage)
  * [func (s *Stage) Clean()](#Stage.Clean)
  * [func (s *Stage) JSONString() (string, error)](#Stage.JSONString)
  * [func (s *Stage) Validate() error](#Stage.Validate)
* [type StageContainer](#StageContainer)
* [type StagesService](#StagesService)
* [type StringResponse](#StringResponse)
* [type Tab](#Tab)
* [type Task](#Task)
  * [func (t *Task) Validate() error](#Task.Validate)
* [type TaskAttributes](#TaskAttributes)
  * [func (t *TaskAttributes) ValidateAnt() error](#TaskAttributes.ValidateAnt)
  * [func (t *TaskAttributes) ValidateExec() error](#TaskAttributes.ValidateExec)
* [type TaskPluginConfiguration](#TaskPluginConfiguration)
* [type Version](#Version)
* [type Versioned](#Versioned)


#### <a name="pkg-files">Package files</a>
[agent.go](/src/github.com/drewsonne/go-gocd/gocd/agent.go) [approval.go](/src/github.com/drewsonne/go-gocd/gocd/approval.go) [authentication.go](/src/github.com/drewsonne/go-gocd/gocd/authentication.go) [client.go](/src/github.com/drewsonne/go-gocd/gocd/client.go) [config.go](/src/github.com/drewsonne/go-gocd/gocd/config.go) [configuration.go](/src/github.com/drewsonne/go-gocd/gocd/configuration.go) [configuration_task.go](/src/github.com/drewsonne/go-gocd/gocd/configuration_task.go) [doc.go](/src/github.com/drewsonne/go-gocd/gocd/doc.go) [encryption.go](/src/github.com/drewsonne/go-gocd/gocd/encryption.go) [environment.go](/src/github.com/drewsonne/go-gocd/gocd/environment.go) [genericactions.go](/src/github.com/drewsonne/go-gocd/gocd/genericactions.go) [gocd.go](/src/github.com/drewsonne/go-gocd/gocd/gocd.go) [jobs.go](/src/github.com/drewsonne/go-gocd/gocd/jobs.go) [jobs_validation.go](/src/github.com/drewsonne/go-gocd/gocd/jobs_validation.go) [links.go](/src/github.com/drewsonne/go-gocd/gocd/links.go) [logging.go](/src/github.com/drewsonne/go-gocd/gocd/logging.go) [pipeline.go](/src/github.com/drewsonne/go-gocd/gocd/pipeline.go) [pipeline_material.go](/src/github.com/drewsonne/go-gocd/gocd/pipeline_material.go) [pipelineconfig.go](/src/github.com/drewsonne/go-gocd/gocd/pipelineconfig.go) [pipelinegroups.go](/src/github.com/drewsonne/go-gocd/gocd/pipelinegroups.go) [pipelinetemplate.go](/src/github.com/drewsonne/go-gocd/gocd/pipelinetemplate.go) [plugin.go](/src/github.com/drewsonne/go-gocd/gocd/plugin.go) [properties.go](/src/github.com/drewsonne/go-gocd/gocd/properties.go) [resource.go](/src/github.com/drewsonne/go-gocd/gocd/resource.go) [resource_agent.go](/src/github.com/drewsonne/go-gocd/gocd/resource_agent.go) [resource_approval.go](/src/github.com/drewsonne/go-gocd/gocd/resource_approval.go) [resource_environment.go](/src/github.com/drewsonne/go-gocd/gocd/resource_environment.go) [resource_jobs.go](/src/github.com/drewsonne/go-gocd/gocd/resource_jobs.go) [resource_pipeline.go](/src/github.com/drewsonne/go-gocd/gocd/resource_pipeline.go) [resource_pipeline_material.go](/src/github.com/drewsonne/go-gocd/gocd/resource_pipeline_material.go) [resource_pipeline_material_dependency.go](/src/github.com/drewsonne/go-gocd/gocd/resource_pipeline_material_dependency.go) [resource_pipeline_material_git.go](/src/github.com/drewsonne/go-gocd/gocd/resource_pipeline_material_git.go) [resource_pipeline_material_hg.go](/src/github.com/drewsonne/go-gocd/gocd/resource_pipeline_material_hg.go) [resource_pipeline_material_p4.go](/src/github.com/drewsonne/go-gocd/gocd/resource_pipeline_material_p4.go) [resource_pipeline_material_pkg.go](/src/github.com/drewsonne/go-gocd/gocd/resource_pipeline_material_pkg.go) [resource_pipeline_material_plugin.go](/src/github.com/drewsonne/go-gocd/gocd/resource_pipeline_material_plugin.go) [resource_pipeline_material_svn.go](/src/github.com/drewsonne/go-gocd/gocd/resource_pipeline_material_svn.go) [resource_pipeline_material_tfs.go](/src/github.com/drewsonne/go-gocd/gocd/resource_pipeline_material_tfs.go) [resource_pipelinegroups.go](/src/github.com/drewsonne/go-gocd/gocd/resource_pipelinegroups.go) [resource_pipelinetemplate.go](/src/github.com/drewsonne/go-gocd/gocd/resource_pipelinetemplate.go) [resource_properties.go](/src/github.com/drewsonne/go-gocd/gocd/resource_properties.go) [resource_stages.go](/src/github.com/drewsonne/go-gocd/gocd/resource_stages.go) [resource_task.go](/src/github.com/drewsonne/go-gocd/gocd/resource_task.go) [stages.go](/src/github.com/drewsonne/go-gocd/gocd/stages.go) 


## <a name="pkg-constants">Constants</a>
``` go
const (
    EnvVarDefaultProfile = "GOCD_DEFAULT_PROFILE"
    EnvVarServer         = "GOCD_SERVER"
    EnvVarUsername       = "GOCD_USERNAME"
    EnvVarPassword       = "GOCD_PASSWORD"
    EnvVarSkipSsl        = "GOCD_SKIP_SSL_CHECK"
)
```
Environment variables for configuration.

``` go
const (
    // JobStateTransitionPassed "Passed"
    JobStateTransitionPassed = "Passed"
    // JobStateTransitionScheduled "Scheduled"
    JobStateTransitionScheduled = "Scheduled"
)
```
``` go
const (
    LogLevelEnvVarName = "GOCD_LOG_LEVEL"
    LogLevelDefault    = "WARNING"
    LogTypeEnvVarName  = "GOCD_LOG_TYPE"
    LogTypeDefault     = "TEXT"
)
```
Set logging level and type constants

``` go
const ConfigDirectoryPath = "~/.gocd.conf"
```
ConfigDirectoryPath is the location where the authentication information is stored




## <a name="CheckResponse">func</a> [CheckResponse](/src/target/gocd.go?s=7605:7654#L309)
``` go
func CheckResponse(response *http.Response) error
```
CheckResponse asserts that the http response status code was 2xx.



## <a name="ConfigFilePath">func</a> [ConfigFilePath](/src/target/config.go?s=1981:2033#L86)
``` go
func ConfigFilePath() (configPath string, err error)
```
ConfigFilePath specifies the default path to a config file



## <a name="LoadConfigByName">func</a> [LoadConfigByName](/src/target/config.go?s=891:957#L34)
``` go
func LoadConfigByName(name string, cfg *Configuration) (err error)
```
LoadConfigByName loads configurations from yaml at default file location



## <a name="LoadConfigFromFile">func</a> [LoadConfigFromFile](/src/target/config.go?s=1516:1585#L64)
``` go
func LoadConfigFromFile() (cfgs map[string]*Configuration, err error)
```
LoadConfigFromFile on disk and return it as a Config item



## <a name="SetupLogging">func</a> [SetupLogging](/src/target/logging.go?s=885:922#L46)
``` go
func SetupLogging(log *logrus.Logger)
```
SetupLogging based on Environment Variables


	Set Logging level with $GOCD_LOG_LEVEL
	Allowed Values:
	  - DEBUG
	  - INFO
	  - WARNING
	  - ERROR
	  - FATAL
	  - PANIC
	
	Set Logging type  with $GOCD_LOG_TYPE
	Allowed Values:
	  - JSON
	  - TEXT




## <a name="APIClientRequest">type</a> [APIClientRequest](/src/target/genericactions.go?s=175:375#L14)
``` go
type APIClientRequest struct {
    Method       string
    Path         string
    APIVersion   string
    RequestBody  interface{}
    ResponseType string
    ResponseBody interface{}
    Headers      map[string]string
}
```
APIClientRequest helper struct to reduce amount of code.










## <a name="APIRequest">type</a> [APIRequest](/src/target/gocd.go?s=1392:1451#L63)
``` go
type APIRequest struct {
    HTTP *http.Request
    Body string
}
```
APIRequest encapsulates the net/http.Request object, and a string representing the Body.










## <a name="APIResponse">type</a> [APIResponse](/src/target/gocd.go?s=1210:1298#L56)
``` go
type APIResponse struct {
    HTTP    *http.Response
    Body    string
    Request *APIRequest
}
```
APIResponse encapsulates the net/http.Response object, a string representing the Body, and a gocd.Request object
encapsulating the response from the API.










## <a name="Agent">type</a> [Agent](/src/target/agent.go?s=443:1448#L20)
``` go
type Agent struct {
    UUID             string        `json:"uuid,omitempty"`
    Hostname         string        `json:"hostname,omitempty"`
    ElasticAgentID   string        `json:"elastic_agent_id,omitempty"`
    ElasticPluginID  string        `json:"elastic_plugin_id,omitempty"`
    IPAddress        string        `json:"ip_address,omitempty"`
    Sandbox          string        `json:"sandbox,omitempty"`
    OperatingSystem  string        `json:"operating_system,omitempty"`
    FreeSpace        int           `json:"free_space,omitempty"`
    AgentConfigState string        `json:"agent_config_state,omitempty"`
    AgentState       string        `json:"agent_state,omitempty"`
    Resources        []string      `json:"resources,omitempty"`
    Environments     []string      `json:"environments,omitempty"`
    BuildState       string        `json:"build_state,omitempty"`
    BuildDetails     *BuildDetails `json:"build_details,omitempty"`
    Links            *HALLinks     `json:"_links,omitempty,omitempty"`
    // contains filtered or unexported fields
}
```
Agent represents agent in GoCD.










### <a name="Agent.GetLinks">func</a> (\*Agent) [GetLinks](/src/target/resource_agent.go?s=54:90#L4)
``` go
func (a *Agent) GetLinks() *HALLinks
```
GetLinks returns HAL links for agent




### <a name="Agent.RemoveLinks">func</a> (\*Agent) [RemoveLinks](/src/target/resource_agent.go?s=210:239#L9)
``` go
func (a *Agent) RemoveLinks()
```
RemoveLinks sets the `Link` attribute as `nil`. Used when rendering an `Agent` struct to JSON.




## <a name="AgentBulkOperationUpdate">type</a> [AgentBulkOperationUpdate](/src/target/agent.go?s=2240:2363#L54)
``` go
type AgentBulkOperationUpdate struct {
    Add    []string `json:"add,omitempty"`
    Remove []string `json:"remove,omitempty"`
}
```
AgentBulkOperationUpdate describes an action to be performed on an Environment or Resource during an agent update.










## <a name="AgentBulkOperationsUpdate">type</a> [AgentBulkOperationsUpdate](/src/target/agent.go?s=1938:2120#L48)
``` go
type AgentBulkOperationsUpdate struct {
    Environments *AgentBulkOperationUpdate `json:"environments,omitempty"`
    Resources    *AgentBulkOperationUpdate `json:"resources,omitempty"`
}
```
AgentBulkOperationsUpdate describes the structure for a single Operation in AgentBulkUpdate the PUT payload when
updating multiple agents










## <a name="AgentBulkUpdate">type</a> [AgentBulkUpdate](/src/target/agent.go?s=1543:1792#L40)
``` go
type AgentBulkUpdate struct {
    Uuids            []string                   `json:"uuids"`
    Operations       *AgentBulkOperationsUpdate `json:"operations,omitempty"`
    AgentConfigState string                     `json:"agent_config_state,omitempty"`
}
```
AgentBulkUpdate describes the structure for the PUT payload when updating multiple agents










## <a name="AgentsResponse">type</a> [AgentsResponse](/src/target/agent.go?s=244:406#L12)
``` go
type AgentsResponse struct {
    Links    *HALLinks `json:"_links,omitempty"`
    Embedded *struct {
        Agents []*Agent `json:"agents"`
    } `json:"_embedded,omitempty"`
}
```
AgentsResponse describes the structure of the API response when listing collections of agent object.










## <a name="AgentsService">type</a> [AgentsService](/src/target/agent.go?s=112:138#L9)
``` go
type AgentsService service
```
AgentsService describes Actions which can be performed on agents










### <a name="AgentsService.BulkUpdate">func</a> (\*AgentsService) [BulkUpdate](/src/target/agent.go?s=3849:3958#L99)
``` go
func (s *AgentsService) BulkUpdate(ctx context.Context, agents AgentBulkUpdate) (string, *APIResponse, error)
```
BulkUpdate will change the configuration for multiple agents in a single request.




### <a name="AgentsService.Delete">func</a> (\*AgentsService) [Delete](/src/target/agent.go?s=3606:3700#L94)
``` go
func (s *AgentsService) Delete(ctx context.Context, uuid string) (string, *APIResponse, error)
```
Delete will remove an existing agent. Note: The agent must be disabled, and not currently building to be deleted.




### <a name="AgentsService.Get">func</a> (\*AgentsService) [Get](/src/target/agent.go?s=3106:3197#L84)
``` go
func (s *AgentsService) Get(ctx context.Context, uuid string) (*Agent, *APIResponse, error)
```
Get will retrieve a single agent based on the provided UUID.




### <a name="AgentsService.JobRunHistory">func</a> (\*AgentsService) [JobRunHistory](/src/target/agent.go?s=4238:4339#L111)
``` go
func (s *AgentsService) JobRunHistory(ctx context.Context, uuid string) ([]*Job, *APIResponse, error)
```
JobRunHistory will return a list of Jobs run on this agent.




### <a name="AgentsService.List">func</a> (\*AgentsService) [List](/src/target/agent.go?s=2688:2769#L68)
``` go
func (s *AgentsService) List(ctx context.Context) ([]*Agent, *APIResponse, error)
```
List will retrieve all agents, their status, and metadata from the GoCD Server.




### <a name="AgentsService.Update">func</a> (\*AgentsService) [Update](/src/target/agent.go?s=3319:3427#L89)
``` go
func (s *AgentsService) Update(ctx context.Context, uuid string, agent *Agent) (*Agent, *APIResponse, error)
```
Update will modify the configuration for an existing agents.




## <a name="Approval">type</a> [Approval](/src/target/approval.go?s=116:257#L4)
``` go
type Approval struct {
    Type          string         `json:"type,omitempty"`
    Authorization *Authorization `json:"authorization,omitempty"`
}
```
Approval represents a request/response object describing the approval configuration for a GoCD Job










### <a name="Approval.Clean">func</a> (\*Approval) [Clean](/src/target/resource_approval.go?s=113:139#L5)
``` go
func (a *Approval) Clean()
```
Clean ensures integrity of the schema by making sure
empty elements are not printed to json.




## <a name="Artifact">type</a> [Artifact](/src/target/jobs.go?s=2071:2206#L44)
``` go
type Artifact struct {
    Type        string `json:"type"`
    Source      string `json:"source"`
    Destination string `json:"destination"`
}
```
Artifact describes the result of a job










## <a name="Auth">type</a> [Auth](/src/target/gocd.go?s=2763:2817#L113)
``` go
type Auth struct {
    Username string
    Password string
}
```
Auth structure wrapping the Username and Password variables, which are used to get an Auth cookie header used for
subsequent requests.










## <a name="Authorization">type</a> [Authorization](/src/target/approval.go?s=432:543#L11)
``` go
type Authorization struct {
    Users []string `json:"users,omitempty"`
    Roles []string `json:"roles,omitempty"`
}
```
Authorization describes the access control for a "manual" approval type. Specifies whoe (role or users) can approve
the job to move to the next stage of the pipeline.










## <a name="BuildCause">type</a> [BuildCause](/src/target/pipeline.go?s=2791:3074#L77)
``` go
type BuildCause struct {
    Approver          string             `json:"approver,omitempty"`
    MaterialRevisions []MaterialRevision `json:"material_revisions"`
    TriggerForced     bool               `json:"trigger_forced"`
    TriggerMessage    string             `json:"trigger_message"`
}
```
BuildCause describes the triggers which caused the build to start.










## <a name="BuildDetails">type</a> [BuildDetails](/src/target/agent.go?s=2433:2603#L60)
``` go
type BuildDetails struct {
    Links    *HALLinks `json:"_links"`
    Pipeline string    `json:"pipeline"`
    Stage    string    `json:"stage"`
    Job      string    `json:"job"`
}
```
BuildDetails describes the builds being performed on this agent.










## <a name="CipherText">type</a> [CipherText](/src/target/encryption.go?s=247:366#L11)
``` go
type CipherText struct {
    EncryptedValue string    `json:"encrypted_value"`
    Links          *HALLinks `json:"_links"`
}
```
CipherText sescribes the response from the api with an encrypted value.










## <a name="Client">type</a> [Client](/src/target/gocd.go?s=1552:2264#L69)
``` go
type Client struct {
    BaseURL  *url.URL
    Username string
    Password string

    UserAgent string

    Log *logrus.Logger

    Agents            *AgentsService
    PipelineGroups    *PipelineGroupsService
    Stages            *StagesService
    Jobs              *JobsService
    PipelineTemplates *PipelineTemplatesService
    Pipelines         *PipelinesService
    PipelineConfigs   *PipelineConfigsService
    Configuration     *ConfigurationService
    Encryption        *EncryptionService
    Plugins           *PluginsService
    Environments      *EnvironmentsService
    Properties        *PropertiesService
    // contains filtered or unexported fields
}
```
Client struct which acts as an interface to the GoCD Server. Exposes resource service handlers.







### <a name="NewClient">func</a> [NewClient](/src/target/gocd.go?s=3317:3384#L130)
``` go
func NewClient(cfg *Configuration, httpClient *http.Client) *Client
```
NewClient creates a new client based on the provided configuration payload, and optionally a custom httpClient to
allow overriding of http client structures.





### <a name="Client.Do">func</a> (\*Client) [Do](/src/target/gocd.go?s=6426:6541#L256)
``` go
func (c *Client) Do(ctx context.Context, req *APIRequest, v interface{}, responseType string) (*APIResponse, error)
```
Do takes an HTTP request and resposne the response from the GoCD API endpoint.




### <a name="Client.Lock">func</a> (\*Client) [Lock](/src/target/gocd.go?s=4551:4574#L177)
``` go
func (c *Client) Lock()
```
Lock the client until release




### <a name="Client.Login">func</a> (\*Client) [Login](/src/target/authentication.go?s=167:216#L7)
``` go
func (c *Client) Login(ctx context.Context) error
```
Login sends basic auth to the GoCD Server and sets an auth cookie in the client to enable cookie based auth
for future requests.




### <a name="Client.NewRequest">func</a> (\*Client) [NewRequest](/src/target/gocd.go?s=4758:4874#L187)
``` go
func (c *Client) NewRequest(method, urlStr string, body interface{}, apiVersion string) (req *APIRequest, err error)
```
NewRequest creates an HTTP requests to the GoCD API endpoints.




### <a name="Client.Unlock">func</a> (\*Client) [Unlock](/src/target/gocd.go?s=4640:4665#L182)
``` go
func (c *Client) Unlock()
```
Unlock the client after a lock action




## <a name="ConfigApproval">type</a> [ConfigApproval](/src/target/configuration.go?s=2573:2662#L60)
``` go
type ConfigApproval struct {
    Type string `xml:"type,attr,omitempty" json:",omitempty"`
}
```
ConfigApproval part of cruise-control.xml. @TODO better documentation










## <a name="ConfigArtifact">type</a> [ConfigArtifact](/src/target/configuration.go?s=2365:2498#L54)
``` go
type ConfigArtifact struct {
    Src         string `xml:"src,attr"`
    Destination string `xml:"dest,attr,omitempty" json:",omitempty"`
}
```
ConfigArtifact part of cruise-control.xml. @TODO better documentation










## <a name="ConfigAuthConfig">type</a> [ConfigAuthConfig](/src/target/configuration.go?s=7269:7443#L182)
``` go
type ConfigAuthConfig struct {
    ID         string           `xml:"id,attr"`
    PluginID   string           `xml:"pluginId,attr"`
    Properties []ConfigProperty `xml:"property"`
}
```
ConfigAuthConfig part of cruise-control.xml. @TODO better documentation










## <a name="ConfigElastic">type</a> [ConfigElastic](/src/target/configuration.go?s=7517:7604#L189)
``` go
type ConfigElastic struct {
    Profiles []ConfigElasticProfile `xml:"profiles>profile"`
}
```
ConfigElastic part of cruise-control.xml. @TODO better documentation










## <a name="ConfigElasticProfile">type</a> [ConfigElasticProfile](/src/target/configuration.go?s=7685:7863#L194)
``` go
type ConfigElasticProfile struct {
    ID         string           `xml:"id,attr"`
    PluginID   string           `xml:"pluginId,attr"`
    Properties []ConfigProperty `xml:"property"`
}
```
ConfigElasticProfile part of cruise-control.xml. @TODO better documentation










## <a name="ConfigEnvironmentVariable">type</a> [ConfigEnvironmentVariable](/src/target/configuration.go?s=2748:2849#L65)
``` go
type ConfigEnvironmentVariable struct {
    Name  string `xml:"name,attr"`
    Value string `xml:"value"`
}
```
ConfigEnvironmentVariable part of cruise-control.xml. @TODO better documentation










## <a name="ConfigFilter">type</a> [ConfigFilter](/src/target/configuration.go?s=3385:3459#L84)
``` go
type ConfigFilter struct {
    Ignore string `xml:"pattern,attr,omitempty"`
}
```
ConfigFilter part of cruise-control.xml. @TODO better documentation










## <a name="ConfigJob">type</a> [ConfigJob](/src/target/configuration.go?s=1837:2290#L45)
``` go
type ConfigJob struct {
    Name                 string                      `xml:"name,attr"`
    EnvironmentVariables []ConfigEnvironmentVariable `xml:"environmentvariables>variable" json:",omitempty"`
    Tasks                ConfigTasks                 `xml:"tasks"`
    Resources            []string                    `xml:"resources>resource" json:",omitempty"`
    Artifacts            []ConfigArtifact            `xml:"artifacts>artifact" json:",omitempty"`
}
```
ConfigJob part of cruise-control.xml. @TODO better documentation










## <a name="ConfigMaterialRepository">type</a> [ConfigMaterialRepository](/src/target/configuration.go?s=4468:4861#L115)
``` go
type ConfigMaterialRepository struct {
    ID                  string                    `xml:"id,attr"`
    Name                string                    `xml:"name,attr"`
    PluginConfiguration ConfigPluginConfiguration `xml:"pluginConfiguration"`
    Configuration       []ConfigProperty          `xml:"configuration>property"`
    Packages            []ConfigPackage           `xml:"packages>package"`
}
```
ConfigMaterialRepository part of cruise-control.xml. @TODO better documentation










## <a name="ConfigPackage">type</a> [ConfigPackage](/src/target/configuration.go?s=4935:5125#L124)
``` go
type ConfigPackage struct {
    ID            string           `xml:"id,attr"`
    Name          string           `xml:"name,attr"`
    Configuration []ConfigProperty `xml:"configuration>property"`
}
```
ConfigPackage part of cruise-control.xml. @TODO better documentation










## <a name="ConfigParam">type</a> [ConfigParam](/src/target/configuration.go?s=3531:3622#L89)
``` go
type ConfigParam struct {
    Name  string `xml:"name,attr"`
    Value string `xml:",chardata"`
}
```
ConfigParam part of cruise-control.xml. @TODO better documentation










## <a name="ConfigPipeline">type</a> [ConfigPipeline](/src/target/configuration.go?s=882:1513#L26)
``` go
type ConfigPipeline struct {
    Name                 string                      `xml:"name,attr"`
    LabelTemplate        string                      `xml:"labeltemplate,attr"`
    Params               []ConfigParam               `xml:"params>param"`
    GitMaterials         []GitRepositoryMaterial     `xml:"materials>git,omitempty"`
    PipelineMaterials    []PipelineMaterial          `xml:"materials>pipeline,omitempty"`
    Timer                string                      `xml:"timer"`
    EnvironmentVariables []ConfigEnvironmentVariable `xml:"environmentvariables>variable"`
    Stages               []ConfigStage               `xml:"stage"`
}
```
ConfigPipeline part of cruise-control.xml. @TODO better documentation










## <a name="ConfigPipelineGroup">type</a> [ConfigPipelineGroup](/src/target/configuration.go?s=680:807#L20)
``` go
type ConfigPipelineGroup struct {
    Name      string           `xml:"group,attr"`
    Pipelines []ConfigPipeline `xml:"pipeline"`
}
```
ConfigPipelineGroup contains a single pipeline groups










## <a name="ConfigPluginConfiguration">type</a> [ConfigPluginConfiguration](/src/target/configuration.go?s=5211:5321#L131)
``` go
type ConfigPluginConfiguration struct {
    ID      string `xml:"id,attr"`
    Version string `xml:"version,attr"`
}
```
ConfigPluginConfiguration part of cruise-control.xml. @TODO better documentation










## <a name="ConfigProperty">type</a> [ConfigProperty](/src/target/configuration.go?s=7938:8022#L201)
``` go
type ConfigProperty struct {
    Key   string `xml:"key"`
    Value string `xml:"value"`
}
```
ConfigProperty part of cruise-control.xml. @TODO better documentation










## <a name="ConfigRepository">type</a> [ConfigRepository](/src/target/configuration.go?s=3699:3863#L95)
``` go
type ConfigRepository struct {
    Plugin string              `xml:"plugin,attr"`
    ID     string              `xml:"id,attr"`
    Git    ConfigRepositoryGit `xml:"git"`
}
```
ConfigRepository part of cruise-control.xml. @TODO better documentation










## <a name="ConfigRepositoryGit">type</a> [ConfigRepositoryGit](/src/target/configuration.go?s=3943:4007#L102)
``` go
type ConfigRepositoryGit struct {
    URL string `xml:"url,attr"`
}
```
ConfigRepositoryGit part of cruise-control.xml. @TODO better documentation










## <a name="ConfigRole">type</a> [ConfigRole](/src/target/configuration.go?s=7097:7192#L176)
``` go
type ConfigRole struct {
    Name  string   `xml:"name,attr"`
    Users []string `xml:"users>user"`
}
```
ConfigRole part of cruise-control.xml. @TODO better documentation










## <a name="ConfigSCM">type</a> [ConfigSCM](/src/target/configuration.go?s=4077:4383#L107)
``` go
type ConfigSCM struct {
    ID                  string                    `xml:"id,attr"`
    Name                string                    `xml:"name,attr"`
    PluginConfiguration ConfigPluginConfiguration `xml:"pluginConfiguration"`
    Configuration       []ConfigProperty          `xml:"configuration>property"`
}
```
ConfigSCM part of cruise-control.xml. @TODO better documentation










## <a name="ConfigSecurity">type</a> [ConfigSecurity](/src/target/configuration.go?s=6632:6885#L163)
``` go
type ConfigSecurity struct {
    AuthConfigs  []ConfigAuthConfig `xml:"authConfigs>authConfig"`
    Roles        []ConfigRole       `xml:"roles>role"`
    Admins       []string           `xml:"admins>user"`
    PasswordFile PasswordFilePath   `xml:"passwordFile"`
}
```
ConfigSecurity part of cruise-control.xml. @TODO better documentation










## <a name="ConfigServer">type</a> [ConfigServer](/src/target/configuration.go?s=5394:6285#L137)
``` go
type ConfigServer struct {
    MailHost                  MailHost       `xml:"mailhost"`
    Security                  ConfigSecurity `xml:"security"`
    Elastic                   ConfigElastic  `xml:"elastic"`
    ArtifactsDir              string         `xml:"artifactsdir,attr"`
    SiteURL                   string         `xml:"siteUrl,attr"`
    SecureSiteURL             string         `xml:"secureSiteUrl,attr"`
    PurgeStart                string         `xml:"purgeStart,attr"`
    PurgeUpTo                 string         `xml:"purgeUpto,attr"`
    JobTimeout                int            `xml:"jobTimeout,attr"`
    AgentAutoRegisterKey      string         `xml:"agentAutoRegisterKey,attr"`
    WebhookSecret             string         `xml:"webhookSecret,attr"`
    CommandRepositoryLocation string         `xml:"commandRepositoryLocation,attr"`
    ServerID                  string         `xml:"serverId,attr"`
}
```
ConfigServer part of cruise-control.xml. @TODO better documentation










## <a name="ConfigStage">type</a> [ConfigStage](/src/target/configuration.go?s=1585:1767#L38)
``` go
type ConfigStage struct {
    Name     string         `xml:"name,attr"`
    Approval ConfigApproval `xml:"approval,omitempty" json:",omitempty"`
    Jobs     []ConfigJob    `xml:"jobs>job"`
}
```
ConfigStage part of cruise-control.xml. @TODO better documentation










## <a name="ConfigTask">type</a> [ConfigTask](/src/target/configuration_task.go?s=243:1076#L13)
``` go
type ConfigTask struct {
    // Because we need to preserve the order of tasks, and we have an array of elements with mixed types,
    // we need to use this generic xml type for tasks.
    XMLName  xml.Name        `json:",omitempty"`
    Type     string          `xml:"type,omitempty"`
    RunIf    ConfigTaskRunIf `xml:"runif"`
    Command  string          `xml:"command,attr,omitempty"  json:",omitempty"`
    Args     []string        `xml:"arg,omitempty"  json:",omitempty"`
    Pipeline string          `xml:"pipeline,attr,omitempty"  json:",omitempty"`
    Stage    string          `xml:"stage,attr,omitempty"  json:",omitempty"`
    Job      string          `xml:"job,attr,omitempty"  json:",omitempty"`
    SrcFile  string          `xml:"srcfile,attr,omitempty"  json:",omitempty"`
    SrcDir   string          `xml:"srcdir,attr,omitempty"  json:",omitempty"`
}
```
ConfigTask part of cruise-control.xml. @TODO better documentation










## <a name="ConfigTaskRunIf">type</a> [ConfigTaskRunIf](/src/target/configuration_task.go?s=1152:1218#L29)
``` go
type ConfigTaskRunIf struct {
    Status string `xml:"status,attr"`
}
```
ConfigTaskRunIf part of cruise-control.xml. @TODO better documentation










## <a name="ConfigTasks">type</a> [ConfigTasks](/src/target/configuration_task.go?s=112:172#L8)
``` go
type ConfigTasks struct {
    Tasks []ConfigTask `xml:",any"`
}
```
ConfigTasks part of cruise-control.xml. @TODO better documentation










## <a name="ConfigXML">type</a> [ConfigXML](/src/target/configuration.go?s=246:621#L11)
``` go
type ConfigXML struct {
    Repositories       []ConfigMaterialRepository `xml:"repositories>repository"`
    Server             ConfigServer               `xml:"server"`
    SCMS               []ConfigSCM                `xml:"scms>scm"`
    ConfigRepositories []ConfigRepository         `xml:"config-repos>config-repo"`
    PipelineGroups     []ConfigPipelineGroup      `xml:"pipelines"`
}
```
ConfigXML part of cruise-control.xml. @TODO better documentation










## <a name="Configuration">type</a> [Configuration](/src/target/config.go?s=586:813#L26)
``` go
type Configuration struct {
    Server       string
    Username     string `yaml:"username,omitempty"`
    Password     string `yaml:"password,omitempty"`
    SkipSslCheck bool   `yaml:"skip_ssl_check,omitempty" survey:"skip_ssl_check"`
}
```
Configuration object used to initialise a gocd lib client to interact with the GoCD server.










### <a name="Configuration.Client">func</a> (\*Configuration) [Client](/src/target/gocd.go?s=3081:3121#L124)
``` go
func (c *Configuration) Client() *Client
```
Client returns a client which allows us to interact with the GoCD Server.




### <a name="Configuration.HasAuth">func</a> (\*Configuration) [HasAuth](/src/target/gocd.go?s=2911:2949#L119)
``` go
func (c *Configuration) HasAuth() bool
```
HasAuth checks whether or not we have the required Username/Password variables provided.




## <a name="ConfigurationService">type</a> [ConfigurationService](/src/target/configuration.go?s=143:176#L8)
``` go
type ConfigurationService service
```
ConfigurationService describes the HAL _link resource for the api response object for a pipelineconfig










### <a name="ConfigurationService.Get">func</a> (\*ConfigurationService) [Get](/src/target/configuration.go?s=8526:8616#L218)
``` go
func (cs *ConfigurationService) Get(ctx context.Context) (*ConfigXML, *APIResponse, error)
```
Get will retrieve all agents, their status, and metadata from the GoCD Server.
Get returns a list of pipeline instanves describing the pipeline history.




### <a name="ConfigurationService.GetVersion">func</a> (\*ConfigurationService) [GetVersion](/src/target/configuration.go?s=8888:8983#L229)
``` go
func (cs *ConfigurationService) GetVersion(ctx context.Context) (*Version, *APIResponse, error)
```
GetVersion will return version information about the GoCD server.




## <a name="EmbeddedEnvironments">type</a> [EmbeddedEnvironments](/src/target/environment.go?s=440:527#L17)
``` go
type EmbeddedEnvironments struct {
    Environments []*Environment `json:"environments"`
}
```
EmbeddedEnvironments encapsulates the environment struct










## <a name="EncryptionService">type</a> [EncryptionService](/src/target/encryption.go?s=140:170#L8)
``` go
type EncryptionService service
```
EncryptionService describes the HAL _link resource for the api response object for a pipelineconfig










### <a name="EncryptionService.Encrypt">func</a> (\*EncryptionService) [Encrypt](/src/target/encryption.go?s=430:540#L17)
``` go
func (es *EncryptionService) Encrypt(ctx context.Context, plaintext string) (*CipherText, *APIResponse, error)
```
Encrypt takes a plaintext value and returns a cipher text.




## <a name="Environment">type</a> [Environment](/src/target/environment.go?s=586:1036#L22)
``` go
type Environment struct {
    Links                *HALLinks              `json:"_links,omitempty"`
    Name                 string                 `json:"name"`
    Pipelines            []*Pipeline            `json:"pipelines,omitempty"`
    Agents               []*Agent               `json:"agents,omitempty"`
    EnvironmentVariables []*EnvironmentVariable `json:"environment_variables,omitempty"`
    Version              string                 `json:"version"`
}
```
Environment describes a group of pipelines and agents










### <a name="Environment.GetLinks">func</a> (\*Environment) [GetLinks](/src/target/resource_environment.go?s=629:673#L28)
``` go
func (env *Environment) GetLinks() *HALLinks
```
GetLinks from the Environment




### <a name="Environment.GetVersion">func</a> (\*Environment) [GetVersion](/src/target/resource_environment.go?s=889:942#L38)
``` go
func (env *Environment) GetVersion() (version string)
```
GetVersion retrieves a version string for this pipeline




### <a name="Environment.RemoveLinks">func</a> (\*Environment) [RemoveLinks](/src/target/resource_environment.go?s=427:464#L17)
``` go
func (env *Environment) RemoveLinks()
```
RemoveLinks gets the Environment ready to be submitted to the GoCD API.




### <a name="Environment.SetVersion">func</a> (\*Environment) [SetVersion](/src/target/resource_environment.go?s=751:801#L33)
``` go
func (env *Environment) SetVersion(version string)
```
SetVersion sets a version string for this pipeline




## <a name="EnvironmentPatchRequest">type</a> [EnvironmentPatchRequest](/src/target/environment.go?s=1116:1371#L32)
``` go
type EnvironmentPatchRequest struct {
    Pipelines            *PatchStringAction          `json:"pipelines"`
    Agents               *PatchStringAction          `json:"agents"`
    EnvironmentVariables *EnvironmentVariablesAction `json:"environment_variables"`
}
```
EnvironmentPatchRequest describes the actions to perform on an environment










## <a name="EnvironmentVariable">type</a> [EnvironmentVariable](/src/target/jobs.go?s=2550:2767#L64)
``` go
type EnvironmentVariable struct {
    Name           string `json:"name"`
    Value          string `json:"value,omitempty"`
    EncryptedValue string `json:"encrypted_value,omitempty"`
    Secure         bool   `json:"secure"`
}
```
EnvironmentVariable describes an environment variable key/pair.










## <a name="EnvironmentVariablesAction">type</a> [EnvironmentVariablesAction](/src/target/environment.go?s=1469:1602#L39)
``` go
type EnvironmentVariablesAction struct {
    Add    []*EnvironmentVariable `json:"add"`
    Remove []*EnvironmentVariable `json:"remove"`
}
```
EnvironmentVariablesAction describes a collection of Environment Variables to add or remove.










## <a name="EnvironmentsResponse">type</a> [EnvironmentsResponse](/src/target/environment.go?s=243:378#L11)
``` go
type EnvironmentsResponse struct {
    Links    *HALLinks             `json:"_links"`
    Embedded *EmbeddedEnvironments `json:"_embedded"`
}
```
EnvironmentsResponse describes the response obejct for a plugin API call.










### <a name="EnvironmentsResponse.GetLinks">func</a> (\*EnvironmentsResponse) [GetLinks](/src/target/resource_environment.go?s=277:329#L12)
``` go
func (er *EnvironmentsResponse) GetLinks() *HALLinks
```
GetLinks from the EnvironmentResponse




### <a name="EnvironmentsResponse.RemoveLinks">func</a> (\*EnvironmentsResponse) [RemoveLinks](/src/target/resource_environment.go?s=98:143#L4)
``` go
func (er *EnvironmentsResponse) RemoveLinks()
```
RemoveLinks gets the EnvironmentsResponse ready to be submitted to the GoCD API.




## <a name="EnvironmentsService">type</a> [EnvironmentsService](/src/target/environment.go?s=132:164#L8)
``` go
type EnvironmentsService service
```
EnvironmentsService exposes calls for interacting with Environment objects in the GoCD API.










### <a name="EnvironmentsService.Create">func</a> (\*EnvironmentsService) [Create](/src/target/environment.go?s=2335:2442#L68)
``` go
func (es *EnvironmentsService) Create(ctx context.Context, name string) (*Environment, *APIResponse, error)
```
Create an environment




### <a name="EnvironmentsService.Delete">func</a> (\*EnvironmentsService) [Delete](/src/target/environment.go?s=2132:2233#L63)
``` go
func (es *EnvironmentsService) Delete(ctx context.Context, name string) (string, *APIResponse, error)
```
Delete an environment




### <a name="EnvironmentsService.Get">func</a> (\*EnvironmentsService) [Get](/src/target/environment.go?s=2715:2819#L83)
``` go
func (es *EnvironmentsService) Get(ctx context.Context, name string) (*Environment, *APIResponse, error)
```
Get a single environment by name




### <a name="EnvironmentsService.List">func</a> (\*EnvironmentsService) [List](/src/target/environment.go?s=1802:1903#L51)
``` go
func (es *EnvironmentsService) List(ctx context.Context) (*EnvironmentsResponse, *APIResponse, error)
```
List all environments




### <a name="EnvironmentsService.Patch">func</a> (\*EnvironmentsService) [Patch](/src/target/environment.go?s=3124:3262#L95)
``` go
func (es *EnvironmentsService) Patch(ctx context.Context, name string, patch *EnvironmentPatchRequest) (*Environment, *APIResponse, error)
```
Patch an environments configuration by adding or removing pipelines, agents, environment variables




## <a name="GitRepositoryMaterial">type</a> [GitRepositoryMaterial](/src/target/configuration.go?s=3178:3312#L78)
``` go
type GitRepositoryMaterial struct {
    URL     string         `xml:"url,attr"`
    Filters []ConfigFilter `xml:"filter>ignore,omitempty"`
}
```
GitRepositoryMaterial part of cruise-control.xml. @TODO better documentation










## <a name="HALContainer">type</a> [HALContainer](/src/target/resource.go?s=387:455#L17)
``` go
type HALContainer interface {
    RemoveLinks()
    GetLinks() *HALLinks
}
```
HALContainer represents objects with HAL _link and _embedded resources.










## <a name="HALLink">type</a> [HALLink](/src/target/resource.go?s=632:718#L29)
``` go
type HALLink struct {
    Name string
    URL  *url.URL
}
```
HALLink describes a HAL link










## <a name="HALLinks">type</a> [HALLinks](/src/target/links.go?s=206:248#L15)
``` go
type HALLinks struct {
    // contains filtered or unexported fields
}
```
HALLinks describes a collection of HALLinks










### <a name="HALLinks.Add">func</a> (\*HALLinks) [Add](/src/target/links.go?s=264:302#L20)
``` go
func (al *HALLinks) Add(link *HALLink)
```
Add a link




### <a name="HALLinks.Get">func</a> (HALLinks) [Get](/src/target/links.go?s=368:412#L25)
``` go
func (al HALLinks) Get(name string) *HALLink
```
Get a HALLink by name




### <a name="HALLinks.GetOk">func</a> (HALLinks) [GetOk](/src/target/links.go?s=524:578#L31)
``` go
func (al HALLinks) GetOk(name string) (*HALLink, bool)
```
GetOk a HALLink by name, and if it doesn't exist, return false




### <a name="HALLinks.Keys">func</a> (HALLinks) [Keys](/src/target/links.go?s=767:801#L41)
``` go
func (al HALLinks) Keys() []string
```
Keys returns a string list of link names




### <a name="HALLinks.MarshallJSON">func</a> (HALLinks) [MarshallJSON](/src/target/links.go?s=966:1015#L50)
``` go
func (al HALLinks) MarshallJSON() ([]byte, error)
```
MarshallJSON allows the encoding of links into JSON




### <a name="HALLinks.UnmarshalJSON">func</a> (\*HALLinks) [UnmarshalJSON](/src/target/links.go?s=1218:1271#L59)
``` go
func (al *HALLinks) UnmarshalJSON(j []byte) (e error)
```
UnmarshalJSON allows the decoding of links from JSON




## <a name="Job">type</a> [Job](/src/target/jobs.go?s=353:2027#L18)
``` go
type Job struct {
    AgentUUID            string                 `json:"agent_uuid,omitempty"`
    Name                 string                 `json:"name"`
    JobStateTransitions  []*JobStateTransition  `json:"job_state_transitions,omitempty"`
    ScheduledDate        int                    `json:"scheduled_date,omitempty"`
    OriginalJobID        string                 `json:"original_job_id,omitempty"`
    PipelineCounter      int                    `json:"pipeline_counter,omitempty"`
    Rerun                bool                   `json:"rerun,omitempty"`
    PipelineName         string                 `json:"pipeline_name,omitempty"`
    Result               string                 `json:"result,omitempty"`
    State                string                 `json:"state,omitempty"`
    ID                   int                    `json:"id,omitempty"`
    StageCounter         string                 `json:"stage_counter,omitempty"`
    StageName            string                 `json:"stage_name,omitempty"`
    RunInstanceCount     int                    `json:"run_instance_count,omitempty"`
    Timeout              int                    `json:"timeout,omitempty"`
    EnvironmentVariables []*EnvironmentVariable `json:"environment_variables,omitempty"`
    Properties           []*JobProperty         `json:"properties,omitempty"`
    Resources            []string               `json:"resources,omitempty"`
    Tasks                []*Task                `json:"tasks,omitempty"`
    Tabs                 []*Tab                 `json:"tabs,omitempty"`
    Artifacts            []*Artifact            `json:"artifacts,omitempty"`
    ElasticProfileID     string                 `json:"elastic_profile_id,omitempty"`
}
```
Job describes a job object.










### <a name="Job.JSONString">func</a> (\*Job) [JSONString](/src/target/resource_jobs.go?s=116:158#L9)
``` go
func (j *Job) JSONString() (string, error)
```
JSONString returns a string of this stage as a JSON object.




### <a name="Job.Validate">func</a> (\*Job) [Validate](/src/target/resource_jobs.go?s=362:392#L20)
``` go
func (j *Job) Validate() error
```
Validate a job structure has non-nil values on correct attributes




## <a name="JobProperty">type</a> [JobProperty](/src/target/jobs.go?s=2364:2481#L57)
``` go
type JobProperty struct {
    Name   string `json:"name"`
    Source string `json:"source"`
    XPath  string `json:"xpath"`
}
```
JobProperty describes the property for a job










## <a name="JobRunHistoryResponse">type</a> [JobRunHistoryResponse](/src/target/jobs.go?s=5356:5511#L129)
``` go
type JobRunHistoryResponse struct {
    Jobs       []*Job              `json:"jobs,omitempty"`
    Pagination *PaginationResponse `json:"pagination,omitempty"`
}
```
JobRunHistoryResponse describes the api response from










## <a name="JobSchedule">type</a> [JobSchedule](/src/target/jobs.go?s=5688:6116#L140)
``` go
type JobSchedule struct {
    Name                 string               `xml:"name,attr"`
    ID                   string               `xml:"id,attr"`
    Link                 JobScheduleLink      `xml:"link"`
    BuildLocator         string               `xml:"buildLocator"`
    Resources            []string             `xml:"resources>resource"`
    EnvironmentVariables *[]JobScheduleEnvVar `xml:"environmentVariables,omitempty>variable"`
}
```
JobSchedule describes the event causes for a job










## <a name="JobScheduleEnvVar">type</a> [JobScheduleEnvVar](/src/target/jobs.go?s=6194:6291#L150)
``` go
type JobScheduleEnvVar struct {
    Name  string `xml:"name,attr"`
    Value string `xml:",innerxml"`
}
```
JobScheduleEnvVar describes the environmnet variables for a job schedule










## <a name="JobScheduleLink">type</a> [JobScheduleLink](/src/target/jobs.go?s=6355:6447#L156)
``` go
type JobScheduleLink struct {
    Rel  string `xml:"rel,attr"`
    HRef string `xml:"href,attr"`
}
```
JobScheduleLink describes the HAL links for a job schedule










## <a name="JobScheduleResponse">type</a> [JobScheduleResponse](/src/target/jobs.go?s=5566:5634#L135)
``` go
type JobScheduleResponse struct {
    Jobs []*JobSchedule `xml:"job"`
}
```
JobScheduleResponse contains a collection of jobs










## <a name="JobStateTransition">type</a> [JobStateTransition](/src/target/jobs.go?s=5107:5297#L122)
``` go
type JobStateTransition struct {
    StateChangeTime int    `json:"state_change_time,omitempty"`
    ID              int    `json:"id,omitempty"`
    State           string `json:"state,omitempty"`
}
```
JobStateTransition describes a State Transition object in a GoCD api response










## <a name="JobsService">type</a> [JobsService](/src/target/jobs.go?s=296:320#L15)
``` go
type JobsService service
```
JobsService describes the HAL _link resource for the api response object for a job










### <a name="JobsService.ListScheduled">func</a> (\*JobsService) [ListScheduled](/src/target/jobs.go?s=6488:6583#L162)
``` go
func (js *JobsService) ListScheduled(ctx context.Context) ([]*JobSchedule, *APIResponse, error)
```
ListScheduled lists Pipeline groups




## <a name="MailHost">type</a> [MailHost](/src/target/configuration.go?s=6354:6557#L154)
``` go
type MailHost struct {
    Hostname string `xml:"hostname,attr"`
    Port     int    `xml:"port,attr"`
    TLS      bool   `xml:"tls,attr"`
    From     string `xml:"from,attr"`
    Admin    string `xml:"admin,attr"`
}
```
MailHost part of cruise-control.xml. @TODO better documentation










## <a name="Material">type</a> [Material](/src/target/pipeline.go?s=1869:2113#L49)
``` go
type Material struct {
    Type        string            `json:"type"`
    Fingerprint string            `json:"fingerprint,omitempty"`
    Description string            `json:"description,omitempty"`
    Attributes  MaterialAttribute `json:"attributes"`
}
```
Material describes an artifact dependency for a pipeline object.










### <a name="Material.Equal">func</a> (Material) [Equal](/src/target/resource_pipeline_material.go?s=158:220#L10)
``` go
func (m Material) Equal(a *Material) (isEqual bool, err error)
```
Equal is true if the two materials are logically equivalent. Not neccesarily literally equal.




### <a name="Material.UnmarshalJSON">func</a> (\*Material) [UnmarshalJSON](/src/target/resource_pipeline_material.go?s=367:415#L21)
``` go
func (m *Material) UnmarshalJSON(b []byte) error
```
UnmarshalJSON string into a Material struct




## <a name="MaterialAttribute">type</a> [MaterialAttribute](/src/target/pipeline_material.go?s=106:193#L4)
``` go
type MaterialAttribute interface {
    // contains filtered or unexported methods
}
```
MaterialAttribute describes the behaviour of the GoCD material structures for a pipeline










## <a name="MaterialAttributesDependency">type</a> [MaterialAttributesDependency](/src/target/pipeline_material.go?s=3021:3219#L86)
``` go
type MaterialAttributesDependency struct {
    Name       string `json:"name"`
    Pipeline   string `json:"pipeline"`
    Stage      string `json:"stage"`
    AutoUpdate bool   `json:"auto_update,omitempty"`
}
```
MaterialAttributesDependency describes a Pipeline dependency material










## <a name="MaterialAttributesGit">type</a> [MaterialAttributesGit](/src/target/pipeline_material.go?s=245:750#L9)
``` go
type MaterialAttributesGit struct {
    Name   string `json:"name,omitempty"`
    URL    string `json:"url,omitempty"`
    Branch string `json:"branch,omitempty"`

    SubmoduleFolder string `json:"submodule_folder,omitempty"`
    ShallowClone    bool   `json:"shallow_clone,omitempty"`

    Destination  string          `json:"destination,omitempty"`
    Filter       *MaterialFilter `json:"filter,omitempty"`
    InvertFilter bool            `json:"invert_filter"`
    AutoUpdate   bool            `json:"auto_update,omitempty"`
}
```
MaterialAttributesGit describes a git material










## <a name="MaterialAttributesHg">type</a> [MaterialAttributesHg](/src/target/pipeline_material.go?s=1422:1733#L40)
``` go
type MaterialAttributesHg struct {
    Name string `json:"name"`
    URL  string `json:"url"`

    Destination  string          `json:"destination"`
    Filter       *MaterialFilter `json:"filter,omitempty"`
    InvertFilter bool            `json:"invert_filter"`
    AutoUpdate   bool            `json:"auto_update,omitempty"`
}
```
MaterialAttributesHg describes a Mercurial material type










## <a name="MaterialAttributesP4">type</a> [MaterialAttributesP4](/src/target/pipeline_material.go?s=1794:2334#L51)
``` go
type MaterialAttributesP4 struct {
    Name       string `json:"name"`
    Port       string `json:"port"`
    UseTickets bool   `json:"use_tickets"`
    View       string `json:"view"`

    Username          string `json:"username"`
    Password          string `json:"password"`
    EncryptedPassword string `json:"encrypted_password"`

    Destination  string          `json:"destination"`
    Filter       *MaterialFilter `json:"filter,omitempty"`
    InvertFilter bool            `json:"invert_filter"`
    AutoUpdate   bool            `json:"auto_update,omitempty"`
}
```
MaterialAttributesP4 describes a Perforce material type










## <a name="MaterialAttributesPackage">type</a> [MaterialAttributesPackage](/src/target/pipeline_material.go?s=3280:3346#L94)
``` go
type MaterialAttributesPackage struct {
    Ref string `json:"ref"`
}
```
MaterialAttributesPackage describes a package reference










## <a name="MaterialAttributesPlugin">type</a> [MaterialAttributesPlugin](/src/target/pipeline_material.go?s=3404:3630#L99)
``` go
type MaterialAttributesPlugin struct {
    Ref string `json:"ref"`

    Destination  string          `json:"destination"`
    Filter       *MaterialFilter `json:"filter,omitempty"`
    InvertFilter bool            `json:"invert_filter"`
}
```
MaterialAttributesPlugin describes a plugin material










## <a name="MaterialAttributesSvn">type</a> [MaterialAttributesSvn](/src/target/pipeline_material.go?s=803:1360#L24)
``` go
type MaterialAttributesSvn struct {
    Name              string `json:"name,omitempty"`
    URL               string `json:"url,omitempty"`
    Username          string `json:"username"`
    Password          string `json:"password"`
    EncryptedPassword string `json:"encrypted_password"`

    CheckExternals bool `json:"check_externals"`

    Destination  string          `json:"destination,omitempty"`
    Filter       *MaterialFilter `json:"filter,omitempty"`
    InvertFilter bool            `json:"invert_filter"`
    AutoUpdate   bool            `json:"auto_update,omitempty"`
}
```
MaterialAttributesSvn describes a material type










## <a name="MaterialAttributesTfs">type</a> [MaterialAttributesTfs](/src/target/pipeline_material.go?s=2405:2946#L68)
``` go
type MaterialAttributesTfs struct {
    Name string `json:"name"`

    URL         string `json:"url"`
    ProjectPath string `json:"project_path"`
    Domain      string `json:"domain"`

    Username          string `json:"username"`
    Password          string `json:"password"`
    EncryptedPassword string `json:"encrypted_password"`

    Destination  string          `json:"destination"`
    Filter       *MaterialFilter `json:"filter,omitempty"`
    InvertFilter bool            `json:"invert_filter"`
    AutoUpdate   bool            `json:"auto_update,omitempty"`
}
```
MaterialAttributesTfs describes a Team Foundation Server material










## <a name="MaterialFilter">type</a> [MaterialFilter](/src/target/pipeline.go?s=2165:2228#L57)
``` go
type MaterialFilter struct {
    Ignore []string `json:"ignore"`
}
```
MaterialFilter describes which globs to ignore










## <a name="MaterialRevision">type</a> [MaterialRevision](/src/target/pipeline.go?s=3189:3502#L85)
``` go
type MaterialRevision struct {
    Modifications []Modification `json:"modifications"`
    Material      struct {
        Description string `json:"description"`
        Fingerprint string `json:"fingerprint"`
        Type        string `json:"type"`
        ID          int    `json:"id"`
    } `json:"material"`
    Changed bool `json:"changed"`
}
```
MaterialRevision describes the uniquely identifiable version for the material which was pulled for this build










## <a name="Modification">type</a> [Modification](/src/target/pipeline.go?s=3595:3861#L97)
``` go
type Modification struct {
    EmailAddress string `json:"email_address"`
    ID           int    `json:"id"`
    ModifiedTime int    `json:"modified_time"`
    UserName     string `json:"user_name"`
    Comment      string `json:"comment"`
    Revision     string `json:"revision"`
}
```
Modification describes the commit/revision for the material which kicked off the build.










## <a name="PaginationResponse">type</a> [PaginationResponse](/src/target/gocd.go?s=2341:2467#L99)
``` go
type PaginationResponse struct {
    Offset   int `json:"offset"`
    Total    int `json:"total"`
    PageSize int `json:"page_size"`
}
```
PaginationResponse is a struct used to handle paging through resposnes.










## <a name="Parameter">type</a> [Parameter](/src/target/pipeline.go?s=1546:1628#L37)
``` go
type Parameter struct {
    Name  string `json:"name"`
    Value string `json:"value"`
}
```
Parameter represents a key/value










## <a name="PasswordFilePath">type</a> [PasswordFilePath](/src/target/configuration.go?s=6963:7026#L171)
``` go
type PasswordFilePath struct {
    Path string `xml:"path,attr"`
}
```
PasswordFilePath describes the location to set of user/passwords on disk










## <a name="PatchStringAction">type</a> [PatchStringAction](/src/target/environment.go?s=1679:1775#L45)
``` go
type PatchStringAction struct {
    Add    []string `json:"add"`
    Remove []string `json:"remove"`
}
```
PatchStringAction describes a collection of resources to add or remove.










## <a name="Pipeline">type</a> [Pipeline](/src/target/pipeline.go?s=378:1508#L18)
``` go
type Pipeline struct {
    Group                 string                 `json:"group,omitempty"`
    Links                 *HALLinks              `json:"_links,omitempty"`
    Name                  string                 `json:"name"`
    LabelTemplate         string                 `json:"label_template,omitempty"`
    EnablePipelineLocking bool                   `json:"enable_pipeline_locking,omitempty"`
    Template              string                 `json:"template,omitempty"`
    Origin                *PipelineConfigOrigin  `json:"origin,omitempty"`
    Parameters            []*Parameter           `json:"parameters,omitempty"`
    EnvironmentVariables  []*EnvironmentVariable `json:"environment_variables,omitempty"`
    Materials             []Material             `json:"materials,omitempty"`
    Label                 string                 `json:"label,omitempty"`
    Stages                []*Stage               `json:"stages,omitempty"`
    Version               string                 `json:"version,omitempty"`
}
```
Pipeline describes a pipeline object










### <a name="Pipeline.AddStage">func</a> (\*Pipeline) [AddStage](/src/target/resource_pipeline.go?s=748:789#L39)
``` go
func (p *Pipeline) AddStage(stage *Stage)
```
AddStage appends a stage to this pipeline




### <a name="Pipeline.GetLinks">func</a> (\*Pipeline) [GetLinks](/src/target/resource_pipeline.go?s=445:484#L24)
``` go
func (p *Pipeline) GetLinks() *HALLinks
```
GetLinks from pipeline




### <a name="Pipeline.GetName">func</a> (\*Pipeline) [GetName](/src/target/resource_pipeline.go?s=533:568#L29)
``` go
func (p *Pipeline) GetName() string
```
GetName of the pipeline




### <a name="Pipeline.GetStage">func</a> (\*Pipeline) [GetStage](/src/target/resource_pipeline.go?s=146:198#L9)
``` go
func (p *Pipeline) GetStage(stageName string) *Stage
```
GetStage from the pipeline template




### <a name="Pipeline.GetStages">func</a> (\*Pipeline) [GetStages](/src/target/resource_pipeline.go?s=45:84#L4)
``` go
func (p *Pipeline) GetStages() []*Stage
```
GetStages from the pipeline




### <a name="Pipeline.GetVersion">func</a> (\*Pipeline) [GetVersion](/src/target/resource_pipeline.go?s=1250:1298#L60)
``` go
func (p *Pipeline) GetVersion() (version string)
```
GetVersion retrieves a version string for this pipeline




### <a name="Pipeline.RemoveLinks">func</a> (\*Pipeline) [RemoveLinks](/src/target/resource_pipeline.go?s=366:398#L19)
``` go
func (p *Pipeline) RemoveLinks()
```
RemoveLinks from the pipeline object for json marshalling.




### <a name="Pipeline.SetStage">func</a> (\*Pipeline) [SetStage](/src/target/resource_pipeline.go?s=881:925#L44)
``` go
func (p *Pipeline) SetStage(newStage *Stage)
```
SetStage replaces a stage if it already exists




### <a name="Pipeline.SetStages">func</a> (\*Pipeline) [SetStages](/src/target/resource_pipeline.go?s=633:678#L34)
``` go
func (p *Pipeline) SetStages(stages []*Stage)
```
SetStages overwrites any existing stages




### <a name="Pipeline.SetVersion">func</a> (\*Pipeline) [SetVersion](/src/target/resource_pipeline.go?s=1119:1164#L55)
``` go
func (p *Pipeline) SetVersion(version string)
```
SetVersion sets a version string for this pipeline




## <a name="PipelineConfigOrigin">type</a> [PipelineConfigOrigin](/src/target/pipeline.go?s=1709:1799#L43)
``` go
type PipelineConfigOrigin struct {
    Type string `json:"type"`
    File string `json:"file"`
}
```
PipelineConfigOrigin describes where a pipeline config is being loaded from










## <a name="PipelineConfigRequest">type</a> [PipelineConfigRequest](/src/target/pipelineconfig.go?s=278:398#L12)
``` go
type PipelineConfigRequest struct {
    Group    string    `json:"group,omitempty"`
    Pipeline *Pipeline `json:"pipeline"`
}
```
PipelineConfigRequest describes a request object for creating or updating pipelines










### <a name="PipelineConfigRequest.GetVersion">func</a> (\*PipelineConfigRequest) [GetVersion](/src/target/resource_pipeline.go?s=1355:1417#L65)
``` go
func (pr *PipelineConfigRequest) GetVersion() (version string)
```
GetVersion of pipeline config




### <a name="PipelineConfigRequest.SetVersion">func</a> (\*PipelineConfigRequest) [SetVersion](/src/target/resource_pipeline.go?s=1489:1548#L70)
``` go
func (pr *PipelineConfigRequest) SetVersion(version string)
```
SetVersion of pipeline config




## <a name="PipelineConfigsService">type</a> [PipelineConfigsService](/src/target/pipelineconfig.go?s=154:189#L9)
``` go
type PipelineConfigsService service
```
PipelineConfigsService describes the HAL _link resource for the api response object for a pipelineconfig










### <a name="PipelineConfigsService.Create">func</a> (\*PipelineConfigsService) [Create](/src/target/pipelineconfig.go?s=1209:1331#L46)
``` go
func (pcs *PipelineConfigsService) Create(ctx context.Context, group string, p *Pipeline) (*Pipeline, *APIResponse, error)
```
Create a pipeline configuration




### <a name="PipelineConfigsService.Delete">func</a> (\*PipelineConfigsService) [Delete](/src/target/pipelineconfig.go?s=1637:1742#L63)
``` go
func (pcs *PipelineConfigsService) Delete(ctx context.Context, name string) (string, *APIResponse, error)
```
Delete a pipeline configuration




### <a name="PipelineConfigsService.Get">func</a> (\*PipelineConfigsService) [Get](/src/target/pipelineconfig.go?s=457:562#L18)
``` go
func (pcs *PipelineConfigsService) Get(ctx context.Context, name string) (*Pipeline, *APIResponse, error)
```
Get a single PipelineTemplate object in the GoCD API.




### <a name="PipelineConfigsService.Update">func</a> (\*PipelineConfigsService) [Update](/src/target/pipelineconfig.go?s=795:916#L30)
``` go
func (pcs *PipelineConfigsService) Update(ctx context.Context, name string, p *Pipeline) (*Pipeline, *APIResponse, error)
```
Update a pipeline configuration




## <a name="PipelineGroup">type</a> [PipelineGroup](/src/target/pipelinegroups.go?s=342:450#L12)
``` go
type PipelineGroup struct {
    Name      string      `json:"name"`
    Pipelines []*Pipeline `json:"pipelines"`
}
```
PipelineGroup describes a pipeline group API response.










## <a name="PipelineGroups">type</a> [PipelineGroups](/src/target/pipelinegroups.go?s=246:282#L9)
``` go
type PipelineGroups []*PipelineGroup
```
PipelineGroups represents a collection of pipeline groups










### <a name="PipelineGroups.GetGroupByPipeline">func</a> (\*PipelineGroups) [GetGroupByPipeline](/src/target/resource_pipelinegroups.go?s=443:522#L16)
``` go
func (pg *PipelineGroups) GetGroupByPipeline(pipeline *Pipeline) *PipelineGroup
```
GetGroupByPipeline finds the pipeline group for the pipeline supplied




### <a name="PipelineGroups.GetGroupByPipelineName">func</a> (\*PipelineGroups) [GetGroupByPipelineName](/src/target/resource_pipelinegroups.go?s=103:187#L4)
``` go
func (pg *PipelineGroups) GetGroupByPipelineName(pipelineName string) *PipelineGroup
```
GetGroupByPipelineName finds the pipeline group for the name of the pipeline supplied




## <a name="PipelineGroupsService">type</a> [PipelineGroupsService](/src/target/pipelinegroups.go?s=149:183#L6)
``` go
type PipelineGroupsService service
```
PipelineGroupsService describes the HAL _link resource for the api response object for a pipeline group response.










### <a name="PipelineGroupsService.List">func</a> (\*PipelineGroupsService) [List](/src/target/pipelinegroups.go?s=476:587#L18)
``` go
func (pgs *PipelineGroupsService) List(ctx context.Context, name string) (*PipelineGroups, *APIResponse, error)
```
List Pipeline groups




## <a name="PipelineHistory">type</a> [PipelineHistory](/src/target/pipeline.go?s=2294:2375#L62)
``` go
type PipelineHistory struct {
    Pipelines []*PipelineInstance `json:"pipelines"`
}
```
PipelineHistory describes the history of runs for a pipeline










## <a name="PipelineInstance">type</a> [PipelineInstance](/src/target/pipeline.go?s=2429:2719#L67)
``` go
type PipelineInstance struct {
    BuildCause   BuildCause `json:"build_cause"`
    CanRun       bool       `json:"can_run"`
    Name         string     `json:"name"`
    NaturalOrder int        `json:"natural_order"`
    Comment      string     `json:"comment"`
    Stages       []*Stage   `json:"stages"`
}
```
PipelineInstance describes a single pipeline run










## <a name="PipelineMaterial">type</a> [PipelineMaterial](/src/target/configuration.go?s=2926:3096#L71)
``` go
type PipelineMaterial struct {
    Name         string `xml:"pipelineName,attr"`
    StageName    string `xml:"stageName,attr"`
    MaterialName string `xml:"materialName,attr"`
}
```
PipelineMaterial part of cruise-control.xml. @TODO better documentation










## <a name="PipelineRequest">type</a> [PipelineRequest](/src/target/pipeline.go?s=232:336#L12)
``` go
type PipelineRequest struct {
    Group    string    `json:"group"`
    Pipeline *Pipeline `json:"pipeline"`
}
```
PipelineRequest describes a pipeline request object










## <a name="PipelineStatus">type</a> [PipelineStatus](/src/target/pipeline.go?s=3935:4072#L107)
``` go
type PipelineStatus struct {
    Locked      bool `json:"locked"`
    Paused      bool `json:"paused"`
    Schedulable bool `json:"schedulable"`
}
```
PipelineStatus describes whether a pipeline can be run or scheduled.










## <a name="PipelineTemplate">type</a> [PipelineTemplate](/src/target/pipelinetemplate.go?s=1135:1468#L41)
``` go
type PipelineTemplate struct {
    Links    *HALLinks                 `json:"_links,omitempty"`
    Name     string                    `json:"name"`
    Embedded *embeddedPipelineTemplate `json:"_embedded,omitempty"`
    Version  string                    `json:"template_version"`
    Stages   []*Stage                  `json:"stages,omitempty"`
}
```
PipelineTemplate describes a response from the API for a pipeline template object.










### <a name="PipelineTemplate.AddStage">func</a> (\*PipelineTemplate) [AddStage](/src/target/resource_pipelinetemplate.go?s=601:651#L29)
``` go
func (pt *PipelineTemplate) AddStage(stage *Stage)
```
AddStage appends a stage to this pipeline




### <a name="PipelineTemplate.GetName">func</a> (PipelineTemplate) [GetName](/src/target/resource_pipelinetemplate.go?s=367:410#L19)
``` go
func (pt PipelineTemplate) GetName() string
```
GetName of the pipeline template




### <a name="PipelineTemplate.GetStage">func</a> (PipelineTemplate) [GetStage](/src/target/resource_pipelinetemplate.go?s=164:224#L9)
``` go
func (pt PipelineTemplate) GetStage(stageName string) *Stage
```
GetStage from the pipeline template




### <a name="PipelineTemplate.GetStages">func</a> (PipelineTemplate) [GetStages](/src/target/resource_pipelinetemplate.go?s=54:101#L4)
``` go
func (pt PipelineTemplate) GetStages() []*Stage
```
GetStages from the pipeline template




### <a name="PipelineTemplate.GetVersion">func</a> (PipelineTemplate) [GetVersion](/src/target/resource_pipelinetemplate.go?s=1448:1504#L60)
``` go
func (pt PipelineTemplate) GetVersion() (version string)
```
GetVersion retrieves a version string for this pipeline




### <a name="PipelineTemplate.Pipelines">func</a> (PipelineTemplate) [Pipelines](/src/target/resource_pipelinetemplate.go?s=921:971#L39)
``` go
func (pt PipelineTemplate) Pipelines() []*Pipeline
```
Pipelines returns a list of Pipelines attached to this PipelineTemplate object.




### <a name="PipelineTemplate.RemoveLinks">func</a> (\*PipelineTemplate) [RemoveLinks](/src/target/resource_pipelinetemplate.go?s=775:816#L34)
``` go
func (pt *PipelineTemplate) RemoveLinks()
```
RemoveLinks gets the PipelineTemplate ready to be submitted to the GoCD API.




### <a name="PipelineTemplate.SetStage">func</a> (\*PipelineTemplate) [SetStage](/src/target/resource_pipelinetemplate.go?s=1057:1110#L44)
``` go
func (pt *PipelineTemplate) SetStage(newStage *Stage)
```
SetStage replaces a stage if it already exists




### <a name="PipelineTemplate.SetStages">func</a> (\*PipelineTemplate) [SetStages](/src/target/resource_pipelinetemplate.go?s=476:530#L24)
``` go
func (pt *PipelineTemplate) SetStages(stages []*Stage)
```
SetStages overwrites any existing stages




### <a name="PipelineTemplate.SetVersion">func</a> (\*PipelineTemplate) [SetVersion](/src/target/resource_pipelinetemplate.go?s=1307:1361#L55)
``` go
func (pt *PipelineTemplate) SetVersion(version string)
```
SetVersion sets a version string for this pipeline




## <a name="PipelineTemplateRequest">type</a> [PipelineTemplateRequest](/src/target/pipelinetemplate.go?s=266:406#L12)
``` go
type PipelineTemplateRequest struct {
    Name    string   `json:"name"`
    Stages  []*Stage `json:"stages"`
    Version string   `json:"version"`
}
```
PipelineTemplateRequest describes a PipelineTemplate










### <a name="PipelineTemplateRequest.GetVersion">func</a> (PipelineTemplateRequest) [GetVersion](/src/target/resource_pipelinetemplate.go?s=1731:1794#L70)
``` go
func (pt PipelineTemplateRequest) GetVersion() (version string)
```
GetVersion retrieves a version string for this pipeline




### <a name="PipelineTemplateRequest.SetVersion">func</a> (\*PipelineTemplateRequest) [SetVersion](/src/target/resource_pipelinetemplate.go?s=1583:1644#L65)
``` go
func (pt *PipelineTemplateRequest) SetVersion(version string)
```
SetVersion sets a version string for this pipeline




## <a name="PipelineTemplateResponse">type</a> [PipelineTemplateResponse](/src/target/pipelinetemplate.go?s=494:674#L19)
``` go
type PipelineTemplateResponse struct {
    Name     string `json:"name"`
    Embedded *struct {
        Pipelines []*struct {
            Name string `json:"name"`
        }
    } `json:"_embedded,omitempty"`
}
```
PipelineTemplateResponse describes an api response for a single pipeline templates










## <a name="PipelineTemplatesResponse">type</a> [PipelineTemplatesResponse](/src/target/pipelinetemplate.go?s=763:953#L29)
``` go
type PipelineTemplatesResponse struct {
    Links    *HALLinks `json:"_links,omitempty"`
    Embedded *struct {
        Templates []*PipelineTemplate `json:"templates"`
    } `json:"_embedded,omitempty"`
}
```
PipelineTemplatesResponse describes an api response for multiple pipeline templates










## <a name="PipelineTemplatesService">type</a> [PipelineTemplatesService](/src/target/pipelinetemplate.go?s=171:208#L9)
``` go
type PipelineTemplatesService service
```
PipelineTemplatesService describes the HAL _link resource for the api response object for a pipeline configuration objects.










### <a name="PipelineTemplatesService.Create">func</a> (\*PipelineTemplatesService) [Create](/src/target/pipelinetemplate.go?s=2401:2532#L76)
``` go
func (pts *PipelineTemplatesService) Create(ctx context.Context, name string, st []*Stage) (*PipelineTemplate, *APIResponse, error)
```
Create a new PipelineTemplate object in the GoCD API.




### <a name="PipelineTemplatesService.Delete">func</a> (\*PipelineTemplatesService) [Delete](/src/target/pipelinetemplate.go?s=3417:3524#L116)
``` go
func (pts *PipelineTemplatesService) Delete(ctx context.Context, name string) (string, *APIResponse, error)
```
Delete a PipelineTemplate from the GoCD API.




### <a name="PipelineTemplatesService.Get">func</a> (\*PipelineTemplatesService) [Get](/src/target/pipelinetemplate.go?s=1527:1642#L50)
``` go
func (pts *PipelineTemplatesService) Get(ctx context.Context, name string) (*PipelineTemplate, *APIResponse, error)
```
Get a single PipelineTemplate object in the GoCD API.




### <a name="PipelineTemplatesService.List">func</a> (\*PipelineTemplatesService) [List](/src/target/pipelinetemplate.go?s=2007:2112#L63)
``` go
func (pts *PipelineTemplatesService) List(ctx context.Context) ([]*PipelineTemplate, *APIResponse, error)
```
List all PipelineTemplate objects in the GoCD API.




### <a name="PipelineTemplatesService.Update">func</a> (\*PipelineTemplatesService) [Update](/src/target/pipelinetemplate.go?s=2879:3025#L96)
``` go
func (pts *PipelineTemplatesService) Update(ctx context.Context, name string, template *PipelineTemplate) (*PipelineTemplate, *APIResponse, error)
```
Update an PipelineTemplate object in the GoCD API.




## <a name="PipelinesService">type</a> [PipelinesService](/src/target/pipeline.go?s=146:175#L9)
``` go
type PipelinesService service
```
PipelinesService describes the HAL _link resource for the api response object for a pipelineconfig










### <a name="PipelinesService.GetHistory">func</a> (\*PipelinesService) [GetHistory](/src/target/pipeline.go?s=5599:5724#L153)
``` go
func (pgs *PipelinesService) GetHistory(ctx context.Context, name string, offset int) (*PipelineHistory, *APIResponse, error)
```
GetHistory returns a list of pipeline instances describing the pipeline history.




### <a name="PipelinesService.GetInstance">func</a> (\*PipelinesService) [GetInstance](/src/target/pipeline.go?s=5145:5272#L140)
``` go
func (pgs *PipelinesService) GetInstance(ctx context.Context, name string, offset int) (*PipelineInstance, *APIResponse, error)
```
GetInstance of a pipeline run.




### <a name="PipelinesService.GetStatus">func</a> (\*PipelinesService) [GetStatus](/src/target/pipeline.go?s=4157:4280#L114)
``` go
func (pgs *PipelinesService) GetStatus(ctx context.Context, name string, offset int) (*PipelineStatus, *APIResponse, error)
```
GetStatus returns a list of pipeline instanves describing the pipeline history.




### <a name="PipelinesService.Pause">func</a> (\*PipelinesService) [Pause](/src/target/pipeline.go?s=4533:4629#L125)
``` go
func (pgs *PipelinesService) Pause(ctx context.Context, name string) (bool, *APIResponse, error)
```
Pause allows a pipeline to handle new build events




### <a name="PipelinesService.ReleaseLock">func</a> (\*PipelinesService) [ReleaseLock](/src/target/pipeline.go?s=4950:5052#L135)
``` go
func (pgs *PipelinesService) ReleaseLock(ctx context.Context, name string) (bool, *APIResponse, error)
```
ReleaseLock frees a pipeline to handle new build events




### <a name="PipelinesService.Unpause">func</a> (\*PipelinesService) [Unpause](/src/target/pipeline.go?s=4738:4836#L130)
``` go
func (pgs *PipelinesService) Unpause(ctx context.Context, name string) (bool, *APIResponse, error)
```
Unpause allows a pipeline to handle new build events




## <a name="PluggableInstanceSettings">type</a> [PluggableInstanceSettings](/src/target/plugin.go?s=1017:1172#L31)
``` go
type PluggableInstanceSettings struct {
    Configurations []PluginConfiguration `json:"configurations"`
    View           PluginView            `json:"view"`
}
```
PluggableInstanceSettings describes plugin configuration










## <a name="Plugin">type</a> [Plugin](/src/target/plugin.go?s=430:955#L20)
``` go
type Plugin struct {
    Links                     *HALLinks                 `json:"_links"`
    ID                        string                    `json:"id"`
    Name                      string                    `json:"name"`
    DisplayName               string                    `json:"display_name"`
    Version                   string                    `json:"version"`
    Type                      string                    `json:"type"`
    PluggableInstanceSettings PluggableInstanceSettings `json:"pluggable_instance_settings"`
}
```
Plugin describes a single plugin resource.










## <a name="PluginConfiguration">type</a> [PluginConfiguration](/src/target/jobs.go?s=2829:2971#L72)
``` go
type PluginConfiguration struct {
    Key      string                      `json:"key"`
    Metadata PluginConfigurationMetadata `json:"metadata"`
}
```
PluginConfiguration describes how to reference a plugin.










## <a name="PluginConfigurationKVPair">type</a> [PluginConfigurationKVPair](/src/target/jobs.go?s=3322:3419#L85)
``` go
type PluginConfigurationKVPair struct {
    Key   string `json:"key"`
    Value string `json:"value"`
}
```
PluginConfigurationKVPair describes a key/value pair of plugin configurations.










## <a name="PluginConfigurationMetadata">type</a> [PluginConfigurationMetadata](/src/target/jobs.go?s=3072:3238#L78)
``` go
type PluginConfigurationMetadata struct {
    Secure         bool `json:"secure"`
    Required       bool `json:"required"`
    PartOfIdentity bool `json:"part_of_identity"`
}
```
PluginConfigurationMetadata describes the schema for a single configuration option for a plugin










## <a name="PluginView">type</a> [PluginView](/src/target/plugin.go?s=1229:1290#L37)
``` go
type PluginView struct {
    Template string `json:"template"`
}
```
PluginView describes any view attached to a plugin.










## <a name="PluginsResponse">type</a> [PluginsResponse](/src/target/plugin.go?s=230:382#L12)
``` go
type PluginsResponse struct {
    Links    *HALLinks `json:"_links"`
    Embedded struct {
        PluginInfo []*Plugin `json:"plugin_info"`
    } `json:"_embedded"`
}
```
PluginsResponse describes the response obejct for a plugin API call.










## <a name="PluginsService">type</a> [PluginsService](/src/target/plugin.go?s=129:156#L9)
``` go
type PluginsService service
```
PluginsService exposes calls for interacting with Plugin objects in the GoCD API.










### <a name="PluginsService.Get">func</a> (\*PluginsService) [Get](/src/target/plugin.go?s=1668:1762#L54)
``` go
func (ps *PluginsService) Get(ctx context.Context, name string) (*Plugin, *APIResponse, error)
```
Get retrieves information about a specific plugin.




### <a name="PluginsService.List">func</a> (\*PluginsService) [List](/src/target/plugin.go?s=1322:1413#L42)
``` go
func (ps *PluginsService) List(ctx context.Context) (*PluginsResponse, *APIResponse, error)
```
List retrieves all plugins




## <a name="Properties">type</a> [Properties](/src/target/resource_properties.go?s=161:305#L13)
``` go
type Properties struct {
    UnmarshallWithHeader bool
    IsDatum              bool
    Header               []string
    DataFrame            [][]string
}
```
Properties describes a properties resource in the GoCD API.







### <a name="NewPropertiesFrame">func</a> [NewPropertiesFrame](/src/target/resource_properties.go?s=385:438#L21)
``` go
func NewPropertiesFrame(frame [][]string) *Properties
```
NewPropertiesFrame generate a new data frame for properties on a gocd job.





### <a name="Properties.AddRow">func</a> (\*Properties) [AddRow](/src/target/resource_properties.go?s=875:915#L45)
``` go
func (pr *Properties) AddRow(r []string)
```
AddRow to an existing properties data frame




### <a name="Properties.Get">func</a> (Properties) [Get](/src/target/resource_properties.go?s=633:688#L34)
``` go
func (pr Properties) Get(row int, column string) string
```
Get a single parameter value for a given run of the job.




### <a name="Properties.MarshalJSON">func</a> (\*Properties) [MarshalJSON](/src/target/resource_properties.go?s=2312:2363#L108)
``` go
func (pr *Properties) MarshalJSON() ([]byte, error)
```
MarshalJSON converts the properties structure to a list of maps




### <a name="Properties.MarshallCSV">func</a> (Properties) [MarshallCSV](/src/target/resource_properties.go?s=1203:1253#L58)
``` go
func (pr Properties) MarshallCSV() (string, error)
```
MarshallCSV returns the data frame as a string




### <a name="Properties.SetRow">func</a> (\*Properties) [SetRow](/src/target/resource_properties.go?s=990:1039#L50)
``` go
func (pr *Properties) SetRow(row int, r []string)
```
SetRow in an existing data frame




### <a name="Properties.UnmarshallCSV">func</a> (\*Properties) [UnmarshallCSV](/src/target/resource_properties.go?s=1588:1641#L75)
``` go
func (pr *Properties) UnmarshallCSV(raw string) error
```
UnmarshallCSV returns the data frame from a string




### <a name="Properties.Write">func</a> (\*Properties) [Write](/src/target/resource_properties.go?s=2025:2081#L96)
``` go
func (pr *Properties) Write(p []byte) (n int, err error)
```
Write the data frame to a byte stream as a csv.




## <a name="PropertiesService">type</a> [PropertiesService](/src/target/properties.go?s=136:166#L11)
``` go
type PropertiesService service
```
PropertiesService describes Actions which can be performed on agents










### <a name="PropertiesService.Create">func</a> (\*PropertiesService) [Create](/src/target/properties.go?s=1604:1736#L54)
``` go
func (ps *PropertiesService) Create(ctx context.Context, name string, value string, pr *PropertyRequest) (bool, *APIResponse, error)
```
Create a specific property for the given job/pipeline/stage run.




### <a name="PropertiesService.Get">func</a> (\*PropertiesService) [Get](/src/target/properties.go?s=1139:1261#L43)
``` go
func (ps *PropertiesService) Get(ctx context.Context, name string, pr *PropertyRequest) (*Properties, *APIResponse, error)
```
Get a specific property for the given job/pipeline/stage run.




### <a name="PropertiesService.List">func</a> (\*PropertiesService) [List](/src/target/properties.go?s=692:802#L32)
``` go
func (ps *PropertiesService) List(ctx context.Context, pr *PropertyRequest) (*Properties, *APIResponse, error)
```
List the properties for the given job/pipeline/stage run.




### <a name="PropertiesService.ListHistorical">func</a> (\*PropertiesService) [ListHistorical](/src/target/properties.go?s=2508:2628#L82)
``` go
func (ps *PropertiesService) ListHistorical(ctx context.Context, pr *PropertyRequest) (*Properties, *APIResponse, error)
```
ListHistorical properties for a given pipeline, stage, job.




## <a name="PropertyCreateResponse">type</a> [PropertyCreateResponse](/src/target/properties.go?s=563:629#L26)
``` go
type PropertyCreateResponse struct {
    Name  string
    Value string
}
```
PropertyCreateResponse handles the parsing of the response when creating a property










## <a name="PropertyRequest">type</a> [PropertyRequest](/src/target/properties.go?s=262:474#L14)
``` go
type PropertyRequest struct {
    Pipeline        string
    PipelineCounter int
    Stage           string
    StageCounter    int
    Job             string
    LimitPipeline   string
    Limit           int
    Single          bool
}
```
PropertyRequest describes the parameters to be submitted when calling/creating properties.










## <a name="Stage">type</a> [Stage](/src/target/stages.go?s=166:781#L7)
``` go
type Stage struct {
    Name                  string                 `json:"name"`
    FetchMaterials        bool                   `json:"fetch_materials"`
    CleanWorkingDirectory bool                   `json:"clean_working_directory"`
    NeverCleanupArtifacts bool                   `json:"never_cleanup_artifacts"`
    Approval              *Approval              `json:"approval,omitempty"`
    EnvironmentVariables  []*EnvironmentVariable `json:"environment_variables,omitempty"`
    Resources             []string               `json:"resource,omitempty"`
    Jobs                  []*Job                 `json:"jobs,omitempty"`
}
```
Stage represents a GoCD Stage object.










### <a name="Stage.Clean">func</a> (\*Stage) [Clean](/src/target/resource_stages.go?s=740:763#L39)
``` go
func (s *Stage) Clean()
```
Clean the approvel step.




### <a name="Stage.JSONString">func</a> (\*Stage) [JSONString](/src/target/resource_stages.go?s=116:160#L9)
``` go
func (s *Stage) JSONString() (string, error)
```
JSONString returns a string of this stage as a JSON object.




### <a name="Stage.Validate">func</a> (\*Stage) [Validate](/src/target/resource_stages.go?s=409:441#L20)
``` go
func (s *Stage) Validate() error
```
Validate ensures the attributes attached to this structure are ready for submission to the GoCD API.




## <a name="StageContainer">type</a> [StageContainer](/src/target/resource.go?s=125:310#L6)
``` go
type StageContainer interface {
    GetName() string
    SetStage(stage *Stage)
    GetStage(string) *Stage
    SetStages(stages []*Stage)
    GetStages() []*Stage
    AddStage(stage *Stage)
    Versioned
}
```
StageContainer describes structs which contain stages, eg Pipelines and PipelineTemplates










## <a name="StagesService">type</a> [StagesService](/src/target/stages.go?s=97:123#L4)
``` go
type StagesService service
```
StagesService exposes calls for interacting with Stage objects in the GoCD API.










## <a name="StringResponse">type</a> [StringResponse](/src/target/gocd.go?s=985:1048#L50)
``` go
type StringResponse struct {
    Message string `json:"message"`
}
```
StringResponse handles the unmarshaling of the single string response from DELETE requests.










## <a name="Tab">type</a> [Tab](/src/target/jobs.go?s=2241:2314#L51)
``` go
type Tab struct {
    Name string `json:"name"`
    Path string `json:"path"`
}
```
Tab description in a gocd job










## <a name="Task">type</a> [Task](/src/target/jobs.go?s=3470:3578#L91)
``` go
type Task struct {
    Type       string         `json:"type"`
    Attributes TaskAttributes `json:"attributes"`
}
```
Task Describes a Task object in the GoCD api.










### <a name="Task.Validate">func</a> (\*Task) [Validate](/src/target/resource_task.go?s=76:107#L6)
``` go
func (t *Task) Validate() error
```
Validate each of the possible task types.




## <a name="TaskAttributes">type</a> [TaskAttributes](/src/target/jobs.go?s=3639:4850#L97)
``` go
type TaskAttributes struct {
    RunIf               []string                    `json:"run_if,omitempty"`
    Command             string                      `json:"command,omitempty"`
    WorkingDirectory    string                      `json:"working_directory,omitempty"`
    Arguments           []string                    `json:"arguments,omitempty"`
    BuildFile           string                      `json:"build_file,omitempty"`
    Target              string                      `json:"target,omitempty"`
    NantPath            string                      `json:"nant_path,omitempty"`
    Pipeline            string                      `json:"pipeline,omitempty"`
    Stage               string                      `json:"stage,omitempty"`
    Job                 string                      `json:"job,omitempty"`
    Source              string                      `json:"source,omitempty"`
    IsSourceAFile       bool                        `json:"is_source_a_file,omitempty"`
    Destination         string                      `json:"destination,omitempty"`
    PluginConfiguration *TaskPluginConfiguration    `json:"plugin_configuration,omitempty"`
    Configuration       []PluginConfigurationKVPair `json:"configuration,omitempty"`
}
```
TaskAttributes describes all the properties for a Task.










### <a name="TaskAttributes.ValidateAnt">func</a> (\*TaskAttributes) [ValidateAnt](/src/target/jobs_validation.go?s=623:667#L24)
``` go
func (t *TaskAttributes) ValidateAnt() error
```
ValidateAnt checks that the specified values for the Task struct are correct for a an Ant task




### <a name="TaskAttributes.ValidateExec">func</a> (\*TaskAttributes) [ValidateExec](/src/target/jobs_validation.go?s=132:177#L6)
``` go
func (t *TaskAttributes) ValidateExec() error
```
ValidateExec checks that the specified values for the Task struct are correct for a cli exec task




## <a name="TaskPluginConfiguration">type</a> [TaskPluginConfiguration](/src/target/jobs.go?s=4924:5024#L116)
``` go
type TaskPluginConfiguration struct {
    ID      string `json:"id"`
    Version string `json:"version"`
}
```
TaskPluginConfiguration is for specifying options for pluggable task










## <a name="Version">type</a> [Version](/src/target/configuration.go?s=8090:8365#L207)
``` go
type Version struct {
    Links       *HALLinks `json:"_links"`
    Version     string    `json:"version"`
    BuildNumber string    `json:"build_number"`
    GitSHA      string    `json:"git_sha"`
    FullVersion string    `json:"full_version"`
    CommitURL   string    `json:"commit_url"`
}
```
Version part of cruise-control.xml. @TODO better documentation










## <a name="Versioned">type</a> [Versioned](/src/target/resource.go?s=521:598#L23)
``` go
type Versioned interface {
    GetVersion() string
    SetVersion(version string)
}
```
Versioned describes resources which can get and set versions














- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
