package glue


type GluePartitionStorageDescriptorColumns struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_partition#name GluePartition#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_partition#comment GluePartition#comment}.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_partition#type GluePartition#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

