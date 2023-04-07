package lightsaildistribution


type LightsailDistributionDefaultCacheBehavior struct {
	// The cache behavior of the distribution.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lightsail_distribution#behavior LightsailDistribution#behavior}
	Behavior *string `field:"required" json:"behavior" yaml:"behavior"`
}

