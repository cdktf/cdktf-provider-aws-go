package keyspacestable


type KeyspacesTableTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/keyspaces_table#create KeyspacesTable#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/keyspaces_table#delete KeyspacesTable#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/keyspaces_table#update KeyspacesTable#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

