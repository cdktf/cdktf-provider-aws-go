package lightsaildistribution


type LightsailDistributionTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lightsail_distribution#create LightsailDistribution#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lightsail_distribution#delete LightsailDistribution#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lightsail_distribution#update LightsailDistribution#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

