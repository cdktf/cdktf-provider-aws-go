// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kmsexternalkey

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KmsExternalKeyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/kms_external_key#bypass_policy_lockout_safety_check KmsExternalKey#bypass_policy_lockout_safety_check}.
	BypassPolicyLockoutSafetyCheck interface{} `field:"optional" json:"bypassPolicyLockoutSafetyCheck" yaml:"bypassPolicyLockoutSafetyCheck"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/kms_external_key#deletion_window_in_days KmsExternalKey#deletion_window_in_days}.
	DeletionWindowInDays *float64 `field:"optional" json:"deletionWindowInDays" yaml:"deletionWindowInDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/kms_external_key#description KmsExternalKey#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/kms_external_key#enabled KmsExternalKey#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/kms_external_key#id KmsExternalKey#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/kms_external_key#key_material_base64 KmsExternalKey#key_material_base64}.
	KeyMaterialBase64 *string `field:"optional" json:"keyMaterialBase64" yaml:"keyMaterialBase64"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/kms_external_key#multi_region KmsExternalKey#multi_region}.
	MultiRegion interface{} `field:"optional" json:"multiRegion" yaml:"multiRegion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/kms_external_key#policy KmsExternalKey#policy}.
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/kms_external_key#tags KmsExternalKey#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/kms_external_key#tags_all KmsExternalKey#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/kms_external_key#valid_to KmsExternalKey#valid_to}.
	ValidTo *string `field:"optional" json:"validTo" yaml:"validTo"`
}

