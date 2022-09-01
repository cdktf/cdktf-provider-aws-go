// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type KeyspacesTableEncryptionSpecification struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/keyspaces_table#kms_key_identifier KeyspacesTable#kms_key_identifier}.
	KmsKeyIdentifier *string `field:"optional" json:"kmsKeyIdentifier" yaml:"kmsKeyIdentifier"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/keyspaces_table#type KeyspacesTable#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

