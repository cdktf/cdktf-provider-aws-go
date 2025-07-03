// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerdomain


type SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsEmrSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/sagemaker_domain#assumable_role_arns SagemakerDomain#assumable_role_arns}.
	AssumableRoleArns *[]*string `field:"optional" json:"assumableRoleArns" yaml:"assumableRoleArns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/sagemaker_domain#execution_role_arns SagemakerDomain#execution_role_arns}.
	ExecutionRoleArns *[]*string `field:"optional" json:"executionRoleArns" yaml:"executionRoleArns"`
}

