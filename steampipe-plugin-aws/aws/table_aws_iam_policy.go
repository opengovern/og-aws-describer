package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"strings"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION
//error
//check the TODO comments

func tableAwsIamPolicy(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_iam_policy",
		Description: "AWS IAM Policy",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"arn"}),
			Hydrate:    kaytu.GetIAMPolicy,
		},
		List: &plugin.ListConfig{
			Hydrate:    kaytu.ListIAMPolicy,
			KeyColumns: plugin.KeyColumnSlice{
				//TODO The code generator doesn't support any of these type of filters. For now, don't pass them to the backend
				// instead, let PostgreSQL handle them

				// {Name: "is_aws_managed", Require: plugin.Optional, Operators: []string{"<>", "="}},
				// {Name: "is_attached", Require: plugin.Optional, Operators: []string{"<>", "="}},
				// {Name: "path", Require: plugin.Optional},
			},
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The friendly name that identifies the iam policy.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Policy.PolicyName"),
			},
			{
				Name:        "policy_id",
				Description: "The stable and unique string identifying the policy.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Policy.PolicyId"),
			},
			{
				Name:        "path",
				Description: "The path to the policy.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Policy.Path"),
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) specifying the iam policy.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Policy.Arn"),
			},
			{
				Name:        "is_aws_managed",
				Description: "Specifies whether the policy is AWS Managed or Customer Managed. If true policy is aws managed otherwise customer managed.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.From(isPolicyAwsManaged),
			},
			{
				Name:        "is_attachable",
				Description: "Specifies whether the policy can be attached to an IAM user, group, or role.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Policy.IsAttachable"),
			},
			{
				Name:        "create_date",
				Description: "The date and time, when the policy was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Policy.CreateDate"),
			},
			{
				Name:        "update_date",
				Description: "The date and time, when the policy was last updated.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Policy.UpdateDate"),
			},
			{
				Name:        "attachment_count",
				Description: "The number of entities (users, groups, and roles) that the policy is attached to.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Policy.AttachmentCount"),
			},
			{
				Name:        "is_attached",
				Description: "Specifies whether the policy is attached to at least one IAM user, group, or role.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Policy.AttachmentCount").Transform(attachementCountToBool),
			},
			{
				Name:        "default_version_id",
				Description: "The identifier for the version of the policy that is set as the default version.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Policy.DefaultVersionId"),
			},
			{
				Name:        "permissions_boundary_usage_count",
				Description: "The number of entities (users and roles) for which the policy is used to set the permissions boundary.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Policy.PermissionsBoundaryUsageCount"),
			},
			{
				Name:        "policy",
				Description: "Contains the details about the policy.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.PolicyVersion.Document").Transform(unescape),
			},
			{
				Name:        "policy_std",
				Description: "Contains the policy in a canonical form for easier searching.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.PolicyVersion.Document").Transform(unescape).Transform(policyToCanonical),
			},
			{
				Name:        "tags_src",
				Description: "A list of tags attached with the IAM policy.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Policy.Tags"),
			},

			// Standard columns for all tables
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(iamPolicyTurbotTags),
			},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Policy.PolicyName"),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Policy.Arn").Transform(arnToAkas),
			},
		}),
	}
}

//// LIST FUNCTION

//// HYDRATE FUNCTIONS

// isPolicyAwsManaged returns true if policy is aws managed
func isPolicyAwsManaged(_ context.Context, d *transform.TransformData) (interface{}, error) {
	policy := d.HydrateItem.(kaytu.IAMPolicy).Description.Policy
	metadata := d.HydrateItem.(kaytu.IAMPolicy).Metadata

	// policy arn for aws managed policy
	// arn:aws-us-gov:iam::aws:policy/aws-service-role/AccessAnalyzerServiceRolePolicy in us gov cloud
	// arn:aws:iam::aws:policy/aws-service-role/AccessAnalyzerServiceRolePolicy in commercial cloud
	if strings.HasPrefix(*policy.Arn, "arn:"+metadata.Partition+":iam::aws:policy") {
		return true, nil
	}

	return false, nil
}

//// TRANSFORM FUNCTIONS

func iamPolicyTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	policy := d.HydrateItem.(kaytu.IAMPolicy).Description.Policy
	var turbotTagsMap map[string]string
	if policy.Tags == nil {
		return nil, nil
	}

	turbotTagsMap = map[string]string{}
	for _, i := range policy.Tags {
		turbotTagsMap[*i.Key] = *i.Value
	}

	return &turbotTagsMap, nil
}

func attachementCountToBool(_ context.Context, d *transform.TransformData) (interface{}, error) {
	if attachementCount, ok := d.Value.(int32); ok {
		if attachementCount == 0 {
			return false, nil
		}
	} else if v, ok := d.Value.(*int32); ok {
		attachementCount := *v
		if attachementCount == 0 {
			return false, nil
		}
	}
	return true, nil
}
