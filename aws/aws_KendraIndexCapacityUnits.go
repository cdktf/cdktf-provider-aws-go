// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type KendraIndexCapacityUnits struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kendra_index#query_capacity_units KendraIndex#query_capacity_units}.
	QueryCapacityUnits *float64 `field:"optional" json:"queryCapacityUnits" yaml:"queryCapacityUnits"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kendra_index#storage_capacity_units KendraIndex#storage_capacity_units}.
	StorageCapacityUnits *float64 `field:"optional" json:"storageCapacityUnits" yaml:"storageCapacityUnits"`
}

