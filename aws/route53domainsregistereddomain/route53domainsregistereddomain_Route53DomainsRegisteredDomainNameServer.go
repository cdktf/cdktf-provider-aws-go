package route53domainsregistereddomain


type Route53DomainsRegisteredDomainNameServer struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/route53domains_registered_domain#name Route53DomainsRegisteredDomain#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/route53domains_registered_domain#glue_ips Route53DomainsRegisteredDomain#glue_ips}.
	GlueIps *[]*string `field:"optional" json:"glueIps" yaml:"glueIps"`
}

