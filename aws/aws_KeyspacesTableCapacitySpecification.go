// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type KeyspacesTableCapacitySpecification struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/keyspaces_table#read_capacity_units KeyspacesTable#read_capacity_units}.
	ReadCapacityUnits *float64 `field:"optional" json:"readCapacityUnits" yaml:"readCapacityUnits"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/keyspaces_table#throughput_mode KeyspacesTable#throughput_mode}.
	ThroughputMode *string `field:"optional" json:"throughputMode" yaml:"throughputMode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/keyspaces_table#write_capacity_units KeyspacesTable#write_capacity_units}.
	WriteCapacityUnits *float64 `field:"optional" json:"writeCapacityUnits" yaml:"writeCapacityUnits"`
}

