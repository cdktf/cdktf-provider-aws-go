package neptunecluster


type NeptuneClusterServerlessV2ScalingConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/neptune_cluster#max_capacity NeptuneCluster#max_capacity}.
	MaxCapacity *float64 `field:"optional" json:"maxCapacity" yaml:"maxCapacity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/neptune_cluster#min_capacity NeptuneCluster#min_capacity}.
	MinCapacity *float64 `field:"optional" json:"minCapacity" yaml:"minCapacity"`
}

