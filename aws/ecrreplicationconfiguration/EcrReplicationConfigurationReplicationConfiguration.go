package ecrreplicationconfiguration


type EcrReplicationConfigurationReplicationConfiguration struct {
	// rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/ecr_replication_configuration#rule EcrReplicationConfiguration#rule}
	Rule interface{} `field:"required" json:"rule" yaml:"rule"`
}

