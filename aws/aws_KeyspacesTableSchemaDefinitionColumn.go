// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type KeyspacesTableSchemaDefinitionColumn struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/keyspaces_table#name KeyspacesTable#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/keyspaces_table#type KeyspacesTable#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

