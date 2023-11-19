# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>directory_id</td><td>The directory identifier.</td></tr>
	<tr><td>certificate_id</td><td>The identifier of the certificate.</td></tr>
	<tr><td>common_name</td><td>The common name for the certificate.</td></tr>
	<tr><td>type</td><td>The function that the registered certificate performs. Valid values include ClientLDAPS or ClientCertAuth. The default value is ClientLDAPS.</td></tr>
	<tr><td>state</td><td>The state of the certificate. Valid values: Registering | Registered | RegisterFailed | Deregistering | Deregistered | DeregisterFailed.</td></tr>
	<tr><td>expiry_date_time</td><td>The date and time when the certificate will expire.</td></tr>
	<tr><td>registered_date_time</td><td>The date and time that the certificate was registered.</td></tr>
	<tr><td>state_reason</td><td>Describes a state change for the certificate.</td></tr>
	<tr><td>client_cert_auth_settings</td><td>A ClientCertAuthSettings object that contains client certificate authentication settings.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
</table>