package maps

import (
	describer "github.com/opengovern/og-describer-render/discovery/describers"
	model "github.com/opengovern/og-describer-render/discovery/pkg/models"
	"github.com/opengovern/og-describer-render/discovery/provider"
	"github.com/opengovern/og-describer-render/platform/constants"
	"github.com/opengovern/og-util/pkg/integration/interfaces"
)

var ResourceTypes = map[string]model.ResourceType{

	"Render/Blueprint": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Render/Blueprint",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByRender(describer.ListBlueprints),
		GetDescriber:    provider.DescribeSingleByRender(describer.GetBlueprint),
	},

	"Render/Deploy": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Render/Deploy",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByRender(describer.ListDeploys),
		GetDescriber:    nil,
	},

	"Render/Disk": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Render/Disk",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByRender(describer.ListDisks),
		GetDescriber:    provider.DescribeSingleByRender(describer.GetDisk),
	},

	"Render/EnvGroup": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Render/EnvGroup",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByRender(describer.ListEnvGroups),
		GetDescriber:    provider.DescribeSingleByRender(describer.GetEnvGroup),
	},

	"Render/Environment": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Render/Environment",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByRender(describer.ListEnvironments),
		GetDescriber:    provider.DescribeSingleByRender(describer.GetEnvironment),
	},

	"Render/Header": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Render/Header",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByRender(describer.ListHeaders),
		GetDescriber:    nil,
	},

	"Render/Job": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Render/Job",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByRender(describer.ListJobs),
		GetDescriber:    nil,
	},

	"Render/PostgresInstance": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Render/PostgresInstance",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByRender(describer.ListPostgresInstances),
		GetDescriber:    provider.DescribeSingleByRender(describer.GetPostgresInstance),
	},

	"Render/Project": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Render/Project",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByRender(describer.ListProjects),
		GetDescriber:    provider.DescribeSingleByRender(describer.GetProject),
	},

	"Render/Route": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Render/Route",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByRender(describer.ListRoutes),
		GetDescriber:    nil,
	},

	"Render/Service": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Render/Service",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByRender(describer.ListServices),
		GetDescriber:    provider.DescribeSingleByRender(describer.GetService),
	},

	"Render/Postgresql/Backup": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Render/Postgresql/Backup",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByRender(describer.ListPostgresqlBackups),
		GetDescriber:    nil,
	},
}

var ResourceTypeConfigs = map[string]*interfaces.ResourceTypeConfiguration{

	"Render/Blueprint": {
		Name:            "Render/Blueprint",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"Render/Deploy": {
		Name:            "Render/Deploy",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"Render/Disk": {
		Name:            "Render/Disk",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"Render/EnvGroup": {
		Name:            "Render/EnvGroup",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"Render/Environment": {
		Name:            "Render/Environment",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"Render/Header": {
		Name:            "Render/Header",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"Render/Job": {
		Name:            "Render/Job",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"Render/PostgresInstance": {
		Name:            "Render/PostgresInstance",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"Render/Project": {
		Name:            "Render/Project",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"Render/Route": {
		Name:            "Render/Route",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"Render/Service": {
		Name:            "Render/Service",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"Render/Postgresql/Backup": {
		Name:            "Render/Postgresql/Backup",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},
}

var ResourceTypesList = []string{
	"Render/Blueprint",
	"Render/Deploy",
	"Render/Disk",
	"Render/EnvGroup",
	"Render/Environment",
	"Render/Header",
	"Render/Job",
	"Render/PostgresInstance",
	"Render/Project",
	"Render/Route",
	"Render/Service",
	"Render/Postgresql/Backup",
}
