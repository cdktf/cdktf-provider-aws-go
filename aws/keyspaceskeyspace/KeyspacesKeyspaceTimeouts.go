package keyspaceskeyspace


type KeyspacesKeyspaceTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/keyspaces_keyspace#create KeyspacesKeyspace#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/keyspaces_keyspace#delete KeyspacesKeyspace#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

