// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudfrontvpcorigin

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudfrontVpcOriginConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/cloudfront_vpc_origin#tags CloudfrontVpcOrigin#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/cloudfront_vpc_origin#timeouts CloudfrontVpcOrigin#timeouts}
	Timeouts *CloudfrontVpcOriginTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// vpc_origin_endpoint_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/cloudfront_vpc_origin#vpc_origin_endpoint_config CloudfrontVpcOrigin#vpc_origin_endpoint_config}
	VpcOriginEndpointConfig interface{} `field:"optional" json:"vpcOriginEndpointConfig" yaml:"vpcOriginEndpointConfig"`
}

