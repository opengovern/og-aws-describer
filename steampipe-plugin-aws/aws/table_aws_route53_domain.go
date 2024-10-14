package aws

import (
	"context"
	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsRoute53Domain(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_route53_domain",
		Description: "AWS Route53 Domain",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("domain_name"),
			Hydrate:    opengovernance.GetRoute53Domain,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListRoute53Domain,
		},
		Columns: awsKaytuColumns([]*plugin.Column{
			{
				Name:        "domain_name",
				Description: "The name of the domain.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Domain.DomainName")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) specifying the domain.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ARN"),
			},
			{
				Name:        "abuse_contact_email",
				Description: "Email address to contact to report incorrect contact information for a domain,to report that the domain is being used to send spam, to report that someone is cyber squatting on a domain name, or report some other type of abuse.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Domain.AbuseContactEmail")},
			{
				Name:        "abuse_contact_phone",
				Description: "Phone number for reporting abuse.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Domain.AbuseContactPhone")},
			{
				Name:        "admin_privacy",
				Description: "Specifies whether contact information is concealed from WHOIS queries.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Domain.AdminPrivacy")},
			{
				Name:        "auto_renew",
				Description: "Indicates whether the domain is automatically renewed upon expiration.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Domain.AutoRenew")},
			// As of May 25, 2021, API doesn't return the DNSSEC configuration in response.
			// {
			// 	Name:        "dns_sec",
			// 	Description: "Reserved for future use.",
			// 	Type:        proto.ColumnType_STRING,
			// 	Hydrate:     getRoute53Domain,
			// },
			{
				Name:        "creation_date",
				Description: "The date when the domain was created as found in the response to a WHOIS query.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Domain.CreationDate")},
			{
				Name:        "expiration_date",
				Description: "The date when the registration for the domain is set to expire. The date and time is in Unix time format and Coordinated Universal time (UTC).",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Domain.ExpirationDate")},
			{
				Name:        "registrant_privacy",
				Description: "Specifies whether contact information is concealed from WHOIS queries.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Domain.RegistrantPrivacy")},
			{
				Name:        "registrar_name",
				Description: "Name of the registrar of the domain as identified in the registry. Domains with a .com, .net, or .org TLD are registered by Amazon Registrar.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Domain.RegistrarName")},
			{
				Name:        "registrar_url",
				Description: "Web address of the registrar.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Domain.RegistrarUrl")},
			{
				Name:        "registry_domain_id",
				Description: "Reserved for future use.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Domain.RegistryDomainId")},
			{
				Name:        "reseller",
				Description: "Reseller of the domain. Domains registered or transferred using Route 53 domains will have Amazon as the reseller.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Domain.Reseller")},
			{
				Name:        "tech_privacy",
				Description: "Specifies whether contact information is concealed from WHOIS queries.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Domain.TechPrivacy")},
			{
				Name:        "transfer_lock",
				Description: "Indicates whether a domain is locked from unauthorized transfer to another party.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.DomainSummary.TransferLock")},
			{
				Name:        "updated_date",
				Description: "The last updated date of the domain as found in the response to a WHOIS query.The date and time is in Unix time format and Coordinated Universal time (UTC).",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Domain.UpdatedDate")},
			{
				Name:        "who_is_server",
				Description: "The fully qualified name of the WHOIS server that can answer the WHOIS query for the domain.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Domain.WhoIsServer")},
			{
				Name:        "nameservers",
				Description: "The name of the domain.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Domain.Nameservers")},
			{
				Name:        "registrant_contact",
				Description: "Provides details about the domain registrant.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Domain.RegistrantContact")},
			{
				Name:        "status_list",
				Description: "An array of domain name status codes, also known as Extensible Provisioning Protocol (EPP) status codes.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Domain.StatusList")},
			{
				Name:        "tech_contact",
				Description: "Provides details about the domain technical contact.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Domain.TechContact")},
			{
				Name:        "admin_contact",
				Description: "Provides details about the domain administrative contact.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Domain.AdminContact")},
			{
				Name:        "tags_src",
				Description: "A list of tags assigned to the resource.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags")},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Domain.DomainName")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags").Transform(route53DomainTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(transform.EnsureStringArray),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func route53DomainTurbotTags(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.HydrateItem.(opengovernance.Route53Domain).Description.Tags
	// Mapping the resource tags inside turbotTags
	var turbotTagsMap map[string]string
	if tags != nil {
		turbotTagsMap = map[string]string{}
		for _, i := range tags {
			turbotTagsMap[*i.Key] = *i.Value
		}
	}
	return turbotTagsMap, nil
}
