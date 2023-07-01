package ec2managedprefixlist


type Ec2ManagedPrefixListEntry struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/ec2_managed_prefix_list#cidr Ec2ManagedPrefixList#cidr}.
	Cidr *string `field:"required" json:"cidr" yaml:"cidr"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/ec2_managed_prefix_list#description Ec2ManagedPrefixList#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

