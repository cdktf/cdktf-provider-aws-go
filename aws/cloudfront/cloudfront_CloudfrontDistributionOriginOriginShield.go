package cloudfront


type CloudfrontDistributionOriginOriginShield struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution#enabled CloudfrontDistribution#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution#origin_shield_region CloudfrontDistribution#origin_shield_region}.
	OriginShieldRegion *string `field:"required" json:"originShieldRegion" yaml:"originShieldRegion"`
}

