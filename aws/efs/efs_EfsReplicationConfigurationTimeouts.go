package efs


type EfsReplicationConfigurationTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/efs_replication_configuration#create EfsReplicationConfiguration#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/efs_replication_configuration#delete EfsReplicationConfiguration#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

