package lightsaildistribution


type LightsailDistributionCacheBehaviorSettingsForwardedHeaders struct {
	// The specific headers to forward to your distribution's origin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/lightsail_distribution#headers_allow_list LightsailDistribution#headers_allow_list}
	HeadersAllowList *[]*string `field:"optional" json:"headersAllowList" yaml:"headersAllowList"`
	// The headers that you want your distribution to forward to your origin and base caching on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/lightsail_distribution#option LightsailDistribution#option}
	Option *string `field:"optional" json:"option" yaml:"option"`
}

