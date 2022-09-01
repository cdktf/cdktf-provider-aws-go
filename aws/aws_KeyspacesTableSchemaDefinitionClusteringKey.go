// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type KeyspacesTableSchemaDefinitionClusteringKey struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/keyspaces_table#name KeyspacesTable#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/keyspaces_table#order_by KeyspacesTable#order_by}.
	OrderBy *string `field:"required" json:"orderBy" yaml:"orderBy"`
}

