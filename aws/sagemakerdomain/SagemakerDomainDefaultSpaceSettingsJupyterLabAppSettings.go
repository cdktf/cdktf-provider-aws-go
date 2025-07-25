// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerdomain


type SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettings struct {
	// app_lifecycle_management block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/sagemaker_domain#app_lifecycle_management SagemakerDomain#app_lifecycle_management}
	AppLifecycleManagement *SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsAppLifecycleManagement `field:"optional" json:"appLifecycleManagement" yaml:"appLifecycleManagement"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/sagemaker_domain#built_in_lifecycle_config_arn SagemakerDomain#built_in_lifecycle_config_arn}.
	BuiltInLifecycleConfigArn *string `field:"optional" json:"builtInLifecycleConfigArn" yaml:"builtInLifecycleConfigArn"`
	// code_repository block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/sagemaker_domain#code_repository SagemakerDomain#code_repository}
	CodeRepository interface{} `field:"optional" json:"codeRepository" yaml:"codeRepository"`
	// custom_image block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/sagemaker_domain#custom_image SagemakerDomain#custom_image}
	CustomImage interface{} `field:"optional" json:"customImage" yaml:"customImage"`
	// default_resource_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/sagemaker_domain#default_resource_spec SagemakerDomain#default_resource_spec}
	DefaultResourceSpec *SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsDefaultResourceSpec `field:"optional" json:"defaultResourceSpec" yaml:"defaultResourceSpec"`
	// emr_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/sagemaker_domain#emr_settings SagemakerDomain#emr_settings}
	EmrSettings *SagemakerDomainDefaultSpaceSettingsJupyterLabAppSettingsEmrSettings `field:"optional" json:"emrSettings" yaml:"emrSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/sagemaker_domain#lifecycle_config_arns SagemakerDomain#lifecycle_config_arns}.
	LifecycleConfigArns *[]*string `field:"optional" json:"lifecycleConfigArns" yaml:"lifecycleConfigArns"`
}

