package lightsaildistribution


type LightsailDistributionCacheBehaviorSettingsForwardedQueryStrings struct {
	// Indicates whether the distribution forwards and caches based on query strings.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lightsail_distribution#option LightsailDistribution#option}
	Option interface{} `field:"optional" json:"option" yaml:"option"`
	// The specific query strings that the distribution forwards to the origin.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lightsail_distribution#query_strings_allowed_list LightsailDistribution#query_strings_allowed_list}
	QueryStringsAllowedList *[]*string `field:"optional" json:"queryStringsAllowedList" yaml:"queryStringsAllowedList"`
}

