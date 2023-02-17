package ecrreplicationconfiguration


type EcrReplicationConfigurationReplicationConfigurationRuleDestination struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ecr_replication_configuration#region EcrReplicationConfiguration#region}.
	Region *string `field:"required" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ecr_replication_configuration#registry_id EcrReplicationConfiguration#registry_id}.
	RegistryId *string `field:"required" json:"registryId" yaml:"registryId"`
}

