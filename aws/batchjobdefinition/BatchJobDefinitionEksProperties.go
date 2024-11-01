// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package batchjobdefinition


type BatchJobDefinitionEksProperties struct {
	// pod_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/batch_job_definition#pod_properties BatchJobDefinition#pod_properties}
	PodProperties *BatchJobDefinitionEksPropertiesPodProperties `field:"required" json:"podProperties" yaml:"podProperties"`
}

