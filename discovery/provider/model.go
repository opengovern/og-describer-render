//go:generate go run ../../pkg/sdk/runable/steampipe_es_client_generator/main.go -pluginPath ../../steampipe-plugin-REPLACEME/REPLACEME -file $GOFILE -output ../../pkg/sdk/es/resources_clients.go -resourceTypesFile ../resource_types/resource-types.json

// Implement types for each resource

package provider

import "time"

type Metadata struct{}

type OwnerJSON struct {
	ID                   string `json:"id"`
	Name                 string `json:"name"`
	Email                string `json:"email"`
	TwoFactorAuthEnabled bool   `json:"twoFactorAuthEnabled"` // converted to  from "twoFactorAuthEnabled"
	Type                 string `json:"type"`
}

type Owner struct {
	ID                   string
	Name                 string
	Email                string
	TwoFactorAuthEnabled bool
	Type                 string
}

type ProjectResponse struct {
	Project ProjectJSON `json:"project"`
	Cursor  string      `json:"cursor"`
}

type ProjectJSON struct {
	ID             string    `json:"id"`
	CreatedAt      time.Time `json:"createdAt"`      // converted to  from "createdAt"
	UpdatedAt      time.Time `json:"updatedAt"`      // converted to  from "updatedAt"
	Name           string    `json:"name"`
	Owner          OwnerJSON `json:"owner"`
	EnvironmentIDs []string  `json:"environmentIds"` // converted to  from "environmentIds"
}

type ProjectDescription struct {
	ID             string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Name           string
	Owner          Owner
	EnvironmentIDs []string
}

type EnvironmentResponse struct {
	Environment EnvironmentJSON `json:"environment"`
	Cursor      string          `json:"cursor"`
}

type EnvironmentJSON struct {
	ID              string   `json:"id"`
	Name            string   `json:"name"`
	ProjectID       string   `json:"projectId"`       // converted to  from "projectId"
	DatabasesIDs    []string `json:"databasesIds"`    // converted to  from "databasesIds"
	RedisIDs        []string `json:"redisIds"`        // converted to  from "redisIds"
	ServiceIDs      []string `json:"serviceIds"`      // converted to  from "serviceIds"
	EnvGroupIDs     []string `json:"envGroupIds"`    // converted to  from "envGroupIds"
	ProtectedStatus string   `json:"protectedStatus"` // converted to  from "protectedStatus"
}

type EnvironmentDescription struct {
	ID              string   `json:"id"`
	Name            string   `json:"name"`
	ProjectID       string   `json:"projectId"`       // converted to  from "projectId"
	DatabasesIDs    []string `json:"databasesIds"`    // converted to  from "databasesIds"
	RedisIDs        []string `json:"redisIds"`        // converted to  from "redisIds"
	ServiceIDs      []string `json:"serviceIds"`      // converted to  from "serviceIds"
	EnvGroupIDs     []string `json:"envGroupIds"`    // converted to  from "envGroupIds"
	ProtectedStatus string   `json:"protectedStatus"` // converted to  from "protectedStatus"
}

type IPAllowJSON struct {
	CIDRBlock   string `json:"cidrBlock"` // converted to  from "cidrBlock"
	Description string `json:"description"`
}

type IPAllow struct {
	CIDRBlock   string `json:"cidrBlock"` // converted to  from "cidrBlock"
	Description string `json:"description"`
}

type ReadReplicaJSON struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ReadReplica struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type PostgresResponse struct {
	Postgres PostgresJSON `json:"postgres"`
	Cursor   string       `json:"cursor"`
}

type PostgresJSON struct {
	ID                      string            `json:"id"`
	IPAllowList             []IPAllowJSON     `json:"ipAllowList"`              // converted to  from "ipAllowList"
	CreatedAt               time.Time         `json:"createdAt"`                 // converted to  from "createdAt"
	UpdatedAt               time.Time         `json:"updatedAt"`                 // converted to  from "updatedAt"
	ExpiresAt               time.Time         `json:"expiresAt"`                 // converted to  from "expiresAt"
	DatabaseName            string            `json:"databaseName"`              // converted to  from "databaseName"
	DatabaseUser            string            `json:"databaseUser"`              // converted to  from "databaseUser"
	EnvironmentID           string            `json:"environmentId"`             // converted to  from "environmentId"
	HighAvailabilityEnabled bool              `json:"highAvailabilityEnabled"`  // converted to  from "highAvailabilityEnabled"
	Name                    string            `json:"name"`
	Owner                   OwnerJSON         `json:"owner"`
	Plan                    string            `json:"plan"`
	DiskSizeGB              int               `json:"diskSizeGb"`       // converted to  from "diskSizeGB"
	PrimaryPostgresID       string            `json:"primaryPostgresId"`// converted to  from "primaryPostgresID"
	Region                  string            `json:"region"`
	ReadReplicas            []ReadReplicaJSON `json:"readReplicas"`      // converted to  from "readReplicas"
	Role                    string            `json:"role"`
	Status                  string            `json:"status"`
	Version                 string            `json:"version"`
	Suspended               string            `json:"suspended"`
	Suspenders              []string          `json:"suspenders"`
	DashboardURL            string            `json:"dashboardUrl"` // converted to  from "dashboardUrl"
}

