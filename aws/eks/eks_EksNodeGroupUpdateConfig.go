package eks


type EksNodeGroupUpdateConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/eks_node_group#max_unavailable EksNodeGroup#max_unavailable}.
	MaxUnavailable *float64 `field:"optional" json:"maxUnavailable" yaml:"maxUnavailable"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/eks_node_group#max_unavailable_percentage EksNodeGroup#max_unavailable_percentage}.
	MaxUnavailablePercentage *float64 `field:"optional" json:"maxUnavailablePercentage" yaml:"maxUnavailablePercentage"`
}

