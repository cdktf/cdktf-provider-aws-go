// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type KeyspacesTableSchemaDefinitionPartitionKey struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/keyspaces_table#name KeyspacesTable#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
}

