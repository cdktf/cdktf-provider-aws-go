// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accountregion

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AccountRegionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/account_region#enabled AccountRegion#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/account_region#region_name AccountRegion#region_name}.
	RegionName *string `field:"required" json:"regionName" yaml:"regionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/account_region#account_id AccountRegion#account_id}.
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/account_region#id AccountRegion#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/account_region#timeouts AccountRegion#timeouts}
	Timeouts *AccountRegionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

