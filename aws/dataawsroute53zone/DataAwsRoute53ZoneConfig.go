// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsroute53zone

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsRoute53ZoneConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/data-sources/route53_zone#id DataAwsRoute53Zone#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/data-sources/route53_zone#name DataAwsRoute53Zone#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/data-sources/route53_zone#private_zone DataAwsRoute53Zone#private_zone}.
	PrivateZone interface{} `field:"optional" json:"privateZone" yaml:"privateZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/data-sources/route53_zone#tags DataAwsRoute53Zone#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/data-sources/route53_zone#vpc_id DataAwsRoute53Zone#vpc_id}.
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/data-sources/route53_zone#zone_id DataAwsRoute53Zone#zone_id}.
	ZoneId *string `field:"optional" json:"zoneId" yaml:"zoneId"`
}

