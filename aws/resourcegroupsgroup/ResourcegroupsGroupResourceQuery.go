package resourcegroupsgroup


type ResourcegroupsGroupResourceQuery struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/resourcegroups_group#query ResourcegroupsGroup#query}.
	Query *string `field:"required" json:"query" yaml:"query"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/resourcegroups_group#type ResourcegroupsGroup#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

