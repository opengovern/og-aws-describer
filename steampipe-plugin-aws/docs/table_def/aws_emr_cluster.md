# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The name of the cluster.</td></tr>
	<tr><td>id</td><td>The unique identifier for the cluster.</td></tr>
	<tr><td>cluster_arn</td><td>The Amazon Resource Name of the cluster.</td></tr>
	<tr><td>state</td><td>The current state of the cluster.</td></tr>
	<tr><td>status</td><td>The current status details about the cluster.</td></tr>
	<tr><td>auto_scaling_role</td><td>An IAM role for automatic scaling policies.</td></tr>
	<tr><td>auto_terminate</td><td>Specifies whether the cluster should terminate after completing all steps.</td></tr>
	<tr><td>custom_ami_id</td><td>Available only in Amazon EMR version 5.7.0 and later. The ID of a custom Amazon EBS-backed Linux AMI if the cluster uses a custom AMI.</td></tr>
	<tr><td>ebs_root_volume_size</td><td>The size of the Amazon EBS root device volume of the Linux AMI that is used for each EC2 instance, in GiB. Available in Amazon EMR version 4.x and later.</td></tr>
	<tr><td>instance_collection_type</td><td>The instance group configuration of the cluster.</td></tr>
	<tr><td>log_encryption_kms_key_id</td><td>The AWS KMS customer master key (CMK) used for encrypting log files. This attribute is only available with EMR version 5.30.0 and later, excluding EMR 6.0.0.</td></tr>
	<tr><td>log_uri</td><td>The path to the Amazon S3 location where logs for this cluster are stored.</td></tr>
	<tr><td>outpost_arn</td><td>The Amazon Resource Name (ARN) of the Outpost where the cluster is launched.</td></tr>
	<tr><td>master_public_dns_name</td><td>The DNS name of the master node.</td></tr>
	<tr><td>normalized_instance_hours</td><td>An approximation of the cost of the cluster, represented in m1.small/hours.</td></tr>
	<tr><td>release_label</td><td>The Amazon EMR release label, which determines the version of open-source application packages installed on the cluster.</td></tr>
	<tr><td>repo_upgrade_on_boot</td><td>Applies only when CustomAmiID is used. Specifies the type of updates that are applied from the Amazon Linux AMI package repositories when an instance boots using the AMI.</td></tr>
	<tr><td>requested_ami_version</td><td>Applies only when CustomAmiID is used. Specifies the type of updates that are applied from the Amazon Linux AMI package repositories when an instance boots using the AMI.</td></tr>
	<tr><td>running_ami_version</td><td>The AMI version running on this cluster.</td></tr>
	<tr><td>scale_down_behavior</td><td>The way that individual Amazon EC2 instances terminate when an automatic scale-in activity occurs or an instance group is resized.</td></tr>
	<tr><td>security_configuration</td><td>The name of the security configuration applied to the cluster.</td></tr>
	<tr><td>service_role</td><td>The IAM role that will be assumed by the Amazon EMR service to access AWS resources on your behalf.</td></tr>
	<tr><td>step_concurrency_level</td><td>Specifies the number of steps that can be executed concurrently.</td></tr>
	<tr><td>termination_protected</td><td>Indicates whether Amazon EMR will lock the cluster to prevent the EC2 instances from being terminated by an API call or user intervention, or in the event of a cluster error.</td></tr>
	<tr><td>visible_to_all_users</td><td>Indicates whether the cluster is visible to all IAM users of the AWS account associated with the cluster.</td></tr>
	<tr><td>applications</td><td>The applications installed on this cluster.</td></tr>
	<tr><td>configurations</td><td>Applies only to Amazon EMR releases 4.x and later. The list of Configurations supplied to the EMR cluster.</td></tr>
	<tr><td>ec2_instance_attributes</td><td>Provides information about the EC2 instances in a cluster grouped by category.</td></tr>
	<tr><td>placement_groups</td><td>Placement group configured for an Amazon EMR cluster.</td></tr>
	<tr><td>kerberos_attributes</td><td>Attributes for Kerberos configuration when Kerberos authentication is enabled using a security configuration.</td></tr>
	<tr><td>tags_src</td><td>A list of tags associated with a cluster.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>tags</td><td>A map of tags for the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in Kaytu.</td></tr>
	<tr><td>kaytu_metadata</td><td>kaytu Metadata of the AWS resource.</td></tr>
</table>