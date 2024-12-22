package render

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-render/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableRenderEnvGroup(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "render_env_group",
		Description: "Information about environment group descriptions, including ID, name, owner, service links, and timestamps.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListEnvGroup,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    opengovernance.GetEnvGroup,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The unique identifier for the environment group.", Transform: transform.FromField("Description.ID")},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the environment group.", Transform: transform.FromField("Description.Name")},
			{Name: "ownerId", Type: proto.ColumnType_STRING, Description: "The ID of the owner of the environment group.", Transform: transform.FromField("Description.OwnerID")},
			{Name: "createdAt", Type: proto.ColumnType_TIMESTAMP, Description: "The timestamp of when the environment group was created.", Transform: transform.FromField("Description.CreatedAt")},
			{Name: "updatedAt", Type: proto.ColumnType_TIMESTAMP, Description: "The timestamp of the last update to the environment group.", Transform: transform.FromField("Description.UpdatedAt")},
			{Name: "serviceLinks", Type: proto.ColumnType_JSON, Description: "A list of service links associated with the environment group.", Transform: transform.FromField("Description.ServiceLinks")},
			{Name: "environmentId", Type: proto.ColumnType_STRING, Description: "The ID of the associated environment.", Transform: transform.FromField("Description.EnvironmentID")},
		},
	}
}
