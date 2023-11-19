# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The name of the image.</td></tr>
	<tr><td>arn</td><td>The Amazon Resource Name (ARN) of the image.</td></tr>
	<tr><td>appstream_agent_version</td><td>The version of the AppStream 2.0 agent to use for instances that are launched from this image.</td></tr>
	<tr><td>base_image_arn</td><td>The ARN of the image from which this image was created.</td></tr>
	<tr><td>created_time</td><td>The time the image was created.</td></tr>
	<tr><td>description</td><td>The description to display.</td></tr>
	<tr><td>display_name</td><td>The image name to display.</td></tr>
	<tr><td>image_builder_name</td><td>The name of the image builder that was used to create the private image. If the image is shared, this value is null.</td></tr>
	<tr><td>image_builder_supported</td><td>Indicates whether an image builder can be launched from this image.</td></tr>
	<tr><td>platform</td><td>The operating system platform of the image.</td></tr>
	<tr><td>public_base_image_released_date</td><td>The release date of the public base image. For private images, this date is the release date of the base image from which the image was created.</td></tr>
	<tr><td>state</td><td>The image starts in the PENDING state. If image creation succeeds, the state is AVAILABLE. If image creation fails, the state is FAILED.</td></tr>
	<tr><td>visibility</td><td>Indicates whether the image is public or private.</td></tr>
	<tr><td>applications</td><td>The applications associated with the image.</td></tr>
	<tr><td>image_errors</td><td>Describes the errors that are returned when a new image can&#39;t be created.</td></tr>
	<tr><td>image_permissions</td><td>The permissions to provide to the destination AWS account for the specified image.</td></tr>
	<tr><td>state_change_reason</td><td>The reason why the last state change occurred.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>tags</td><td>A map of tags for the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
</table>