type PostgresDescription struct {
	ID                      string
	IPAllowList             []IPAllow
	CreatedAt               time.Time
	UpdatedAt               time.Time
	ExpiresAt               time.Time
	DatabaseName            string
	DatabaseUser            string
	EnvironmentID           string
	HighAvailabilityEnabled bool
	Name                    string
	Owner                   Owner
	Plan                    string
	DiskSizeGB              int
	PrimaryPostgresID       string
	Region                  string
	ReadReplicas            []ReadReplica
	Role                    string
	Status                  string
	Version                 string
	Suspended               string
	Suspenders              []string
	DashboardURL            string
}

type BuildFilterJSON struct {
	Paths        []string `json:"paths"`
	IgnoredPaths []string `json:"ignoredPaths"` // converted to  from "ignoredPaths"
}

type BuildFilter struct {
	Paths        []string
	IgnoredPaths []string
}

type RegistryCredentialJSON struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type RegistryCredential struct {
	ID   string
	Name string
}

type ServiceDetailsJSON struct {
	BuildCommand string           `json:"buildCommand"` // converted to  from "buildCommand"
	ParentServer ParentServerJSON `json:"parentServer"` // converted to  from "parentServer"
	PublishPath  string           `json:"publishPath"`  // converted to  from "publishPath"
	Previews     PreviewsJSON     `json:"previews"`
	URL          string           `json:"url"`
	BuildPlan    string           `json:"buildPlan"` // converted to  from "buildPlan"
}

type ServiceDetails struct {
	BuildCommand string
	ParentServer ParentServer
	PublishPath  string
	Previews     Previews
	URL          string
	BuildPlan    string
}

type ParentServerJSON struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ParentServer struct {
	ID   string
	Name string
}

type PreviewsJSON struct {
	Generation string `json:"generation"`
}

type Previews struct {
	Generation string
}

type ServiceResponse struct {
	Service ServiceJSON `json:"service"`
	Cursor  string      `json:"cursor"`
}

type ServiceJSON struct {
	ID                 string                 `json:"id"`
	AutoDeploy         string                 `json:"autoDeploy"`         // converted to  from "autoDeploy"
	Branch             string                 `json:"branch"`
	BuildFilter        BuildFilterJSON        `json:"buildFilter"`        // converted to  from "buildFilter"
	CreatedAt          time.Time              `json:"createdAt"`          // converted to  from "createdAt"
	DashboardURL       string                 `json:"dashboardUrl"`       // converted to  from "dashboardUrl"
	EnvironmentID      string                 `json:"environmentId"`      // converted to  from "environmentId"
	ImagePath          string                 `json:"imagePath"`          // converted to  from "imagePath"
	Name               string                 `json:"name"`
	NotifyOnFail       string                 `json:"notifyOnFail"`      // converted to  from "notifyOnFail"
	OwnerID            string                 `json:"ownerId"`            // converted to  from "ownerId"
	RegistryCredential RegistryCredentialJSON `json:"registryCredential"` // converted to  from "registryCredential"
	Repo               string                 `json:"repo"`
	RootDir            string                 `json:"rootDir"`            // converted to  from "rootDir"
	Slug               string                 `json:"slug"`
	Suspended          string                 `json:"suspended"`
	Suspenders         []string               `json:"suspenders"`
	Type               string                 `json:"type"`
	UpdatedAt          time.Time              `json:"updatedAt"`      // converted to  from "updatedAt"

	ServiceDetails     ServiceDetailsJSON     `json:"serviceDetails"` // converted to  from "serviceDetails"
}

type ServiceDescription struct {
	ID                 string
	AutoDeploy         string
	Branch             string
	BuildFilter        BuildFilter
	CreatedAt          time.Time
	DashboardURL       string
	EnvironmentID      string
	ImagePath          string
	Name               string
	NotifyOnFail       string
	OwnerID            string
	RegistryCredential RegistryCredential
	Repo               string
	RootDir            string
	Slug               string
	Suspended          string
	Suspenders         []string
	Type               string
	UpdatedAt          time.Time
	ServiceDetails     ServiceDetails
}

type JobResponse struct {
	Job    JobJSON `json:"job"`
	Cursor string  `json:"cursor"`
}

