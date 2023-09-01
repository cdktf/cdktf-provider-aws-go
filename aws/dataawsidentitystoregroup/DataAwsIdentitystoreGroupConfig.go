// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsidentitystoregroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsIdentitystoreGroupConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/data-sources/identitystore_group#identity_store_id DataAwsIdentitystoreGroup#identity_store_id}.
	IdentityStoreId *string `field:"required" json:"identityStoreId" yaml:"identityStoreId"`
	// alternate_identifier block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/data-sources/identitystore_group#alternate_identifier DataAwsIdentitystoreGroup#alternate_identifier}
	AlternateIdentifier *DataAwsIdentitystoreGroupAlternateIdentifier `field:"optional" json:"alternateIdentifier" yaml:"alternateIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/data-sources/identitystore_group#group_id DataAwsIdentitystoreGroup#group_id}.
	GroupId *string `field:"optional" json:"groupId" yaml:"groupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/data-sources/identitystore_group#id DataAwsIdentitystoreGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

