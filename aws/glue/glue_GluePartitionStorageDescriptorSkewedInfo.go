package glue


type GluePartitionStorageDescriptorSkewedInfo struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_partition#skewed_column_names GluePartition#skewed_column_names}.
	SkewedColumnNames *[]*string `field:"optional" json:"skewedColumnNames" yaml:"skewedColumnNames"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_partition#skewed_column_value_location_maps GluePartition#skewed_column_value_location_maps}.
	SkewedColumnValueLocationMaps *map[string]*string `field:"optional" json:"skewedColumnValueLocationMaps" yaml:"skewedColumnValueLocationMaps"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_partition#skewed_column_values GluePartition#skewed_column_values}.
	SkewedColumnValues *[]*string `field:"optional" json:"skewedColumnValues" yaml:"skewedColumnValues"`
}

