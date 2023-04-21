package cloudhsmv2cluster


type CloudhsmV2ClusterTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/cloudhsm_v2_cluster#create CloudhsmV2Cluster#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/cloudhsm_v2_cluster#delete CloudhsmV2Cluster#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/cloudhsm_v2_cluster#update CloudhsmV2Cluster#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

