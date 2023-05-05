package dataawsec2managedprefixlist


type DataAwsEc2ManagedPrefixListFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/data-sources/ec2_managed_prefix_list#name DataAwsEc2ManagedPrefixList#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/data-sources/ec2_managed_prefix_list#values DataAwsEc2ManagedPrefixList#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

