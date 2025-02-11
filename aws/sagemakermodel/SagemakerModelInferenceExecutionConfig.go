// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakermodel


type SagemakerModelInferenceExecutionConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.86.1/docs/resources/sagemaker_model#mode SagemakerModel#mode}.
	Mode *string `field:"required" json:"mode" yaml:"mode"`
}

