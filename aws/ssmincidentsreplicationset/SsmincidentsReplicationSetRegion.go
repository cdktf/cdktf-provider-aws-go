package ssmincidentsreplicationset


type SsmincidentsReplicationSetRegion struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssmincidents_replication_set#name SsmincidentsReplicationSet#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssmincidents_replication_set#kms_key_arn SsmincidentsReplicationSet#kms_key_arn}.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

