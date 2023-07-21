package resourcegroupsgroup


type ResourcegroupsGroupConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/resourcegroups_group#type ResourcegroupsGroup#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/resourcegroups_group#parameters ResourcegroupsGroup#parameters}
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
}