type JobJSON struct {
	ID           string    `json:"id"`
	ServiceID    string    `json:"serviceId"`    // converted to  from "serviceId"
	StartCommand string    `json:"startCommand"` // converted to  from "startCommand"
	PlanID       string    `json:"planId"`       // converted to  from "planId"
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"createdAt"`   // converted to  from "createdAt"
	StartedAt    time.Time `json:"startedAt"`   // converted to  from "startedAt"
	FinishedAt   time.Time `json:"finishedAt"`  // converted to  from "finishedAt"
}

type JobDescription struct {
	ID           string
	ServiceID    string
	StartCommand string
	PlanID       string
	Status       string
	CreatedAt    time.Time
	StartedAt    time.Time
	FinishedAt   time.Time
}

type DiskResponse struct {
	Disk   DiskJSON `json:"disk"`
	Cursor string   `json:"cursor"`
}

type DiskJSON struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	SizeGB    int       `json:"sizeGb"`    // converted to  from "sizeGB"
	MountPath string    `json:"mountPath"` // converted to  from "mountPath"
	ServiceID string    `json:"serviceId"` // converted to  from "serviceId"
	CreatedAt time.Time `json:"createdAt"` // converted to  from "createdAt"
	UpdatedAt time.Time `json:"updatedAt"` // converted to  from "updatedAt"
}

type DiskDescription struct {
	ID        string
	Name      string
	SizeGB    int
	MountPath string
	ServiceID string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CommitJSON struct {
	ID        string    `json:"id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"createdAt"` // converted to  from "createdAt"
}

type Commit struct {
	ID        string
	Message   string
	CreatedAt time.Time
}

type ImageJSON struct {
	Ref                string `json:"ref"`
	SHA                string `json:"sha"`
	RegistryCredential string `json:"registryCredential"` // converted to  from "registryCredential"
}

type Image struct {
	Ref                string
	SHA                string
	RegistryCredential string
}

type DeployResponse struct {
	Deploy DeployJSON `json:"deploy"`
	Cursor string     `json:"cursor"`
}

type DeployJSON struct {
	ID         string     `json:"id"`
	Commit     CommitJSON `json:"commit"`
	Image      ImageJSON  `json:"image"`
	Status     string     `json:"status"`
	Trigger    string     `json:"trigger"`
	FinishedAt time.Time  `json:"finishedAt"` // converted to  from "finishedAt"
	CreatedAt  time.Time  `json:"createdAt"`  // converted to  from "createdAt"
	UpdatedAt  time.Time  `json:"updatedAt"`  // converted to  from "updatedAt"
}

type DeployDescription struct {
	ID         string
	ServiceID 	string
	Commit     Commit
	Image      Image
	Status     string
	Trigger    string
	FinishedAt time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type BlueprintResponse struct {
	BluePrint BlueprintJSON `json:"blueprint"`
	Cursor    string        `json:"cursor"`
}

type BlueprintJSON struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Status   string    `json:"status"`
	AutoSync bool      `json:"autoSync"` // converted to  from "autoSync"
	Repo     string    `json:"repo"`
	Branch   string    `json:"branch"`
	LastSync time.Time `json:"lastSync"` // converted to  from "lastSync"
}

type BlueprintDescription struct {
	ID       string
	Name     string
	Status   string
	AutoSync bool
	Repo     string
	Branch   string
	LastSync time.Time
}

type ServiceLinkJSON struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type ServiceLink struct {
	ID   string
	Name string
	Type string
}

type EnvGroupResponse struct {
	EnvGroup EnvGroupJSON `json:"envGroup"`
	Cursor   string       `json:"cursor"`
}

type EnvGroupJSON struct {
	ID            string            `json:"id"`
	Name          string            `json:"name"`
	OwnerID       string            `json:"ownerId"`      // converted to  from "ownerId"
	CreatedAt     time.Time         `json:"createdAt"`    // converted to  from "createdAt"
	UpdatedAt     time.Time         `json:"updatedAt"`    // converted to  from "updatedAt"
	ServiceLinks  []ServiceLinkJSON `json:"serviceLinks"` // converted to  from "serviceLinks"
	EnvironmentID string            `json:"environmentId"`// converted to  from "environmentId"
}

type EnvGroupDescription struct {
	ID            string
	Name          string
	OwnerID       string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	ServiceLinks  []ServiceLink
	EnvironmentID string
}

type HeaderResponse struct {
	Header HeaderJSON `json:"header"`
	Cursor string     `json:"cursor"`
}

type HeaderJSON struct {
	ID    string `json:"id"`
	Path  string `json:"path"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

type HeaderDescription struct {
	ID    string
	Path  string
	Name  string
	Value string
}

type RouteResponse struct {
	Route  RouteJSON `json:"route"`
	Cursor string    `json:"cursor"`
}

type RouteJSON struct {
	ID          string `json:"id"`
	Type        string `json:"type"`
	Source      string `json:"source"`
	Destination string `json:"destination"`
	Priority    int    `json:"priority"`
}

type RouteDescription struct {
	ID          string
	Type        string
	Source      string
	Destination string
	Priority    int
}

type PostgresqlBackupJSON struct {
	ID        string `json:"id"`
	CreatedAt string `json:"createdAt"` // converted to  from "createdAt"
	URL       string `json:"url"`
}

type PostgresqlBackupDescription struct {
	ID        string
	CreatedAt string
	URL       string
}
