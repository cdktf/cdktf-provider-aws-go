package rds


type DbInstanceAutomatedBackupsReplicationTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/db_instance_automated_backups_replication#create DbInstanceAutomatedBackupsReplication#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/db_instance_automated_backups_replication#delete DbInstanceAutomatedBackupsReplication#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

