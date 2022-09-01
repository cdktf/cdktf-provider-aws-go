package ecr


type EcrRegistryScanningConfigurationRule struct {
	// repository_filter block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ecr_registry_scanning_configuration#repository_filter EcrRegistryScanningConfiguration#repository_filter}
	RepositoryFilter interface{} `field:"required" json:"repositoryFilter" yaml:"repositoryFilter"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ecr_registry_scanning_configuration#scan_frequency EcrRegistryScanningConfiguration#scan_frequency}.
	ScanFrequency *string `field:"required" json:"scanFrequency" yaml:"scanFrequency"`
}

