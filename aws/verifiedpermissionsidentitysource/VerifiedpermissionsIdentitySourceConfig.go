// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package verifiedpermissionsidentitysource

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VerifiedpermissionsIdentitySourceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/verifiedpermissions_identity_source#policy_store_id VerifiedpermissionsIdentitySource#policy_store_id}.
	PolicyStoreId *string `field:"required" json:"policyStoreId" yaml:"policyStoreId"`
	// configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/verifiedpermissions_identity_source#configuration VerifiedpermissionsIdentitySource#configuration}
	Configuration interface{} `field:"optional" json:"configuration" yaml:"configuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/verifiedpermissions_identity_source#principal_entity_type VerifiedpermissionsIdentitySource#principal_entity_type}.
	PrincipalEntityType *string `field:"optional" json:"principalEntityType" yaml:"principalEntityType"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/verifiedpermissions_identity_source#region VerifiedpermissionsIdentitySource#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

