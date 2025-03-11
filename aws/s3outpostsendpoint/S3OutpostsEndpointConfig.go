// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3outpostsendpoint

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type S3OutpostsEndpointConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/s3outposts_endpoint#outpost_id S3OutpostsEndpoint#outpost_id}.
	OutpostId *string `field:"required" json:"outpostId" yaml:"outpostId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/s3outposts_endpoint#security_group_id S3OutpostsEndpoint#security_group_id}.
	SecurityGroupId *string `field:"required" json:"securityGroupId" yaml:"securityGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/s3outposts_endpoint#subnet_id S3OutpostsEndpoint#subnet_id}.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/s3outposts_endpoint#access_type S3OutpostsEndpoint#access_type}.
	AccessType *string `field:"optional" json:"accessType" yaml:"accessType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/s3outposts_endpoint#customer_owned_ipv4_pool S3OutpostsEndpoint#customer_owned_ipv4_pool}.
	CustomerOwnedIpv4Pool *string `field:"optional" json:"customerOwnedIpv4Pool" yaml:"customerOwnedIpv4Pool"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/s3outposts_endpoint#id S3OutpostsEndpoint#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

