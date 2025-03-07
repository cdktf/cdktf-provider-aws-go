// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudwatchlogdeliverydestinationpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudwatchLogDeliveryDestinationPolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/cloudwatch_log_delivery_destination_policy#delivery_destination_name CloudwatchLogDeliveryDestinationPolicy#delivery_destination_name}.
	DeliveryDestinationName *string `field:"required" json:"deliveryDestinationName" yaml:"deliveryDestinationName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/cloudwatch_log_delivery_destination_policy#delivery_destination_policy CloudwatchLogDeliveryDestinationPolicy#delivery_destination_policy}.
	DeliveryDestinationPolicy *string `field:"required" json:"deliveryDestinationPolicy" yaml:"deliveryDestinationPolicy"`
}

