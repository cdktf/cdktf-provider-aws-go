package eks


type EksNodeGroupTaint struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/eks_node_group#effect EksNodeGroup#effect}.
	Effect *string `field:"required" json:"effect" yaml:"effect"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/eks_node_group#key EksNodeGroup#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/eks_node_group#value EksNodeGroup#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

