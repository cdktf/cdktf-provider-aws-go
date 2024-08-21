// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package codecommittrigger


type CodecommitTriggerTrigger struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/codecommit_trigger#destination_arn CodecommitTrigger#destination_arn}.
	DestinationArn *string `field:"required" json:"destinationArn" yaml:"destinationArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/codecommit_trigger#events CodecommitTrigger#events}.
	Events *[]*string `field:"required" json:"events" yaml:"events"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/codecommit_trigger#name CodecommitTrigger#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/codecommit_trigger#branches CodecommitTrigger#branches}.
	Branches *[]*string `field:"optional" json:"branches" yaml:"branches"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/codecommit_trigger#custom_data CodecommitTrigger#custom_data}.
	CustomData *string `field:"optional" json:"customData" yaml:"customData"`
}

