package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsEmrCluster(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_emr_cluster",
		Description: "AWS EMR Cluster",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidRequestException"}),
			},
			Hydrate: kaytu.GetEMRCluster,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListEMRCluster,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.Name")},
			{
				Name:        "id",
				Description: "The unique identifier for the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.Id")},
			{
				Name:        "cluster_arn",
				Description: "The Amazon Resource Name of the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.ClusterArn")},
			{
				Name:        "status",
				Description: "The current status details about the cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.Status")},
			{
				Name:        "auto_scaling_role",
				Description: "An IAM role for automatic scaling policies.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.AutoScalingRole")},
			{
				Name:        "auto_terminate",
				Description: "Specifies whether the cluster should terminate after completing all steps.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Cluster.AutoTerminate")},
			{
				Name:        "custom_ami_id",
				Description: "Available only in Amazon EMR version 5.7.0 and later. The ID of a custom Amazon EBS-backed Linux AMI if the cluster uses a custom AMI.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.CustomAmiId")},
			{
				Name:        "ebs_root_volume_size",
				Description: "The size of the Amazon EBS root device volume of the Linux AMI that is used for each EC2 instance, in GiB. Available in Amazon EMR version 4.x and later.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.EbsRootVolumeSize")},
			{
				Name:        "instance_collection_type",
				Description: "The instance group configuration of the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.InstanceCollectionType")},
			{
				Name:        "log_encryption_kms_key_id",
				Description: "The AWS KMS customer master key (CMK) used for encrypting log files. This attribute is only available with EMR version 5.30.0 and later, excluding EMR 6.0.0.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.LogEncryptionKmsKeyId")},
			{
				Name:        "log_uri",
				Description: "The path to the Amazon S3 location where logs for this cluster are stored.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.LogUri")},
			{
				Name:        "outpost_arn",
				Description: "The Amazon Resource Name (ARN) of the Outpost where the cluster is launched.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.OutpostArn")},
			{
				Name:        "master_public_dns_name",
				Description: "The DNS name of the master node.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.MasterPublicDnsName")},
			{
				Name:        "normalized_instance_hours",
				Description: "An approximation of the cost of the cluster, represented in m1.small/hours.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Cluster.NormalizedInstanceHours")},
			{
				Name:        "release_label",
				Description: "The Amazon EMR release label, which determines the version of open-source application packages installed on the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.ReleaseLabel")},
			{
				Name:        "repo_upgrade_on_boot",
				Description: "Applies only when CustomAmiID is used. Specifies the type of updates that are applied from the Amazon Linux AMI package repositories when an instance boots using the AMI.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.RepoUpgradeOnBoot")},
			{
				Name:        "requested_ami_version",
				Description: "Applies only when CustomAmiID is used. Specifies the type of updates that are applied from the Amazon Linux AMI package repositories when an instance boots using the AMI.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.RequestedAmiVersion")},
			{
				Name:        "running_ami_version",
				Description: "The AMI version running on this cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.RunningAmiVersion")},
			{
				Name:        "scale_down_behavior",
				Description: "The way that individual Amazon EC2 instances terminate when an automatic scale-in activity occurs or an instance group is resized.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.ScaleDownBehavior")},
			{
				Name:        "security_configuration",
				Description: "The name of the security configuration applied to the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.SecurityConfiguration")},
			{
				Name:        "service_role",
				Description: "The IAM role that will be assumed by the Amazon EMR service to access AWS resources on your behalf.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.ServiceRole")},
			{
				Name:        "step_concurrency_level",
				Description: "Specifies the number of steps that can be executed concurrently.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Cluster.StepConcurrencyLevel")},
			{
				Name:        "termination_protected",
				Description: "Indicates whether Amazon EMR will lock the cluster to prevent the EC2 instances from being terminated by an API call or user intervention, or in the event of a cluster error.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Cluster.TerminationProtected")},
			{
				Name:        "visible_to_all_users",
				Description: "Indicates whether the cluster is visible to all IAM users of the AWS account associated with the cluster.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Cluster.VisibleToAllUsers")},
			{
				Name:        "applications",
				Description: "The applications installed on this cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.Applications")},
			{
				Name:        "configurations",
				Description: "Applies only to Amazon EMR releases 4.x and later. The list of Configurations supplied to the EMR cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.Configurations")},
			{
				Name:        "ec2_instance_attributes",
				Description: "Provides information about the EC2 instances in a cluster grouped by category.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.Ec2InstanceAttributes")},
			{
				Name:        "placement_groups",
				Description: "Placement group configured for an Amazon EMR cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.PlacementGroups")},
			{
				Name:        "kerberos_attributes",
				Description: "Attributes for Kerberos configuration when Kerberos authentication is enabled using a security configuration.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.KerberosAttributes")},
			{
				Name:        "tags_src",
				Description: "A list of tags associated with a cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.Tags")},

			// Standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.Name")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.Tags").Transform(getEmrClusterTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cluster.ClusterArn").Transform(arnToAkas),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func getEmrClusterTurbotTags(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getEmrClusterTurbotTags")
	clusterTags := d.HydrateItem.(kaytu.EMRCluster).Description.Cluster

	if clusterTags == nil {
		return nil, nil
	}

	// Mapping the resource tags inside turbotTags
	var turbotTagsMap map[string]string
	if clusterTags.Tags != nil {
		turbotTagsMap = map[string]string{}
		for _, i := range clusterTags.Tags {
			turbotTagsMap[*i.Key] = *i.Value
		}
	}

	return turbotTagsMap, nil
}
