# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The name of the Availability Zone, Local Zone, or Wavelength Zone</td></tr>
	<tr><td>zone_id</td><td>The ID of the Availability Zone, Local Zone, or Wavelength Zone.</td></tr>
	<tr><td>zone_type</td><td>The type of zone. The valid values are availability-zone, local-zone, and wavelength-zone.</td></tr>
	<tr><td>opt_in_status</td><td>For Availability Zones, this parameter always has the value of opt-in-not-required. For Local Zones and Wavelength Zones, this parameter is the opt-in status. The possible values are opted-in, and not-opted-in.</td></tr>
	<tr><td>group_name</td><td>For Availability Zones, this parameter has the same value as the Region name. For Local Zones, the name of the associated group, for example us-west-2-lax-1. For Wavelength Zones, the name of the associated group, for example us-east-1-wl1-bos-wlz-1.</td></tr>
	<tr><td>region_name</td><td>The name of the Region.</td></tr>
	<tr><td>parent_zone_name</td><td>The name of the zone that handles some of the Local Zone or Wavelength Zone control plane operations, such as API calls.</td></tr>
	<tr><td>parent_zone_id</td><td>The ID of the zone that handles some of the Local Zone or Wavelength Zone control plane operations, such as API calls</td></tr>
	<tr><td>messages</td><td>Any messages about the Availability Zone, Local Zone, or Wavelength Zone.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
</table>