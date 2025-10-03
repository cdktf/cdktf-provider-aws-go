// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerdomain


type SagemakerDomainDefaultUserSettingsKernelGatewayAppSettings struct {
	// custom_image block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/sagemaker_domain#custom_image SagemakerDomain#custom_image}
	CustomImage interface{} `field:"optional" json:"customImage" yaml:"customImage"`
	// default_resource_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/sagemaker_domain#default_resource_spec SagemakerDomain#default_resource_spec}
	DefaultResourceSpec *SagemakerDomainDefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpec `field:"optional" json:"defaultResourceSpec" yaml:"defaultResourceSpec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/sagemaker_domain#lifecycle_config_arns SagemakerDomain#lifecycle_config_arns}.
	LifecycleConfigArns *[]*string `field:"optional" json:"lifecycleConfigArns" yaml:"lifecycleConfigArns"`
}

