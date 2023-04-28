package gluepartitionindex


type GluePartitionIndexTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/glue_partition_index#create GluePartitionIndex#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/glue_partition_index#delete GluePartitionIndex#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

