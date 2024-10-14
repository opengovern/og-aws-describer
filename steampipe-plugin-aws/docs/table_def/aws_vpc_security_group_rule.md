# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>security_group_rule_id</td><td>The ID of the security group rule.</td></tr>
	<tr><td>group_name</td><td>The name of the security group to which rule belongs.</td></tr>
	<tr><td>group_id</td><td>The ID of the security group to which rule belongs.</td></tr>
	<tr><td>is_egress</td><td>Indicates whether the security group rule is an outbound rule.</td></tr>
	<tr><td>type</td><td>Type of the rule ( ingress | egress).</td></tr>
	<tr><td>vpc_id</td><td>The ID of the VPC for the security group.</td></tr>
	<tr><td>owner_id</td><td>The AWS account ID of the owner of the security group to which rule belongs.</td></tr>
	<tr><td>group_owner_id</td><td>The ID of the Amazon Web Services account that owns the security group.</td></tr>
	<tr><td>description</td><td>The security group rule description.</td></tr>
	<tr><td>ip_protocol</td><td>The IP protocol name (tcp, udp, icmp, icmpv6) or number [see Protocol Numbers ](http://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml). Use -1 to specify all protocols. When authorizing security group rules, specifying -1 or a protocol number other than tcp, udp, icmp, or icmpv6 allows traffic on all ports, regardless of any port range specified. For tcp, udp, and icmp, a port range is specified. For icmpv6, the port range is optional. If port range is omitted, traffic for all types and codes is allowed.</td></tr>
	<tr><td>from_port</td><td>The start of port range for the TCP and UDP protocols, or an ICMP/ICMPv6 type number. A value of -1 indicates all ICMP/ICMPv6 types.</td></tr>
	<tr><td>to_port</td><td>The end of port range for the TCP and UDP protocols, or an ICMP/ICMPv6 code. A value of -1 indicates all ICMP/ICMPv6 codes.</td></tr>
	<tr><td>cidr_ip</td><td>The IPv4 CIDR range. It can be either a CIDR range or a source security group, not both. A single IPv4 address is denoted by /32 prefix length.</td></tr>
	<tr><td>cidr_ipv4</td><td>The IPv4 CIDR range.</td></tr>
	<tr><td>cidr_ipv6</td><td>The IPv6 CIDR range. It can be either a CIDR range or a source security group, not both. A single IPv6 address is denoted by /128 prefix length.</td></tr>
	<tr><td>pair_group_id</td><td>The ID of the security group that references this user ID group pair.</td></tr>
	<tr><td>referenced_group_id</td><td>The ID of the referenced security group.</td></tr>
	<tr><td>pair_group_name</td><td>The name of the security group that references this user ID group pair.</td></tr>
	<tr><td>pair_peering_status</td><td>The status of a VPC peering connection, if applicable.</td></tr>
	<tr><td>referenced_peering_status</td><td>The status of a VPC peering connection, if applicable.</td></tr>
	<tr><td>pair_user_id</td><td>The ID of an AWS account. For a referenced security group in another VPC, the account ID of the referenced security group is returned in the response. If the referenced security group is deleted, this value is not returned.</td></tr>
	<tr><td>referenced_user_id</td><td>The ID of an AWS account. For a referenced security group in another VPC, the account ID of the referenced security group is returned in the response. If the referenced security group is deleted, this value is not returned.</td></tr>
	<tr><td>pair_vpc_id</td><td>The ID of the VPC for the referenced security group, if applicable.</td></tr>
	<tr><td>referenced_vpc_id</td><td>The ID of the VPC for the referenced security group, if applicable.</td></tr>
	<tr><td>pair_vpc_peering_connection_id</td><td>The ID of the VPC peering connection, if applicable.</td></tr>
	<tr><td>referenced_vpc_peering_connection_id</td><td>The ID of the VPC peering connection, if applicable.</td></tr>
	<tr><td>prefix_list_id</td><td>The ID of the referenced prefix list.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>region</td><td>The AWS Region in which the resource is located.</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in opengovernance.</td></tr>
	<tr><td>kaytu_metadata</td><td>kaytu Metadata of the AWS resource.</td></tr>
</table>