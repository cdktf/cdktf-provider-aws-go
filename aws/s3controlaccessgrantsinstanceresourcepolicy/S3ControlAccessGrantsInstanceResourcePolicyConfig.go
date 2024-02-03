// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3controlaccessgrantsinstanceresourcepolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type S3ControlAccessGrantsInstanceResourcePolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/s3control_access_grants_instance_resource_policy#policy S3ControlAccessGrantsInstanceResourcePolicy#policy}.
	Policy *string `field:"required" json:"policy" yaml:"policy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/s3control_access_grants_instance_resource_policy#account_id S3ControlAccessGrantsInstanceResourcePolicy#account_id}.
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
}

