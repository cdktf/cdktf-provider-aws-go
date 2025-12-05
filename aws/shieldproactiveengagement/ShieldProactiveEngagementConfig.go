// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package shieldproactiveengagement

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ShieldProactiveEngagementConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/shield_proactive_engagement#enabled ShieldProactiveEngagement#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// emergency_contact block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/shield_proactive_engagement#emergency_contact ShieldProactiveEngagement#emergency_contact}
	EmergencyContact interface{} `field:"optional" json:"emergencyContact" yaml:"emergencyContact"`
}

