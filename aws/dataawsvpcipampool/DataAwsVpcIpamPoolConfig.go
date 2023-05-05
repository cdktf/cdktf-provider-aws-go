package dataawsvpcipampool

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsVpcIpamPoolConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/data-sources/vpc_ipam_pool#allocation_resource_tags DataAwsVpcIpamPool#allocation_resource_tags}.
	AllocationResourceTags *map[string]*string `field:"optional" json:"allocationResourceTags" yaml:"allocationResourceTags"`
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/data-sources/vpc_ipam_pool#filter DataAwsVpcIpamPool#filter}
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/data-sources/vpc_ipam_pool#id DataAwsVpcIpamPool#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/data-sources/vpc_ipam_pool#ipam_pool_id DataAwsVpcIpamPool#ipam_pool_id}.
	IpamPoolId *string `field:"optional" json:"ipamPoolId" yaml:"ipamPoolId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/data-sources/vpc_ipam_pool#tags DataAwsVpcIpamPool#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/data-sources/vpc_ipam_pool#timeouts DataAwsVpcIpamPool#timeouts}
	Timeouts *DataAwsVpcIpamPoolTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

