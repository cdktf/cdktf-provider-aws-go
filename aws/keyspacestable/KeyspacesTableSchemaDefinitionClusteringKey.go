package keyspacestable


type KeyspacesTableSchemaDefinitionClusteringKey struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/keyspaces_table#name KeyspacesTable#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/keyspaces_table#order_by KeyspacesTable#order_by}.
	OrderBy *string `field:"required" json:"orderBy" yaml:"orderBy"`
}

