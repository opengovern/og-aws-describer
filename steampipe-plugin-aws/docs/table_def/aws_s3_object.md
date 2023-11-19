# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>key</td><td>The name that you assign to an object. You use the object key to retrieve the object.</td></tr>
	<tr><td>arn</td><td>The ARN of the AWS S3 Object.</td></tr>
	<tr><td>bucket_name</td><td>The name of the container bucket of this object.</td></tr>
	<tr><td>last_modified</td><td>Last modified time of the object.</td></tr>
	<tr><td>storage_class</td><td>The class of storage used to store the object.</td></tr>
	<tr><td>version_id</td><td>The version ID of the object.</td></tr>
	<tr><td>accept_ranges</td><td>Indicates that a range of bytes was specified.</td></tr>
	<tr><td>body</td><td>The raw bytes of the object data as a string. If the bytes entirely consists of valid UTF8 runes, an UTF8 is sent otherwise the bas64 encoding of the bytes is sent.</td></tr>
	<tr><td>bucket_key_enabled</td><td>Indicates whether the object uses an S3 Bucket Key for server-side encryption with Amazon Web Services KMS (SSE-KMS)</td></tr>
	<tr><td>cache_control</td><td>Specifies caching behavior along the request/reply chain.</td></tr>
	<tr><td>checksum_crc32</td><td>The base64-encoded, 32-bit CRC32 checksum of the object. This will only be present if it was uploaded with the object. With multipart uploads, this may not be a checksum value of the object.</td></tr>
	<tr><td>checksum_crc32c</td><td>The base64-encoded, 32-bit CRC32C checksum of the object. This will only be present if it was uploaded with the object. With multipart uploads, this may not be a checksum value of the object.</td></tr>
	<tr><td>checksum_sha1</td><td>The base64-encoded, 160-bit SHA-1 digest of the object. This will only be present if it was uploaded with the object. With multipart uploads, this may not be a checksum value of the object.</td></tr>
	<tr><td>checksum_sha256</td><td>The base64-encoded, 256-bit SHA-256 digest of the object. This will only be present if it was uploaded with the object. With multipart uploads, this may not be a checksum value of the object.</td></tr>
	<tr><td>content_disposition</td><td>Specifies presentational information for the object.</td></tr>
	<tr><td>content_encoding</td><td>Specifies what content encodings have been applied to the object.</td></tr>
	<tr><td>content_language</td><td>The language the content is in.</td></tr>
	<tr><td>content_length</td><td>Size of the body in bytes.</td></tr>
	<tr><td>content_range</td><td>The portion of the object returned in the response.</td></tr>
	<tr><td>content_type</td><td>A standard MIME type describing the format of the object data.</td></tr>
	<tr><td>delete_marker</td><td>Specifies whether the object retrieved was (true) or was not (false) a delete marker.</td></tr>
	<tr><td>expiration</td><td>If the object expiration is configured (see PUT Bucket lifecycle), the response includes this header. It includes the expiry-date and rule-id key-value pairs providing object expiration information. The value of the rule-id is URL-encoded.</td></tr>
	<tr><td>expires</td><td>The date and time at which the object is no longer cacheable.</td></tr>
	<tr><td>etag</td><td>The entity tag of the object.</td></tr>
	<tr><td>object_lock_legal_hold_status</td><td>Like a retention period, a legal hold prevents an object version from being overwritten or deleted. A legal hold remains in effect until removed.</td></tr>
	<tr><td>object_lock_mode</td><td>The Object Lock mode currently in place for this object.</td></tr>
	<tr><td>object_lock_retain_until_date</td><td>The date and time when this object&#39;s Object Lock will expire.</td></tr>
	<tr><td>parts_count</td><td>The count of parts this object has. This value is only returned if you specify partNumber in your request and the object was uploaded as a multipart upload.</td></tr>
	<tr><td>prefix</td><td>The prefix of the key of the object.</td></tr>
	<tr><td>replication_status</td><td>Amazon S3 can return this if your request involves a bucket that is either a source or destination in a replication rule.</td></tr>
	<tr><td>request_charged</td><td>If present, indicates that the requester was successfully charged for the request.</td></tr>
	<tr><td>restore</td><td>Provides information about object restoration action and expiration time of the restored object copy.</td></tr>
	<tr><td>server_side_encryption</td><td>The server-side encryption algorithm used when storing this object in Amazon S3.</td></tr>
	<tr><td>size</td><td>Size in bytes of the object.</td></tr>
	<tr><td>sse_customer_algorithm</td><td>If server-side encryption with a customer-provided encryption key was requested, the response will include this header confirming the encryption algorithm used.</td></tr>
	<tr><td>sse_customer_key_md5</td><td>If server-side encryption with a customer-provided encryption key was requested, the response will include this header to provide round-trip message integrity verification of the customer-provided encryption key.</td></tr>
	<tr><td>sse_kms_key_id</td><td>If present, specifies the ID of the Amazon Web Services Key Management Service(Amazon Web Services KMS) symmetric customer managed key that was used for the object.</td></tr>
	<tr><td>tag_count</td><td>The number of tags, if any, on the object.</td></tr>
	<tr><td>website_redirection_location</td><td>If the bucket is configured as a website, redirects requests for this object  to another object in the same bucket or to an external URL.</td></tr>
	<tr><td>acl</td><td>ACLs define which AWS accounts or groups are granted access along with the type of access.</td></tr>
	<tr><td>checksum</td><td>The checksum or digest of the object.</td></tr>
	<tr><td>metadata</td><td>A map of metadata to store with the object in S3.</td></tr>
	<tr><td>object_parts</td><td>A collection of parts associated with a multipart upload.</td></tr>
	<tr><td>owner</td><td>The owner of the object.</td></tr>
	<tr><td>tags_src</td><td>A list of tags assigned to the object.</td></tr>
	<tr><td>tags</td><td>A map of tags for the resource.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>region</td><td>The AWS Region in which the object is located.</td></tr>
	<tr><td>partition</td><td>The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).</td></tr>
	<tr><td>account_id</td><td>The AWS Account ID in which the resource is located.</td></tr>
</table>