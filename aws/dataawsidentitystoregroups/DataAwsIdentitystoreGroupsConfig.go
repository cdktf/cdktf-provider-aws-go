// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsidentitystoregroups

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsIdentitystoreGroupsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/data-sources/identitystore_groups#identity_store_id DataAwsIdentitystoreGroups#identity_store_id}.
	IdentityStoreId *string `field:"required" json:"identityStoreId" yaml:"identityStoreId"`
}

