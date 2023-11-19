# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The fully qualified name of the directory.</td></tr>
	<tr><td>directory_id</td><td>The directory identifier.</td></tr>
	<tr><td>arn</td><td>The Amazon Resource Name (ARN) that uniquely identifies the directory.</td></tr>
	<tr><td>stage</td><td>The current stage of the directory.</td></tr>
	<tr><td>type</td><td>The directory type.</td></tr>
	<tr><td>access_url</td><td>The access URL for the directory, such as http://<alias>.awsapps.com.</td></tr>
	<tr><td>alias</td><td>The alias for the directory.</td></tr>
	<tr><td>description</td><td>The description for the directory.</td></tr>
	<tr><td>desired_number_of_domain_controllers</td><td>The desired number of domain controllers in the directory if the directory is Microsoft AD.</td></tr>
	<tr><td>edition</td><td>The edition associated with this directory.</td></tr>
	<tr><td>launch_time</td><td>Specifies when the directory was created.</td></tr>
	<tr><td>radius_status</td><td>The status of the RADIUS MFA server connection.</td></tr>
	<tr><td>share_method</td><td>The method used when sharing a directory to determine whether the directory should be shared within your AWS organization (ORGANIZATIONS) or with any AWS account by sending a shared directory request (HANDSHAKE).</td></tr>
	<tr><td>share_notes</td><td>A directory share request that is sent by the directory owner to the directory consumer.</td></tr>
	<tr><td>share_status</td><td>Current directory status of the shared AWS Managed Microsoft AD directory.</td></tr>
	<tr><td>short_name</td><td>The short name of the directory.</td></tr>
	<tr><td>size</td><td>The directory size.</td></tr>
	<tr><td>sso_enabled</td><td>Indicates if single sign-on is enabled for the directory. For more information, see EnableSso and DisableSso.</td></tr>
	<tr><td>stage_last_updated_date_time</td><td>The date and time that the stage was last updated.</td></tr>
	<tr><td>stage_reason</td><td>Additional information about the directory stage.</td></tr>
	<tr><td>connect_settings</td><td>A DirectoryConnectSettingsDescription object that contains additional information about an AD Connector directory.</td></tr>
	<tr><td>dns_ip_addrs</td><td>The IP addresses of the DNS servers for the directory.</td></tr>
	<tr><td>event_topics</td><td>Amazon SNS topic names that receive status messages from the specified Directory ID.</td></tr>
	<tr><td>snapshot_limit</td><td>Obtains the manual snapshot limits for a directory.</td></tr>
	<tr><td>owner_directory_description</td><td>Describes the AWS Managed Microsoft AD directory in the directory owner account.</td></tr>
	<tr><td>radius_settings</td><td>A RadiusSettings object that contains information about the RADIUS server.</td></tr>
	<tr><td>regions_info</td><td>Lists the Regions where the directory has replicated.</td></tr>
	<tr><td>shared_directories</td><td>Details about the shared directory in the directory owner account for which the share request in the directory consumer account has been accepted.</td></tr>
	<tr><td>vpc_settings</td><td>A DirectoryVpcSettingsDescription object that contains additional information about a directory.</td></tr>
	<tr><td>tags_src</td><td>A list of tags currently associated with the Directory Service Directory.</td></tr>
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