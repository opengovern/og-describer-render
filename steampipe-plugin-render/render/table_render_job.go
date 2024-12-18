package render

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-render/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableRenderJob(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "render_job",
		Description: "Information about job descriptions, including ID, service details, status, and timestamps.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListJob,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    opengovernance.GetJob,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The unique identifier for the job."},
			{Name: "serviceId", Type: proto.ColumnType_STRING, Description: "The ID of the service associated with the job."},
			{Name: "startCommand", Type: proto.ColumnType_STRING, Description: "The start command for the job."},
			{Name: "planId", Type: proto.ColumnType_STRING, Description: "The ID of the plan associated with the job."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "The current status of the job (e.g., running, completed)."},
			{Name: "createdAt", Type: proto.ColumnType_TIMESTAMP, Description: "The timestamp of when the job was created."},
			{Name: "startedAt", Type: proto.ColumnType_TIMESTAMP, Description: "The timestamp of when the job started."},
			{Name: "finishedAt", Type: proto.ColumnType_TIMESTAMP, Description: "The timestamp of when the job finished."},
		},
	}
}
