// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kmsreplicaexternalkey

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KmsReplicaExternalKeyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/kms_replica_external_key#primary_key_arn KmsReplicaExternalKey#primary_key_arn}.
	PrimaryKeyArn *string `field:"required" json:"primaryKeyArn" yaml:"primaryKeyArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/kms_replica_external_key#bypass_policy_lockout_safety_check KmsReplicaExternalKey#bypass_policy_lockout_safety_check}.
	BypassPolicyLockoutSafetyCheck interface{} `field:"optional" json:"bypassPolicyLockoutSafetyCheck" yaml:"bypassPolicyLockoutSafetyCheck"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/kms_replica_external_key#deletion_window_in_days KmsReplicaExternalKey#deletion_window_in_days}.
	DeletionWindowInDays *float64 `field:"optional" json:"deletionWindowInDays" yaml:"deletionWindowInDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/kms_replica_external_key#description KmsReplicaExternalKey#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/kms_replica_external_key#enabled KmsReplicaExternalKey#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/kms_replica_external_key#id KmsReplicaExternalKey#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/kms_replica_external_key#key_material_base64 KmsReplicaExternalKey#key_material_base64}.
	KeyMaterialBase64 *string `field:"optional" json:"keyMaterialBase64" yaml:"keyMaterialBase64"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/kms_replica_external_key#policy KmsReplicaExternalKey#policy}.
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/kms_replica_external_key#tags KmsReplicaExternalKey#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/kms_replica_external_key#tags_all KmsReplicaExternalKey#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/kms_replica_external_key#valid_to KmsReplicaExternalKey#valid_to}.
	ValidTo *string `field:"optional" json:"validTo" yaml:"validTo"`
}

