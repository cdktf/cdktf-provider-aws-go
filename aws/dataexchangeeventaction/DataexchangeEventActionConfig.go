// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataexchangeeventaction

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataexchangeEventActionConfig struct {
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
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/dataexchange_event_action#action DataexchangeEventAction#action}
	Action interface{} `field:"optional" json:"action" yaml:"action"`
	// event block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/dataexchange_event_action#event DataexchangeEventAction#event}
	Event interface{} `field:"optional" json:"event" yaml:"event"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/dataexchange_event_action#region DataexchangeEventAction#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

