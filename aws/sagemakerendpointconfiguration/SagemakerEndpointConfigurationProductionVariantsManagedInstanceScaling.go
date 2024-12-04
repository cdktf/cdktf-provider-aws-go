// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerendpointconfiguration


type SagemakerEndpointConfigurationProductionVariantsManagedInstanceScaling struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.79.0/docs/resources/sagemaker_endpoint_configuration#max_instance_count SagemakerEndpointConfiguration#max_instance_count}.
	MaxInstanceCount *float64 `field:"optional" json:"maxInstanceCount" yaml:"maxInstanceCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.79.0/docs/resources/sagemaker_endpoint_configuration#min_instance_count SagemakerEndpointConfiguration#min_instance_count}.
	MinInstanceCount *float64 `field:"optional" json:"minInstanceCount" yaml:"minInstanceCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.79.0/docs/resources/sagemaker_endpoint_configuration#status SagemakerEndpointConfiguration#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

