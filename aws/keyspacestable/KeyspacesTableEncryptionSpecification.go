package keyspacestable


type KeyspacesTableEncryptionSpecification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.10.0/docs/resources/keyspaces_table#kms_key_identifier KeyspacesTable#kms_key_identifier}.
	KmsKeyIdentifier *string `field:"optional" json:"kmsKeyIdentifier" yaml:"kmsKeyIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.10.0/docs/resources/keyspaces_table#type KeyspacesTable#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

