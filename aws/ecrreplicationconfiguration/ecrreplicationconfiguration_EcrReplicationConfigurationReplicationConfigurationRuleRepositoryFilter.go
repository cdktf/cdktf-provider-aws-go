package ecrreplicationconfiguration


type EcrReplicationConfigurationReplicationConfigurationRuleRepositoryFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ecr_replication_configuration#filter EcrReplicationConfiguration#filter}.
	Filter *string `field:"required" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ecr_replication_configuration#filter_type EcrReplicationConfiguration#filter_type}.
	FilterType *string `field:"required" json:"filterType" yaml:"filterType"`
}

