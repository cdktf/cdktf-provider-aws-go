package keyspacestable


type KeyspacesTableClientSideTimestamps struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/keyspaces_table#status KeyspacesTable#status}.
	Status *string `field:"required" json:"status" yaml:"status"`
}

