<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The name of the AMI that was provided during image creation.</td></tr>
	<tr><td>image_id</td><td>The ID of the AMI.</td></tr>
	<tr><td>state</td><td>The current state of the AMI. If the state is available, the image is successfully registered and can be used to launch an instance.</td></tr>
	<tr><td>image_type</td><td>The type of image.</td></tr>
	<tr><td>image_location</td><td>The location of the AMI.</td></tr>
	<tr><td>creation_date</td><td>The date and time when the image was created.</td></tr>
	<tr><td>architecture</td><td>The architecture of the image.</td></tr>
	<tr><td>description</td><td>The description of the AMI that was provided during image creation.</td></tr>
	<tr><td>ena_support</td><td>Specifies whether enhanced networking with ENA is enabled.</td></tr>
	<tr><td>hypervisor</td><td>The hypervisor type of the image.</td></tr>
	<tr><td>image_owner_alias</td><td>The AWS account alias (for example, amazon, self) or the AWS account ID of the AMI owner.</td></tr>
	<tr><td>imds_support</td><td>If v2.0, it indicates that IMDSv2 is specified in the AMI.</td></tr>
	<tr><td>kernel_id</td><td>The kernel associated with the image, if any. Only applicable for machine images.</td></tr>
	<tr><td>owner_id</td><td>The AWS account ID of the image owner.</td></tr>
	<tr><td>platform</td><td>This value is set to windows for Windows AMIs; otherwise, it is blank.</td></tr>
	<tr><td>platform_details</td><td>The platform details associated with the billing code of the AMI. For more information, see Obtaining Billing Information (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ami-billing-info.html) in the Amazon Elastic Compute Cloud User Guide.</td></tr>
	<tr><td>public</td><td>Indicates whether the image has public launch permissions. The value is true if this image has public launch permissions or false if it has only implicit and explicit launch permissions.</td></tr>
	<tr><td>ramdisk_id</td><td>The RAM disk associated with the image, if any. Only applicable for machine images.</td></tr>
	<tr><td>root_device_name</td><td>The device name of the root device volume (for example, /dev/sda1).</td></tr>
	<tr><td>root_device_type</td><td>The type of root device used by the AMI. The AMI can use an EBS volume or an instance store volume.</td></tr>
	<tr><td>sriov_net_support</td><td>Specifies whether enhanced networking with the Intel 82599 Virtual Function interface is enabled.</td></tr>
	<tr><td>usage_operation</td><td>The operation of the Amazon EC2 instance and the billing code that is associated with the AMI. For the list of UsageOperation codes, see Platform Details and [Usage Operation Billing Codes](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ami-billing-info.html#billing-info) in the Amazon Elastic Compute Cloud User Guide.</td></tr>
	<tr><td>virtualization_type</td><td>The type of virtualization of the AMI.</td></tr>
	<tr><td>block_device_mappings</td><td>Any block device mapping entries.</td></tr>
	<tr><td>product_codes</td><td>Any product codes associated with the AMI.</td></tr>
	<tr><td>tags_src</td><td>A list of tags attached to the AMI.</td></tr>
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