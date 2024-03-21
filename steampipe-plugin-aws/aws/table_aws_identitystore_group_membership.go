package aws

import (
	"context"
	identitystorev1 "github.com/aws/aws-sdk-go/service/identitystore"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsIdentityStoreGroupMembership(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_identitystore_group_membership",
		Description: "AWS Identity Store Group Membership",
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListIdentityStoreGroupMembership,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ResourceNotFoundException"}),
			},
		},
		GetMatrixItemFunc: SupportedRegionMatrix(identitystorev1.EndpointsID),
		Columns: awsRegionalColumns([]*plugin.Column{
			{
				Name:        "membership_id",
				Description: "The identifier for a GroupMembership object in an identity store.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MembershipId"),
			},
			{
				Name:        "identity_store_id",
				Description: "The globally unique identifier for the identity store.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.IdentityStoreId"),
			},
			{
				Name:        "group_id",
				Description: "The identifier for a group in the identity store.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.GroupId"),
			},
			{
				Name:        "member_id",
				Description: "Specific identifier for a user indicates that the user is a member of the group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MemberId.Value"),
			},

			// Standard columns for all tables
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MembershipId"),
			},
		}),
	}
}
