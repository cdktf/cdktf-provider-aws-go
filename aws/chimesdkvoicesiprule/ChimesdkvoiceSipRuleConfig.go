// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package chimesdkvoicesiprule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ChimesdkvoiceSipRuleConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/chimesdkvoice_sip_rule#name ChimesdkvoiceSipRule#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// target_applications block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/chimesdkvoice_sip_rule#target_applications ChimesdkvoiceSipRule#target_applications}
	TargetApplications interface{} `field:"required" json:"targetApplications" yaml:"targetApplications"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/chimesdkvoice_sip_rule#trigger_type ChimesdkvoiceSipRule#trigger_type}.
	TriggerType *string `field:"required" json:"triggerType" yaml:"triggerType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/chimesdkvoice_sip_rule#trigger_value ChimesdkvoiceSipRule#trigger_value}.
	TriggerValue *string `field:"required" json:"triggerValue" yaml:"triggerValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/chimesdkvoice_sip_rule#disabled ChimesdkvoiceSipRule#disabled}.
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/chimesdkvoice_sip_rule#id ChimesdkvoiceSipRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

