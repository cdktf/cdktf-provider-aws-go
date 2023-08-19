package dataawsimagebuilderdistributionconfigurations


type DataAwsImagebuilderDistributionConfigurationsFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/data-sources/imagebuilder_distribution_configurations#name DataAwsImagebuilderDistributionConfigurations#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/data-sources/imagebuilder_distribution_configurations#values DataAwsImagebuilderDistributionConfigurations#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

