package flowlog


type FlowLogDestinationOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/flow_log#file_format FlowLog#file_format}.
	FileFormat *string `field:"optional" json:"fileFormat" yaml:"fileFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/flow_log#hive_compatible_partitions FlowLog#hive_compatible_partitions}.
	HiveCompatiblePartitions interface{} `field:"optional" json:"hiveCompatiblePartitions" yaml:"hiveCompatiblePartitions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/flow_log#per_hour_partition FlowLog#per_hour_partition}.
	PerHourPartition interface{} `field:"optional" json:"perHourPartition" yaml:"perHourPartition"`
}

