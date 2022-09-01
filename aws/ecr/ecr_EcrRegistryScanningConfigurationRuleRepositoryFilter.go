package ecr


type EcrRegistryScanningConfigurationRuleRepositoryFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ecr_registry_scanning_configuration#filter EcrRegistryScanningConfiguration#filter}.
	Filter *string `field:"required" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ecr_registry_scanning_configuration#filter_type EcrRegistryScanningConfiguration#filter_type}.
	FilterType *string `field:"required" json:"filterType" yaml:"filterType"`
}

