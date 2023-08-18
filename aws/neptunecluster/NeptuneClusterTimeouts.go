package neptunecluster


type NeptuneClusterTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/neptune_cluster#create NeptuneCluster#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/neptune_cluster#delete NeptuneCluster#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/neptune_cluster#update NeptuneCluster#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

