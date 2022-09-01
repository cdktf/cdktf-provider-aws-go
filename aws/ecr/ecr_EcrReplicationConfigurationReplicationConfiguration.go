package ecr


type EcrReplicationConfigurationReplicationConfiguration struct {
	// rule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ecr_replication_configuration#rule EcrReplicationConfiguration#rule}
	Rule interface{} `field:"required" json:"rule" yaml:"rule"`
}

