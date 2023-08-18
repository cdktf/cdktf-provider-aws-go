package docdbcluster


type DocdbClusterTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/docdb_cluster#create DocdbCluster#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/docdb_cluster#delete DocdbCluster#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/docdb_cluster#update DocdbCluster#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

