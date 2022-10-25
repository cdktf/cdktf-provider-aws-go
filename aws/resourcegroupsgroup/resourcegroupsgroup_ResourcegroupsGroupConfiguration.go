package resourcegroupsgroup


type ResourcegroupsGroupConfiguration struct {
	// parameters block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/resourcegroups_group#parameters ResourcegroupsGroup#parameters}
	Parameters interface{} `field:"required" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/resourcegroups_group#type ResourcegroupsGroup#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

