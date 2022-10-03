package dataawsec2managedprefixlists


type DataAwsEc2ManagedPrefixListsFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ec2_managed_prefix_lists#name DataAwsEc2ManagedPrefixLists#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ec2_managed_prefix_lists#values DataAwsEc2ManagedPrefixLists#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

