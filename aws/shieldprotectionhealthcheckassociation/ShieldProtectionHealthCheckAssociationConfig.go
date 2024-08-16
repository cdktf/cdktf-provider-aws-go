// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package shieldprotectionhealthcheckassociation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ShieldProtectionHealthCheckAssociationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/shield_protection_health_check_association#health_check_arn ShieldProtectionHealthCheckAssociation#health_check_arn}.
	HealthCheckArn *string `field:"required" json:"healthCheckArn" yaml:"healthCheckArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/shield_protection_health_check_association#shield_protection_id ShieldProtectionHealthCheckAssociation#shield_protection_id}.
	ShieldProtectionId *string `field:"required" json:"shieldProtectionId" yaml:"shieldProtectionId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/shield_protection_health_check_association#id ShieldProtectionHealthCheckAssociation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

