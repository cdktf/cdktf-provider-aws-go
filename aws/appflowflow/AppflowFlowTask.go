// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appflowflow


type AppflowFlowTask struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/appflow_flow#source_fields AppflowFlow#source_fields}.
	SourceFields *[]*string `field:"required" json:"sourceFields" yaml:"sourceFields"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/appflow_flow#task_type AppflowFlow#task_type}.
	TaskType *string `field:"required" json:"taskType" yaml:"taskType"`
	// connector_operator block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/appflow_flow#connector_operator AppflowFlow#connector_operator}
	ConnectorOperator interface{} `field:"optional" json:"connectorOperator" yaml:"connectorOperator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/appflow_flow#destination_field AppflowFlow#destination_field}.
	DestinationField *string `field:"optional" json:"destinationField" yaml:"destinationField"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/appflow_flow#task_properties AppflowFlow#task_properties}.
	TaskProperties *map[string]*string `field:"optional" json:"taskProperties" yaml:"taskProperties"`
}

