package glue


type GluePartitionIndexTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_partition_index#create GluePartitionIndex#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_partition_index#delete GluePartitionIndex#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

