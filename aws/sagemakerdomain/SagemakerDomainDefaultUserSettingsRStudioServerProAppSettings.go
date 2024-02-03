// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerdomain


type SagemakerDomainDefaultUserSettingsRStudioServerProAppSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/sagemaker_domain#access_status SagemakerDomain#access_status}.
	AccessStatus *string `field:"optional" json:"accessStatus" yaml:"accessStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/sagemaker_domain#user_group SagemakerDomain#user_group}.
	UserGroup *string `field:"optional" json:"userGroup" yaml:"userGroup"`
}

