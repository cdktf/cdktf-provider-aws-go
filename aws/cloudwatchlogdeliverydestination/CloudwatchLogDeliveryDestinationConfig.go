// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudwatchlogdeliverydestination

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudwatchLogDeliveryDestinationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/cloudwatch_log_delivery_destination#name CloudwatchLogDeliveryDestination#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// delivery_destination_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/cloudwatch_log_delivery_destination#delivery_destination_configuration CloudwatchLogDeliveryDestination#delivery_destination_configuration}
	DeliveryDestinationConfiguration interface{} `field:"optional" json:"deliveryDestinationConfiguration" yaml:"deliveryDestinationConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/cloudwatch_log_delivery_destination#delivery_destination_type CloudwatchLogDeliveryDestination#delivery_destination_type}.
	DeliveryDestinationType *string `field:"optional" json:"deliveryDestinationType" yaml:"deliveryDestinationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/cloudwatch_log_delivery_destination#output_format CloudwatchLogDeliveryDestination#output_format}.
	OutputFormat *string `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/cloudwatch_log_delivery_destination#region CloudwatchLogDeliveryDestination#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/cloudwatch_log_delivery_destination#tags CloudwatchLogDeliveryDestination#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

