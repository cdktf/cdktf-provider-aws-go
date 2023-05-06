package dataawsvpclatticeservicenetwork

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsVpclatticeServiceNetworkConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/data-sources/vpclattice_service_network#service_network_identifier DataAwsVpclatticeServiceNetwork#service_network_identifier}.
	ServiceNetworkIdentifier *string `field:"required" json:"serviceNetworkIdentifier" yaml:"serviceNetworkIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/data-sources/vpclattice_service_network#tags DataAwsVpclatticeServiceNetwork#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

