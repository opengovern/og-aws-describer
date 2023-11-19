# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>log_group_name</td><td>The name of the log group to which this event belongs.</td></tr>
	<tr><td>log_stream_name</td><td>The name of the log stream to which this event belongs.</td></tr>
	<tr><td>timestamp</td><td>The time when the event occurred.</td></tr>
	<tr><td>version</td><td>The VPC Flow Logs version. If you use the default format, the version is 2. If you use a custom format, the version is the highest version among the specified fields. For example, if you specify only fields from version 2, the version is 2. If you specify a mixture of fields from versions 2, 3, and 4, the version is 4.</td></tr>
	<tr><td>interface_account_id</td><td>The AWS account ID of the owner of the source network interface for which traffic is recorded. If the network interface is created by an AWS service, for example when creating a VPC endpoint or Network Load Balancer, the record may display unknown for this field.</td></tr>
	<tr><td>interface_id</td><td>The ID of the network interface for which the traffic is recorded.</td></tr>
	<tr><td>src_addr</td><td>The source address for incoming traffic, or the IPv4 or IPv6 address of the network interface for outgoing traffic on the network interface. The IPv4 address of the network interface is always its private IPv4 address. See also pkt-srcaddr.</td></tr>
	<tr><td>dst_addr</td><td>The destination address for outgoing traffic, or the IPv4 or IPv6 address of the network interface for incoming traffic on the network interface. The IPv4 address of the network interface is always its private IPv4 address. See also pkt-dstaddr.</td></tr>
	<tr><td>src_port</td><td>The source port of the traffic.</td></tr>
	<tr><td>dst_port</td><td>The destination port of the traffic.</td></tr>
	<tr><td>protocol</td><td>The IANA protocol number of the traffic. For more information, see Assigned Internet Protocol Numbers.</td></tr>
	<tr><td>packets</td><td>The number of packets transferred during the flow.</td></tr>
	<tr><td>bytes</td><td>The number of bytes transferred during the flow.</td></tr>
	<tr><td>start</td><td>The time when the first packet of the flow was received within the aggregation interval. This might be up to 60 seconds after the packet was transmitted or received on the network interface.</td></tr>
	<tr><td>end</td><td>The time when the last packet of the flow was received within the aggregation interval. This might be up to 60 seconds after the packet was transmitted or received on the network interface.</td></tr>
	<tr><td>action</td><td>The action that is associated with the traffic: ACCEPT — The recorded traffic was permitted by the security groups and network ACLs. REJECT — The recorded traffic was not permitted by the security groups or network ACLs.</td></tr>
	<tr><td>log_status</td><td>The logging status of the flow log: OK — Data is logging normally to the chosen destinations. NODATA — There was no network traffic to or from the network interface during the aggregation interval. SKIPDATA — Some flow log records were skipped during the aggregation interval. This may be because of an internal capacity constraint, or an internal error.</td></tr>
	<tr><td>event_id</td><td>The ID of the event.</td></tr>
	<tr><td>filter</td><td>Filter pattern for the search.</td></tr>
	<tr><td>ingestion_time</td><td>The time when the event was ingested.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in Kaytu.</td></tr>
	<tr><td>kaytu_metadata</td><td>kaytu Metadata of the AWS resource.</td></tr>
</table>