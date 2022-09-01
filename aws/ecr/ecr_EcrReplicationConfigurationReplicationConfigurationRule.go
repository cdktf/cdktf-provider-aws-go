package ecr


type EcrReplicationConfigurationReplicationConfigurationRule struct {
	// destination block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ecr_replication_configuration#destination EcrReplicationConfiguration#destination}
	Destination interface{} `field:"required" json:"destination" yaml:"destination"`
	// repository_filter block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ecr_replication_configuration#repository_filter EcrReplicationConfiguration#repository_filter}
	RepositoryFilter interface{} `field:"optional" json:"repositoryFilter" yaml:"repositoryFilter"`
}

